require 'date'
require 'csv'

# Звездочкой (*) отмечены предпраздничные (сокращенные) дни.
# Плюсом (+) отмечены перенесенные выходные дни
CALENDAR_PATH = ARGV[0]
SHORTENED_POSTFIX = '*'
WEEKEND_POSTPONED_POSTFIX = '+'
DAY_SEP = ','
YEAR_COL = 'Год/Месяц'
MONTHS_MAP = {
  1 => 'Январь',
  2 => 'Февраль',
  3 => 'Март',
  4 => 'Апрель',
  5 => 'Май',
  6 => 'Июнь',
  7 => 'Июль',
  8 => 'Август',
  9 => 'Сентябрь',
  10 => 'Октябрь',
  11 => 'Ноябрь',
  12 => 'Декабрь'
}.freeze

csv = CSV.open(CALENDAR_PATH, headers: true).map(&:to_h)

index = csv.each_with_object({}) do |row, object|
  year = Integer(row.fetch(YEAR_COL))

  object[year] = MONTHS_MAP.each_with_object({}) do |(ix, name), obj|
    obj[ix] = row.fetch(name).split(DAY_SEP).each_with_object(Hash.new) do |day, result|
      result[Integer(day.delete(WEEKEND_POSTPONED_POSTFIX))] = true unless day[SHORTENED_POSTFIX]
    end
  end
end

def hash2go(hash)
  out = StringIO.new
  out << ?{

  hash.each do |key, value|
    out << "#{key}: \n"

    if value.is_a?(TrueClass)
        out << "true, \n"
    else
        out << hash2go(value)
        out << ?,
    end
  end

  out << ?}
  out.string
end

puts <<~CODE
package russian_calendar

// CODE GENERATED AT #{Time.now.utc}. DO NOT EDIT.

func init() {
    Source = SourceMap#{hash2go(index)}
}
CODE
