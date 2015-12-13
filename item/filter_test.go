package item_test

import (
  "reflect"
  "testing"
  "github.com/svenfuchs/todo.go/item"
)

var (
  lines = []string{ "# Comment", "- foo +bar due:2015-12-13 [1]", "x baz done:2015-12-13 [2]" }
  items = item.ParseList(lines).Items
)

func TestFilterEmpty(t *testing.T) {
  filter   := item.Filter{}
  actual   := filter.Apply(items[1])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByIdSuccess(t *testing.T) {
  filter   := item.Filter{ Id: 1 }
  actual   := filter.Apply(items[1])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByIdComment(t *testing.T) {
  filter   := item.Filter{ Id: 1 }
  actual   := filter.Apply(items[0])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByIdFailure(t *testing.T) {
  filter   := item.Filter{ Id: 2 }
  actual   := filter.Apply(items[1])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByIdSuccessTextIgnored(t *testing.T) {
  filter   := item.Filter{ Id: 1, Text: "ignored" }
  actual   := filter.Apply(items[1])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByTextSuccess(t *testing.T) {
  filter   := item.Filter{ Text: "fo" }
  actual   := filter.Apply(items[1])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByTextFailure(t *testing.T) {
  filter   := item.Filter{ Text: "unknown" }
  actual   := filter.Apply(items[1])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByTextComment(t *testing.T) {
  filter   := item.Filter{ Text: "Comment" }
  actual   := filter.Apply(items[0])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByStatusPendSuccess(t *testing.T) {
  filter   := item.Filter{ Status: item.Pend }
  actual   := filter.Apply(items[1])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByStatusPendFailure(t *testing.T) {
  filter   := item.Filter{ Status: item.Pend }
  actual   := filter.Apply(items[2])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByStatusDone(t *testing.T) {
  filter   := item.Filter{ Status: item.Pend }
  actual   := filter.Apply(items[1])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByStatusDoneFailure(t *testing.T) {
  filter   := item.Filter{ Status: item.Done }
  actual   := filter.Apply(items[2])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByAfterSuccess(t *testing.T) {
  filter   := item.Filter{ After: "2015-12-01" }
  actual   := filter.Apply(items[2])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByAfterFailure(t *testing.T) {
  filter   := item.Filter{ After: "2016-01-01" }
  actual   := filter.Apply(items[2])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByAfterFailureDoneMissing(t *testing.T) {
  filter   := item.Filter{ After: "2015-12-01" }
  actual   := filter.Apply(items[1])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByBeforeSuccess(t *testing.T) {
  filter   := item.Filter{ Before: "2016-01-01" }
  actual   := filter.Apply(items[2])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByBeforeFailure(t *testing.T) {
  filter   := item.Filter{ Before: "2015-12-01" }
  actual   := filter.Apply(items[2])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByBeforeMissing(t *testing.T) {
  filter   := item.Filter{ Before: "2016-01-01" }
  actual   := filter.Apply(items[1])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByProjectSuccess(t *testing.T) {
  filter   := item.Filter{ Projects: "bar,bam" }
  actual   := filter.Apply(items[1])
  expected := true

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func TestFilterByProjectFailure(t *testing.T) {
  filter   := item.Filter{ Projects: "missing,unknown" }
  actual   := filter.Apply(items[1])
  expected := false

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}
