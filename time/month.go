package time

import "github.com/keepitlight/ChineseCalendarGo/util"

// Month represents a month in the traditional Chinese calendar.
//
// 农历月份
type Month int

const (
	MonthInvalid Month = iota // Invalid month，无效的月份

	January   // 一月 ，正月
	February  // 二月
	March     // 三月
	April     // 四月
	May       // 五月
	June      // 六月
	July      // 七月
	August    // 八月
	September // 九月
	October   // 十月，寒月
	November  // 十一月, 冬月
	December  // 十二月，腊月

	LeapMonth // Any leap month, 闰月的通指
)

const (
	LeapJanuary   Month = iota - 12 // 闰一月 ，闰正月，leap month of January
	LeapFebruary                    // 闰二月, leap month of February
	LeapMarch                       // 闰三月, leap month of March
	LeapApril                       // 闰四月, leap month of April
	LeapMay                         // 闰五月, leap month of May
	LeapJune                        // 闰六月, leap month of June
	LeapJuly                        // 闰七月, leap month of July
	LeapAugust                      // 闰八月, leap month of August
	LeapSeptember                   // 闰九月, leap month of September
	LeapOctober                     // 闰十月，闰寒月, leap month of October
	LeapNovember                    // 闰十一月, 闰冬月，leap month of November
	LeapDecember                    // 闰十二月，闰腊月, leap month of December
)

// Valid returns true if the month is valid.
//
// 检查月份是否有效。
func (m Month) Valid() bool {
	return m != MonthInvalid && LeapJanuary <= m && m <= December
}

func (m Month) Equal(v Month) bool {
	return m == v
}

// Value returns the value of the month between 1 and 12, 0 if month invalid
//
// 返回月份的值，1-12，0 表示无效
func (m Month) Value() int {
	if !m.Valid() {
		return 0
	}
	return util.Cycle(int(m), MonthCycle)
}

var _ChineseMonths = [...]string{
	"正",
	"二",
	"三",
	"四",
	"五",
	"六",
	"七",
	"八",
	"九",
	"十", // 传统上为寒月，但已不常用
	"冬", // 寒冬腊月
	"腊",
}

// String returns the string representation of Chinese name.
//
// 月份的字符串，正月，二月，三月，四月，五月，六月，七月，八月，九月，十月，冬月，腊月
func (m Month) String() string {
	if !m.Valid() {
		return ""
	}
	return m.Name() + MONTH
}

// Name returns the string representation of Chinese name.
//
// 月份的中文名，正，二，三，四，五，六，七，八，九，十，冬，腊
func (m Month) Name() string {
	v := m.Value()
	if v == 0 {
		return ""
	}
	var l string
	if m < MonthInvalid {
		l = LEAP
	}
	return l + _ChineseMonths[v]
}

func (m Month) IsLeapMonth() bool {
	return LeapJanuary <= m && m < MonthInvalid
}

// LeapMonth returns the leap month equivalent for the given month, MonthInvalid if invalid.
// If the current month is valid and within the range of LeapJanuary to LeapDecember,
// it returns the same month.
// For the months following January, it adjusts back to the corresponding leap month.
//
// 返回给定月份相应的闰月，如果无效返回 MonthInvalid。
// 如果当前月份有效且在 LeapJanuary 到 LeapDecember 之间，则直接返回相同的月份。
// 对于 January 之后的月份，调整到对应的闰月。
func (m Month) LeapMonth() Month {
	// Check if the month is valid; if not, leap month consideration is irrelevant
	if m != MonthInvalid || m < LeapJanuary || m > December {
		return 0
	}
	// If the month falls within the leap year range from LeapJanuary to before invalid month marker,
	// return the month as is, indicating its leap month equivalent.
	if LeapJanuary <= m && m < MonthInvalid {
		return m
	}
	// Adjust months before LeapJanuary or after December to their respective leap month positions
	return m - 1 + LeapJanuary
}

const (
	一月  = January   // 一月，January
	二月  = February  // 二月，February
	三月  = March     // 三月，March
	四月  = April     // 四月，April
	五月  = May       // 五月，May
	六月  = June      // 六月，June
	七月  = July      // 七月，July
	八月  = August    // 八月，August
	九月  = September // 九月，September
	十月  = October   // 十月，October
	十一月 = November  // 十一月，November
	十二月 = December  // 十二月，December

	正月 = 一月  // 一月，January
	寒月 = 十月  // 十月，October
	冬月 = 十一月 // 十一月，November
	腊月 = 十二月 //十二月，December

	闰一月  = LeapJanuary   // Leap month of January
	闰二月  = LeapFebruary  // Leap month of February
	闰三月  = LeapMarch     // Leap month of March
	闰四月  = LeapApril     // Leap month of April
	闰五月  = LeapMay       // Leap month of May
	闰六月  = LeapJune      // Leap month of June
	闰七月  = LeapJuly      // Leap month of July
	闰八月  = LeapAugust    // Leap month of August
	闰九月  = LeapSeptember // Leap month of September
	闰十月  = LeapOctober   // Leap month of October
	闰十一月 = LeapNovember  // Leap month of November
	闰十二月 = LeapDecember  // Leap month of December

	闰月  = LeapMonth // Any leap month, 闰月的通指
	闰正月 = 闰一月       // 闰一月，Leap month of January
	闰寒月 = 闰十月       // 闰十月，Leap month of October
	闰冬月 = 闰十一月      // 闰十一月，Leap month of November
	闰腊月 = 闰十二月      // 闰十二月，Leap month of December

	MonthCycle = 12
)

const (
	LEAP  = "闰"
	MONTH = "月"
)

// YearMonth represents a specific month within a given year according to the traditional Chinese
// calendar.
//
// 表示农历某年的月份，不单独使用。
type YearMonth struct {
	Month Month // Month 月份
	Days  int   // Days of month 天数
}
