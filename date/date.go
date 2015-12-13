package date

import (
  "time"
  "regexp"
  "strings"
)

type Clock interface { Now() time.Time }

type t struct {}
func (t) Now() time.Time { return time.Now() }

type stub struct { now time.Time }
func (s stub) Now() time.Time { return s.now }

var (
  Time = t {}
  date = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
)

func Stub(str string) Clock {
  now, _ := time.Parse("2006-01-02", str)
  return stub { now }
}

func Normalize(str string, t Clock) string {
  if date.MatchString(str) {
    return str
  }

  str = strings.ToLower(str)
  time, err := byName(str, t)
  if err != nil {
    time, err = byDistance(str, t)
  }
  if err != nil {
    time, err = byMark(str, t)
  }
  if err != nil {
    return str
  }

  return time.Format("2006-01-02")
}

func rjoin(strs []string) string {
  return strings.Join(strs, "|")
}
