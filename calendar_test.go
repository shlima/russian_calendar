package russian_calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCalendar_Implementation(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		require.Implements(t, (*ICalendar)(nil), Calendar{})
	})
}

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

func TestCalendar_IsHoliday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  bool
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it is true",
			input: Date(2000, 1, 5),
			want:  true,
		}, {
			desc:  "it is false",
			input: Date(2000, 1, 6),
			want:  false,
		}, {
			desc:  "it is error",
			input: Date(2000, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.IsHoliday()

			require.Equal(t, out, tt.want)

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
			}
		})
	}
}

func TestCalendar_IsWorkday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  bool
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it is true",
			input: Date(2000, 1, 4),
			want:  true,
		}, {
			desc:  "it is false",
			input: Date(2000, 1, 5),
			want:  false,
		}, {
			desc:  "it is error",
			input: Date(2000, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.IsWorkday()

			require.Equal(t, out, tt.want)

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
			}
		})
	}
}

func TestCalendar_GteHoliday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it works for the same date",
			input: Date(2000, 1, 5),
			want:  Date(2000, 1, 5),
		}, {
			desc:  "it works for the forthcoming date",
			input: Date(2000, 1, 1),
			want:  Date(2000, 1, 5),
		}, {
			desc:  "it is error",
			input: Date(2000, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.GteHoliday()

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, out.Time(), tt.want)
		})
	}
}

func TestCalendar_GtHoliday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5:  true,
				10: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it works for the same date",
			input: Date(2000, 1, 5),
			want:  Date(2000, 1, 10),
		}, {
			desc:  "it works for the forthcoming date",
			input: Date(2000, 1, 6),
			want:  Date(2000, 1, 10),
		}, {
			desc:  "it is error",
			input: Date(2000, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.GtHoliday()

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, out.Time(), tt.want)
		})
	}
}

func TestCalendar_LteHoliday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it works for the same date",
			input: Date(2000, 1, 5),
			want:  Date(2000, 1, 5),
		}, {
			desc:  "it works for the forthcoming date",
			input: Date(2000, 1, 10),
			want:  Date(2000, 1, 5),
		}, {
			desc:  "it is error",
			input: Date(1999, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.LteHoliday()

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, out.Time(), tt.want)
		})
	}
}

func TestCalendar_LtHoliday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5:  true,
				10: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it works for the same date",
			input: Date(2000, 1, 10),
			want:  Date(2000, 1, 5),
		}, {
			desc:  "it works for the forthcoming date",
			input: Date(2000, 1, 8),
			want:  Date(2000, 1, 5),
		}, {
			desc:  "it is error",
			input: Date(1999, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.LtHoliday()

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, out.Time(), tt.want)
		})
	}
}

func TestCalendar_GteWorkday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5:  true,
				10: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it works for the same date",
			input: Date(2000, 1, 6),
			want:  Date(2000, 1, 6),
		}, {
			desc:  "it works for the forthcoming date",
			input: Date(2000, 1, 5),
			want:  Date(2000, 1, 6),
		}, {
			desc:  "it is error",
			input: Date(2000, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.GteWorkday()

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, out.Time(), tt.want)
		})
	}
}

func TestCalendar_GtWorkday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it works for the same date",
			input: Date(2000, 1, 5),
			want:  Date(2000, 1, 6),
		}, {
			desc:  "it works for the forthcoming date",
			input: Date(2000, 1, 4),
			want:  Date(2000, 1, 6),
		}, {
			desc:  "it is error",
			input: Date(2000, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.GtWorkday()

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, out.Time(), tt.want)
		})
	}
}

func TestCalendar_LteWorkday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5:  true,
				10: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it works for the same date",
			input: Date(2000, 1, 6),
			want:  Date(2000, 1, 6),
		}, {
			desc:  "it works for the forthcoming date",
			input: Date(2000, 1, 10),
			want:  Date(2000, 1, 9),
		}, {
			desc:  "it is error",
			input: Date(2000, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.LteWorkday()

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, out.Time(), tt.want)
		})
	}
}

func TestCalendar_LtWorkday(t *testing.T) {
	type args struct {
		desc  string
		input time.Time
		want  time.Time
		err   error
	}

	source := &SourceMap{
		2000: {
			1: {
				5:  true,
				10: true,
			},
		},
	}

	tests := []args{
		{
			desc:  "it works for the same date",
			input: Date(2000, 1, 3),
			want:  Date(2000, 1, 2),
		}, {
			desc:  "it works for the forthcoming date",
			input: Date(2000, 1, 10),
			want:  Date(2000, 1, 9),
		}, {
			desc:  "it is error",
			input: Date(1999, 2, 1),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			calendar := NewSourced(tt.input, source)
			out, err := calendar.LtWorkday()

			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, out.Time(), tt.want)
		})
	}
}
