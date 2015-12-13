package item

import (
  "errors"
  "time"
)

const (
  Pend string = "pend"
  Done string = "done"
  None string = "none"
)

func Parse(line string) Item {
  p := parser{ line }
  return Item { line, p.id(), p.status(), p.tags(), p.text(), p.projects() }
}

type Item struct {
  Line     string
  Id       int
  Status   string
  Tags     map[string]string
  Text     string
  Projects []string
}

func (i Item) Toggle() (Item, error) {
  switch i.Status {
    case None:
      return i, errors.New("Cannot toggle an item that is not either pending or done.")
    case Done:
      i.SetPend()
      return i, nil
    default:
      i.SetDone()
      return i, nil
  }
}

func (i *Item) SetPend() {
  i.Status = Pend
  delete(i.Tags, "done")
}

func (i *Item) SetDone() {
  i.Status = Done
  i.Tags["done"] = time.Now().Format("2006-01-02")
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
