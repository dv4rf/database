package compute

import (
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

func (p *Parser) ParseInput(input string) (*[]string, error) {
	p.logger.Info(fmt.Sprintf("input for parse is: %s ", input))

	st := newStateMachine()
	tokens, err := st.parse(input)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
