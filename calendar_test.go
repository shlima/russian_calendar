package russian_calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCalendar_Next(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
	}

	tests := []args{
		{
			desc:  "it works for a middle of a month",
			input: Date(2021, 6, 15),
			want:  Date(2021, 6, 16),
		}, {
			desc:  "it works for the end of a year",
			input: Date(2021, 12, 31),
			want:  Date(2022, 1, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := New(tt.input)
			require.Equal(t, tt.want, calendar.Next().Time())
		})
	}
}

func TestCalendar_Prev(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
	}

	tests := []args{
		{
			desc:  "it works for a middle of a month",
			input: Date(2021, 6, 15),
			want:  Date(2021, 6, 14),
		}, {
			desc:  "it works for the beginning of a year",
			input: Date(2022, 1, 1),
			want:  Date(2021, 12, 31),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := New(tt.input)
			require.Equal(t, tt.want, calendar.Prev().Time())
		})
	}
}
