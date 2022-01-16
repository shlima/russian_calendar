package russian_calendar

import "time"

func Date(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func DaysOfYear(year int) []time.Time {
	out := []time.Time{Date(year, 1, 1)}

	for {
		last := out[len(out)-1]
		next := New(last).Next().Time()
		if next.Year() > year {
			break
		}

		out = append(out, next)
	}

	return out
}

func Diff(uno, dos []time.Time) (out []time.Time) {
	ix := make(map[time.Time]bool)

	for _, el := range uno {
		ix[el] = true
	}

	for _, el := range dos {
		if _, found := ix[el]; !found {
			out = append(out, el)
		}
	}

	return
}

func YearTimetable(year int) (holidays []time.Time, workdays []time.Time, err error) {
	for _, day := range DaysOfYear(year) {
		calendar := New(day)
		isHoliday, err := calendar.IsHoliday()
		if err != nil {
			break
		}

		isWorkday, err := calendar.IsWorkday()
		if err != nil {
			break
		}

		if isHoliday {
			holidays = append(holidays, day)
		}

		if isWorkday {
			workdays = append(workdays, day)
		}
	}

	return
}
