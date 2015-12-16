package cmd

import (
  . "github.com/svenfuchs/todo"
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
  lines := c.in.MustReadLines()
  list  := ParseItemList(lines)
  return list
}

func (c Cmd) output(list List, format string) {
  lines := c.formatted(list.Items, format)
  c.out.WriteLines(lines)
}

func (c Cmd) formatted(items []Item, format string) []string {
  return NewFormat(format).Apply(items)
}


