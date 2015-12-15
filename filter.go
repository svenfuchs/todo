package todo

import (
  "strings"
  "github.com/svenfuchs/todo/date"
)

func NewFilter(id int, status string, text string, projects []string, date FilterDate) Filter {
  if len(status) > 4 {
    status = status[0:4]
  }
  return Filter{ id, status, text, projects, date }
}

func ParseFilter(line string) Filter {
  p := NewParser(line)
  return NewFilter(p.Id(), p.Status(), p.Text(), []string{}, FilterDate{})
}

type Filter struct {
  id int
  status string
  text string
  projects []string
  date FilterDate
}

func (f Filter) Apply(i Item) bool {
  if i.IsNone() {
    return false
  } else if f.id != 0 {
    return f.id == i.id
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
  return f.text == "" || strings.Contains(i.text, f.text)
}

func (f Filter) matchesStatus(i Item) bool {
  return f.status == "" || f.status == i.status
}

func (f Filter) matchesProjects(i Item) bool {
  return len(f.projects) == 0  || len(intersect(i.projects, f.projects)) > 0
}

func (f Filter) matchesDate(i Item) bool {
  return f.date.matches(f.statusDate(i))
}

func (f Filter) statusDate(i Item) string {
  switch f.status {
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

func NewFilterDate(d string, m string) FilterDate {
  return FilterDate{ date.Normalize(d, date.Time), m }
}

type FilterDate struct {
  date string
  mode string
}

func (d FilterDate) matches(date string) bool {
  modes := map[string][]int {
    "date":   []int { 0    },
    "after":  []int { 1    },
    "since":  []int { 0, 1 },
    "before": []int { -1   },
  }

  cmp, ok := modes[d.mode]
  if !ok {
    return true
  } else if date == "" {
    return false
  }
  return includes(cmp, strings.Compare(date, d.date))
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
