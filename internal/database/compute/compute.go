package compute

import (
	"context"
	"go.uber.org/zap"
)

type parser interface {
	ParseInput(string) (*[]string, error)
}

type analyzer interface {
	AnalyzeInput([]string) (string, []string, error)
}

type Compute struct {
	parser   parser
	analyzer analyzer
	logger   *zap.Logger
}

func NewCompute(parser parser, analyzer analyzer, logger *zap.Logger) (*Compute, error) {

	return &Compute{
		parser:   parser,
		analyzer: analyzer,
		logger:   logger,
	}, nil
}

func (d *Compute) HandleQuery(ctx context.Context, input string) (string, []string, error) {
	tokens, err := d.parser.ParseInput(input)
	if err != nil {
		return "", nil, err
	}

	command, args, err := d.analyzer.AnalyzeInput(*tokens)
	if err != nil {
		return "", nil, err
	}

	return command, args, nil
}
