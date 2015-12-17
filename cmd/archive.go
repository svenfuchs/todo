package cmd

import (
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)

func NewArchiveCmd(path string, filter Filter, format string, config map[string]string) ArchiveCmd {
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

