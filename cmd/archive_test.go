package cmd

import (
	. "github.com/svenfuchs/todo/io"
	. "github.com/svenfuchs/todo/test"
	"testing"
)

func TestCmdArchiveAppendsToArchive(t *testing.T) {
	args := &Args{Ids: []int{1}}
	src := NewMemoryIo("x foo [1]\nx bar [2]")
	out := NewMemoryIo("")
	arch := NewMemoryIo("x baz [3]\n")

	ArchiveCmd{Cmd{args, src, out}, arch}.Run()
	actual := arch.ReadLines()
	expected := []string{"x baz [3]", "x foo [1]"}

	AssertEqual(t, actual, expected)
}

func TestCmdArchiveRemovesFromSource(t *testing.T) {
	args := &Args{Ids: []int{1}}
	src := NewMemoryIo("x foo [1]\nx bar [2]")
	out := NewMemoryIo("")
	arch := NewMemoryIo("")

	ArchiveCmd{Cmd{args, src, out}, arch}.Run()
	actual := src.ReadLines()
	expected := []string{"x bar [2]"}

	AssertEqual(t, actual, expected)
}

func TestCmdArchiveReportsArchived(t *testing.T) {
	args := &Args{Ids: []int{1}, Report: true}
	src := NewMemoryIo("x foo [1]\nx bar [2]")
	out := NewMemoryIo("")
	arch := NewMemoryIo("")

	ArchiveCmd{Cmd{args, src, out}, arch}.Run()
	actual := out.ReadLines()
	expected := []string{"Archived 1 item:", "", "x foo [1]"}

	AssertEqual(t, actual, expected)
}
