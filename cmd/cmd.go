package cmd

import (
  "github.com/svenfuchs/todo.go/format"
  "github.com/svenfuchs/todo.go/item"
  "github.com/svenfuchs/todo.go/source"
)

type Runnable interface {
  Run()
}

type Cmd struct {
  In source.Io
  Out source.Io
  Filter item.Filter
  Format string
}

func (c Cmd) list() item.List {
  lines := c.In.MustReadLines()
  list  := item.ParseList(lines)
  return list
}

func (c Cmd) output(list item.List) {
  lines := c.formatted(list.Items)
  c.Out.WriteLines(lines)
}

func (c Cmd) formatted(items []item.Item) []string {
  return format.New(c.Format).Apply(items)
}

func NewList(path string, filter item.Filter, format string) List {
  in   := source.New(path)
  out  := source.New("")
  return List{ Cmd { in, out, filter, format } }
}

type List struct {
  Cmd
}

func (c List) Run() {
  list  := c.list()
  list   = list.Select(c.Filter)
  // list   = list.SortBy(func (item item.Item) string { return item.DoneDate() }) TODO
  c.output(list)
}

type Toggle struct {
  Cmd
}

func (c Toggle) Run() {
  c.Format = "full"
  list := c.list()
  list  = list.Toggle(c.Filter)
  c.output(list)
}
