package service

func NewService(config map[string]string) Service {
	var service Service
	switch config["service"] {
	case "idonethis":
		service = NewIdonethis(config)
	}
	return service
}

type Service interface {
	ReadLines() []string
	WriteLines([]string)
}
