package date_test

import (
	"github.com/svenfuchs/todo/date"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	date.Time = date.NewStub(today)
}

func TestDateParse(t *testing.T) {
	actual := date.Normalize("2015-12-13")
	expected := today

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}
