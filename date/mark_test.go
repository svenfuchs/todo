package date_test

import (
	"github.com/svenfuchs/todo/date"
	"reflect"
	"testing"
)

var (
	last_sun = "2015-12-06"
	last_mon = "2015-12-07"
	last_sat = "2015-12-12"
	next_mon = "2015-12-14"
	next_sat = "2015-12-19"
	next_sun = "2015-12-20"

	last_nov = "2015-11-01"
	next_jan = "2016-01-01"

	last_year = "2014-01-01"
	next_year = "2016-01-01"
)

func TestDateMarkLastYear(t *testing.T) {
	actual := date.Normalize("last year")
	expected := last_year

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkNextYear(t *testing.T) {
	actual := date.Normalize("next year")
	expected := next_year

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkLastNovember(t *testing.T) {
	actual := date.Normalize("last november")
	expected := last_nov

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkNextJanuar(t *testing.T) {
	actual := date.Normalize("next januar")
	expected := next_jan

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkLastSunday(t *testing.T) {
	actual := date.Normalize("last sunday")
	expected := last_sun

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkLastMonday(t *testing.T) {
	actual := date.Normalize("last monday")
	expected := last_mon

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkLastSaturday(t *testing.T) {
	actual := date.Normalize("last saturday")
	expected := last_sat

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkNextMonday(t *testing.T) {
	actual := date.Normalize("next monday")
	expected := next_mon

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkNextSunday(t *testing.T) {
	actual := date.Normalize("next sunday")
	expected := next_sun

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateMarkNextSaturday(t *testing.T) {
	actual := date.Normalize("next saturday")
	expected := next_sat

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}
