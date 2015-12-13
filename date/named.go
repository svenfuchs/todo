package date

import (
  "fmt"
  "time"
  "errors"
)

func byName(str string, t Clock) (time.Time, error) {
  f, ok := nameMap[str]
  if !ok {
    return t.Now(), errors.New(fmt.Sprintf("Not a named date: %s", str))
  }
  return f(str, t), nil
}

var nameMap = map[string]func(string, Clock) time.Time {
  "yesterday": yesterday,
  "today":     today,
  "tomorrow":  tomorrow,
}

func today(str string, t Clock) time.Time {
  return t.Now()
}

func yesterday(str string, t Clock) time.Time {
  return t.Now().AddDate(0, 0, -1)
}

func tomorrow(str string, t Clock) time.Time {
  return t.Now().AddDate(0, 0, 1)
}

