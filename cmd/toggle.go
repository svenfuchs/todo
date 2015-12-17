package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)


func NewToggleCmd(path string, filter Filter) ToggleCmd {
  src := NewIo(path)
  // out := NewFileIo("")
  return ToggleCmd{ Cmd { src, nil, filter, "full" } }
}

type ToggleCmd struct {
  Cmd
}

func (c ToggleCmd) Run() {
  list := c.list()
  list = list.Toggle(c.filter)
  c.write(c.src, list, "full")
}

