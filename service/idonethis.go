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

func (s Idonethis) Push(line string) error {
  _, err := s.post(line)
  return err
}

func (s Idonethis) Fetch() ([]string, error) {
  var lines []string

  body, err := s.get()
  if err != nil { return lines, err }
  page := s.decode(body)

  for _, done := range page.Results {
    lines = append(lines, done.Raw_text)
  }
  return lines, nil
}

func (s Idonethis) decode(body []byte) api.Page {
  var page api.Page
  json.Unmarshal(body, &page)
  return page
}

func (s Idonethis) get() ([]byte, error) {
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
  return NewHttpRequest("GET", idonethis_uri, query, headers, nil).run()
}

func (s Idonethis) post(line string) ([]byte, error) {
  headers := &map[string]string {
    "Authorization": "Token " + s.token,
    "Content-Type":  "application/x-www-form-urlencoded",
  }
  body := bytes.NewBufferString(toQueryString(map[string]string {
    "team": s.team,
    "raw_text": line,
  }));
  return NewHttpRequest("POST", idonethis_uri, nil, headers, body).run()
}
