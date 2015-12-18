package io_test

import (
  "testing"
  . "github.com/svenfuchs/todo/io"
  . "github.com/svenfuchs/todo/test"
)

func TestFileIoReadLines(t *testing.T) {
  writeFile([]byte("- foo\nx bar\n"))
  actual := NewFileIo(path).ReadLines()
  expected := []string{ "- foo", "x bar" }
  AssertEqual(t, actual, expected)
}

func TestFileIoWriteLinesCreatesFile(t *testing.T) {
  io := NewFileIo(path)
  io.WriteLines([]string{ "x foo", "- bar" })

  actual := io.ReadLines()
  expected := []string{ "x foo", "- bar" }
  AssertEqual(t, actual, expected)
}

func TestFileIoWriteLinesReplacesContent(t *testing.T) {
  writeFile([]byte("- baz\n"))
  io := NewFileIo(path)
  io.WriteLines([]string{ "x foo", "- bar" })

  actual := io.ReadLines()
  expected := []string{ "x foo", "- bar" }
  AssertEqual(t, actual, expected)
}


func TestFileIoAppendLinesAppendsToContent(t *testing.T) {
  writeFile([]byte("- baz\n"))
  io := NewFileIo(path)
  io.AppendLines([]string{ "x foo", "- bar" })

  actual := io.ReadLines()
  expected := []string{ "- baz", "x foo", "- bar" }
  AssertEqual(t, actual, expected)
}

