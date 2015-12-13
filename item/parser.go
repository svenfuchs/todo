package item

import (
  "regexp"
  "strconv"
  "strings"
  "github.com/svenfuchs/todo.go/date"
)

var (
  statusPattern   = regexp.MustCompile(`^([-x]{1})`)
  idPattern       = regexp.MustCompile(` \[(\d+)\]`)
  tagsPattern     = regexp.MustCompile(` ([^\s]+):([^\s]+)`)
  projectsPattern = regexp.MustCompile(` \+([\w\-]+)`)
)

type parser struct {
  line string
}

func (p parser) id() int {
  matches:= idPattern.FindStringSubmatch(p.line)
  if matches == nil {
    return 0
  }
  id, _ := strconv.Atoi(matches[1])
  return id
}

func (p parser) status() string {
  match := statusPattern.FindString(p.line)
  switch match {
    case "-":
      return Pend
    case "x":
      return Done
    default:
      return None
  }
}

func (p parser) tags() map[string]string {
  tags := map[string]string{}
  for _, match := range tagsPattern.FindAllStringSubmatch(p.line, -1) {
    tags[match[1]] = date.Normalize(match[2], date.Time)
  }
  return tags
}

func (p parser) text() string {
  text := p.line
  text = idPattern.ReplaceAllString(text, "")
  text = statusPattern.ReplaceAllString(text, "")
  text = tagsPattern.ReplaceAllString(text, "")
  return strings.TrimSpace(text)
}

func (p parser) projects() []string {
  matches := projectsPattern.FindAllStringSubmatch(p.text(), -1)
  projects := []string{}
  for _, match := range matches {
    projects = append(projects, match[1])
  }
  return projects
}
