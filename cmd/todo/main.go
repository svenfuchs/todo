package main

import (
  "os"
  "gopkg.in/alecthomas/kingpin.v2"
  . "github.com/svenfuchs/todo/cmd"
)

func main() {
  app := kingpin.New("todo", "A command-line todo.txt tool.")

  a := NewArgs()
  c := app.Command("archive", "Archive done todo items."      ).Action(call(a.Run, "archive"))
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&a.Path)
	c.Flag("archive",  "File to archive to."                    ).Short('a').SetValue(&funcValue{ a.SetArchive })
	c.Flag("format",   "Output format."                         ).Short('o').StringVar(&a.Format)
	c.Flag("before",   "Filter by done before."                 ).Short('b').SetValue(&funcValue{ a.SetBefore })

  a = NewArgs()
	c = app.Command("list", "Filter and list todo items."       ).Action(call(a.Run, "list"))
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&a.Path)
	c.Flag("format",   "Output format."                         ).Short('o').StringVar(&a.Format)
	c.Flag("id",       "Filter by id."                          ).Short('i').IntsVar(&a.Ids)
	c.Flag("status",   "Filter by status."                      ).Short('s').StringVar(&a.Status)
	c.Flag("text",     "Filter by text."                        ).Short('t').StringVar(&a.Text)
	c.Flag("projects", "Filter by projects (comma separated)."  ).Short('p').StringsVar(&a.Projects)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ a.SetDate   })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ a.SetAfter  })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ a.SetSince  })
	c.Flag("before",   "Filter by done before."                 ).Short('b').SetValue(&funcValue{ a.SetBefore })
	c.Arg("input",     "Filter by full line."                   ).StringVar(&a.Line)

  a = NewArgs()
	c = app.Command("push", "Push todo items.")
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&a.Path)
	c.Flag("format",   "Output format."                         ).Short('o').StringVar(&a.Format)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ a.SetDate  })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ a.SetAfter })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ a.SetSince })

  s := c.Command("idonethis", "Service Idonethis"             ).Action(call(a.Run, "push")).PreAction(call(a.SetService, "idonethis"))
  s.Flag("username", "Idonethis username"                     ).Envar("TODO_IDONETHIS_USERNAME").SetValue(&funcValue{ a.SetUser })
  s.Flag("token",    "Idonethis token"                        ).Envar("TODO_IDONETHIS_TOKEN"   ).SetValue(&funcValue{ a.SetToken })
  s.Flag("team",     "Idonethis team"                         ).Envar("TODO_IDONETHIS_TEAM"    ).SetValue(&funcValue{ a.SetTeam })
	s.Arg("input",     "Filter by full line."                   ).StringVar(&a.Line)

  a = NewArgs()
	c = app.Command("toggle", "Toggle todo items."              ).Action(call(a.Run, "toggle"))
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&a.Path)
	c.Flag("id",       "Filter by id."                          ).Short('i').IntsVar(&a.Ids)
	c.Flag("status",   "Filter by status."                      ).Short('s').StringVar(&a.Status)
	c.Flag("text",     "Filter by text."                        ).Short('t').StringVar(&a.Text)
	c.Flag("projects", "Filter by projects (comma separated)."  ).Short('p').StringsVar(&a.Projects)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ a.SetDate   })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ a.SetAfter  })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ a.SetSince  })
	c.Flag("before",   "Filter by done before."                 ).Short('b').SetValue(&funcValue{ a.SetBefore })
	c.Arg("input",     "Filter by full line."                   ).StringVar(&a.Line)

  kingpin.MustParse(app.Parse(os.Args[1:]))
}

// For the sub-sub command `todo push [service]` i need to set config["service"]
// as a PreAction.

func call(f func(string), value string) kingpin.Action {
  return func(p *kingpin.ParseContext) error {
    f(value)
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
