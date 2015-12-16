package source

// var stdin *os.File

type Source interface {
  ReadLines() ([]string, error)
  WriteLines([]string) error
  AppendLines([]string) error
}

func NewSource(path string) Source {
  if path == "" {
    return StdioSource{}
  } else {
    return FileSource{ path: path }
  }
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}
