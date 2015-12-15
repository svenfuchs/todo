package todo

import (
  "errors"
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
  line     string
  id       int
  status   string
  tags     map[string]string
  text     string
  projects []string
}

func (i Item) Toggle() (Item, error) {
  switch i.status {
    case Pend:
      i.status = Done
      i.tags["done"] = time.Now().Format("2006-01-02")
      return i, nil
    case Done:
      i.status = Pend
      delete(i.tags, "done")
      return i, nil
    default:
      return i, errors.New("Cannot toggle an item that is not either pending or done.")
  }
}

func (i Item) IsDone() bool {
  return i.status == Done
}

func (i Item) IsPend() bool {
  return i.status == Pend
}

func (i Item) IsNone() bool {
  return i.status == None
}

func (i Item) DueDate() string {
  return i.tags["due"]
}

func (i Item) DoneDate() string {
  return i.tags["done"]
}
