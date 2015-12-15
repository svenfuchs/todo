package todo

import (
  "fmt"
  "testing"
  "time"
  . "github.com/svenfuchs/todo/test"
)

func TestCmdListByIdFound(t *testing.T) {
  in     := NewMemorySource("- foo [1]\nx bar [2]")
  out    := NewMemorySource("")
  filter := Filter{ id: 1 }
  format := ""

  ListCmd { Cmd { &in, &out, filter, format } }.Run()
  actual   := out.MustReadLines()
  expected := []string{ "- foo" }

  AssertEqual(t, actual, expected)
}

func TestCmdListByTextFound(t *testing.T) {
  in     := NewMemorySource("- foo [1]\nx bar [2]")
  out    := NewMemorySource("")
  filter := Filter{ text: "bar" }
  format := ""

  ListCmd { Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "x bar" }

  AssertEqual(t, actual, expected)
}

func TestCmdListByProjectsFound(t *testing.T) {
  in     := NewMemorySource("- foo +baz [1]\nx bar +baz [2]")
  out    := NewMemorySource("")
  filter := Filter{ projects: []string { "baz" } }
  format := ""

  ListCmd { Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "- foo +baz", "x bar +baz" }

  AssertEqual(t, actual, expected)
}

func TestCmdListFormat(t *testing.T) {
  in     := NewMemorySource("- foo [1]\nx bar [2]")
  out    := NewMemorySource("")
  filter := Filter{}
  format := "id,text"

  ListCmd { Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "[1] foo", "[2] bar" }

  AssertEqual(t, actual, expected)
}

func TestCmdToggleByIdFound(t *testing.T) {
  now    := time.Now().Format("2006-01-02")
  in     := NewMemorySource("# Comment\n- foo [1]\nx bar [2]")
  out    := NewMemorySource("")
  filter := Filter{ id: 1 }
  format := ""

  ToggleCmd { Cmd { &in, &out, filter, format } }.Run()
  actual   := out.MustReadLines()
  expected := []string{ "# Comment", fmt.Sprintf("x foo done:%s [1]", now), "x bar [2]" }

  AssertEqual(t, actual, expected)
}

func TestCmdToggleByTextFound(t *testing.T) {
  in     := NewMemorySource("# Comment\n- foo [1]\nx bar done:2015-12-13 [2]")
  out    := NewMemorySource("")
  filter := Filter{ text: "bar" }
  format := ""

  ToggleCmd { Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected := []string{ "# Comment", "- foo [1]", "- bar [2]" }

  AssertEqual(t, actual, expected)
}

func TestCmdToggleByProjectsFound(t *testing.T) {
  now    := time.Now().Format("2006-01-02")
  in     := NewMemorySource("- foo +baz [1]\nx bar +baz [2]")
  out    := NewMemorySource("")
  filter := Filter{ projects: []string { "baz" } }
  format := ""

  ToggleCmd { Cmd { &in, &out, filter, format } }.Run()
  actual, _ := out.ReadLines()
  expected := []string{ fmt.Sprintf("x foo +baz done:%s [1]", now), "- bar +baz [2]" }

  AssertEqual(t, actual, expected)
}
