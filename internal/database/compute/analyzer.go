package compute

import (
	"errors"
	"fmt"
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

var commands = map[string]string{
	"GET": "GET",
	"SET": "SET",
	"DEL": "DEL",
}

func (a *Analyzer) AnalyzeInput(tokens []string) (string, []string, error) {
	a.logger.Info(fmt.Sprintf("tokens for analyze are: %s ", tokens))

	command, args := tokens[0], tokens[1:]
	if _, ok := commands[command]; !ok {
		return "", nil, errors.New("invalid command")
	}
	switch command {
	case commands["GET"]:
		if len(args) != 1 {
			return "", nil, errors.New("invalid len args for GET")
		}

	case commands["SET"]:
		if len(args) != 2 {
			return "", nil, errors.New("invalid len args for SET")
		}
	case commands["DEL"]:
		if len(args) != 1 {
			return "", nil, errors.New("invalid len args for DEL")
		}
	}

	return command, args, nil
}
