package cmd

import (
  // "fmt"
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/io"
)

// rename c.in to c.src
// make Filter work with arrays on each field, remove list.RejectIf

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
  arch := list.Select(c.filter)
  keep := list.Reject(c.filter)

  c.archive.AppendLines(c.formatted(arch.Items, c.format))
  c.src.WriteLines(c.formatted(keep.Items, c.format))
  c.output(arch, c.format)
}

