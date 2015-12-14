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
  tagsPattern     = regexp.MustCompile(` ([\w\-]+):([\w\-]+)`)
  projectsPattern = regexp.MustCompile(` \+([\w\-]+)`)
)

type Parser struct {
  Line string
}

func (p Parser) Id() int {
  matches:= idPattern.FindStringSubmatch(p.Line)
  if matches == nil {
    return 0
  }
  id, _ := strconv.Atoi(matches[1])
  return id
}

func (p Parser) Status() string {
  match := statusPattern.FindString(p.Line)
  switch match {
    case "-":
      return Pend
    case "x":
      return Done
    default:
      return None
  }
}

func (p Parser) Tags() map[string]string {
  tags := map[string]string{}
  for _, match := range tagsPattern.FindAllStringSubmatch(p.Line, -1) {
    tags[match[1]] = date.Normalize(match[2], date.Time)
  }
  return tags
}

func (p Parser) Text() string {
  text := p.Line
  text = idPattern.ReplaceAllString(text, "")
  text = statusPattern.ReplaceAllString(text, "")
  text = tagsPattern.ReplaceAllString(text, "")
  return strings.TrimSpace(text)
}

func (p Parser) Projects() []string {
  matches := projectsPattern.FindAllStringSubmatch(p.Text(), -1)
  projects := []string{}
  for _, match := range matches {
    projects = append(projects, match[1])
  }
  return projects
}
