package russian_calendar

import "time"

type Calendar struct {
	time   time.Time
	source *SourceMap
}

func New(input time.Time) ICalendar {
	return NewSourced(input, &Source)
}

func NewSourced(input time.Time, source *SourceMap) ICalendar {
	return Calendar{
		time:   input,
		source: source,
	}
}

func (c Calendar) Time() time.Time {
	return c.time
}

func (c Calendar) Next() ICalendar {
	return NewSourced(c.time.Add(24*time.Hour), c.source)
}

func (c Calendar) Prev() ICalendar {
	return NewSourced(c.time.Add(-24*time.Hour), c.source)
}

func (c Calendar) IsHoliday() (bool, error) {
	return c.source.Exists(c.time)
}

func (c Calendar) IsWorkday() (bool, error) {
	if holiday, err := c.IsHoliday(); err != nil {
		return false, err
	} else {
		return !holiday, nil
	}
}

func (c Calendar) GteHoliday() (ICalendar, error) {
	if holiday, err := c.IsHoliday(); err != nil || holiday {
		return c, err
	}

	return c.Next().GteHoliday()
}

func (c Calendar) GtHoliday() (ICalendar, error) {
	return c.Next().GteHoliday()
}

func (c Calendar) LteHoliday() (ICalendar, error) {
	if holiday, err := c.IsHoliday(); err != nil || holiday {
		return c, err
	}

	return c.Prev().LteHoliday()
}

func (c Calendar) LtHoliday() (ICalendar, error) {
	return c.Prev().LteHoliday()
}

func (c Calendar) GteWorkday() (ICalendar, error) {
	if workday, err := c.IsWorkday(); err != nil || workday {
		return c, err
	}

	return c.Next().GteWorkday()
}

func (c Calendar) GtWorkday() (ICalendar, error) {
	return c.Next().GteWorkday()
}

func (c Calendar) LteWorkday() (ICalendar, error) {
	if workday, err := c.IsWorkday(); err != nil || workday {
		return c, err
	}

	return c.Prev().LteWorkday()
}

func (c Calendar) LtWorkday() (ICalendar, error) {
	return c.Prev().LteWorkday()
}
