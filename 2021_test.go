package russian_calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_2021(t *testing.T) {
	expectedHolidays := []time.Time{
		Date(2021, 1, 1),
		Date(2021, 1, 2),
		Date(2021, 1, 3),
		Date(2021, 1, 4),
		Date(2021, 1, 5),
		Date(2021, 1, 6),
		Date(2021, 1, 7),
		Date(2021, 1, 8),
		Date(2021, 1, 9),
		Date(2021, 1, 10),
		Date(2021, 1, 16),
		Date(2021, 1, 17),
		Date(2021, 1, 23),
		Date(2021, 1, 24),
		Date(2021, 1, 30),
		Date(2021, 1, 31),

		Date(2021, 2, 6),
		Date(2021, 2, 7),
		Date(2021, 2, 13),
		Date(2021, 2, 14),
		Date(2021, 2, 21),
		Date(2021, 2, 22),
		Date(2021, 2, 23),
		Date(2021, 2, 27),
		Date(2021, 2, 28),

		Date(2021, 3, 6),
		Date(2021, 3, 7),
		Date(2021, 3, 8),
		Date(2021, 3, 13),
		Date(2021, 3, 14),
		Date(2021, 3, 20),
		Date(2021, 3, 21),
		Date(2021, 3, 27),
		Date(2021, 3, 28),

		Date(2021, 4, 3),
		Date(2021, 4, 4),
		Date(2021, 4, 10),
		Date(2021, 4, 11),
		Date(2021, 4, 17),
		Date(2021, 4, 18),
		Date(2021, 4, 24),
		Date(2021, 4, 25),

		Date(2021, 5, 1),
		Date(2021, 5, 2),
		Date(2021, 5, 3),
		Date(2021, 5, 8),
		Date(2021, 5, 9),
		Date(2021, 5, 10),
		Date(2021, 5, 15),
		Date(2021, 5, 16),
		Date(2021, 5, 22),
		Date(2021, 5, 23),
		Date(2021, 5, 29),
		Date(2021, 5, 30),

		Date(2021, 6, 5),
		Date(2021, 6, 6),
		Date(2021, 6, 12),
		Date(2021, 6, 13),
		Date(2021, 6, 14),
		Date(2021, 6, 19),
		Date(2021, 6, 20),
		Date(2021, 6, 26),
		Date(2021, 6, 27),

		Date(2021, 7, 3),
		Date(2021, 7, 4),
		Date(2021, 7, 10),
		Date(2021, 7, 11),
		Date(2021, 7, 17),
		Date(2021, 7, 18),
		Date(2021, 7, 24),
		Date(2021, 7, 25),
		Date(2021, 7, 31),

		Date(2021, 8, 1),
		Date(2021, 8, 7),
		Date(2021, 8, 8),
		Date(2021, 8, 14),
		Date(2021, 8, 15),
		Date(2021, 8, 21),
		Date(2021, 8, 22),
		Date(2021, 8, 28),
		Date(2021, 8, 29),

		Date(2021, 9, 4),
		Date(2021, 9, 5),
		Date(2021, 9, 11),
		Date(2021, 9, 12),
		Date(2021, 9, 18),
		Date(2021, 9, 19),
		Date(2021, 9, 25),
		Date(2021, 9, 26),

		Date(2021, 10, 2),
		Date(2021, 10, 3),
		Date(2021, 10, 9),
		Date(2021, 10, 10),
		Date(2021, 10, 16),
		Date(2021, 10, 17),
		Date(2021, 10, 23),
		Date(2021, 10, 24),
		Date(2021, 10, 30),
		Date(2021, 10, 31),

		Date(2021, 11, 4),
		Date(2021, 11, 5),
		Date(2021, 11, 6),
		Date(2021, 11, 7),
		Date(2021, 11, 13),
		Date(2021, 11, 14),
		Date(2021, 11, 20),
		Date(2021, 11, 21),
		Date(2021, 11, 27),
		Date(2021, 11, 28),

		Date(2021, 12, 4),
		Date(2021, 12, 5),
		Date(2021, 12, 11),
		Date(2021, 12, 12),
		Date(2021, 12, 18),
		Date(2021, 12, 19),
		Date(2021, 12, 25),
		Date(2021, 12, 26),
		Date(2021, 12, 31),
	}

	t.Run("it works", func(t *testing.T) {
		holidays, workdays, err := YearTimetable(2021)

		require.NoError(t, err)
		require.ElementsMatch(t, holidays, expectedHolidays)
		require.ElementsMatch(t, workdays, Diff(holidays, DaysOfYear(2021)))
	})
}
