package source

import (
  "io/ioutil"
  "os"
  "testing"
  . "github.com/svenfuchs/todo/test"
)

const (
  path string = "/tmp/todo.test.txt"
)

func TestMain(m *testing.M) {
  setup()
  res := m.Run()
  teardown()
  os.Exit(res)
}

func setup() {
  data := []byte("- foo [1]\nx bar [2]\n")
  checkErr(ioutil.WriteFile(path, data, 0644))
}

func teardown() {
  checkErr(os.Remove(path))
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func TestFileSourceReadLines(t *testing.T) {
  actual, _ := NewFileSource(path).ReadLines()
  expected  := []string{ "- foo [1]", "x bar [2]" }
  AssertEqual(t, actual, expected)
}

func TestFileSourceWriteLines(t *testing.T) {
  source := NewFileSource(path)
  source.WriteLines([]string{ "- foo [1]", "x bar [2]" })

  actual, _ := source.ReadLines()
  expected  := []string{ "- foo [1]", "x bar [2]" }
  AssertEqual(t, actual, expected)
}

