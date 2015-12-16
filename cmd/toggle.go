package cmd

import (
  . "github.com/svenfuchs/todo"
)


func NewToggleCmd(path string, filter Filter) ToggleCmd {
  in  := NewFileSource(path)
  out := NewFileSource(path)
  // out := NewFileSource("")
  return ToggleCmd{ Cmd { in, out, filter, "full" } }
}

type ToggleCmd struct {
  Cmd
}

func (c ToggleCmd) Run() {
  list := c.list()
  list = list.Toggle(c.filter)
  c.output(list, "full")
}

