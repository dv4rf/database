package compute

import (
	"context"
	"go.uber.org/zap"
)

type Query struct {
	commandID int
	arguments []string
}

func NewQuery(commandID int, arguments []string) Query {
	return Query{
		commandID: commandID,
		arguments: arguments,
	}
}

func (c *Query) CommandID() int {
	return c.commandID
}

func (c *Query) Arguments() []string {
	return c.arguments
}

type parser interface {
	ParseQuery(context.Context, string) ([]string, error)
}

type analyzer interface {
	AnalyzeQuery(context.Context, []string) (Query, error)
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

func (d *Compute) HandleQuery(ctx context.Context, queryStr string) (Query, error) {
	tokens, err := d.parser.ParseQuery(ctx, queryStr)
	if err != nil {
		return Query{}, err
	}

	query, err := d.analyzer.AnalyzeQuery(ctx, tokens)
	if err != nil {
		return Query{}, err
	}

	return query, nil
}
