package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/source"
)

type Runnable interface {
  Run()
}

type Cmd struct {
  in Source
  out Source
  filter Filter
  format string
}

func (c Cmd) list() List {
  lines, _ := c.in.ReadLines()
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


