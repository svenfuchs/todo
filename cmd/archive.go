package cmd

import (
  . "github.com/svenfuchs/todo/data"
  . "github.com/svenfuchs/todo/io"
)

func NewArchiveCmd(args *Args) Runnable {
  args.Status = Done
  if args.Date == "" {
    args.Date = "before:two weeks ago"
  }

  src := NewIo(args.Path)
  out := NewStdErr()
  archive := NewIo(args.Config["archive"])
  return ArchiveCmd{ Cmd { args, src, out }, archive }
}

type ArchiveCmd struct {
  Cmd
  archive Io
}

func (c ArchiveCmd) Run() {
  list := c.list()
  dump := list.Select(c.filter())
  keep := list.Reject(c.filter())

  c.append(c.archive, dump)
  c.write(c.src, keep)
  c.write(c.out, dump)
}

