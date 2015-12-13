package item_test

import (
  "reflect"
  "testing"
  "time"
  "github.com/svenfuchs/todo.go/item"
)

func TestItemPendingToggleStatus(t *testing.T) {
  i := item.Parse("- foo")
  i, _ = i.Toggle()

  actual   := i.Status
  expected := item.Done

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestItemPendingToggleText(t *testing.T) {
  i := item.Parse("- foo")
  i, _ = i.Toggle()

  actual   := i.Text
  expected := "foo"

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestItemPendingToggleTagsDone(t *testing.T) {
  i := item.Parse("- foo")
  i, _ = i.Toggle()

  actual   := i.Tags["done"]
  expected := time.Now().Format("2006-01-02")

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestItemDoneToggleStatus(t *testing.T) {
  i := item.Parse("x foo")
  i, _ = i.Toggle()

  actual   := i.Status
  expected := item.Pend

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestItemDoneToggleText(t *testing.T) {
  i := item.Parse("x foo")
  i, _ = i.Toggle()

  actual   := i.Text
  expected := "foo"

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestItemDoneToggleTagsDone(t *testing.T) {
  i := item.Parse("x foo")
  i, _ = i.Toggle()

  _, ok := i.Tags["done"]

  if ok {
    t.Fatalf("Expected Tags[\"done\"] to be not present, but was: %s", i.Tags)
  }
}

func TestItemNoneToggleError(t *testing.T) {
  i      := item.Parse("# Comment")
  i, err := i.Toggle()

  if err == nil {
    t.Fatalf("Expected an error to be returned but was: %q", err)
  }
}

func TestItemPendingIsDone(t *testing.T) {
  actual   := item.Parse("- foo").IsDone()
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestItemDoneIsDone(t *testing.T) {
  actual   := item.Parse("x foo").IsDone()
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestCommentItemIsDone(t *testing.T) {
  actual   := item.Parse("# Comment").IsDone()
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestItemPendingIsPend(t *testing.T) {
  actual   := item.Parse("- foo").IsPend()
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestItemDoneIsPend(t *testing.T) {
  actual   := item.Parse("x foo").IsPend()
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestCommentItemIsPend(t *testing.T) {
  actual   := item.Parse("# Comment").IsPend()
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestItemPendingIsNone(t *testing.T) {
  actual   := item.Parse("- foo").IsNone()
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestItemDoneIsNone(t *testing.T) {
  actual   := item.Parse("x foo").IsNone()
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestCommentItemIsNone(t *testing.T) {
  actual   := item.Parse("# Comment").IsNone()
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected id to be %q, but was: %q", expected, actual)
  }
}

func TestItemIdValid(t *testing.T) {
  actual   := item.Parse("- foo [1]").Id
  expected := 1

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Id to be %s, but was: %s", expected, actual)
  }
}

func TestItemIdMissing(t *testing.T) {
  actual   := item.Parse("- foo").Id
  expected := 0

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Id to be %s, but was: %s", expected, actual)
  }
}

func TestItemStatusPending(t *testing.T) {
  actual   := item.Parse("- foo").Status
  expected := item.Pend

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestItemStatusDone(t *testing.T) {
  actual   := item.Parse("x foo").Status
  expected := item.Done

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestItemStatusNone(t *testing.T) {
  actual   := item.Parse("# Comment").Status
  expected := item.None

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected Status to be %q, but was: %q", expected, actual)
  }
}

func TestItemDoneDatePresent(t *testing.T) {
  actual   := item.Parse("- foo done:2015-12-13").DoneDate()
  expected := "2015-12-13"

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected DoneDate to be %q, but was: %q", expected, actual)
  }
}

func TestItemDoneDateMissing(t *testing.T) {
  actual   := item.Parse("- foo").DoneDate()
  expected := ""

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected DoneDate to be %q, but was: %q", expected, actual)
  }
}

func TestItemDueDatePresent(t *testing.T) {
  actual   := item.Parse("- foo due:2015-12-13").DueDate()
  expected := "2015-12-13"

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected DueDate to be %q, but was: %q", expected, actual)
  }
}

func TestItemDueDateMissing(t *testing.T) {
  actual   := item.Parse("- foo").DueDate()
  expected := ""

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected DueDate to be %q, but was: %q", expected, actual)
  }
}

func TestItemTags(t *testing.T) {
  actual   := item.Parse("- foo bar:baz bam:bum").Tags
  expected := map[string]string{ "bar": "baz", "bam": "bum" }

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected tags to be %s, but was: %s", expected, actual)
  }
}

func TestItemTagsEmpty(t *testing.T) {
  actual   := item.Parse("- foo").Tags
  expected := map[string]string{}

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected tags to be %s, but was: %s", expected, actual)
  }
}

func TestItemTagsDate(t *testing.T) {
  actual   := item.Parse("- foo done:2015-12-13").Tags
  expected := map[string]string{ "done": "2015-12-13" }

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected tags to be %s, but was: %s", expected, actual)
  }
}

func TestItemTagsNamedDate(t *testing.T) {
  actual   := item.Parse("- foo done:today").Tags
  expected := map[string]string{ "done": time.Now().Format("2006-01-02") }

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected tags to be %s, but was: %s", expected, actual)
  }
}

func TestItemText(t *testing.T) {
  actual   := item.Parse("- foo bar baz").Text
  expected := "foo bar baz"

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected text to be %s, but was: %s", expected, actual)
  }
}

func TestItemProjects(t *testing.T) {
  actual   := item.Parse("- foo +bar +baz").Projects
  expected := []string{ "bar", "baz" }

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected text to be %s, but was: %s", expected, actual)
  }
}
