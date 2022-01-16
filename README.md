## GO производственный календарь

![CI](https://github.com/shlima/russian_calendar/actions/workflows/test.yml/badge.svg)

Пакет проверяет признак праздника/рабочего дня на заданную дату,
[набор данных](https://github.com/shlima/russian_calendar/blob/master/calendar.csv)
предоставлен [Порталом открытых данных](https://data.gov.ru/opendata/7708660670-proizvcalendar)
Минэкономразвития России (включая дополнения
из [Консультант+](http://www.consultant.ru/law/ref/calendar/proizvodstvennye/))

Работает на основе [кодогенерации](https://github.com/shlima/russian_calendar/blob/master/bin/build.rb) из исходного CSV файла.

## Использование

```go
package main

import (
	"time"

	calendar "github.com/shlima/russian_calendar"
)

func main() {
	cal := calendar.New(time.Now())

	// Returns true if today is a day off or a public holiday 
	cal.IsHoliday()
	// Returns true if today is a working day
	cal.IsWorkday()
	// Returns the next business day (including the current one)
	cal.GteWorkday()
	// Returns the previous business day (including the current one)
	cal.LteWorkday()
	// Returns the next business day
	cal.GtWorkday()
	// Returns the previous business day
	cal.LtWorkday()
	// Returns the next weekend day (including the current one)
	cal.GteHoliday()
	// Returns the previous weekend day (including the current one)
	cal.LteHoliday()
	// Returns the next weekend day
	cal.GtHoliday()
	// Returns the previous weekend day
	cal.LtHoliday()

	// Returns the next day
	cal.Next()
	// Returns the previous day
	cal.Prev()
}
```

## Обновление данных

* обновить файл `calendar.csv` из data.gov.ru
* make build (сгенерирует go hash)
* git push
