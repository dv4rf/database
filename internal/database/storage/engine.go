package storage

import (
	"context"
)

type Storage struct {
	m map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		m: make(map[string]string),
	}
}

func (s *Storage) Set(ctx context.Context, key, value string) error {
	s.m[key] = value
	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, bool) {
	value, ok := s.m[key]
	return value, ok
}

func (s *Storage) Del(ctx context.Context, key string) error {
	delete(s.m, key)
	return nil
}
