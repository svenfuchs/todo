package cmd

import (
  . "github.com/svenfuchs/todo/io"
)


func NewToggleCmd(args *Args) Runnable {
  args.Format = "full"
  src := NewIo(args.Path)
  out := NewStdIo()
  return ToggleCmd{ Cmd { args, src, out } }
}

type ToggleCmd struct {
  Cmd
}

func (c ToggleCmd) Run() {
  list := c.list()
  selected := list.Toggle(c.filter())
  c.write(c.out, selected)
  c.write(c.src, list)
}

