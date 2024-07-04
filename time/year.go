package time

import "github.com/keepitlight/ChineseCalendarGo/util"

// Year represents the year of the traditional Chinese calendar, Zero is invalid,
// negative of year means BC
//
// 表示农历年份，农历没有连续的计年传统，借用公历年份表示农历年份，公元前纪年使用负数。
type Year int

const (
	YearInvalid = 0 // Invalid year, no zero 无效的年份，没有 0 年
)

// Valid returns whether the year is valid.
//
// 年份是否有效
func (y Year) Valid() bool {
	return y != YearInvalid
}

const (
	NAME = "农历"
	AD   = "公元"
	BC   = "前"
	YEAR = "年"
)

// String returns the string representation of the year.
//
// 返回年份的字符串表示，如 "二〇二四年"，"公元前二二〇年"
func (y Year) String() string {
	if y == YearInvalid {
		return ""
	}
	p := AD
	return p + y.Name() + YEAR
}

// Name returns the name of the year.
//
// 返回年份的名称，如 "二〇二四"，"前二二〇"
func (y Year) Name() string {
	if y == YearInvalid {
		return ""
	}
	var p string
	v := y
	if v < 0 {
		p = BC
		v = 0 - v
	}
	return p + util.Digits(uint64(v))
}
