package item_test

import (
  "testing"
  "github.com/svenfuchs/todo.go/item"
  . "github.com/svenfuchs/todo.go/test"
)

var (
  lines = []string{ "# Comment", "- foo +bar due:2015-12-13 [1]", "x baz done:2015-12-13 [2]" }
  items = item.ParseList(lines).Items
)

func Test_Filter_Empty(t *testing.T) {
  filter := item.Filter{}
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ById_Succeeds(t *testing.T) {
  filter := item.Filter{ Id: 1 }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ById_Comment(t *testing.T) {
  filter := item.Filter{ Id: 1 }
  AssertFalse(t, filter.Apply(items[0]))
}

func Test_Filter_ById_Fails(t *testing.T) {
  filter := item.Filter{ Id: 2 }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ById_TextIgnored_Succeeds(t *testing.T) {
  filter := item.Filter{ Id: 1, Text: "ignored" }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByText_Succeeds(t *testing.T) {
  filter := item.Filter{ Text: "fo" }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByText_Fails(t *testing.T) {
  filter := item.Filter{ Text: "unknown" }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByText_Comment(t *testing.T) {
  filter := item.Filter{ Text: "Comment" }
  AssertFalse(t, filter.Apply(items[0]))
}

func Test_Filter_ByProject_Succeeds(t *testing.T) {
  filter := item.Filter{ Projects: []string{ "bar", "bam" } }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByProject_Fails(t *testing.T) {
  filter := item.Filter{ Projects: []string{ "missing", "unknown" } }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByStatus_Pend_Succeeds(t *testing.T) {
  filter := item.Filter{ Status: item.Pend }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByStatus_Pend_Fails(t *testing.T) {
  filter := item.Filter{ Status: item.Pend }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByStatus_Done_Succeeds(t *testing.T) {
  filter := item.Filter{ Status: item.Pend }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByStatus_Done_Fails(t *testing.T) {
  filter := item.Filter{ Status: item.Done }
  AssertTrue(t, filter.Apply(items[2]))
}

// ------------------------------------------------------------------------------------

func Test_Filter_ByDate_ModeDate_DoneDatePresent_Succeeds(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-13", Mode: "date" }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeDate_DoneDatePresent_Fails(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-12", Mode: "date" }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeDate_DoneDateMissing(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-13", Mode: "date" }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByDate_ModeBefore_DoneDatePresent_Succeeds(t *testing.T) {
  filter := item.Filter{ Date: "2016-01-01", Mode: "before" }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeBefore_DoneDatePresent_Fails(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-01", Mode: "before" }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeBefore_DoneDateMissing(t *testing.T) {
  filter := item.Filter{ Date: "2016-01-01", Mode: "before" }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByDate_ModeSince_EqualDoneDatePresent_Succeeds(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-13", Mode: "since" }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeSince_GreaterDoneDatePresent_Succeeds(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-01", Mode: "since" }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeSince_DoneDatePresent_Fails(t *testing.T) {
  filter := item.Filter{ Date: "2016-01-01", Mode: "since" }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeSince_NoDoneDatePresent_Fails(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-01", Mode: "since" }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByDate_ModeAfter_DoneDatePresent_Succeeds(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-01", Mode: "after" }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeAfter_DoneDatePresent_Fails(t *testing.T) {
  filter := item.Filter{ Date: "2016-01-01", Mode: "after" }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeAfter_NoDoneDatePresent_Fails(t *testing.T) {
  filter := item.Filter{ Date: "2015-12-01", Mode: "after" }
  AssertFalse(t, filter.Apply(items[1]))
}
