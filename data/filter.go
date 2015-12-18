package data

import (
	"github.com/svenfuchs/todo/date"
	"strings"
)

func NewFilter(ids []int, status string, text string, projects []string, date string) Filter {
	if len(status) > 4 {
		status = status[0:4]
	}
	if len(ids) == 1 && ids[0] == 0 {
		ids = []int{}
	}
	return Filter{ids, status, text, projects, NewFilterDate(date)}
}

func ParseFilter(line string) Filter {
	p := NewParser(line)
	return NewFilter([]int{p.Id()}, p.Status(), p.Text(), []string{}, "")
}

type Filter struct {
	Ids      []int
	Status   string
	Text     string
	Projects []string
	Date     FilterDate
}

func (f Filter) Apply(i Item) bool {
	if i.IsNone() {
		return false
	} else if len(f.Ids) > 0 && i.Id != 0 {
		return includesInt(f.Ids, i.Id)
	} else {
		return f.matchesData(i)
	}
}

func (f Filter) IsEmpty() bool {
	return len(f.Ids) == 0 && len(f.Projects) == 0 && f.Status == "" && f.Text == "" && f.Date.IsEmpty()
}

func (f Filter) matchesData(i Item) bool {
	return f.matchesText(i) &&
		f.matchesStatus(i) &&
		f.matchesProjects(i) &&
		f.matchesDate(i)
}

func (f Filter) matchesText(i Item) bool {
	return f.Text == "" || strings.Contains(i.Text, f.Text)
}

func (f Filter) matchesStatus(i Item) bool {
	return f.Status == "" || f.Status == i.Status
}

func (f Filter) matchesProjects(i Item) bool {
	return len(f.Projects) == 0 || len(intersect(i.Projects, f.Projects)) > 0
}

func (f Filter) matchesDate(i Item) bool {
	return f.Date.matches(f.statusDate(i))
}

func (f Filter) statusDate(i Item) string {
	// switch f.status {
	//   case "added":
	//     return i.AddedDate()
	//   case "due":
	//     return i.DueDate()
	//   case "done":
	//     return i.DoneDate()
	//   default:
	//     return i.DoneDate()
	// }
	return i.DoneDate()
}

func NewFilterDate(date string) FilterDate {
	if date == "" {
		date = ":"
	}
	parts := strings.SplitN(date, ":", 2)
	return FilterDate{parts[0], parts[1]}
}

type FilterDate struct {
	mode string
	date string
}

func (d FilterDate) IsEmpty() bool {
	return d.date == ""
}

func (d FilterDate) matches(date string) bool {
	if d.date == "" {
		return true
	} else if date == "" {
		return false
	}

	cmp := strings.Compare(date, normalizeDate(d.date))
	switch d.mode {
	case "after":
		return cmp == 1
	case "since":
		return cmp == 0 || cmp == 1
	case "before":
		return cmp == -1
	default:
		return cmp == 0
	}
}

func normalizeDate(str string) string {
	return date.Normalize(str)
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

func includesInt(nums []int, num int) bool {
	for _, n := range nums {
		if num == n {
			return true
		}
	}
	return false
}
