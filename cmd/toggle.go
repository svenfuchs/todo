package cmd

import (
  . "github.com/svenfuchs/todo/io"
)


func NewToggleCmd(args *Args) Runnable {
  args.Format = "full"
  src := NewIo(args.Path)
  out := NewStdErr()
  return ToggleCmd{ Cmd { args, src, out } }
}

type ToggleCmd struct {
  Cmd
}

func (c ToggleCmd) Run() {
  list := c.list()
  list.Toggle(c.filter())
  res := list.Toggle(c.filter())

  c.write(c.src, list)
  c.report(c.out, "toggle", res)
}
