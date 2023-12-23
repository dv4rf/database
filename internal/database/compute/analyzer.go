package compute

import (
	"context"
	"go.uber.org/zap"
)

type Analyzer struct {
	logger *zap.Logger
}

func NewAnalyzer(logger *zap.Logger) (*Analyzer, error) {
	analyser := &Analyzer{
		logger: logger,
	}
	return analyser, nil
}

func (a *Analyzer) AnalyzeQuery(ctx context.Context, tokens []string) (Query, error) {
	return Query{}, nil
}
