package service

import (
  "testing"
  . "github.com/svenfuchs/todo/test"
)

func stubGet() {
  uri     := "https://idonethis.com/api/v0.1/dones/?done_date_after=2015-12-11&owner=username&page_size=100&team=team"
  status  := 200
  headers := map[string]string{ "Content-Type": "application/json" }
  body, _ := ReadFile(ExpandRelativePath("../../test/http/idonethis.json"))
  StubRequest(uri, status, headers, body)
}

func TestIdonethis(t *testing.T) {
  stubGet()
  config := map[string]string{ "username": "username", "token": "token", "team": "team"}
  service := NewIdonethis(config)
  lines := service.ReadLines()

  AssertEqual(t, lines[0], "Add a stale job metric to monitor done:2015-12-10 [162]")
}
