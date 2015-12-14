package item_test

import (
  "testing"
  "time"
  "github.com/svenfuchs/todo.go/item"
  . "github.com/svenfuchs/todo.go/test"
)

func TestItemText(t *testing.T) {
  actual   := item.Parse("- foo bar baz").Text
  expected := "foo bar baz"
  AssertEqual(t, actual, expected)
}

func TestItemPendingIsDone(t *testing.T) {
  actual   := item.Parse("- foo").IsDone()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestItemDoneIsDone(t *testing.T) {
  actual   := item.Parse("x foo").IsDone()
  expected := true
  AssertEqual(t, actual, expected)
}

func TestCommentItemIsDone(t *testing.T) {
  actual   := item.Parse("# Comment").IsDone()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestItemPendingIsPend(t *testing.T) {
  actual   := item.Parse("- foo").IsPend()
  expected := true
  AssertEqual(t, actual, expected)
}

func TestItemDoneIsPend(t *testing.T) {
  actual   := item.Parse("x foo").IsPend()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestCommentItemIsPend(t *testing.T) {
  actual   := item.Parse("# Comment").IsPend()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestItemPendingIsNone(t *testing.T) {
  actual   := item.Parse("- foo").IsNone()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestItemDoneIsNone(t *testing.T) {
  actual   := item.Parse("x foo").IsNone()
  expected := false
  AssertEqual(t, actual, expected)
}

func TestCommentItemIsNone(t *testing.T) {
  actual   := item.Parse("# Comment").IsNone()
  expected := true
  AssertEqual(t, actual, expected)
}

func TestItemIdValid(t *testing.T) {
  actual   := item.Parse("- foo [1]").Id
  expected := 1
  AssertEqual(t, actual, expected)
}

func TestItemIdMissing(t *testing.T) {
  actual   := item.Parse("- foo").Id
  expected := 0
  AssertEqual(t, actual, expected)
}

func TestItemStatusPending(t *testing.T) {
  actual   := item.Parse("- foo").Status
  expected := item.Pend
  AssertEqual(t, actual, expected)
}

func TestItemStatusDone(t *testing.T) {
  actual   := item.Parse("x foo").Status
  expected := item.Done
  AssertEqual(t, actual, expected)
}

func TestItemStatusNone(t *testing.T) {
  actual   := item.Parse("# Comment").Status
  expected := item.None
  AssertEqual(t, actual, expected)
}

func TestItemProjects(t *testing.T) {
  actual   := item.Parse("- foo +bar +baz").Projects
  expected := []string{ "bar", "baz" }
  AssertEqual(t, actual, expected)
}

func TestItemDoneDatePresent(t *testing.T) {
  actual   := item.Parse("- foo done:2015-12-13").DoneDate()
  expected := "2015-12-13"
  AssertEqual(t, actual, expected)
}

func TestItemDoneDateMissing(t *testing.T) {
  actual   := item.Parse("- foo").DoneDate()
  expected := ""
  AssertEqual(t, actual, expected)
}

func TestItemDueDatePresent(t *testing.T) {
  actual   := item.Parse("- foo due:2015-12-13").DueDate()
  expected := "2015-12-13"
  AssertEqual(t, actual, expected)
}

func TestItemDueDateMissing(t *testing.T) {
  actual   := item.Parse("- foo").DueDate()
  expected := ""
  AssertEqual(t, actual, expected)
}

func TestItemTags(t *testing.T) {
  actual   := item.Parse("- foo bar:baz bam:bum").Tags
  expected := map[string]string{ "bar": "baz", "bam": "bum" }
  AssertEqual(t, actual, expected)
}

func TestItemTagsEmpty(t *testing.T) {
  actual   := item.Parse("- foo").Tags
  expected := map[string]string{}
  AssertEqual(t, actual, expected)
}

func TestItemTagsDate(t *testing.T) {
  actual   := item.Parse("- foo done:2015-12-13").Tags
  expected := map[string]string{ "done": "2015-12-13" }
  AssertEqual(t, actual, expected)
}

func TestItemTagsNamedDate(t *testing.T) {
  actual   := item.Parse("- foo done:today").Tags
  expected := map[string]string{ "done": time.Now().Format("2006-01-02") }
  AssertEqual(t, actual, expected)
}

func TestItemTagsDoNotMatchUrls(t *testing.T) {
  actual   := item.Parse("- foo http://host.com/path foo:bar [1]").Tags
  expected := map[string]string{ "foo": "bar" }
  AssertEqual(t, actual, expected)
}

func TestItemPendingToggleStatus(t *testing.T) {
  i := item.Parse("- foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.Status, item.Done)
}

func TestItemPendingToggleText(t *testing.T) {
  i := item.Parse("- foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.Text, "foo")
}

func TestItemPendingToggleTagsDone(t *testing.T) {
  i := item.Parse("- foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.Tags["done"], time.Now().Format("2006-01-02"))
}

func TestItemDoneToggleStatus(t *testing.T) {
  i := item.Parse("x foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.Status, item.Pend)
}

func TestItemDoneToggleText(t *testing.T) {
  i := item.Parse("x foo")
  i, _ = i.Toggle()
  AssertEqual(t, i.Text, "foo")
}

func TestItemDoneToggleTagsDone(t *testing.T) {
  i := item.Parse("x foo")
  i, _ = i.Toggle()
  _, ok := i.Tags["done"]
  AssertFalse(t, ok)
}

func TestItemNoneToggleError(t *testing.T) {
  i      := item.Parse("# Comment")
  i, err := i.Toggle()
  AssertNotNil(t, err)
}

