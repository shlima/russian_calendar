package russian_calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_2022(t *testing.T) {
	expectedHolidays := []time.Time{
		Date(2022, 1, 1),
		Date(2022, 1, 2),
		Date(2022, 1, 3),
		Date(2022, 1, 4),
		Date(2022, 1, 5),
		Date(2022, 1, 6),
		Date(2022, 1, 7),
		Date(2022, 1, 8),
		Date(2022, 1, 9),
		Date(2022, 1, 15),
		Date(2022, 1, 16),
		Date(2022, 1, 22),
		Date(2022, 1, 23),
		Date(2022, 1, 29),
		Date(2022, 1, 30),

		Date(2022, 2, 5),
		Date(2022, 2, 6),
		Date(2022, 2, 12),
		Date(2022, 2, 13),
		Date(2022, 2, 19),
		Date(2022, 2, 20),
		Date(2022, 2, 23),
		Date(2022, 2, 26),
		Date(2022, 2, 27),

		Date(2022, 3, 6),
		Date(2022, 3, 7),
		Date(2022, 3, 8),
		Date(2022, 3, 12),
		Date(2022, 3, 13),
		Date(2022, 3, 19),
		Date(2022, 3, 20),
		Date(2022, 3, 26),
		Date(2022, 3, 27),

		Date(2022, 4, 2),
		Date(2022, 4, 3),
		Date(2022, 4, 9),
		Date(2022, 4, 10),
		Date(2022, 4, 16),
		Date(2022, 4, 17),
		Date(2022, 4, 23),
		Date(2022, 4, 24),
		Date(2022, 4, 30),

		Date(2022, 5, 1),
		Date(2022, 5, 2),
		Date(2022, 5, 3),
		Date(2022, 5, 7),
		Date(2022, 5, 8),
		Date(2022, 5, 9),
		Date(2022, 5, 10),
		Date(2022, 5, 14),
		Date(2022, 5, 15),
		Date(2022, 5, 21),
		Date(2022, 5, 22),
		Date(2022, 5, 28),
		Date(2022, 5, 29),

		Date(2022, 6, 4),
		Date(2022, 6, 5),
		Date(2022, 6, 11),
		Date(2022, 6, 12),
		Date(2022, 6, 13),
		Date(2022, 6, 18),
		Date(2022, 6, 19),
		Date(2022, 6, 25),
		Date(2022, 6, 26),

		Date(2022, 7, 2),
		Date(2022, 7, 3),
		Date(2022, 7, 9),
		Date(2022, 7, 10),
		Date(2022, 7, 16),
		Date(2022, 7, 17),
		Date(2022, 7, 23),
		Date(2022, 7, 24),
		Date(2022, 7, 30),
		Date(2022, 7, 31),

		Date(2022, 8, 6),
		Date(2022, 8, 7),
		Date(2022, 8, 13),
		Date(2022, 8, 14),
		Date(2022, 8, 20),
		Date(2022, 8, 21),
		Date(2022, 8, 27),
		Date(2022, 8, 28),

		Date(2022, 9, 3),
		Date(2022, 9, 4),
		Date(2022, 9, 10),
		Date(2022, 9, 11),
		Date(2022, 9, 17),
		Date(2022, 9, 18),
		Date(2022, 9, 24),
		Date(2022, 9, 25),

		Date(2022, 10, 1),
		Date(2022, 10, 2),
		Date(2022, 10, 8),
		Date(2022, 10, 9),
		Date(2022, 10, 15),
		Date(2022, 10, 16),
		Date(2022, 10, 22),
		Date(2022, 10, 23),
		Date(2022, 10, 29),
		Date(2022, 10, 30),

		Date(2022, 11, 4),
		Date(2022, 11, 5),
		Date(2022, 11, 6),
		Date(2022, 11, 12),
		Date(2022, 11, 13),
		Date(2022, 11, 19),
		Date(2022, 11, 20),
		Date(2022, 11, 26),
		Date(2022, 11, 27),

		Date(2022, 12, 3),
		Date(2022, 12, 4),
		Date(2022, 12, 10),
		Date(2022, 12, 11),
		Date(2022, 12, 17),
		Date(2022, 12, 18),
		Date(2022, 12, 24),
		Date(2022, 12, 25),
		Date(2022, 12, 31),
	}

	t.Run("it works", func(t *testing.T) {
		holidays, workdays, err := YearTimetable(2022)

		require.NoError(t, err)
		require.ElementsMatch(t, holidays, expectedHolidays)
		require.ElementsMatch(t, workdays, Diff(holidays, DaysOfYear(2022)))
	})
}
