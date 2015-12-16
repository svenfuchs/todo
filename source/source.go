package source

// var stdin *os.File

type Source interface {
  ReadLines() ([]string, error)
  WriteLines([]string) error
  MustReadLines() ([]string)
  MustWriteLines([]string)
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

func mustReadLines(s Source) []string {
  lines, err := s.ReadLines()
  check(err)
  return lines
}

func mustWriteLines(s Source, lines []string) {
  err := s.WriteLines(lines)
  check(err)
}
