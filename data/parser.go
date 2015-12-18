package data

import (
	"github.com/svenfuchs/todo/date"
	"regexp"
	"strconv"
	"strings"
)

var patterns = map[string]*regexp.Regexp{
	"status":   regexp.MustCompile(`^([-x]{1})`),
	"id":       regexp.MustCompile(` \[(\d+)\]`),
	"tags":     regexp.MustCompile(` ([\w\-]+):([\w\-]+)`),
	"projects": regexp.MustCompile(` \+([\w\-]+)`),
}

func NewParser(line string) Parser {
	return Parser{line}
}

type Parser struct {
	Line string
}

func (p Parser) Id() int {
	matches := patterns["id"].FindStringSubmatch(p.Line)
	if matches == nil {
		return 0
	}
	id, _ := strconv.Atoi(matches[1])
	return id
}

func (p Parser) Status() string {
	match := patterns["status"].FindString(p.Line)
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
	for _, match := range patterns["tags"].FindAllStringSubmatch(p.Line, -1) {
		tags[match[1]] = date.Normalize(match[2])
	}
	return tags
}

func (p Parser) Text() string {
	text := p.Line
	for _, key := range []string{"status", "id", "tags"} {
		text = patterns[key].ReplaceAllString(text, "")
	}
	return strings.TrimSpace(text)
}

func (p Parser) Projects() []string {
	matches := patterns["projects"].FindAllStringSubmatch(p.Text(), -1)
	projects := []string{}
	for _, match := range matches {
		projects = append(projects, match[1])
	}
	return projects
}
