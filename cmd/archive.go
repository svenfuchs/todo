package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)

func NewArchiveCmd(path string, filter Filter, opts map[string]string, config map[string]string) ArchiveCmd {
  filter.Status = Done
  if filter.Date.IsEmpty() {
    filter.Date = NewFilterDate("two weeks ago", "before")
  }

  format, ok := opts["format"]
  if !ok {
    format = "full"
  }

  src := NewIo(path)
  out := NewIo("")
  archive := NewIo(config["archive"])
  return ArchiveCmd{ Cmd { src, out, filter, format }, archive }
}

type ArchiveCmd struct {
  Cmd
  archive Io
}

func (c ArchiveCmd) Run() {
  list := c.list()
  dump := list.Select(c.filter)
  keep := list.Reject(c.filter)

  c.append(c.archive, dump, c.format)
  c.write(c.src, keep, c.format)
  c.write(c.out, dump, c.format)
}

