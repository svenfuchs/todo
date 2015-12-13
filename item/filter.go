package item

import (
  "strings"
  "github.com/svenfuchs/todo.go/date"
)

func NewFilter(line string, id int, status string, text string, projects string, after string, before string, due string) Filter {
  if line != "" {
    p := parser{line}
    id = p.id()
    status = p.status()
    text = p.text()
  }
  after  = date.Normalize(after, date.Time)
  before = date.Normalize(before, date.Time)
  due    = date.Normalize(due, date.Time)
  return Filter { id, status, text, projects, after, before, due }
}

type Filter struct {
  Id int
  Status string
  Text string
  Projects string
  After string
  Before string
  Due string
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
    f.matchesAfter(i)     &&
    f.matchesBefore(i)    &&
    f.matchesDue(i)
}

func (f Filter) matchesText(i Item) bool {
  return f.Text == "" || strings.Contains(i.Text, f.Text)
}

func (f Filter) matchesStatus(i Item) bool {
  return f.Status == "" || f.Status == i.Status
}

func (f Filter) matchesAfter(i Item) bool {
  return f.After == "" || f.After <= i.DoneDate()
}

func (f Filter) matchesBefore(i Item) bool {
  return f.Before == "" || i.DoneDate() != "" && f.Before > i.DoneDate()
}

func (f Filter) matchesDue(i Item) bool {
  return f.Due == "" || i.DueDate() != "" && f.Due <= i.DueDate()
}

func (f Filter) matchesProjects(i Item) bool {
  p := strings.Fields(strings.Replace(f.Projects, ",", " ", -1))
  if len(p) == 0 {
    return true
  }
  if len(intersect(i.Projects, p)) > 0 {
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
