package compute

import (
	"context"
	"go.uber.org/zap"
)

type Parser struct {
	logger *zap.Logger
}

func NewParser(logger *zap.Logger) (*Parser, error) {
	return &Parser{
		logger: logger,
	}, nil
}

func (p *Parser) ParseQuery(ctx context.Context, query string) ([]string, error) {
	return []string{}, nil
}
