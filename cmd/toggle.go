package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/source"
)


func NewToggleCmd(path string, filter Filter) ToggleCmd {
  in  := NewSource(path)
  out := NewSource(path)
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

