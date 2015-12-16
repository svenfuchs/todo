package service

// import (
//   "os"
//   "testing"
//   . "github.com/svenfuchs/todo/test"
// )
//
// var (
//   team  = "personal" // os.Getenv("IDONETHIS_TEAM")
//   user  = os.Getenv("IDONETHIS_USERNAME")
//   token = os.Getenv("IDONETHIS_TOKEN")
// )
//
// func stubGet() {
//   uri     := "https://idonethis.com/api/v0.1/dones/?owner=svenfuchs"
//   status  := 200
//   headers := map[string]string{ "Content-Type": "application/json" }
//   body, _ := ReadFile(ExpandRelativePath("../stubs/idonethis.json"))
//   StubRequest(uri, status, headers, body)
// }
//
//
// func TestIdonethis(t *testing.T) {
//   stubGet()
//   s := NewIdonethis(team, user, token, "")
//   ids, _ := s.ids()
//
//   AssertEqual(t, ids[0], 162)
// }
