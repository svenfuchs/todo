package date_test

import (
	"github.com/svenfuchs/todo/date"
	"reflect"
	"testing"
)

var (
	yesterday = "2015-12-12"
	today     = "2015-12-13"
	tomorrow  = "2015-12-14"
)

func TestDateNamedToday(t *testing.T) {
	actual := date.Normalize("today")
	expected := today

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateNamedYesterday(t *testing.T) {
	actual := date.Normalize("yesterday")
	expected := yesterday

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateNamedTomorrow(t *testing.T) {
	actual := date.Normalize("tomorrow")
	expected := tomorrow

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}
