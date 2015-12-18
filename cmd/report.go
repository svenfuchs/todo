package cmd

import (
  "bytes"
  "log"
  "regexp"
  "strings"
  "text/template"
  . "github.com/svenfuchs/todo/data"
  . "github.com/svenfuchs/todo/io"
)

type Report struct {
  Action string
  Count  int
  Plural bool
  Any    bool
  Items  string
}

var tmpl = template.Must(template.New("report").Parse(
  "{{.Action}} {{.Count}} item{{if .Plural}}s{{end}}{{if .Any}}:\n\n{{.Items}}{{else}}.{{end}}",
))

func (c Cmd) report(out Io, action string, list List) {
  if !c.args.Report { return }
  report := c.reportFor(action, list)
  result := c.render("report", report)
  out.WriteLines(strings.Split(result, "\n"))
}

func (c Cmd) reportFor(action string, list List) Report {
  action = camelize(participle(action))
  count := len(list.Items)
  lines := strings.Join(c.formatted(list.Items), "\n")
  return Report { action, count, count > 1, count > 0, lines }

}

func (c Cmd) render(name string, report Report) string {
  buffer := new(bytes.Buffer)
  err := tmpl.Execute(buffer, report)
  if err != nil { log.Fatal(err) }
  return buffer.String()
}

func participle(str string) string {
  return regexp.MustCompile(`eed$`).ReplaceAllString(str + "ed", "ed")
}

func camelize(str string) string {
	r := regexp.MustCompile(`(\b|\-|_|\s)+(.)?`)
	return r.ReplaceAllStringFunc(str, func(str string) string {
    return regexp.MustCompile(`(\-|_|\s)+`).ReplaceAllString(strings.ToUpper(str), "")
	})
}
