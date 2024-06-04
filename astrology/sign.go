package astrology

import "time"

// Sign denotes the sign of the zodiac in Western astrology.
//
// 星座
type Sign int

const (
	// Aries 白羊座
	Aries Sign = iota
	// Taurus 金牛座
	Taurus
	// Gemini 双子座
	Gemini
	// Cancer 巨蟹座
	Cancer
	// Leo 狮子座
	Leo
	// Virgo 处女座
	Virgo
	// Libra 天秤座
	Libra
	// Scorpio 天蝎座
	Scorpio
	// Sagittarius 射手座
	Sagittarius
	// Capricorn 摩羯座
	Capricorn
	// Aquarius 宝瓶/水瓶座
	Aquarius
	// Pisces 双鱼座
	Pisces
)

var (
	_ChineseNames = [...]string{
		"白羊",
		"金牛",
		"双子",
		"巨蟹",
		"狮子",
		"处女",
		"天秤",
		"天蝎",
		"射手",
		"摩羯",
		"水瓶",
		"双鱼",
	}
	names = [...]string{
		"Aries",
		"Taurus",
		"Gemini",
		"Cancer",
		"Leo",
		"Virgo",
		"Libra",
		"Scorpio",
		"Sagittarius",
		"Capricorn",
		"Aquarius",
		"Pisces",
	}
	signs = [...]string{
		"♈︎",
		"♉︎",
		"♊︎",
		"♋︎",
		"♌︎",
		"♍︎",
		"♎︎",
		"♏︎",
		"♐︎",
		"♑︎",
		"♒︎",
		"♓︎",
	}
	emojis = [...]string{
		"♈️",
		"♉️",
		"♊️",
		"♋️",
		"♌️",
		"♍️",
		"♎️",
		"♏️",
		"♐️",
		"♑️",
		"♒️",
		"♓️",
	}
	days = []*monthDay{
		{3, 21},
		{4, 20},
		{5, 21},
		{6, 22},
		{7, 23},
		{8, 23},
		{9, 23},
		{10, 24},
		{11, 23},
		{12, 22},
		{1, 20},
		{2, 19},
		{3, 21}, // 下一年
	}
)

type monthDay struct {
	month time.Month
	day   int
}

// Get returns the sign of the zodiac in Western astrology
//
// 返回指定时间的星座
func Get(t time.Time) Sign {
	_, month, day := t.Date()
	var s *monthDay
	for i, d := range days {
		if s == nil {
			s = d
			continue
		}
		if month == s.month && day >= s.day ||
			month == d.month && day < d.day {
			return Sign(i - 1)
		}
	}
	return Aries
}

// Emoji returns the emoji of the sign
//
// 返回星座的表情符号
func (s Sign) Emoji() string {
	return emojis[s]
}

// Symbol returns the symbol of the zodiac in Western astrology
//
// 返回星座的符号
func (s Sign) Symbol() string {
	return signs[s]
}

// String returns the name of the zodiac in Western astrology
//
// 返回星座的名称
func (s Sign) String() string {
	return names[s]
}

// ChineseName returns the name of the zodiac in Chinese
//
// 返回星座的中文名称
func (s Sign) ChineseName() string {
	return _ChineseNames[s]
}
