package random

import (
	"log"
	"strings"
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

func TestRandomNoZeroToN(t *testing.T) {
	number := RandomNoZeroToN(2)
	log.Println(number)
	require.GreaterOrEqual(t, number, 0)
	require.LessOrEqual(t, number, 2)
}

func TestNewName(t *testing.T) {
	fullName := NewName()
	require.NotEmpty(t, fullName)
	require.LessOrEqual(t, len(fullName), 15)
	require.GreaterOrEqual(t, len(fullName), 7)

	fullNameSlice := strings.Split(fullName, " ")
	require.Equal(t, 2, len(fullNameSlice))
	require.GreaterOrEqual(t, len(fullNameSlice[0]), 3)
	require.LessOrEqual(t, len(fullNameSlice[0]), 6)
	require.GreaterOrEqual(t, len(fullNameSlice[1]), 4)
	require.LessOrEqual(t, len(fullNameSlice[1]), 9)
}

func TestEmailDomains(t *testing.T) {
	domains := EmailDomains()
	require.NotNil(t, domains)
	require.NotEmpty(t, domains)
	require.Equal(t, 8, len(domains))
}

func TestNewEmail(t *testing.T) {
	fullName := NewName()
	require.NotEmpty(t, fullName)
	require.LessOrEqual(t, len(fullName), 15)
	require.GreaterOrEqual(t, len(fullName), 7)

	fullNameSlice := strings.Split(fullName, " ")
	require.Equal(t, 2, len(fullNameSlice))
	require.GreaterOrEqual(t, len(fullNameSlice[0]), 3)
	require.LessOrEqual(t, len(fullNameSlice[0]), 6)
	require.GreaterOrEqual(t, len(fullNameSlice[1]), 4)

	// GetEmail passing the full name of the users
	email := NewEmail(fullName)
	require.NotEmpty(t, email)
	require.Contains(t, email, "@")
	require.Contains(t, email, ".com")
	require.Contains(t, email, "_")
}
