package russian_calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSourceMap_Implementation(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		require.Implements(t, (*ISource)(nil), SourceMap{})
	})
}

func TestSourceMap_Exists(t *testing.T) {
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
			desc:  "day matched",
			input: Date(2000, 1, 5),
			want:  true,
		}, {
			desc:  "month matched",
			input: Date(2000, 1, 1),
			want:  false,
		}, {
			desc:  "month mismatched",
			input: Date(2000, 2, 1),
			err:   ErrOutOfRange,
		}, {
			desc:  "year mismatched",
			input: Date(1999, 1, 5),
			err:   ErrOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := source.Exists(tt.input)

			require.Equal(t, got, tt.want)

			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			}
		})
	}
}
