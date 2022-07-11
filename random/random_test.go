package random

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAge(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test 1",
		},
		{
			name: "Test 2",
		},
		{
			name: "Test 3",
		},
		{
			name: "Test 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAge()
			log.Println("Got: ", got)
			require.NotNil(t, got)
			require.GreaterOrEqual(t, got, 13)
			require.LessOrEqual(t, got, 90)
		})
	}
}
