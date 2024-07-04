package time

// Year represents the year of the traditional Chinese calendar, Zero is invalid,
// negative of year means BC
//
// 表示农历年份，农历没有连续的计年传统，借用公历年份表示农历年份，公元前纪年使用负数。
type Year int

const (
	YearInvalid = 0 // Invalid year, no zero 无效的年份，没有 0 年
)

// Days of year
//
// 表示该年农历的总天数
func (y Year) Days() int {
	if !y.Valid() {
		return 0
	}
	year := int(y)
	var (
		i, sum int
	)
	sum = 29 * 12
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (calendars[year-YearStart] & i) != 0 {
			sum++
		}
	}
	return sum + DaysOfMonth
}

func (y Year) Valid() bool {
	return y != YearInvalid && y >= YearStart && y <= YearEnd
}

// Month returns the month of the year.
//
// 返回该年的农历月份（1,2,3...12），例如 2018 年 1 月返回 “正月”，无效的月份返回 nil
func (y Year) Month(month Month) *YearMonth {
	if !y.Valid() {
		return nil
	}
	var m, d int
	if month == LeapMonth {
		// 闰月
		if m = calendars[y-YearStart] & yearMask; m == 0 {
			// 无闰月
			return nil
		} else {
			month = Month(m)
		}
		d = DaysOfMonth // 闰月固定 30 天
	} else if !month.Valid() {
		return nil
	} else if month.IsLeapMonth() {
		// 闰月
		if m = calendars[y-YearStart] & yearMask; m == 0 || m != month.Value() {
			// 无闰月
			return nil
		}
		d = DaysOfMonth // 闰月固定 30 天
	} else {
		d = Days(int(y), int(month))
		if d == 0 {
			return nil
		}
	}
	return &YearMonth{
		month: month,
		days:  d,
	}
}
