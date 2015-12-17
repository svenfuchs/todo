package todo

import (
  "time"
)

const (
  Pend string = "pend"
  Done string = "done"
  None string = "none"
)

func ParseItem(line string) Item {
  p := Parser { line }
  return Item { line, p.Id(), p.Status(), p.Tags(), p.Text(), p.Projects() }
}

type Item struct {
  Line     string
  Id       int
  Status   string
  Tags     map[string]string
  Text     string
  Projects []string
}

func (i *Item) Toggle() *Item {
  if i.Status == Pend {
    i.setDone()
  } else {
    i.setPend()
  }
  return i
}

func (i *Item) setPend() {
  i.Status = Pend
  delete(i.Tags, "done")
  i.setLine()
}

func (i *Item) setDone() {
  i.Status = Done
  i.Tags["done"] = time.Now().Format("2006-01-02")
  i.setLine()
}

func (i Item) IsDone() bool {
  return i.Status == Done
}

func (i Item) IsPend() bool {
  return i.Status == Pend
}

func (i Item) IsNone() bool {
  return i.Status == None
}

func (i Item) DueDate() string {
  return i.Tags["due"]
}

func (i Item) DoneDate() string {
  return i.Tags["done"]
}

func (i *Item) setLine() {
  i.Line = NewFormat("full").Apply([]Item { *i })[0] // TODO
}
