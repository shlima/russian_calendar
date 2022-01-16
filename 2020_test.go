package russian_calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_2020(t *testing.T) {
	expectedHolidays := []time.Time{
		Date(2020, 1, 1),
		Date(2020, 1, 2),
		Date(2020, 1, 3),
		Date(2020, 1, 4),
		Date(2020, 1, 5),
		Date(2020, 1, 6),
		Date(2020, 1, 7),
		Date(2020, 1, 8),
		Date(2020, 1, 11),
		Date(2020, 1, 12),
		Date(2020, 1, 18),
		Date(2020, 1, 19),
		Date(2020, 1, 25),
		Date(2020, 1, 26),

		Date(2020, 2, 1),
		Date(2020, 2, 2),
		Date(2020, 2, 8),
		Date(2020, 2, 9),
		Date(2020, 2, 15),
		Date(2020, 2, 16),
		Date(2020, 2, 22),
		Date(2020, 2, 23),
		Date(2020, 2, 24),
		Date(2020, 2, 29),

		Date(2020, 3, 1),
		Date(2020, 3, 7),
		Date(2020, 3, 8),
		Date(2020, 3, 9),
		Date(2020, 3, 14),
		Date(2020, 3, 15),
		Date(2020, 3, 21),
		Date(2020, 3, 22),
		Date(2020, 3, 28),
		Date(2020, 3, 29),

		Date(2020, 4, 04),
		Date(2020, 4, 05),
		Date(2020, 4, 11),
		Date(2020, 4, 12),
		Date(2020, 4, 18),
		Date(2020, 4, 19),
		Date(2020, 4, 25),
		Date(2020, 4, 26),

		Date(2020, 5, 1),
		Date(2020, 5, 2),
		Date(2020, 5, 3),
		Date(2020, 5, 4),
		Date(2020, 5, 5),
		Date(2020, 5, 9),
		Date(2020, 5, 10),
		Date(2020, 5, 11),
		Date(2020, 5, 16),
		Date(2020, 5, 17),
		Date(2020, 5, 23),
		Date(2020, 5, 24),
		Date(2020, 5, 30),
		Date(2020, 5, 31),

		Date(2020, 6, 06),
		Date(2020, 6, 07),
		Date(2020, 6, 12),
		Date(2020, 6, 13),
		Date(2020, 6, 14),
		Date(2020, 6, 20),
		Date(2020, 6, 21),
		Date(2020, 6, 27),
		Date(2020, 6, 28),

		Date(2020, 7, 04),
		Date(2020, 7, 05),
		Date(2020, 7, 11),
		Date(2020, 7, 12),
		Date(2020, 7, 18),
		Date(2020, 7, 19),
		Date(2020, 7, 25),
		Date(2020, 7, 26),

		Date(2020, 8, 1),
		Date(2020, 8, 2),
		Date(2020, 8, 8),
		Date(2020, 8, 9),
		Date(2020, 8, 15),
		Date(2020, 8, 16),
		Date(2020, 8, 22),
		Date(2020, 8, 23),
		Date(2020, 8, 29),
		Date(2020, 8, 30),

		Date(2020, 9, 05),
		Date(2020, 9, 06),
		Date(2020, 9, 12),
		Date(2020, 9, 13),
		Date(2020, 9, 19),
		Date(2020, 9, 20),
		Date(2020, 9, 26),
		Date(2020, 9, 27),

		Date(2020, 10, 3),
		Date(2020, 10, 4),
		Date(2020, 10, 10),
		Date(2020, 10, 11),
		Date(2020, 10, 17),
		Date(2020, 10, 18),
		Date(2020, 10, 24),
		Date(2020, 10, 25),
		Date(2020, 10, 31),

		Date(2020, 11, 1),
		Date(2020, 11, 4),
		Date(2020, 11, 7),
		Date(2020, 11, 8),
		Date(2020, 11, 14),
		Date(2020, 11, 15),
		Date(2020, 11, 21),
		Date(2020, 11, 22),
		Date(2020, 11, 28),
		Date(2020, 11, 29),

		Date(2020, 12, 05),
		Date(2020, 12, 06),
		Date(2020, 12, 12),
		Date(2020, 12, 13),
		Date(2020, 12, 19),
		Date(2020, 12, 20),
		Date(2020, 12, 26),
		Date(2020, 12, 27),
	}

	t.Run("it works", func(t *testing.T) {
		holidays, workdays, err := YearTimetable(2020)

		require.NoError(t, err)
		require.ElementsMatch(t, holidays, expectedHolidays)
		require.ElementsMatch(t, workdays, Diff(holidays, DaysOfYear(2020)))
	})
}
