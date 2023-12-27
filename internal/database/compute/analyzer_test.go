package compute

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAnalyzer_AnalyzeInput(t *testing.T) {
	t.Parallel()
	t.Run("positive", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			name            string
			tokens          []string
			expectedLenArgs int
			expectedCmd     string
		}{
			{
				name:            "get command",
				tokens:          []string{"GET", "key"},
				expectedLenArgs: 1,
				expectedCmd:     "GET",
			},
			{
				name:            "set command",
				tokens:          []string{"SET", "key", "value"},
				expectedLenArgs: 2,
				expectedCmd:     "SET",
			},
			{
				name:            "del command",
				tokens:          []string{"DEL", "key"},
				expectedLenArgs: 1,
				expectedCmd:     "DEL",
			},
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				logger := zap.NewNop()
				a, _ := NewAnalyzer(logger)

				cmd, args, err := a.AnalyzeInput(tc.tokens)

				require.Nil(t, err)
				require.Equal(t, tc.expectedCmd, cmd)
				require.Len(t, args, tc.expectedLenArgs)
				require.Equal(t, tc.tokens[1:], args)
			})

		}

	})
	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			name        string
			tokens      []string
			expectedErr error
		}{
			{
				name:        "invalid command",
				tokens:      []string{"LET"},
				expectedErr: (errors.New("invalid command")),
			},
			{
				name:        "invalid len GET",
				tokens:      []string{"GET", "key", "value"},
				expectedErr: (errors.New("invalid len args for GET")),
			},
			{
				name:        "invalid len SET",
				tokens:      []string{"SET", "key"},
				expectedErr: (errors.New("invalid len args for SET")),
			},
			{
				name:        "invalid len DEL",
				tokens:      []string{"DEL"},
				expectedErr: (errors.New("invalid len args for DEL")),
			},
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				logger := zap.NewNop()
				a, _ := NewAnalyzer(logger)

				cmd, args, err := a.AnalyzeInput(tc.tokens)

				require.Equal(t, tc.expectedErr, err)
				require.Equal(t, "", cmd)
				require.Empty(t, args)
			})

		}

	})
}
