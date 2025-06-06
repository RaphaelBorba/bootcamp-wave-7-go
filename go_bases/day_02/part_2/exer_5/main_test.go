package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnimal(t *testing.T) {
	tests := []struct {
		name      string
		ani       string
		qty       int64
		expected  float64
		expectErr bool
	}{
		{
			name:      "dog returns 10 per unit",
			ani:       dog,
			qty:       3,
			expected:  30.0,
			expectErr: false,
		},
		{
			name:      "cat returns 5 per unit",
			ani:       cat,
			qty:       4,
			expected:  20.0,
			expectErr: false,
		},
		{
			name:      "hamster returns 0.25 per unit",
			ani:       hamster,
			qty:       8,
			expected:  2.0,
			expectErr: false,
		},
		{
			name:      "spider returns 0.15 per unit",
			ani:       spider,
			qty:       10,
			expected:  1.5,
			expectErr: false,
		},
		{
			name:      "invalid animal returns error",
			ani:       "parrot",
			qty:       5,
			expected:  0.0,
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fn, err := animal(tc.ani)
			if tc.expectErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			actual := fn(tc.qty)
			require.Equal(t, tc.expected, actual)
		})
	}
}
