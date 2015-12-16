package service

type Service interface {
  Push(line string) error
  Fetch() ([]string, error)
}
