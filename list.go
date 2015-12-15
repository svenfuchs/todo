package todo

// import "fmt"

func ParseItemList(lines []string) List {
  items := []Item{}
  for _, line := range lines {
    items = append(items, ParseItem(line))
  }
  return List{ Items: items, nextId: maxId(items) }
}

func maxId(items []Item) int {
  id := 0
  for _, item := range items {
    if id < item.id {
      id = item.id
    }
  }
  return id
}

type List struct {
  Items []Item
  nextId int
}

func (l *List) NextId() int {
  l.nextId ++
  return l.nextId
}

func (l *List) Size() int {
  return len(l.Items)
}

func (l *List) Ids() []int {
  ids := []int{}
  for _, item := range l.Items {
    ids = append(ids, item.id)
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

func (l *List) RejectIf(match func(Item) bool) List {
  items := []Item{}
  for _, item := range l.Items {
    if !match(item) {
      items = append(items, item)
    }
  }
  return List{ Items: items, nextId: l.nextId }
}

func (l *List) Toggle(filter Filter) List {
  items := []Item{}
  for _, item := range l.Items {
    if filter.Apply(item) {
      item, _ = item.Toggle()
    }
    items = append(items, item)
  }
  return List{ Items: items, nextId: l.nextId }
}

// func (l *List) Find(filter Filter) (Item, error) {
//   items := l.Select(filter).Items
//   switch len(items) {
//     case 0:
//       return Item{}, errors.New(fmt.Sprintf("Could not find item matching %s", filter))
//     case 1:
//       return items[0], nil
//     default:
//       return Item{}, errors.New(fmt.Sprintf("Multiple items found matching %s", filter))
//   }
// }
