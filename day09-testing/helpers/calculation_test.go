package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	result := Sum(10, 20)
	
	require.Equal(t, 30, result, "Result should be 30")

	fmt.Println("Test sum passed")
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		request int
		expected int
		errMsg string
	}{
		{
			request: Sum(10, 20),
			expected: 30,
			errMsg: "Result should be 30",
		},
		{
			request: Sum(20, 20),
			expected: 40,
			errMsg: "Result should be 40",
		},
		{
			request: Sum(25, 15),
			expected: 40,
			errMsg: "Result should be 40",
		},
	}
	
	for idx, test := range tests {
		t.Run(fmt.Sprintf("test-%d", idx), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}