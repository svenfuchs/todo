package todo

type Runnable interface {
  Run()
}

type Cmd struct {
  In Source
  Out Source
  Filter Filter
  Format string
}

func (c Cmd) list() List {
  lines := c.In.MustReadLines()
  list  := ParseItemList(lines)
  return list
}

func (c Cmd) output(list List) {
  lines := c.formatted(list.Items)
  c.Out.WriteLines(lines)
}

func (c Cmd) formatted(items []Item) []string {
  return NewFormat(c.Format).Apply(items)
}


func NewListCmd(path string, filter Filter, format string) ListCmd {
  in  := NewFileSource(path)
  out := NewFileSource("")
  return ListCmd{ Cmd { in, out, filter, format } }
}

type ListCmd struct {
  Cmd
}

func (c ListCmd) Run() {
  list := c.list()
  list  = list.Select(c.Filter)
  // list  = list.SortBy(func (item Item) string { return item.DoneDate() }) TODO
  c.output(list)
}

func NewToggleCmd(path string, filter Filter) ToggleCmd {
  in  := NewFileSource(path)
  out := NewFileSource(path)
  // out := NewFileSource("")
  return ToggleCmd{ Cmd { in, out, filter, "full" } }
}

type ToggleCmd struct {
  Cmd
}

func (c ToggleCmd) Run() {
  c.Format = "full"
  list := c.list()
  list = list.Toggle(c.Filter)
  c.output(list)
}
