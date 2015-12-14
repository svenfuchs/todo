package source_test

import (
  "io/ioutil"
  "os"
  "testing"
  "github.com/svenfuchs/todo.go/source"
  . "github.com/svenfuchs/todo.go/test"
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

func TestSourceFileReadLines(t *testing.T) {
  actual, _ := source.New(path).ReadLines()
  expected  := []string{ "- foo [1]", "x bar [2]" }

  AssertEqual(t, actual, expected)
}

func TestSourceFileWriteLines(t *testing.T) {
  source    := source.New(path)
  source.WriteLines([]string{ "- foo [1]", "x bar [2]" })

  actual, _ := source.ReadLines()
  expected  := []string{ "- foo [1]", "x bar [2]" }

  AssertEqual(t, actual, expected)
}

