package date

import (
  "fmt"
  "time"
  "errors"
  "strconv"
  "regexp"
)

var (
  nums  = []string{ `\d+`, "one",  "two",  "three",  "four",  "five",  "six",  "seven",  "eight",  "nine",  "ten" }
  units = []string{ "year", "month", "day", "week" }
  sep   = `(?: |_|\-|\.)`
  ago   = regexp.MustCompile(fmt.Sprintf("^(%s)%s(%s)s?%s(ago|ahead)$", rjoin(nums), sep, rjoin(units), sep))
)

func byDistance(str string, t Clock) (time.Time, error) {
  match := ago.FindAllStringSubmatch(str, -1)
  if match == nil {
    return t.Now(), errors.New(fmt.Sprintf("Not a known date expression: %s", str))
  }

  unit, dir, num := match[0][2], match[0][3], parseNum(match[0][1])
  if unit == "week" {
    num, unit = num * 7, "day"
  }
  if dir == "ago" {
    num = num * -1
  }

  args := []int{0, 0, 0}
  i, _ := index(units, unit)
  args[i] = num
  return t.Now().AddDate(args[0], args[1], args[2]), nil
}

func parseNum(str string) int {
  num, ok := index(nums, str)
  if !ok {
    num, _ = strconv.Atoi(str)
  }
  return num
}

func index(strs []string, str string) (int, bool) {
  for ix, s := range strs {
    if s == str { return ix, true }
  }
  return -1, false
}
