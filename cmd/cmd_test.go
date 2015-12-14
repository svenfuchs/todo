package cmd_test

import (
  "fmt"
  "testing"
  "time"
  "github.com/svenfuchs/todo.go/cmd"
  "github.com/svenfuchs/todo.go/item"
  "github.com/svenfuchs/todo.go/source"
  . "github.com/svenfuchs/todo.go/test"
)

func TestCmdListByIdFound(t *testing.T) {
  in     := source.Memory { Content: "- foo [1]\nx bar [2]" }
  out    := source.Memory {}
  filter := item.Filter{ Id: 1 }
  format := ""

  cmd.List { cmd.Cmd { &in, &out, filter, format } }.Run()
  actual   := out.MustReadLines()
  expected := []string{ "- foo" }

  AssertEqual(t, actual, expected)
}

func TestCmdListByTextFound(t *testing.T) {
  in     := source.Memory { Content: "- foo [1]\nx bar [2]" }
  out    := source.Memory {}
  filter := item.Filter{ Text: "bar" }
  format := ""

  cmd.List { cmd.Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "x bar" }

  AssertEqual(t, actual, expected)
}

func TestCmdListByProjectsFound(t *testing.T) {
  in     := source.Memory { Content: "- foo +baz [1]\nx bar +baz [2]" }
  out    := source.Memory {}
  filter := item.Filter{ Projects: []string { "baz" } }
  format := ""

  cmd.List { cmd.Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "- foo +baz", "x bar +baz" }

  AssertEqual(t, actual, expected)
}

func TestCmdListFormat(t *testing.T) {
  in     := source.Memory { Content: "- foo [1]\nx bar [2]" }
  out    := source.Memory {}
  filter := item.Filter{}
  format := "id,text"

  cmd.List { cmd.Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "[1] foo", "[2] bar" }

  AssertEqual(t, actual, expected)
}

func TestCmdToggleByIdFound(t *testing.T) {
  now    := time.Now().Format("2006-01-02")
  in     := source.Memory { Content: "# Comment\n- foo [1]\nx bar [2]" }
  out    := source.Memory {}
  filter := item.Filter{ Id: 1 }
  format := ""

  cmd.Toggle { cmd.Cmd { &in, &out, filter, format } }.Run()
  actual   := out.MustReadLines()
  expected := []string{ "# Comment", fmt.Sprintf("x foo done:%s [1]", now), "x bar [2]" }

  AssertEqual(t, actual, expected)
}

func TestCmdToggleByTextFound(t *testing.T) {
  in     := source.Memory { Content: "# Comment\n- foo [1]\nx bar done:2015-12-13 [2]" }
  out    := source.Memory {}
  filter := item.Filter{ Text: "bar" }
  format := ""

  cmd.Toggle { cmd.Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected := []string{ "# Comment", "- foo [1]", "- bar [2]" }

  AssertEqual(t, actual, expected)
}

func TestCmdToggleByProjectsFound(t *testing.T) {
  now    := time.Now().Format("2006-01-02")
  in     := source.Memory { Content: "- foo +baz [1]\nx bar +baz [2]" }
  out    := source.Memory {}
  filter := item.Filter{ Projects: []string { "baz" } }
  format := ""

  cmd.Toggle { cmd.Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected := []string{ fmt.Sprintf("x foo +baz done:%s [1]", now), "- bar +baz [2]" }

  AssertEqual(t, actual, expected)
}
