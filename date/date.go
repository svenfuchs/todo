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

func NewStub(str string) Clock {
  now, _ := time.Parse("2006-01-02", str)
  return stub { now }
}

var (
  Time Clock = t {}
  pattern = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
)

func Normalize(str string) string {
  if pattern.MatchString(str) {
    return str
  }

  str = strings.ToLower(str)
  time, err := ByName(str, Time)
  if err != nil {
    time, err = ByDistance(str, Time)
  }
  if err != nil {
    time, err = ByMark(str, Time)
  }
  if err != nil {
    return str
  }

  return time.Format("2006-01-02")
}

func rjoin(strs []string) string {
  return strings.Join(strs, "|")
}
