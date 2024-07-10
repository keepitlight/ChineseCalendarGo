package time

import (
	"time"
)

// Unix returns the Time corresponding to the given Unix timestamp.
//
// 计算时间戳对应的农历时间
func Unix(t, nsec int64) *Time {
	v := time.Unix(t, nsec)
	return Convert(v)
}

// Convert makes a Time corresponding to the given time.Time, which is between 1900/1/31 and 2100/12/1.
//
// 转换公历时间为农历时间，仅支持公历 1900/1/31~2100/12/1 之间的农历时间，不支持的时间返回 nil
func Convert(t time.Time) *Time {
	var (
		i, offset, leap         int
		daysOfYear, daysOfMonth int
		isLeap                  bool
	)
	// 与 1900-01-31 相差多少天
	y, m, d := t.Date()
	t1 := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	t2 := time.Date(1900, 1, 31, 0, 0, 0, 0, t.Location())
	v := t1.Unix() - t2.Unix()
	if v < 0 || y > 2100 {
		// 超出范围
		return nil
	}
	offset = int(v / 86400)

	for i = 1900; i < 2101 && offset > 0; i++ {
		daysOfYear = DaysOfYear(i)
		offset -= daysOfYear
	}
	if offset < 0 {
		offset += daysOfYear
		i--
	}

	// 闰月
	leap = LeapMonthOf(i)
	// 农历年
	year := i

	isLeap = false
	// 用当年的天数 offset, 逐个减去每月(农历)的天数, 求出当天是本月的第几天
	for i = 1; i < 13 && offset > 0; i++ {
		// 闰月
		if leap > 0 && i == (leap+1) && !isLeap {
			i--
			isLeap = true
			// 农历闰月天数
			daysOfMonth = DaysOfMonth
		} else {
			// 计算农历普通月天数
			daysOfMonth = Days(year, i)
		}
		// 解除闰月
		if true == isLeap && i == (leap+1) {
			isLeap = false
		}
		offset -= daysOfMonth
	}
	// offset 为 0 时, 并且刚才计算的月份是闰月, 要校正
	if 0 == offset && leap > 0 && i == leap+1 {
		if isLeap {
			isLeap = false
		} else {
			isLeap = true
			i--
		}
	}
	if offset < 0 {
		offset += daysOfMonth
		i--
	}
	day := offset + 1
	hour, _, _ := t.Clock()
	hour -= 1
	hour /= 2
	return &Time{
		year:  Year(year),
		month: Month(i),
		day:   Day(day),
		t:     t,
	}
}

// Date returns a Time of the Chinese calendar corresponding to the given parameters, nil if invalid.
//
// 创建农历日期时间，参数年月日为农历年月日及时辰，仅支持农历一九〇〇年正月初一至二一〇〇年十二月初一，无效时间返回 nil
func Date(year Year, month Month, day Day, hour, minute, second, nanosecond int, loc *time.Location) *Time {
	// 参数区间 1900.1.31 ~ 2100.12.1
	if !year.Valid() || year < YearStart || year > YearEnd ||
		!month.Valid() || !day.Valid() {
		return nil
	}
	if c := ClockOf(hour, minute, second, nanosecond); c == nil {
		return nil
	}

	var days int
	lm := LeapMonthOf(int(year))
	if month == LeapMonth {
		if lm < 1 {
			// 传参要求计算该闰月公历 但该年无闰月
			return nil
		}
		month = Month(lm).LeapMonth()
		days = DaysOfMonth // 闰月固定天数
	} else if month.IsLeapMonth() {
		if lm < 1 || lm != month.Value() {
			// 传参要求计算该闰月公历 但该年得出的闰月与传参的月份并不同
			return nil
		}
		days = DaysOfMonth // 闰月固定天数
	} else {
		days = Days(int(year), int(month))
	}
	// 日期最大值校验
	if int(day) > days {
		return nil
	}

	// 累计法计算
	// STEP 1 计算到指定农历年份合计的天数
	days = 0
	for i := YearStart; i < int(year); i++ {
		days += DaysOfYear(i)
	}
	// STEP 2 计算到指定农历月份合计的天数
	for i := 1; i <= month.Value(); i++ {
		days += DaysOfLittleMonth // 农历小月的天数
		if lm > 0 && lm == i {
			// 处理闰月
			days += DaysOfMonth
		}
	}
	// 1900 年农历正月初一的公历时间为 1900 年 1 月 30 日 0 时 0 分 0 秒 (该时间也是本农历的最开始起始点)
	// epoch := time.Time(1900, 1, 30, 0, 0, 0, 0, time.Local).Unix()
	if loc == nil {
		loc = time.Local
	}
	t := time.Unix(int64((days+int(day))*86400+EPOCH+hour*3600+minute*60+second), int64(nanosecond)).In(loc)
	return &Time{
		t:     t,
		year:  year,
		month: month,
		day:   day,
	}
}

func Now() *Time {
	return Convert(time.Now())
}
