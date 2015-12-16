package todo

import (
  "testing"
  . "github.com/svenfuchs/todo/test"
)

func TestListItems(t *testing.T) {
  items    := ParseItemList([]string{ "- foo" }).Items
  actual   := items[0].Status
  expected := Pend
  AssertEqual(t, actual, expected)
}

func TestListNextId(t *testing.T) {
  list := ParseItemList([]string{ "- foo [1]", "+ bar [2]" })
  AssertEqual(t, list.NextId(), 3)
  AssertEqual(t, list.NextId(), 4)
}

func TestListSize(t *testing.T) {
  list     := ParseItemList([]string{ "- foo", "+ bar" })
  actual   := list.Size()
  expected := 2
  AssertEqual(t, actual, expected)
}

func TestListIds(t *testing.T) {
  list     := ParseItemList([]string{ "- foo [1]", "+ bar [2]" })
  actual   := list.Ids()
  expected := []int{ 1, 2 }
  AssertEqual(t, actual, expected)
}

// func TestListFindByIdFound(t *testing.T) {
//   list     := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   item, _  := list.Find(Filter{ Id: 1 })
//   actual   := item.Text
//   expected := "foo"
//   AssertEqual(t, actual, expected)  }
// }
//
// func TestListFindByIdMissing(t *testing.T) {
//   list      := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   item, err := list.Find(Filter{ Id: 3 })
//   if err == nil {
//     t.Fatalf("Expected no item to be found, but was: %s", item)
//   }
// }
//
// func TestListFindByTextFound(t *testing.T) {
//   list     := ParseItemList([]string{ "# Comment", "- fooooo [1]", "x bar [2]" })
//   item, _  := list.Find(Filter{ Text: "foo" })
//   actual   := item.Text
//   expected := "fooooo"
//   AssertEqual(t, actual, expected)  }
// }
//
// func TestListFindByTextMissing(t *testing.T) {
//   list      := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
//   item, err := list.Find(Filter{ Text: "baz" })
//   if err == nil {
//     t.Fatalf("Expected no item to be found, but was: %s", item)
//   }
// }
//
// func TestListFindByMultipleItems(t *testing.T) {
//   list      := ParseItemList([]string{ "# Comment", "- foo [1]", "x foo [2]" })
//   item, err := list.Find(Filter{ Text: "foo" })
//   if err == nil {
//     t.Fatalf("Expected no item to be found, but was: %s", item)
//   }
// }

func TestListSelectById(t *testing.T) {
  list     := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  item     := list.Select(Filter{ id: 1 }).Items[0]
  AssertEqual(t, item.Text, "foo")
}

func TestListSelectByTextFound(t *testing.T) {
  list     := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  item     := list.Select(Filter{ text: "fo" }).Items[0]
  AssertEqual(t, item.Text, "foo")
}

func TestListRejectIf(t *testing.T) {
  list  := ParseItemList([]string{ "# Comment", "- foo [1]", "x foo [2]" })
  items := list.RejectIf(func(i Item) bool { return i.IsNone() }).Items
  AssertEqual(t, len(items), 2)
}

func TestListToggleFound(t *testing.T) {
  list := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  list  = list.Toggle(Filter{ text: "bar" })
  i    := list.Items[2]
  AssertEqual(t, i.Text, "bar")
  AssertEqual(t, i.Status, Pend)
}
