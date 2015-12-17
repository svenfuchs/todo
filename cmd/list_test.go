package cmd

import (
  "testing"
  . "github.com/svenfuchs/todo/io"
  . "github.com/svenfuchs/todo/test"
)

func TestCmdListByIdFound(t *testing.T) {
  args := &Args{ Ids: []int{ 1 } }
  src  := NewMemoryIo("- foo [1]\nx bar [2]")
  out  := NewMemoryIo("")

  cmd := ListCmd { Cmd { args, src, out } }
  cmd.Run()
  actual, _ := out.ReadLines()
  expected := []string{ "- foo [1]" }

  AssertEqual(t, actual, expected)
}

func TestCmdListByTextFound(t *testing.T) {
  args   := &Args{ Text: "bar" }
  src    := NewMemoryIo("- foo [1]\nx bar [2]")
  out    := NewMemoryIo("")

  ListCmd { Cmd { args, src, out } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "x bar [2]" }

  AssertEqual(t, actual, expected)
}

func TestCmdListByProjectsFound(t *testing.T) {
  args   := &Args{ Projects: []string { "baz" } }
  src    := NewMemoryIo("- foo +baz [1]\nx bar +baz [2]")
  out    := NewMemoryIo("")

  ListCmd { Cmd { args, src, out } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "- foo +baz [1]", "x bar +baz [2]" }

  AssertEqual(t, actual, expected)
}

func TestCmdListFormat(t *testing.T) {
  args   := &Args{ Format: "id,text" }
  src    := NewMemoryIo("- foo [1]\nx bar [2]")
  out    := NewMemoryIo("")

  ListCmd { Cmd { args, src, out } }.Run()
  actual, _ := out.ReadLines()
  expected  := []string{ "[1] foo", "[2] bar" }

  AssertEqual(t, actual, expected)
}
