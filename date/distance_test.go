package date_test

import (
	"github.com/svenfuchs/todo/date"
	"reflect"
	"testing"
)

var (
	two_years_ago    = "2013-12-13"
	one_year_ago     = "2014-12-13"
	two_months_ago   = "2015-10-13"
	one_month_ago    = "2015-11-13"
	two_weeks_ago    = "2015-11-29"
	one_week_ago     = "2015-12-06"
	two_days_ago     = "2015-12-11"
	one_day_ago      = "2015-12-12"
	one_day_ahead    = "2015-12-14"
	two_days_ahead   = "2015-12-15"
	one_week_ahead   = "2015-12-20"
	two_weeks_ahead  = "2015-12-27"
	one_month_ahead  = "2016-01-13"
	two_months_ahead = "2016-02-13"
	one_year_ahead   = "2016-12-13"
	two_years_ahead  = "2017-12-13"
)

func TestDateDistanceTwoYearsAgo(t *testing.T) {
	actual := date.Normalize("two years ago")
	expected := two_years_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceOneYearAgo(t *testing.T) {
	actual := date.Normalize("one year ago")
	expected := one_year_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceTwoMonthsAgo(t *testing.T) {
	actual := date.Normalize("two months ago")
	expected := two_months_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceOneMonthAgo(t *testing.T) {
	actual := date.Normalize("one month ago")
	expected := one_month_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceTwoWeeksAgo(t *testing.T) {
	actual := date.Normalize("two weeks ago")
	expected := two_weeks_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceOneWeekAgo(t *testing.T) {
	actual := date.Normalize("one week ago")
	expected := one_week_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceTwoDaysAgo(t *testing.T) {
	actual := date.Normalize("two days ago")
	expected := two_days_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceOneDayAgo(t *testing.T) {
	actual := date.Normalize("one day ago")
	expected := one_day_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceOneDayAhead(t *testing.T) {
	actual := date.Normalize("one day ahead")
	expected := one_day_ahead

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceTwoDaysAhead(t *testing.T) {
	actual := date.Normalize("two days ahead")
	expected := two_days_ahead

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceOneWeekAhead(t *testing.T) {
	actual := date.Normalize("one week ahead")
	expected := one_week_ahead

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceTwoWeeksAhead(t *testing.T) {
	actual := date.Normalize("two weeks ahead")
	expected := two_weeks_ahead

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceOneMonthAhead(t *testing.T) {
	actual := date.Normalize("one month ahead")
	expected := one_month_ahead

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceTwoMonthsAhead(t *testing.T) {
	actual := date.Normalize("two months ahead")
	expected := two_months_ahead

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceOneYearAhead(t *testing.T) {
	actual := date.Normalize("one year ahead")
	expected := one_year_ahead

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistanceTwoYearsAhead(t *testing.T) {
	actual := date.Normalize("two years ahead")
	expected := two_years_ahead

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistance2DaysAgo(t *testing.T) {
	actual := date.Normalize("2 days ago")
	expected := two_days_ago

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}

func TestDateDistance1DayAgo(t *testing.T) {
	actual := date.Normalize("1 day ago")
	expected := yesterday

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %q, but was: %q", expected, actual)
	}
}
