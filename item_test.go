package todo

import (
  "testing"
  "time"
  . "github.com/svenfuchs/todo/test"
)

func TestItemPendingIsDone(t *testing.T) {
  actual   := ParseItem("- foo").IsDone()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestItemDoneIsDone(t *testing.T) {
  actual   := ParseItem("x foo").IsDone()
  expected := true
  AssertEqual(t, actual, expected)
}

func TestCommentItemIsDone(t *testing.T) {
  actual   := ParseItem("# Comment").IsDone()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestItemPendingIsPend(t *testing.T) {
  actual   := ParseItem("- foo").IsPend()
  expected := true
  AssertEqual(t, actual, expected)
}

func TestItemDoneIsPend(t *testing.T) {
  actual   := ParseItem("x foo").IsPend()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestCommentItemIsPend(t *testing.T) {
  actual   := ParseItem("# Comment").IsPend()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestItemPendingIsNone(t *testing.T) {
  actual   := ParseItem("- foo").IsNone()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestItemDoneIsNone(t *testing.T) {
  actual   := ParseItem("x foo").IsNone()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestCommentItemIsNone(t *testing.T) {
  actual   := ParseItem("# Comment").IsNone()
  expected := true
  AssertEqual(t, actual, expected)
}

func TestItemDoneDatePresent(t *testing.T) {
  actual   := ParseItem("- foo done:2015-12-13").DoneDate()
  expected := "2015-12-13"
  AssertEqual(t, actual, expected)
}

func TestItemDoneDateMissing(t *testing.T) {
  actual   := ParseItem("- foo").DoneDate()
  expected := ""
  AssertEqual(t, actual, expected)
}

func TestItemDueDatePresent(t *testing.T) {
  actual   := ParseItem("- foo due:2015-12-13").DueDate()
  expected := "2015-12-13"
  AssertEqual(t, actual, expected)
}

func TestItemDueDateMissing(t *testing.T) {
  actual   := ParseItem("- foo").DueDate()
  expected := ""
  AssertEqual(t, actual, expected)
}

func TestItemPendingToggleStatus(t *testing.T) {
  i := ParseItem("- foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.status, Done)
}

func TestItemPendingToggleText(t *testing.T) {
  i := ParseItem("- foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.text, "foo")
}

func TestItemPendingToggleTagsDone(t *testing.T) {
  i := ParseItem("- foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.tags["done"], time.Now().Format("2006-01-02"))
}

func TestItemDoneToggleStatus(t *testing.T) {
  i := ParseItem("x foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.status, Pend)
}

func TestItemDoneToggleText(t *testing.T) {
  i := ParseItem("x foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.text, "foo")
}

func TestItemDoneToggleTagsDone(t *testing.T) {
  i := ParseItem("x foo")
  i, _ = i.Toggle()
  _, ok := i.tags["done"]
  AssertFalse(t, ok)
}

func TestItemNoneToggleError(t *testing.T) {
  i      := ParseItem("# Comment")
  i, err := i.Toggle()
  AssertNotNil(t, err)
}
