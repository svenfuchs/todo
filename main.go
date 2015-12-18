package main

import (
  "os"
  "gopkg.in/alecthomas/kingpin.v2"
  . "github.com/svenfuchs/todo/cmd"
)

func main() {
  app := kingpin.New("todo", "A command-line todo.txt tool.")

  a := NewArgs()
  c := app.Command("archive", "Archive done todo items."      ).Action(action(a.Run, "archive"))
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&a.Path)
	c.Flag("archive",  "File to archive to."                    ).Short('a').SetValue(&funcValue{ set(a.SetConfig, "archive") })
	c.Flag("format",   "Output format."                         ).Short('o').StringVar(&a.Format)
	c.Flag("before",   "Filter by done before."                 ).Short('b').SetValue(&funcValue{ set(a.SetDate, "before") })
  c.Flag("report",   "Print a report to stderr."              ).Short('r').Default("true").BoolVar(&a.Report)

  a = NewArgs()
	c = app.Command("list", "Filter and list todo items."       ).Action(action(a.Run, "list"))
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&a.Path)
	c.Flag("format",   "Output format."                         ).Short('o').StringVar(&a.Format)
	c.Flag("id",       "Filter by id."                          ).Short('i').IntsVar(&a.Ids)
	c.Flag("status",   "Filter by status."                      ).Short('s').StringVar(&a.Status)
	c.Flag("text",     "Filter by text."                        ).Short('t').StringVar(&a.Text)
	c.Flag("projects", "Filter by projects (comma separated)."  ).Short('p').StringsVar(&a.Projects)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ set(a.SetDate, "date")   })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ set(a.SetDate, "after")  })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ set(a.SetDate, "since")  })
	c.Flag("before",   "Filter by done before."                 ).Short('b').SetValue(&funcValue{ set(a.SetDate, "before") })
	c.Arg("input",     "Filter by full line."                   ).StringVar(&a.Line)

  a = NewArgs()
	c = app.Command("push", "Push todo items.")
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&a.Path)
	c.Flag("format",   "Output format."                         ).Short('o').StringVar(&a.Format)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ set(a.SetDate, "date")   })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ set(a.SetDate, "after")  })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ set(a.SetDate, "since")  })
  c.Flag("report",   "Print a report to stderr."              ).Short('r').Default("true").BoolVar(&a.Report)

  s := c.Command("idonethis", "Service Idonethis"             ).Action(action(a.Run, "push")).PreAction(action(a.SetConfig, "service", "idonethis"))
  s.Flag("username", "Idonethis username"                     ).Envar("TODO_IDONETHIS_USERNAME").SetValue(&funcValue{ set(a.SetConfig, "username") })
  s.Flag("token",    "Idonethis token"                        ).Envar("TODO_IDONETHIS_TOKEN"   ).SetValue(&funcValue{ set(a.SetConfig, "token") })
  s.Flag("team",     "Idonethis team"                         ).Envar("TODO_IDONETHIS_TEAM"    ).SetValue(&funcValue{ set(a.SetConfig, "team") })
  c.Flag("report",   "Print a report to stderr."              ).Short('r').Default("true").BoolVar(&a.Report)
	s.Arg("input",     "Filter by full line."                   ).StringVar(&a.Line)

  a = NewArgs()
	c = app.Command("toggle", "Toggle todo items."              ).Action(action(a.Run, "toggle"))
	c.Flag("file",     "Todo.txt file to work with."            ).Short('f').StringVar(&a.Path)
	c.Flag("id",       "Filter by id."                          ).Short('i').IntsVar(&a.Ids)
	c.Flag("status",   "Filter by status."                      ).Short('s').StringVar(&a.Status)
	c.Flag("text",     "Filter by text."                        ).Short('t').StringVar(&a.Text)
	c.Flag("projects", "Filter by projects (comma separated)."  ).Short('p').StringsVar(&a.Projects)
	c.Flag("date",     "Filter by done date."                   ).Short('a').SetValue(&funcValue{ set(a.SetDate, "date")   })
	c.Flag("after",    "Filter by done after."                  ).Short('a').SetValue(&funcValue{ set(a.SetDate, "after")  })
	c.Flag("since",    "Filter by done since."                  ).Short('n').SetValue(&funcValue{ set(a.SetDate, "since")  })
	c.Flag("before",   "Filter by done before."                 ).Short('b').SetValue(&funcValue{ set(a.SetDate, "before") })
  c.Flag("report",   "Print a report to stderr."              ).Short('r').Default("true").BoolVar(&a.Report)
	c.Arg("input",     "Filter by full line."                   ).StringVar(&a.Line)

  kingpin.MustParse(app.Parse(os.Args[1:]))
}

// For run methods i'd like to use a dispatch pattern. For the sub-sub command
// `todo push [service]` i need to set config["service"] as a PreAction.

func action(f func(...string), str ...string) kingpin.Action {
  return func(p *kingpin.ParseContext) error {
    f(str...)
    return nil
  }
}

func set(f func(args ...string), key string) func(string) {
  return func(value string) {
    f(key, value)
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
