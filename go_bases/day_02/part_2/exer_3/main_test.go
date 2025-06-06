package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcSalary(t *testing.T) {
	tests := []struct {
		name        string
		workMinutes int
		category    string
		expected    float64
	}{
		{
			name:        "Category C, 60 minutes → 1 hour * 1000",
			workMinutes: 60,
			category:    "C",
			expected:    1000.0,
		},
		{
			name:        "Category B, 120 minutes → 2 hours * 1500 + 20%",
			workMinutes: 120,
			category:    "B",
			expected:    3600.0,
		},
		{
			name:        "Category A lowercase, 30 minutes → 0.5h * 3000 + 50%",
			workMinutes: 30,
			category:    "a",
			expected:    2250.0,
		},
		{
			name:        "Invalid category returns 0",
			workMinutes: 100,
			category:    "X",
			expected:    0.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := calcSalary(tc.workMinutes, tc.category)
			require.Equal(t, tc.expected, actual)
		})
	}
}
