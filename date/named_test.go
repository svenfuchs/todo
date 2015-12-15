package date_test

import (
  "reflect"
  "testing"
  "github.com/svenfuchs/todo/date"
)

var (
  yesterday        = "2015-12-12"
  today            = "2015-12-13"
  tomorrow         = "2015-12-14"
)

func TestDateNamedToday(t *testing.T) {
  actual   := date.Normalize("today", date.Stub(today))
  expected := today

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %q, but was: %q", expected, actual)
  }
}

func TestDateNamedYesterday(t *testing.T) {
  actual   := date.Normalize("yesterday", date.Stub(today))
  expected := yesterday

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %q, but was: %q", expected, actual)
  }
}

func TestDateNamedTomorrow(t *testing.T) {
  actual   := date.Normalize("tomorrow", date.Stub(today))
  expected := tomorrow

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %q, but was: %q", expected, actual)
  }
}
