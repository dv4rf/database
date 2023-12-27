package compute

import (
	"context"
	"fmt"
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

func (p *Parser) ParseInput(ctx context.Context, input string) ([]string, error) {
	p.logger.Info(fmt.Sprintf("input for parse is: %s ", input))

	st := newStateMachine()
	tokens, err := st.parse(input)
	if err != nil {
		return []string{}, err
	}
	return tokens, nil
}
