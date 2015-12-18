package io

import (
	"strings"
)

func NewMemoryIo(content string) Io {
	return MemoryIo{&content}
}

type MemoryIo struct {
	content *string
}

func (s MemoryIo) ReadLines() []string {
	return strings.Split(*s.content, "\n")
}

func (s MemoryIo) WriteLines(lines []string) {
	*s.content = strings.Join(lines, "\n")
}

func (s MemoryIo) AppendLines(lines []string) {
	*s.content = *s.content + strings.Join(lines, "\n")
}
