package russian_calendar

import (
	"time"
)

type ICalendar interface {
	Next() ICalendar
	Prev() ICalendar
	IsHoliday() (bool, error)
	IsWorkday() (bool, error)
	GteHoliday() (ICalendar, error)
	LteHoliday() (ICalendar, error)
	GtHoliday() (ICalendar, error)
	LtHoliday() (ICalendar, error)
	GteWorkday() (ICalendar, error)
	LteWorkday() (ICalendar, error)
	GtWorkday() (ICalendar, error)
	LtWorkday() (ICalendar, error)
	Time() time.Time
}
