package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)


func NewToggleCmd(path string, filter Filter) ToggleCmd {
  src := NewIo(path)
  out := NewIo("")
  return ToggleCmd{ Cmd { src, out, filter, "full" } }
}

type ToggleCmd struct {
  Cmd
}

func (c ToggleCmd) Run() {
  list := c.list()
  selected := list.Toggle(c.filter)
  c.write(c.out, selected, c.format)
  c.write(c.src, list, c.format)
}

