package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperation(t *testing.T) {
	tests := []struct {
		name      string
		op        string
		inputs    []int
		expected  int
		expectErr bool
	}{
		{
			name:      "minimum with values",
			op:        "minimum",
			inputs:    []int{5, 2, 9, -1, 3},
			expected:  -1,
			expectErr: false,
		},
		{
			name:      "minimum with empty slice",
			op:        "minimum",
			inputs:    []int{},
			expected:  0,
			expectErr: false,
		},
		{
			name:      "average with values",
			op:        "average",
			inputs:    []int{2, 3, 4, 5},
			expected:  3,
			expectErr: false,
		},
		{
			name:      "average with empty slice",
			op:        "average",
			inputs:    []int{},
			expected:  0,
			expectErr: false,
		},
		{
			name:      "maximum with values",
			op:        "maximum",
			inputs:    []int{1, 7, 3, 7, 0},
			expected:  7,
			expectErr: false,
		},
		{
			name:      "maximum with empty slice",
			op:        "maximum",
			inputs:    []int{},
			expected:  0,
			expectErr: false,
		},
		{
			name:      "invalid operation returns error",
			op:        "foo",
			inputs:    []int{1, 2, 3},
			expected:  0,
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fn, err := operation(tc.op)
			if tc.expectErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			result := fn(tc.inputs...)
			require.Equal(t, tc.expected, result)
		})
	}
}
