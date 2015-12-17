package cmd

import (
  . "github.com/svenfuchs/todo/io"
)

func NewListCmd(args *Args) Runnable {
  src := NewIo(args.Path)
  out := NewStdErr()
  return ListCmd{ Cmd { args, src, out } }
}

type ListCmd struct {
  Cmd
}

func (c ListCmd) Run() {
  list := c.list()
  list = list.Select(c.filter())
  c.write(c.out, list)
}
