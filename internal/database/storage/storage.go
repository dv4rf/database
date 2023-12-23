package storage

import (
	"go.uber.org/zap"
)

type Storage struct {
	logger *zap.Logger
}

func NewStorage(logger *zap.Logger) (*Storage, error) {
	return &Storage{
		logger: logger,
	}, nil
}
