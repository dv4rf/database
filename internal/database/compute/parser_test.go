package compute

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestParser_ParseInput(t *testing.T) {
	t.Parallel()
	t.Run("positive", func(t *testing.T) {
		t.Parallel()
		logger := zap.NewNop()
		p, _ := NewParser(logger)
		input := "some test input"
		expectedTokens := &[]string{"some", "test", "input"}

		tokens, err := p.ParseInput(input)

		require.Nil(t, err)
		require.Equal(t, expectedTokens, tokens)
	})
	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			name  string
			input string
		}{
			{
				name:  "invalid symbol",
				input: "&",
			},
			{
				name:  "not ASCII",
				input: "Ж",
			},
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				logger := zap.NewNop()
				p, _ := NewParser(logger)
				input := "%%%"
				expecteErr := errors.New("invalid symbol")

				tokens, err := p.ParseInput(input)

				require.Nil(t, tokens)
				require.Error(t, expecteErr, err)
			})
		}
	})
}
