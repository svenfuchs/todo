package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)


func NewToggleCmd(path string, filter Filter) ToggleCmd {
  src := NewIo(path)
  out := NewIo(path)
  // out := NewFileIo("")
  return ToggleCmd{ Cmd { src, out, filter, "full" } }
}

type ToggleCmd struct {
  Cmd
}

func (c ToggleCmd) Run() {
  list := c.list()
  list = list.Toggle(c.filter)
  c.output(list, "full")
}

