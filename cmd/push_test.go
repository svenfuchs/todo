package cmd

import (
	. "github.com/svenfuchs/todo/io"
	. "github.com/svenfuchs/todo/test"
	"testing"
)

func TestCmdPush(t *testing.T) {
	args := &Args{Ids: []int{1}}
	src := NewMemoryIo("x foo key:value [1]\nx bar [2]")
	out := NewMemoryIo("")
	srv := NewMemoryIo("")

	PushCmd{Cmd{args, src, out}, srv}.Run()
	actual := srv.ReadLines()
	expected := []string{"foo key:value [1]"}

	AssertEqual(t, actual, expected)
}

func TestCmdPushReportsPushed(t *testing.T) {
	args := &Args{Ids: []int{1}, Report: true}
	src := NewMemoryIo("x foo [1]\nx bar [2]")
	out := NewMemoryIo("")
	srv := NewMemoryIo("")

	PushCmd{Cmd{args, src, out}, srv}.Run()
	actual := out.ReadLines()
	expected := []string{"Pushed 1 item:", "", "foo [1]"}

	AssertEqual(t, actual, expected)
}
