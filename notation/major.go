package notation

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
	Jia  Major = 1 + iota // Jia 1st 甲
	Yi                    // Yi 2nd 乙
	Bing                  // Bing 3rd 丙
	Ding                  // Ding 4th 丁
	Wu                    // Wu 5th 戊（去声）
	Ji                    // Ji 6th 己
	Geng                  // Geng 7th 庚
	Xin                   // Xin 8th 辛
	Ren                   // Ren 9th 壬
	Gui                   // Gui 10th 癸

	MajorInvalid Major = 0 // 无效值
	MajorMod           = 10
)

// Majors is Chinese name of 10 celestial stem
//
// 全部的天干名称
var Majors = [...]string{
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

// Name returns the name of the celestial stem
//
// 返回天干名称(甲乙丙丁...)，无效天干返回空字符串
func (m Major) Name() string {
	if !m.Valid() {
		return ""
	}
	return Majors[(m-1)%MajorMod]
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
