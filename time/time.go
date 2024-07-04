package time

import (
	"time"
)

// Time represents a Chinese calendar time.
//
// 农历时间
type Time struct {
	t time.Time

	year  Year  // 农历年
	month Month // 农历月
	day   Day   // 农历日
}

// Supported returns whether the year is supported of the Chinese calendar, which is from BC 1900 to 2100.
//
// 指定年份的农历数据是否被支持，仅支持 1900 至 2100 年的农历数据
func (t *Time) Supported() bool {
	return t.year.Supported()
}

// Year returns the current year of the Chinese calendar.
//
// 返回当前时间的农历年份
func (t *Time) Year() Year {
	return t.year
}

// Month returns the current month of the Chinese calendar.
//
// 返回当前时间的农历月份
func (t *Time) Month() *YearMonth {
	m := t.month
	l := t.month == LeapMonth
	var d int
	if l {
		d = DaysOfMonth
		if v := LeapMonthOf(int(t.year)); v > 0 {
			m = Month(v)
		} else {
			return nil
		}
	} else {
		d = Days(int(t.year), int(t.month))
	}
	return &YearMonth{
		month: m,
		days:  d,
	}
}

// Day returns the current day of the Chinese calendar.
//
// 返回当前时间的农历日
func (t *Time) Day() Day {
	return t.day
}

func (t *Time) Date() (year Year, month Month, day Day) {
	return t.year, t.month, t.day
}

// Equal reports whether t and b represent the same time in the Chinese calendar
// without minute and second.
//
// 比较农历时间是否相等，忽略分秒
func (t *Time) Equal(b *Time) bool {
	if t == nil || b == nil {
		return t == b
	}
	if t.t.Equal(b.t) {
		return true
	}
	return t.year == b.year &&
		t.month.Equal(b.month) &&
		t.day == b.day &&
		t.t.Hour() == b.t.Hour()
}

// Time returns the time of the Gregorian calendar.
//
// 返回当前的公历时间。
func (t *Time) Time() time.Time {
	return t.t
}
