package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)

type Runnable interface {
  Run()
}

type Cmd struct {
  src Io
  out Io
  filter Filter
  format string
}

func (c Cmd) list() List {
  lines, _ := c.src.ReadLines()
  list := ParseItemList(lines)
  return list
}

func (c Cmd) output(list List, format string) {
  lines := c.formatted(list.Items, format)
  c.out.WriteLines(lines)
}

func (c Cmd) formatted(items []Item, format string) []string {
  return NewFormat(format).Apply(items)
}


