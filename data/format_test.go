package data

import (
  "testing"
  . "github.com/svenfuchs/todo/test"
)

func TestFormatFull(t *testing.T) {
  items    := ParseItemList([]string{ "# Comment", "- foo", "x bar due:2015-12-01 done:2015-12-13 [1]" }).Items
  actual   := NewFormat("full").Apply(items)
  expected := []string{ "# Comment", "- foo [0]", "x bar done:2015-12-13 due:2015-12-01 [1]" }

  AssertEqual(t, actual, expected)
}

func TestFormatShort(t *testing.T) {
  items    := ParseItemList([]string{ "# Comment", "- foo", "x bar due:2015-12-01 done:2015-12-13 [1]" }).Items
  actual   := NewFormat("short").Apply(items)
  expected := []string{ "# Comment", "- foo", "x 2015-12-13 bar" }

  AssertEqual(t, actual, expected)
}

func TestFormatId(t *testing.T) {
  items    := ParseItemList([]string{ "- foo", "x bar due:2015-12-01 done:2015-12-13 [1]" }).Items
  actual   := NewFormat("id").Apply(items)
  expected := []string{ "[0]", "[1]" }

  AssertEqual(t, actual, expected)
}

func TestFormatStatus(t *testing.T) {
  items    := ParseItemList([]string{ "- foo", "x bar due:2015-12-01 done:2015-12-13 [1]" }).Items
  actual   := NewFormat("status").Apply(items)
  expected := []string{ "-", "x" }

  AssertEqual(t, actual, expected)
}

func TestFormatText(t *testing.T) {
  items    := ParseItemList([]string{ "- foo", "x bar due:2015-12-01 done:2015-12-13 [1]" }).Items
  actual   := NewFormat("text").Apply(items)
  expected := []string{ "foo", "bar" }

  AssertEqual(t, actual, expected)
}

func TestFormatTags(t *testing.T) {
  items    := ParseItemList([]string{ "- foo key:value", "x bar due:2015-12-01 done:2015-12-13 [1]" }).Items
  actual   := NewFormat("tags").Apply(items)
  expected := []string{ "key:value", "done:2015-12-13 due:2015-12-01" }

  AssertEqual(t, actual, expected)
}

func TestFormatDone(t *testing.T) {
  items    := ParseItemList([]string{ "- foo", "x bar due:2015-12-01 done:2015-12-13 [1]" }).Items
  actual   := NewFormat("done").Apply(items)
  expected := []string{ "", "2015-12-13" }

  AssertEqual(t, actual, expected)
}

func TestFormatDue(t *testing.T) {
  items    := ParseItemList([]string{ "- foo", "x bar due:2015-12-01 done:2015-12-13 [1]" }).Items
  actual   := NewFormat("due").Apply(items)
  expected := []string{ "", "2015-12-01" }

  AssertEqual(t, actual, expected)
}
