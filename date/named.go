package date

import (
  "fmt"
  "time"
  "errors"
)

var nameMap = map[string]func(string, Clock) time.Time {
  "yesterday": Yesterday,
  "today":     Today,
  "tomorrow":  Tomorrow,
}

func ByName(str string, t Clock) (time.Time, error) {
  f, ok := nameMap[str]
  if !ok {
    return t.Now(), errors.New(fmt.Sprintf("Not a named date: %s", str))
  }
  return f(str, t), nil
}

func Today(str string, t Clock) time.Time {
  return t.Now()
}

func Yesterday(str string, t Clock) time.Time {
  return t.Now().AddDate(0, 0, -1)
}

func Tomorrow(str string, t Clock) time.Time {
  return t.Now().AddDate(0, 0, 1)
}

