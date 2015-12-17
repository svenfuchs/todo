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

func (c Cmd) write(io Io, list List, format string) {
  lines := c.formatted(list.Items, format)
  io.WriteLines(lines)
}

func (c Cmd) append(io Io, list List, format string) {
  lines := c.formatted(list.Items, format)
  io.AppendLines(lines)
}

func (c Cmd) formatted(items []Item, format string) []string {
  if format == "" {
    format = "full"
  }
  return NewFormat(format).Apply(items)
}


