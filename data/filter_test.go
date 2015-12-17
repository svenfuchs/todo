package data

import (
  "testing"
  . "github.com/svenfuchs/todo/test"
)

var (
  lines = []string{ "# Comment", "- foo +bar due:2015-12-13 [1]", "x baz done:2015-12-13 [2]" }
  items = ParseItemList(lines).Items
)

func Test_Filter_Empty(t *testing.T) {
  filter := Filter{}
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ById_Succeeds(t *testing.T) {
  filter := Filter{ Ids: []int{ 1 } }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ById_Comment(t *testing.T) {
  filter := Filter{ Ids: []int{ 1 } }
  AssertFalse(t, filter.Apply(items[0]))
}

func Test_Filter_ById_Fails(t *testing.T) {
  filter := Filter{ Ids: []int{ 2 } }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ById_TextIgnored_Succeeds(t *testing.T) {
  filter := Filter{ Ids: []int{ 1 }, Text: "ignored" }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByText_Succeeds(t *testing.T) {
  filter := Filter{ Text: "fo" }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByText_Fails(t *testing.T) {
  filter := Filter{ Text: "unknown" }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByText_Comment(t *testing.T) {
  filter := Filter{ Text: "Comment" }
  AssertFalse(t, filter.Apply(items[0]))
}

func Test_Filter_ByProject_Succeeds(t *testing.T) {
  filter := Filter{ Projects: []string{ "bar", "bam" } }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByProject_Fails(t *testing.T) {
  filter := Filter{ Projects: []string{ "missing", "unknown" } }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByStatus_Pend_Succeeds(t *testing.T) {
  filter := Filter{ Status: Pend }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByStatus_Pend_Fails(t *testing.T) {
  filter := Filter{ Status: Pend }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByStatus_Done_Succeeds(t *testing.T) {
  filter := Filter{ Status: Pend }
  AssertTrue(t, filter.Apply(items[1]))
}

func Test_Filter_ByStatus_Done_Fails(t *testing.T) {
  filter := Filter{ Status: Done }
  AssertTrue(t, filter.Apply(items[2]))
}

// ------------------------------------------------------------------------------------

func Test_Filter_ByDate_ModeDate_DoneDatePresent_Succeeds(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-13", "date" } }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeDate_DoneDatePresent_Fails(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-12", "date" } }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeDate_DoneDateMissing(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-13", "date" } }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByDate_ModeBefore_DoneDatePresent_Succeeds(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2016-01-01", "before" } }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeBefore_DoneDatePresent_Fails(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-01", "before" } }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeBefore_DoneDateMissing(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2016-01-01", "before" } }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByDate_ModeSince_EqualDoneDatePresent_Succeeds(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-13", "since" } }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeSince_GreaterDoneDatePresent_Succeeds(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-01", "since" } }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeSince_DoneDatePresent_Fails(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2016-01-01", "since" } }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeSince_NoDoneDatePresent_Fails(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-01", "since" } }
  AssertFalse(t, filter.Apply(items[1]))
}

func Test_Filter_ByDate_ModeAfter_DoneDatePresent_Succeeds(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-01", "after" } }
  AssertTrue(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeAfter_DoneDatePresent_Fails(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2016-01-01", "after" } }
  AssertFalse(t, filter.Apply(items[2]))
}

func Test_Filter_ByDate_ModeAfter_NoDoneDatePresent_Fails(t *testing.T) {
  filter := Filter{ Date: FilterDate { "2015-12-01", "after" } }
  AssertFalse(t, filter.Apply(items[1]))
}
