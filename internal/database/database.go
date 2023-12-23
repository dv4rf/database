package database

import (
	"context"
	"go.uber.org/zap"
)

type computeLayer interface {
}

type storageLayer interface {
}

type Database struct {
	computeLayer computeLayer
	storageLayer storageLayer
	logger       *zap.Logger
}

func NewDatabase(computeLayer computeLayer, storageLayer storageLayer, logger *zap.Logger) (*Database, error) {
	return &Database{
		computeLayer: computeLayer,
		storageLayer: storageLayer,
		logger:       logger,
	}, nil
}

func (d *Database) HandleQuery(ctx context.Context, queryStr string) string {
	return queryStr
}
