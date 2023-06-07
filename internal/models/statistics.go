package models

type Statistics struct {
	Single int64 `json:"single"`
	Min    int64 `json:"min"`
	Max    int64 `json:"max"`
}

func NewStatistics() *Statistics {
	return &Statistics{}
}

func (s *Statistics) IncSingle() {
	s.Single++
}

func (s *Statistics) IncMin() {
	s.Min++
}

func (s *Statistics) IncMax() {
	s.Max++
}
