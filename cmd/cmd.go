package cmd

import (
  . "github.com/svenfuchs/todo/data"
  . "github.com/svenfuchs/todo/io"
)

type Runnable interface {
  Run()
}

type Cmd struct {
  args *Args
  src Io
  out Io
}

func (c Cmd) list() List {
  lines := c.src.ReadLines()
  list  := ParseItemList(lines)
  return list
}

func (c Cmd) write(io Io, list List) {
  lines := c.formatted(list.Items)
  io.WriteLines(lines)
}

func (c Cmd) append(io Io, list List) {
  lines := c.formatted(list.Items)
  io.AppendLines(lines)
}

func (c Cmd) formatted(items []Item) []string {
  format := c.args.Format
  if format == "" {
    format = "full"
  }
  return NewFormat(format).Apply(items)
}

func (c Cmd) filter() Filter {
  if c.args.Line == "" {
    return NewFilter(c.args.Ids, c.args.Status, c.args.Text, c.args.Projects, c.args.Date)
  } else {
    parser := NewParser(c.args.Line)
    return NewFilter([]int{ parser.Id() }, parser.Status(), parser.Text(), []string{}, "")
  }
}
