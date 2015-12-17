package cmd

import (
  . "github.com/svenfuchs/todo"
)

func NewArgs() *Args {
  return &Args{ Config: map[string]string{} }
}

type Args struct {
  Path string
  Line string
  Ids []int
  Status string
  Text string
  Projects []string
  Date string
  Format string
  Config map[string]string
}

func (a *Args) SetDate(date string)     { a.Date = "date:"   + date }
func (a *Args) SetBefore(date string)   { a.Date = "before:" + date }
func (a *Args) SetSince(date string)    { a.Date = "since:"  + date }
func (a *Args) SetAfter(date string)    { a.Date = "after:"  + date }

func (a *Args) SetUser(value string)    { a.Config["username"] = value }
func (a *Args) SetToken(value string)   { a.Config["token"]    = value }
func (a *Args) SetTeam(value string)    { a.Config["team"]     = value }
func (a *Args) SetService(value string) { a.Config["service"]  = value }
func (a *Args) SetArchive(value string) { a.Config["archive"]  = value }

var cmdFactories = map[string]func(string, Filter, string, map[string]string) Runnable {
  "archive": NewArchiveCmd,
  "list":    NewListCmd,
  "push":    NewPushCmd,
  "toggle":  NewToggleCmd,
}

func (a *Args) Run(cmd string) {
  cmdFactories[cmd](a.Path, a.filter(), a.Format, a.Config).Run()
}

func (a Args) filter() Filter {
  return NewFilter(a.Ids, a.Status, a.Text, a.Projects, a.Date)
}
