package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)

func NewListCmd(path string, filter Filter, format string) ListCmd {
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
  // list = list.SortBy(func (item Item) string { return item.DoneDate() }) TODO
  c.output(list, c.format)
}

