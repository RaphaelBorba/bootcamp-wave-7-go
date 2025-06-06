package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnSalaryAfterTax_WithTestify(t *testing.T) {
	cases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "salário ≤ 50000 (sem imposto)",
			input:    50000,
			expected: 50000,
		},
		{
			name:     "salário > 50000 e ≤ 150000 (13% de imposto)",
			input:    50001,
			expected: 50001 * 0.87,
		},
		{
			name:     "salário exatamente 150000 (13% de imposto)",
			input:    150000,
			expected: 150000 * 0.87,
		},
		{
			name:     "salário > 150000 (27% de imposto)",
			input:    150001,
			expected: 150001 * 0.73,
		},
		{
			name:     "salário bem acima de 150000",
			input:    200000,
			expected: 200000 * 0.73,
		},
	}

	const delta = 0.0001
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := math.Round(returnSalaryAfterTax(tc.input)*100) / 100
			assert.Equal(t, tc.expected, got,
				"returnSalaryAfterTax(%.2f) deveria ser %.5f, obteve %.5f",
			)
		})
	}
}
