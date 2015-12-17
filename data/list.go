package data

// import "fmt"

func ParseItemList(lines []string) List {
  items := []Item{}
  for _, line := range lines {
    items = append(items, ParseItem(line))
  }
  return List{ Items: items, nextId: maxId(items) }
}

type List struct {
  Items []Item
  nextId int
}

func (l *List) NextId() int {
  l.nextId ++
  return l.nextId
}

func (l *List) Ids() []int {
  ids := []int{}
  for _, item := range l.Items {
    ids = append(ids, item.Id)
  }
  return ids
}

func (l *List) Select(filter Filter) List {
  items := []Item{}
  for _, item := range l.Items {
    if filter.Apply(item) {
      items = append(items, item)
    }
  }
  return List{ Items: items, nextId: l.nextId }
}

func (l *List) Reject(filter Filter) List {
  items := []Item{}
  for _, item := range l.Items {
    if filter.IsEmpty() || !filter.Apply(item) {
      items = append(items, item)
    }
  }
  return List{ Items: items, nextId: l.nextId }
}

func (l *List) Toggle(filter Filter) List {
  items := []Item{}
  for i, _ := range l.Items {
    item := &l.Items[i]
    if filter.Apply(*item) {
      items = append(items, *item.Toggle())
    }
  }
  return List{ Items: items, nextId: l.nextId }
}

func maxId(items []Item) int {
  id := 0
  for _, item := range items {
    if id < item.Id {
      id = item.Id
    }
  }
  return id
}
