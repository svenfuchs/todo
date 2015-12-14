package test

import (
  "reflect"
  "testing"
)

func AssertEqual(t *testing.T, actual interface{}, expected interface{}) {
  if !reflect.DeepEqual(actual, expected) {
    t.Fatalf("Expected %s, but was: %s", expected, actual)
  }
}

func AssertTrue(t *testing.T, actual interface{}) {
  AssertEqual(t, actual, true)
}

func AssertFalse(t *testing.T, actual interface{}) {
  AssertEqual(t, actual, false)
}

func AssertNil(t *testing.T, actual interface{}) {
  AssertEqual(t, actual, nil)
}

func AssertNotNil(t *testing.T, actual interface{}) {
  if actual == nil {
    t.Fatalf("Expected %s to be not nil, but it was", actual)
  }
}

