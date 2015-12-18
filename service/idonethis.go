package service

import (
	"bytes"
	"encoding/json"
	api "github.com/svenfuchs/todo/service/idonethis"
	"time"
)

var (
	idonethis_uri = "https://idonethis.com/api/v0.1/dones/"
)

func NewIdonethis(config map[string]string) Idonethis {
	after := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	http := NewHttpClient()
	return Idonethis{&http, config["team"], config["username"], config["token"], after}
}

type Idonethis struct {
	http  *HttpClient
	team  string
	user  string
	token string
	after string
}

func (s Idonethis) WriteLines(lines []string) {
	for _, line := range lines {
		s.post(line)
	}
}

func (s Idonethis) ReadLines() []string {
	body := s.get()
	page := s.decode(body)
	lines := []string{}

	for _, done := range page.Results {
		lines = append(lines, done.Raw_text)
	}
	return lines
}

func (s Idonethis) decode(body []byte) api.Page {
	var page api.Page
	json.Unmarshal(body, &page)
	return page
}

func (s Idonethis) get() []byte {
	query := &map[string]string{
		"owner":           s.user,
		"team":            s.team,
		"done_date_after": s.after,
		"page_size":       "100",
	}
	headers := &map[string]string{
		"Authorization": "Token " + s.token,
		"Content-Type":  "application/json",
	}
	return s.http.run("GET", idonethis_uri, query, headers, nil)
}

func (s Idonethis) post(line string) {
	headers := &map[string]string{
		"Authorization": "Token " + s.token,
		"Content-Type":  "application/x-www-form-urlencoded",
	}
	body := bytes.NewBufferString(toQueryString(map[string]string{
		"team":     s.team,
		"raw_text": line,
	}))
	s.http.run("POST", idonethis_uri, nil, headers, body)
}
