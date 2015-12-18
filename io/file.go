package io

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func NewFileIo(path string) Io {
	if path == "" {
		path = "./todo.txt"
	}
	return FileIo{path}
}

type FileIo struct {
	path string
}

func (s FileIo) ReadLines() []string {
	file := s.openFile()
	lines := []string{}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func (s FileIo) WriteLines(lines []string) {
	s.writeLines(lines, os.O_TRUNC)
}

func (s FileIo) AppendLines(lines []string) {
	s.writeLines(lines, os.O_APPEND)
}

func (s FileIo) openFile() *os.File {
	file, err := os.Open(s.path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (s FileIo) writeLines(lines []string, mode int) {
	if len(lines) == 0 {
		return
	}
	file, err := os.OpenFile(s.path, os.O_WRONLY|os.O_CREATE|mode, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Join(lines, "\n") + "\n"
	file.Write([]byte(str))
}
