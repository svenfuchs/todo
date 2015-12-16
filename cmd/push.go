package cmd

import (
  "time"
  "os"
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/service"
)


func NewPushCmd(path string, filter Filter) PushCmd {
  in  := NewFileSource(path)
  out := NewFileSource("")
  return PushCmd{ Cmd { in, out, filter, "full" } }
}

type PushCmd struct {
  Cmd
}

func (c PushCmd) Run() {
  team    := "personal-2jAM" // os.Getenv("IDONETHIS_TEAM") // TODO
  user    := os.Getenv("IDONETHIS_USERNAME")
  token   := os.Getenv("IDONETHIS_TOKEN")
  after   := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
  service := NewIdonethis(team, user, token, after)
  ids, _  := c.ids(service) // TODO report error

  list := c.list()
  list = list.Select(c.filter)
  list = list.RejectIf(func(i Item) bool { return includesInt(ids, i.Id) })

  c.Push(list, service)
  c.output(list, "full")
}

func (c PushCmd) Push(list List, service Idonethis) {
  for _, item := range list.Items {
    service.Push(item.Line) // TODO report error
  }
}

func (c PushCmd) ids(service Idonethis) ([]int, error) {
  var ids []int
  lines, err := service.Lines()
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
