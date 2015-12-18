package service

import (
  "bytes"
  "encoding/json"
  "time"
  api "github.com/svenfuchs/todo/service/idonethis"
)

var (
  idonethis_uri = "https://idonethis.com/api/v0.1/dones/"
)

func NewIdonethis(config map[string]string) Idonethis {
  after := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
  return Idonethis { config["team"], config["username"], config["token"], after }
}

type Idonethis struct {
  team  string
  user  string
  token string
  after string
}

func (s Idonethis) Push(line string) {
  s.post(line)
}

func (s Idonethis) Fetch() []string {
  body  := s.get()
  page  := s.decode(body)
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
  query  := &map[string]string {
    "owner": s.user,
    "team": s.team,
    "done_date_after": s.after,
    "page_size": "100",
  }
  headers := &map[string]string {
    "Authorization": "Token " + s.token,
    "Content-Type":  "application/json",
  }
  request := NewHttpRequest("GET", idonethis_uri, query, headers, nil)
  return request.Run()
}

func (s Idonethis) post(line string) {
  headers := &map[string]string {
    "Authorization": "Token " + s.token,
    "Content-Type":  "application/x-www-form-urlencoded",
  }
  body := bytes.NewBufferString(toQueryString(map[string]string {
    "team": s.team,
    "raw_text": line,
  }));
  request := NewHttpRequest("POST", idonethis_uri, nil, headers, body)
  request.Run()
}
