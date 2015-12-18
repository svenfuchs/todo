package cmd

import (
  . "github.com/svenfuchs/todo/data"
  . "github.com/svenfuchs/todo/io"
)

func NewArchiveCmd(args *Args) Runnable {
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
  c.initArgs(c.args)

  list := c.list()
  arch := list.Select(c.filter())
  keep := list.Reject(c.filter())

  c.append(c.archive, arch)
  c.write(c.src, keep)
  c.report(c.out, "archive", arch)
}

func (c ArchiveCmd) initArgs(args *Args) {
  args.Status = Done
  if args.Date == "" {
    args.Date = "before:two weeks ago"
  }
}
