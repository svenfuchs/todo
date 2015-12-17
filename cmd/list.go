package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)

func NewListCmd(path string, filter Filter, opts map[string]string, config map[string]string) ListCmd {
  format, ok := opts["format"]
  if !ok {
    format = "full"
  }

  src := NewIo(path)
  out := NewIo("")
  return ListCmd{ Cmd { src, out, filter, format } }
}

type ListCmd struct {
  Cmd
}

func (c ListCmd) Run() {
  list := c.list()
  list = list.Select(c.filter)
  c.write(c.out, list, c.format)
}

