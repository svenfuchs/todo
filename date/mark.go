package date

import (
  "fmt"
  "time"
  "errors"
  "regexp"
  "strings"
)

var (
  months = []string{ "jan",  "feb",  "mar",  "apr",  "may",  "jun",  "jul",  "aug",  "sep",  "oct",  "nov",  "dec" }
  wdays  = []string{ "sun",  "mon",  "tue",  "wed",  "thu",  "fri",  "sat" }
  year   = regexp.MustCompile(fmt.Sprintf("^(last|next)%s(year)*$", sep))
  month  = regexp.MustCompile(fmt.Sprintf("^(last|next)%s(%s)[^\b]*$", sep, rjoin(months)))
  wday   = regexp.MustCompile(fmt.Sprintf("^(last|next)%s(%s)[^\b]*$", sep, rjoin(wdays)))
)

func byMark(str string, t Clock) (time.Time, error) {
  steps := map[string]int { "last": -1, "next": 1 }

  match := year.FindAllStringSubmatch(str, -1)
  if match != nil {
    step := steps[match[0][1]]
    date := t.Now()
    date  = date.AddDate(step, int(date.Month()) * -1 + 1, date.Day() * -1 + 1)
    return date, nil
  }

  match = month.FindAllStringSubmatch(str, -1)
  if match != nil {
    step := steps[match[0][1]]
    name := match[0][2]
    date := t.Now()
    date  = date.AddDate(0, step, date.Day() * -1 + 1)
    for strings.ToLower(date.Month().String())[0:3] != name {
      date = date.AddDate(0, step, 0)
    }
    return date, nil
  }

  match = wday.FindAllStringSubmatch(str, -1)
  if match != nil {
    step := steps[match[0][1]]
    name := match[0][2]
    date := t.Now().AddDate(0, 0, step)
    for strings.ToLower(date.Weekday().String())[0:3] != name {
      date = date.AddDate(0, 0, step)
    }
    return date, nil
  }

  return t.Now(), errors.New(fmt.Sprintf("Not a known date expression: %s", str))
}

