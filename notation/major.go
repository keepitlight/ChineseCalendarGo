package notation

import "github.com/keepitlight/ChineseCalendarGo/util"

// Major celestial stem, representing the primary symbols in the
// traditional calendar/cultural symbol system, with the names of the 10 Heavenly Stems
// expressed in pinyin
//
// 天干，主符号，值范围 1~10，无效用 0 指代
type Major byte

// CelestialStem is alias of Major
//
// Major 别名，天干，主符号，值范围 1~10，无效用 0 指代
type CelestialStem = Major

const (
	Jia  Major = 1 + iota // Jia 1st, Chinese: 甲, Pinyin: jiǎ
	Yi                    // Yi 2nd, Chinese: 乙, Pinyin: yǐ
	Bing                  // Bing 3rd, Chinese: 丙, Pinyin: bǐng
	Ding                  // Ding 4th, Chinese: 丁, Pinyin: dīng
	Wu                    // Wu 5th, Chinese: 戊（去声）, Pinyin: wù
	Ji                    // Ji 6th, Chinese: 己, Pinyin: jǐ
	Geng                  // Geng 7th, Chinese: 庚, Pinyin: gēng
	Xin                   // Xin 8th, Chinese: 辛, Pinyin: xīn
	Ren                   // Ren 9th, Chinese: 壬, Pinyin: rén
	Gui                   // Gui 10th, Chinese: 癸, Pinyin: guǐ

	MajorInvalid Major = 0  // 无效值
	MajorCycle         = 10 // 天干周期

	甲 = Jia  // Jia 1st, Chinese: 甲, Pinyin: jiǎ
	乙 = Yi   // Yi 2nd, Chinese: 乙, Pinyin: yǐ
	丙 = Bing // Bing 3rd, Chinese: 丙, Pinyin: bǐng
	丁 = Ding // Ding 4th, Chinese: 丁, Pinyin: dīng
	戊 = Wu   // Wu 5th, Chinese: 戊（去声）, Pinyin: wù
	己 = Ji   // Ji 6th, Chinese: 己, Pinyin: jǐ
	庚 = Geng // Geng 7th, Chinese: 庚, Pinyin: gēng
	辛 = Xin  // Xin 8th, Chinese: 辛, Pinyin: xīn
	壬 = Ren  // Ren 9th, Chinese: 壬, Pinyin: rén
	癸 = Gui  // Gui 10th, Chinese: 癸, Pinyin: guǐ
)

var majors = [...]string{
	"甲",
	"乙",
	"丙",
	"丁",
	"戊",
	"己",
	"庚",
	"辛",
	"壬",
	"癸",
}

// Majors is Chinese name of 10 celestial stem
//
// 全部的天干名称
func Majors() [10]string {
	return majors
}

// Name returns the name of the celestial stem
//
// 返回天干名称(甲乙丙丁...)，无效天干返回空字符串
func (m Major) Name() string {
	if !m.Valid() {
		return ""
	}
	return majors[(m-1)%MajorCycle]
}

func (m Major) String() string {
	return m.Name()
}

// Value returns the value of the celestial stem
//
// 返回天干序数值(1~10)，无效天干返回 0
func (m Major) Value() byte {
	if !m.Valid() {
		return 0
	}
	return byte(m)
}

// Valid returns whether the celestial stem is valid
func (m Major) Valid() bool {
	return Jia <= m && m <= Gui
}

func (m Major) YinYang() YinYang {
	return YinYang(m % 2)
}

func (m Major) Wuxing() Wuxing {
	return Wuxing((m + 1) / 2)
}

// MajorOf returns the major(celestial stem) of the given value v, notation.MajorInvalid if 0.
//
// 返回任意序数返回对应的天干，0 无效，返回 0
func MajorOf(v int) Major {
	return Major(util.Cycle(v, MajorCycle))
}

// Hour returns the major(celestial stem) of the given hour, day, notation.MajorInvalid if 0.
//
// 根据五鼠遁元计算日干某天某小时的天干，小时或日干无效时返回 notation.MajorInvalid
func Hour(hour int, day Major) Major {
	if !day.Valid() || hour < 0 || hour > 23 {
		return MajorInvalid
	}
	// 甲己还加甲，乙庚丙作初
	// 丙辛从戊起，丁壬庚子居
	// 戊癸起壬子，周而复始求
	switch day {
	case 甲, 己:
		day = 甲
	case 乙, 庚:
		day = 丙
	case 丙, 辛:
		day = 戊
	case 丁, 壬:
		day = 庚
	case 戊, 癸:
		day = 壬
	default:
		return MajorInvalid
	}
	return Major(util.Cycle(int(day)+hour, MajorCycle))
}
