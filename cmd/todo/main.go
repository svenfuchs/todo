package main

import (
  "os"
  "gopkg.in/alecthomas/kingpin.v2"
  "github.com/svenfuchs/todo.go/cmd"
  "github.com/svenfuchs/todo.go/item"
)

type Opts struct {
  path string
  format string
  line string
  id int
  status string
  text string
  projects string
  after string
  before string
  due string
}

func (o Opts) filter() item.Filter {
  if o.status == "pending" {
    o.status = "pend"
  }
  return item.NewFilter(o.line, o.id, o.status, o.text, o.projects, o.after, o.before, o.due)
}

type ListCmd struct {
  Opts
}

func (l *ListCmd) run(c *kingpin.ParseContext) error {
  cmd.NewList(l.path, l.filter(), l.format).Run()
	return nil
}

func initListCmd(app *kingpin.Application) {
	h := &ListCmd{}
	c := app.Command("list", "Filter and list todo items.").Action(h.run)
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&h.path)
	c.Flag("format",   "Output format."                         ).Short('o').StringVar(&h.format)
	c.Flag("id",       "Filter by id."                          ).Short('i').IntVar(&h.id)
	c.Flag("status",   "Filter by id."                          ).Short('s').StringVar(&h.status)
	c.Flag("text",     "Filter by id."                          ).Short('t').StringVar(&h.text)
	c.Flag("projects", "Filter by projects (comma separated)."  ).Short('p').StringVar(&h.projects)
	c.Flag("after",    "Filter by done after."                  ).Short('a').StringVar(&h.after)
	c.Flag("since",    "Filter by done since (alias for after).").StringVar(&h.after)
	c.Flag("before",   "Filter by done before."                 ).Short('b').StringVar(&h.before)
	c.Flag("due",      "Filter by due date."                    ).Short('d').StringVar(&h.due)
	c.Arg("input",     "Filter by full line."                   ).StringVar(&h.line)
}

func main() {
  app := kingpin.New("todo", "A command-line todo.txt tool.")
  initListCmd(app)
  kingpin.MustParse(app.Parse(os.Args[1:]))
}

