package notation

import (
	"github.com/keepitlight/ChineseCalendarGo/util"
)

// Clock represents a clock in the traditional Chinese calendar.
//
// 表示时钟/时辰
type Clock struct {
	hour       int      // 0-23
	minute     int      // 0-59
	second     int      // 0-59
	nanosecond int      // 0-999999999
	notation   Notation // 干支，时辰符号，仅精确到小时
}

// ClockOf makes a new clock.
//
// 创建一个农历时钟，无效时间返回 nil。
func ClockOf(hour, minute, second, nanosecond int, day Major) *Clock {
	if hour < 0 || hour > 23 ||
		minute < 0 || minute > 59 ||
		second < 0 || second > 59 ||
		nanosecond < 0 || nanosecond > 999999999 ||
		!day.Valid() {
		return nil
	}
	j := Hour(hour, day)
	i := MinorOfHour(hour)
	if n, p := Pair(j, i); !p {
		return nil
	} else {
		return &Clock{
			hour:       hour,
			minute:     minute,
			second:     second,
			nanosecond: nanosecond,
			notation:   n,
		}
	}
}

// Notation 返回时辰的干支，非特定/具体的时间/时辰，返回 nil
func (c *Clock) Notation() Notation {
	return c.notation
}
func (c *Clock) Clock() (hour, minute, second int) {
	return c.hour, c.minute, c.second
}
func (c *Clock) Nanosecond() int {
	return c.nanosecond
}
func (c *Clock) Equal(b *Clock) bool {
	if c == nil || b == nil {
		return c == b
	}
	return c.hour == b.hour &&
		c.minute == b.minute &&
		c.second == b.second &&
		c.nanosecond == b.nanosecond
}
func (c *Clock) Minor() Minor {
	return MinorOfHour(c.hour)
}

var clocks = [...]string{
	"零",
	"一",
	"二",
	"三",
	"四",
	"五",
	"六",
	"七",
	"八",
	"九",
	"十",
	"十一",
	"十二",
}

const (
	hourChinese   = "点"
	minuteChinese = "分"
	zeroChinese   = "零"
)

// String returns the string representation of the clock.
//
// 返回时钟的字符串表示，精确到分（分为 0 则忽略），例如上午八点，凌晨五点，晚上九点二十分
func (c *Clock) String() string {
	p := Period(c.hour)
	i := 0
	if c.hour > 0 {
		i = c.hour % 12
	}
	h := clocks[i] + hourChinese
	var m string
	if c.minute > 0 {
		m = util.Small(uint16(c.minute)) + minuteChinese
		if c.minute < 10 {
			m = zeroChinese + m
		}
	}
	return p + h + m // 上午八点，凌晨五点
}

// Periods returns the names of each time period in a day.
//
// 一天 24 个小时各时间段名字表
func Periods() [24]string {
	return periods
}

var (
	periods = [24]string{
		"凌晨",
		"凌晨",
		"凌晨",
		"凌晨",
		"凌晨",
		"凌晨",
		"早上",
		"早上",
		"上午",
		"上午",
		"上午",
		"上午",
		"中午",
		"中午",
		"下午",
		"下午",
		"下午",
		"下午",
		"下午",
		"晚上",
		"晚上",
		"晚上",
		"晚上",
		"晚上",
	}
)

// Period returns the name of each time period in a day.
//
// 返回一天中各时间段的名称，例如凌晨，早上，中午等
func Period(hour int) string {
	return periods[hour%24]
}
