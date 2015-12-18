package data

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
  actual   := len(list.Items)
  expected := 2
  AssertEqual(t, actual, expected)
}

func TestListIds(t *testing.T) {
  list     := ParseItemList([]string{ "- foo [1]", "+ bar [2]" })
  actual   := list.Ids()
  expected := []int{ 1, 2 }
  AssertEqual(t, actual, expected)
}

func TestListSelectById(t *testing.T) {
  list     := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  item     := list.Select(Filter{ Ids: []int{ 1 } }).Items[0]
  AssertEqual(t, item.Text, "foo")
}

func TestListSelectByTextFound(t *testing.T) {
  list     := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  item     := list.Select(Filter{ Text: "fo" }).Items[0]
  AssertEqual(t, item.Text, "foo")
}

func TestListReject(t *testing.T) {
  list  := ParseItemList([]string{ "# Comment", "- foo [1]", "x foo [2]" })
  items := list.Reject(Filter{ Ids: []int{ 1 } }).Items
  AssertEqual(t, len(items), 2)
}

func TestListToggleChangesExistingList(t *testing.T) {
  list := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  list.Toggle(Filter{ Ids: []int{ 1 } })
  AssertEqual(t, list.Items[1].Status, Done)
}

func TestListToggleReturnsResultsList(t *testing.T) {
  list := ParseItemList([]string{ "# Comment", "- foo [1]", "x bar [2]" })
  res := list.Toggle(Filter{ Ids: []int{ 1 } })
  AssertEqual(t, len(res.Items), 1)
  AssertEqual(t, res.Items[0].Status, Done)
}
