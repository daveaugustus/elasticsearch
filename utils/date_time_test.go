package utils_test

import (
	"elasticsearch/utils"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPossibleDates(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
		past  bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test 7 days past",
			args: args{
				year:  0,
				month: 0,
				day:   7,
				past:  true,
			},
			want: 8,
		},
		{
			name: "Test 1 year past",
			args: args{
				year:  1,
				month: 0,
				day:   0,
				past:  true,
			},
			want: 366,
		},

		{
			name: "Test 7 days future",
			args: args{
				year:  0,
				month: 0,
				day:   7,
				past:  false,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetPossibleDates(tt.args.year, tt.args.month, tt.args.day, tt.args.past); !reflect.DeepEqual(got, tt.want) {
				require.Equal(t, len(got), tt.want)
				require.NotEmpty(t, got)
				require.NotNil(t, got)
			}
		})
	}
}
