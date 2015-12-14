package item

import (
  "strings"
)

var modes = map[string][]int {
  "date":   []int { 0    },
  "after":  []int { 1    },
  "since":  []int { 0, 1 },
  "before": []int { -1   },
}

type Filter struct {
  Id int
  Status string
  Text string
  Projects []string
  Date string
  Mode string
}

func (f Filter) Apply(i Item) bool {
  if i.IsNone() {
    return false
  } else if f.Id != 0 {
    return f.Id == i.Id
  } else {
    return f.matchesData(i)
  }
}

func (f Filter) matchesData(i Item) bool {
  return f.matchesText(i) &&
    f.matchesStatus(i)    &&
    f.matchesProjects(i)  &&
    f.matchesDate(i)
}

func (f Filter) matchesText(i Item) bool {
  return f.Text == "" || strings.Contains(i.Text, f.Text)
}

func (f Filter) matchesStatus(i Item) bool {
  return f.Status == "" || f.Status == i.Status
}

func (f Filter) matchesDate(i Item) bool {
  mode, ok := modes[f.Mode]
  if !ok { return true }

  date := f.statusDate(i)
  if date == "" { return false }

  cmp := strings.Compare(date, f.Date)
  return includes(mode, cmp)
}

func (f Filter) statusDate(i Item) string {
  switch f.Status {
    // case "added":
    //   return i.AddedDate()
    // case "due":
    //   return i.DueDate()
    // case "done":
    //   return i.DoneDate()
    default:
      return i.DoneDate()
  }
}

func (f Filter) matchesProjects(i Item) bool {
  if len(f.Projects) == 0 {
    return true
  }
  if len(intersect(i.Projects, f.Projects)) > 0 {
    return true
  }
  return false
}

func intersect(strs1 []string, strs2 []string) []string {
  res := []string{}
  for _, str1 := range strs1 {
    for _, str2 := range strs2 {
      if str1 == str2 {
        res = append(res, str1)
      }
    }
  }
  return res
}

func includes(nums []int, num int) bool {
  for _, n := range nums {
    if num == n { return true }
  }
  return false
}

