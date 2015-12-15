package date_test

import (
  "reflect"
  "testing"
  "github.com/svenfuchs/todo/date"
)

func TestDateParse(t *testing.T) {
  actual   := date.Normalize("2015-12-13", date.Stub(today))
  expected := today

  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %q, but was: %q", expected, actual)
  }
}

