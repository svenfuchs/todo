package cmd

func NewArgs() *Args {
  return &Args{ Config: map[string]string{} }
}

type Args struct {
  Path string
  Line string
  Ids []int
  Status string
  Text string
  Projects []string
  Date string
  Format string
  Config map[string]string
  Report bool
}

func (a *Args) SetConfig(args ...string) {
  a.Config[args[0]] = args[1]
}

func (a *Args) SetDate(args ...string) {
  a.Date = args[0] + ":" + args[1]
}

var cmdFactories = map[string]func(*Args) Runnable {
  "archive": NewArchiveCmd,
  "list":    NewListCmd,
  "push":    NewPushCmd,
  "toggle":  NewToggleCmd,
}

func (a *Args) Run(args ...string) {
  cmdFactories[args[0]](a).Run()
}
