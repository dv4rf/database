package storage

import (
	"context"
	"go.uber.org/zap"
)

type Storage struct {
	logger *zap.Logger
	m      map[string]string
}

func NewStorage(logger *zap.Logger) (*Storage, error) {
	return &Storage{
		logger: logger,
		m:      make(map[string]string),
	}, nil
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
