package cmd

import (
  "time"
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/service"
  . "github.com/svenfuchs/todo/source"
)


func NewPushCmd(path string, filter Filter, config map[string]string) PushCmd {
  in  := NewSource(path)
  out := NewSource("")
  return PushCmd{ Cmd { in, out, filter, "full" }, config }
}

type PushCmd struct {
  Cmd
  config map[string]string
}

func (c PushCmd) Run() {
  service := c.service(c.config)
  ids, _  := c.ids(service) // TODO report error

  list := c.list()
  list = list.Select(c.filter)
  list = list.RejectIf(func(i Item) bool { return includesInt(ids, i.Id) })

  c.push(list, service)
  c.output(list, "full")
}

func (c PushCmd) service(config map[string]string) Service {
  var service Service
  switch config["service"] {
    case "idonethis":
      after := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
      service = NewIdonethis(config["team"], config["username"], config["token"], after)
  }
  return service
}

func (c PushCmd) push(list List, service Service) {
  for _, item := range list.Items {
    service.Push(item.Line) // TODO report error
  }
}

func (c PushCmd) ids(service Service) ([]int, error) {
  var ids []int
  lines, err := service.Fetch()
  if err != nil { return ids, err }

  for _, line := range lines {
    ids = append(ids, ParseItem(line).Id)
  }
  return ids, nil
}

func includesInt(nums []int, num int) bool {
  for _, n := range nums {
    if num == n { return true }
  }
  return false
}
