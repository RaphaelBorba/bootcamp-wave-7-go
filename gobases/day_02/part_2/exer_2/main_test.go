package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReturnAvg(t *testing.T) {

	grades := []int{5, 4, 3, 6, 7, 8}
	avgExpected := 5.5

	result := average(grades)

	require.Equal(t, avgExpected, result)

}
