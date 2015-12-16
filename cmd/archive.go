package cmd

import (
  // "fmt"
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/source"
)

// rename Source to Io? because it's not just a "source"
// rename c.in to c.src
// make Filter work with arrays on each field, remove list.RejectIf

func NewArchiveCmd(path string, filter Filter, format string, config map[string]string) ArchiveCmd {
  in  := NewSource(path)
  out := NewSource("")
  archive := NewSource(config["archive"])
  return ArchiveCmd{ Cmd { in, out, filter, format }, archive }
}

type ArchiveCmd struct {
  Cmd
  archive Source
}

func (c ArchiveCmd) Run() {
  list := c.list()
  arch := list.Select(c.filter)
  keep := list.Reject(c.filter)

  c.archive.AppendLines(c.formatted(arch.Items, c.format))
  c.in.WriteLines(c.formatted(keep.Items, c.format))
  c.output(arch, c.format)
}

