package storage

import (
	"bufio"
	"os"
	"strings"
)

type Storage struct {
	Domains []string
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) ParseDomainsPromFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	var dataSlice []string

	for scanner.Scan() {
		dataSlice = append(dataSlice, strings.TrimSpace(scanner.Text()))
	}
	s.Domains = dataSlice
	return nil
}
