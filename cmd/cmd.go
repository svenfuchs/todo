package todo

import (
  "os"
  "github.com/svenfuchs/todo/date"
  . "github.com/svenfuchs/todo"
  . "github.com/svenfuchs/todo/service"
)

type Runnable interface {
  Run()
}

type Cmd struct {
  in Source
  out Source
  filter Filter
  format string
}

func (c Cmd) list() List {
  lines := c.in.MustReadLines()
  list  := ParseItemList(lines)
  return list
}

func (c Cmd) output(list List, format string) {
  lines := c.formatted(list.Items, format)
  c.out.WriteLines(lines)
}

func (c Cmd) formatted(items []Item, format string) []string {
  return NewFormat(format).Apply(items)
}


func NewListCmd(path string, filter Filter, format string) ListCmd {
  in  := NewFileSource(path)
  out := NewFileSource("")
  return ListCmd{ Cmd { in, out, filter, format } }
}

type ListCmd struct {
  Cmd
}

func (c ListCmd) Run() {
  list := c.list()
  list  = list.Select(c.filter)
  // list  = list.SortBy(func (item Item) string { return item.DoneDate() }) TODO
  c.output(list, c.format)
}

func NewToggleCmd(path string, filter Filter) ToggleCmd {
  in  := NewFileSource(path)
  out := NewFileSource(path)
  // out := NewFileSource("")
  return ToggleCmd{ Cmd { in, out, filter, "full" } }
}

type ToggleCmd struct {
  Cmd
}

func (c ToggleCmd) Run() {
  list := c.list()
  list = list.Toggle(c.filter)
  c.output(list, "full")
}

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
  after   := date.Normalize("yesterday", date.Time)
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
