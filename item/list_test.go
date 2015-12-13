package item_test

import (
  "reflect"
  "testing"
  "github.com/svenfuchs/todo.go/item"
)

func TestListItems(t *testing.T) {
  items    := item.ParseList([]string{ "- foo" }).Items
  actual   := items[0].Status
  expected := item.Pend

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestListNextId(t *testing.T) {
  list := item.ParseList([]string{ "- foo [1]", "+ bar [2]" })

  id := list.NextId()
  if !reflect.DeepEqual(id, 3) {
    t.Fatalf("Expected MaxId to be %s, but was: %s", 3, id)
  }

  id = list.NextId()
  if !reflect.DeepEqual(id, 4) {
    t.Fatalf("Expected MaxId to be %s, but was: %s", 4, id)
  }
}

func TestListSize(t *testing.T) {
  list     := item.ParseList([]string{ "- foo", "+ bar" })
  actual   := list.Size()
  expected := 2

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Size to be %q, but was: %q", expected, actual)
  }
}

func TestListIds(t *testing.T) {
  list     := item.ParseList([]string{ "- foo [1]", "+ bar [2]" })
  actual   := list.Ids()
  expected := []int{ 1, 2 }

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected ids to be %s, but was: %s", expected, actual)
  }
}

// func TestListFindByIdFound(t *testing.T) {
//   list     := item.ParseList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   item, _  := list.Find(item.Filter{ Id: 1 })
//   actual   := item.Text
//   expected := "foo"
//
//   if !reflect.DeepEqual(actual, expected) {
//     t.Fatalf("Expected Text to be %s, but was: %s", expected, actual)
//   }
// }
//
// func TestListFindByIdMissing(t *testing.T) {
//   list      := item.ParseList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   item, err := list.Find(item.Filter{ Id: 3 })
//
//   if err == nil {
//     t.Fatalf("Expected no item to be found, but was: %s", item)
//   }
// }
//
// func TestListFindByTextFound(t *testing.T) {
//   list     := item.ParseList([]string{ "# Comment", "- fooooo [1]", "x bar [2]" })
//   item, _  := list.Find(item.Filter{ Text: "foo" })
//   actual   := item.Text
//   expected := "fooooo"
//
//   if !reflect.DeepEqual(actual, expected) {
//     t.Fatalf("Expected Text to be %s, but was: %s", expected, actual)
//   }
// }
//
// func TestListFindByTextMissing(t *testing.T) {
//   list      := item.ParseList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   item, err := list.Find(item.Filter{ Text: "baz" })
//
//   if err == nil {
//     t.Fatalf("Expected no item to be found, but was: %s", item)
//   }
// }
//
// func TestListFindByMultipleItems(t *testing.T) {
//   list      := item.ParseList([]string{ "# Comment", "- foo [1]", "x foo [2]" })
//   item, err := list.Find(item.Filter{ Text: "foo" })
//
//   if err == nil {
//     t.Fatalf("Expected no item to be found, but was: %s", item)
//   }
// }

func TestListSelectById(t *testing.T) {
  list     := item.ParseList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  item     := list.Select(item.Filter{ Id: 1 }).Items[0]
  actual   := item.Text
  expected := "foo"

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestListSelectByTextFound(t *testing.T) {
  list     := item.ParseList([]string{ "# Comment", "- fooooo [1]", "x bar [2]" })
  item     := list.Select(item.Filter{ Text: "foo" }).Items[0]
  actual   := item.Text
  expected := "fooooo"

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestListRejectIf(t *testing.T) {
  list  := item.ParseList([]string{ "# Comment", "- foo [1]", "x foo [2]" })
  items := list.RejectIf(func(i item.Item) bool { return i.IsNone() }).Items

  if len(items) != 2 {
    t.Fatalf("Expected 2 items to be found, but found %s: %s", len(items), items)
  }
}

func TestListToggleFound(t *testing.T) {
  list := item.ParseList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  list  = list.Toggle(item.Filter{ Text: "bar" })
  i    := list.Items[2]

  if !reflect.DeepEqual(i.Text, "bar") {
    t.Fatalf("Expected Text to be %s, but was: %s", "bar", i.Text)
  }

  if !reflect.DeepEqual(i.Status, item.Pend) {
    t.Fatalf("Expected Status to be %s, but was: %s", item.Pend, i.Status)
  }
}

// func TestListToggleNotFound(t *testing.T) {
//   list   := item.ParseList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   i, err := list.Toggle(item.Filter{ Text: "baz" })
//
//   if err == nil {
//     t.Fatalf("Expected no item to be found, but was: %s", i)
//   }
// }
//
// func TestListToggleInvalidStatus(t *testing.T) {
//   list   := item.ParseList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   i, err := list.Toggle(item.Filter{ Text: "Comment" })
//
//   if err == nil {
//     t.Fatalf("Expected no item to be found, but was: %s", i)
//   }
// }
