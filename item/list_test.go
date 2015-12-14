package item_test

import (
  "testing"
  "github.com/svenfuchs/todo.go/item"
  . "github.com/svenfuchs/todo.go/test"
)

func TestListItems(t *testing.T) {
  items    := item.ParseList([]string{ "- foo" }).Items
  actual   := items[0].Status
  expected := item.Pend

  AssertEqual(t, actual, expected)
}

func TestListNextId(t *testing.T) {
  list := item.ParseList([]string{ "- foo [1]", "+ bar [2]" })

  AssertEqual(t, list.NextId(), 3)
  AssertEqual(t, list.NextId(), 4)
}

func TestListSize(t *testing.T) {
  list     := item.ParseList([]string{ "- foo", "+ bar" })
  actual   := list.Size()
  expected := 2

  AssertEqual(t, actual, expected)
}

func TestListIds(t *testing.T) {
  list     := item.ParseList([]string{ "- foo [1]", "+ bar [2]" })
  actual   := list.Ids()
  expected := []int{ 1, 2 }

  AssertEqual(t, actual, expected)
}

// func TestListFindByIdFound(t *testing.T) {
//   list     := item.ParseList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   item, _  := list.Find(item.Filter{ Id: 1 })
//   actual   := item.Text
//   expected := "foo"
//
//   AssertEqual(t, actual, expected)  }
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
//   AssertEqual(t, actual, expected)  }
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

  AssertEqual(t, actual, expected)
}

func TestListSelectByTextFound(t *testing.T) {
  list     := item.ParseList([]string{ "# Comment", "- fooooo [1]", "x bar [2]" })
  item     := list.Select(item.Filter{ Text: "foo" }).Items[0]
  actual   := item.Text
  expected := "fooooo"

  AssertEqual(t, actual, expected)
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

  AssertEqual(t, i.Text, "bar")
  AssertEqual(t, i.Status, item.Pend)
}
