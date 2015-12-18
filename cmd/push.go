package cmd

import (
  . "github.com/svenfuchs/todo/data"
  . "github.com/svenfuchs/todo/service"
  . "github.com/svenfuchs/todo/io"
)


func NewPushCmd(args *Args) Runnable {
  args.Status = Done
  if args.Date == "" {
    args.Date = "since:yesterday"
  }

  src := NewIo(args.Path)
  out := NewStdErr()
  return PushCmd{ Cmd { args, src, out } }
}

type PushCmd struct {
  Cmd
}

func (c PushCmd) Run() {
  service := NewService(c.args.Config)
  ids := c.ids(service)

  list := c.list()
  list = list.Select(c.filter())
  list = list.Reject(Filter{ Ids: ids })

  c.push(list, service)
  c.report(c.out, "push", list)
}

func (c PushCmd) push(list List, service Service) {
  for _, item := range list.Items {
    service.Push(item.Line) // TODO use format
  }
}

func (c PushCmd) ids(service Service) []int {
  ids   := []int{}
  lines := service.Fetch()
  for _, line := range lines {
    ids = append(ids, ParseItem(line).Id)
  }
  return ids
}

func includesInt(nums []int, num int) bool {
  for _, n := range nums {
    if num == n { return true }
  }
  return false
}
