package io_test

import (
  "bytes"
  "io/ioutil"
  "os"
  "testing"
  "github.com/svenfuchs/todo/io"
)

var (
  path = "/tmp/todo.test.txt"
)


func TestMain(m *testing.M) {
  setup()
  res := m.Run()
  teardown()
  os.Exit(res)
}

func setup() {
  io.Stdin  = bytes.NewBuffer([]byte{})
  io.Stdout = bytes.NewBuffer([]byte{})
  io.Stderr = bytes.NewBuffer([]byte{})
}

func teardown() {
  os.Remove(path)
}

func writeFile(data []byte) {
  err := ioutil.WriteFile(path, data, 0644)
  if err != nil { panic(err) }
}
