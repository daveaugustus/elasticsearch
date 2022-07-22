package random

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomNoNToN(t *testing.T) {
	tests := []struct {
		name string
		max  int
		min  int
	}{
		{
			name: "Test 1",
			max:  90,
			min:  13,
		},
		{
			name: "Test 2",
			max:  90,
			min:  13,
		},
		{
			name: "Test 3",
			max:  90,
			min:  13,
		},
		{
			name: "Test 4",
			max:  90,
			min:  13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomNoNToN(tt.min, tt.max)
			log.Println("Got: ", got)
			require.NotNil(t, got)
			require.GreaterOrEqual(t, got, 13)
			require.LessOrEqual(t, got, 90)
		})
	}
}
