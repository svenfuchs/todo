package main

import (
  "os"
  "gopkg.in/alecthomas/kingpin.v2"
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/cmd"
)

type Opts struct {
  path string
  format string
  line string
  id int
  status string
  text string
  projects []string
  date FilterDate
  config map[string]string
}

func (o Opts) filter() Filter {
  return NewFilter(o.id, o.status, o.text, o.projects, o.date)
}

func (o *Opts) setDate(date string)     { o.date = NewFilterDate("date", date)   }
func (o *Opts) setBefore(date string)   { o.date = NewFilterDate("before", date) }
func (o *Opts) setSince(date string)    { o.date = NewFilterDate("since", date)  }
func (o *Opts) setAfter(date string)    { o.date = NewFilterDate("after", date)  }

func (o *Opts) setUser(value string)    { o.config["username"] = value }
func (o *Opts) setToken(value string)   { o.config["token"] = value }
func (o *Opts) setTeam(value string)    { o.config["team"] = value }
func (o *Opts) setService(value string) { o.config["service"] = value }

func (o *Opts) runList(c *kingpin.ParseContext) error {
  NewListCmd(o.path, o.filter(), o.format).Run()
	return nil
}

func (o *Opts) runToggle(c *kingpin.ParseContext) error {
  NewToggleCmd(o.path, o.filter()).Run()
	return nil
}

func (o *Opts) runPush(c *kingpin.ParseContext) error {
  o.status = Done
  if o.date.IsEmpty() {
    o.date = NewFilterDate("yesterday", "since")
  }
  NewPushCmd(o.path, o.filter(), o.config).Run()
	return nil
}

func main() {
  a := kingpin.New("todo", "A command-line todo.txt tool.")

	h := &Opts{}
	c := a.Command("list", "Filter and list todo items."        ).Action(h.runList)
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&h.path)
	c.Flag("format",   "Output format."                         ).Short('o').StringVar(&h.format)
	c.Flag("id",       "Filter by id."                          ).Short('i').IntVar(&h.id)
	c.Flag("status",   "Filter by status."                      ).Short('s').StringVar(&h.status)
	c.Flag("text",     "Filter by text."                        ).Short('t').StringVar(&h.text)
	c.Flag("projects", "Filter by projects (comma separated)."  ).Short('p').StringsVar(&h.projects)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ h.setDate   })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ h.setAfter  })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ h.setSince  })
	c.Flag("before",   "Filter by done before."                 ).Short('b').SetValue(&funcValue{ h.setBefore })
	c.Arg("input",     "Filter by full line."                   ).StringVar(&h.line)

	h = &Opts{}
	c = a.Command("toggle", "Toggle todo items."                ).Action(h.runToggle)
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&h.path)
	c.Flag("id",       "Filter by id."                          ).Short('i').IntVar(&h.id)
	c.Flag("status",   "Filter by status."                      ).Short('s').StringVar(&h.status)
	c.Flag("text",     "Filter by text."                        ).Short('t').StringVar(&h.text)
	c.Flag("projects", "Filter by projects (comma separated)."  ).Short('p').StringsVar(&h.projects)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ h.setDate   })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ h.setAfter  })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ h.setSince  })
	c.Flag("before",   "Filter by done before."                 ).Short('b').SetValue(&funcValue{ h.setBefore })
	c.Arg("input",     "Filter by full line."                   ).StringVar(&h.line)

  h = &Opts{ config: map[string]string{} }
	c = a.Command("push", "Push todo items.")
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&h.path)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ h.setDate  })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ h.setAfter })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ h.setSince })

  s := c.Command("idonethis", "Service Idonethis"             ).Action(h.runPush).PreAction(setFunc(h.setService, "idonethis"))
  s.Flag("username", "Idonethis username"                     ).Envar("TODO_IDONETHIS_USERNAME").SetValue(&funcValue{ h.setUser })
  s.Flag("token",    "Idonethis token"                        ).Envar("TODO_IDONETHIS_TOKEN"   ).SetValue(&funcValue{ h.setToken })
  s.Flag("team",     "Idonethis team"                         ).Envar("TODO_IDONETHIS_TEAM"    ).SetValue(&funcValue{ h.setTeam })
	s.Arg("input",     "Filter by full line."                   ).StringVar(&h.line)

  kingpin.MustParse(a.Parse(os.Args[1:]))
}

// For the sub-sub command `todo push [service]` i need to set config["service"]
// as a PreAction.

func setFunc(f func(string), name string) kingpin.Action {
  return func(p *kingpin.ParseContext) error {
    f(name)
    return nil
  }
}

// This uses an internal (?) kingpin api to set values through function
// callbacks instead of directly forcing them onto a struct field. Not sure if
// I'm supposed to write a custom parser instead?

type funcValue struct{
  f func(string)
}

func (f *funcValue) Set(value string) error {
	value, err := value, error(nil)
	if err != nil { return err }
  f.f((string)(value))
  return nil
}

func (f *funcValue) String() string {
  return "no idea?"
}
