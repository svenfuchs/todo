package todo

import (
  "testing"
  "time"
  . "github.com/svenfuchs/todo/test"
  // "fmt"
)

func TestParserText(t *testing.T) {
  actual   := Parser{ "- foo bar baz" }.Text()
  expected := "foo bar baz"
  AssertEqual(t, actual, expected)
}

func TestParserIdValid(t *testing.T) {
  actual   := Parser{ "- foo [1]" }.Id()
  expected := 1
  AssertEqual(t, actual, expected)
}

func TestParserIdMissing(t *testing.T) {
  actual   := Parser{ "- foo" }.Id()
  expected := 0
  AssertEqual(t, actual, expected)
}

func TestParserStatusPending(t *testing.T) {
  actual   := Parser{ "- foo" }.Status()
  expected := Pend
  AssertEqual(t, actual, expected)
}

func TestParserStatusDone(t *testing.T) {
  actual   := Parser{ "x foo" }.Status()
  expected := Done
  AssertEqual(t, actual, expected)
}

func TestParserStatusNone(t *testing.T) {
  actual   := Parser{ "# Comment" }.Status()
  expected := None
  AssertEqual(t, actual, expected)
}

func TestParserProjects(t *testing.T) {
  actual   := Parser{ "- foo +bar +baz" }.Projects()
  expected := []string{ "bar", "baz" }
  AssertEqual(t, actual, expected)
}

func TestParserTags(t *testing.T) {
  actual   := Parser{ "- foo bar:baz bam:bum" }.Tags()
  expected := map[string]string{ "bar": "baz", "bam": "bum" }
  AssertEqual(t, actual, expected)
}

func TestParserTagsEmpty(t *testing.T) {
  actual   := Parser{ "- foo" }.Tags()
  expected := map[string]string{}
  AssertEqual(t, actual, expected)
}

func TestParserTagsDate(t *testing.T) {
  actual   := Parser{ "- foo done:2015-12-13" }.Tags()
  expected := map[string]string{ "done": "2015-12-13" }
  AssertEqual(t, actual, expected)
}

func TestParserTagsNamedDate(t *testing.T) {
  actual   := Parser{ "- foo done:today" }.Tags()
  expected := map[string]string{ "done": time.Now().Format("2006-01-02") }
  AssertEqual(t, actual, expected)
}

func TestParserTagsDoNotMatchUrls(t *testing.T) {
  actual   := Parser{ "- foo http://host.com/path foo:bar [1]" }.Tags()
  expected := map[string]string{ "foo": "bar" }
  AssertEqual(t, actual, expected)
}
