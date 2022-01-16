package russian_calendar

import (
	"fmt"
	"time"
)

// SourceMap map represents year, month, day holidays
type SourceMap map[int]map[int]map[int]bool

func (i SourceMap) Exists(input time.Time) (bool, error) {
	y, m, d := input.Year(), int(input.Month()), input.Day()

	yearMap, exists := i[y]
	if !exists {
		return false, fmt.Errorf("%w: year %d", ErrOutOfRange, y)
	}

	monthMap, exists := yearMap[m]
	if !exists {
		return false, fmt.Errorf("%w: %d-%d", ErrOutOfRange, y, m)
	}

	return monthMap[d], nil
}
