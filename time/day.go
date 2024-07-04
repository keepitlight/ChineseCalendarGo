package time

import (
	"github.com/keepitlight/ChineseCalendarGo/util"
)

// Day represents a day in the traditional Chinese calendar
//
// 表示农历日
type Day byte

const (
	nameTenth     = "初十"
	nameTwentieth = "二十"
	nameThirtieth = "三十"
)

var dayPrefixes = [...]string{
	"初", "十", "廿", "卅",
}

// String converts the day to a string of the Chinese calendar.
//
// 将日转换为农历日期字符串，初九日
func (d Day) String() string {
	if d == 0 || d > 30 {
		return ""
	}
	return d.Name() + DAY
}

// Name converts the day to a string of the Chinese calendar.
//
// 将日转换为农历日期字符串，初九，……
func (d Day) Name() string {
	if d == 0 || d > 30 {
		return ""
	}
	var name string
	switch d {
	case 10:
		name = nameTenth
	case 20:
		name = nameTwentieth
	case 30:
		name = nameThirtieth
	default:
		name = dayPrefixes[(d/10)] + util.Digits(uint64(d%10))
	}
	return name
}

func (d Day) Valid() bool {
	return d > 0 && d <= 30
}
