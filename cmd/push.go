package cmd

import (
	. "github.com/svenfuchs/todo/data"
	. "github.com/svenfuchs/todo/io"
	. "github.com/svenfuchs/todo/service"
)

func NewPushCmd(args *Args) Runnable {
	src := NewIo(args.Path)
	out := NewStdErr()
	srv := NewService(args.Config)
	return PushCmd{Cmd{args, src, out}, srv}
}

type PushCmd struct {
	Cmd
	srv Service
}

func (c PushCmd) Run() {
	c.initArgs(c.args)

	list := c.list()
	list = list.Select(c.filter())
	list = list.Reject(Filter{Ids: c.ids()})

	c.srv.WriteLines(c.formatted(list.Items))
	c.report(c.out, "push", list)
}

func (c PushCmd) initArgs(args *Args) {
	args.Status = Done
	if args.Date == "" {
		args.Date = "since:yesterday"
	}
	if args.Format == "" {
		args.Format = "text,tags,id"
	}
}

func (c PushCmd) ids() []int {
	ids := []int{}
	lines := c.srv.ReadLines()
	for _, line := range lines {
		ids = append(ids, ParseItem(line).Id)
	}
	return ids
}

func includesInt(nums []int, num int) bool {
	for _, n := range nums {
		if num == n {
			return true
		}
	}
	return false
}
