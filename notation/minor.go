package notation

import "github.com/keepitlight/ChineseCalendarGo/util"

// Minor  terrestrial branch, representing the secondary symbols
// in the traditional calendar/cultural symbol system, with the names of the 12 Earthly
// Branches expressed in pinyin
//
// 地支，从符号，值范围 1~12，其它值无效，用 0 指代
type Minor byte

// TerrestrialBranch is the alias of Minor
//
// Minor 别名，地支，从符号，值范围 1~12，其它值无效，用 0 指代
type TerrestrialBranch = Minor

const (
	Zi   Minor = 1 + iota // Zi 1st, Chinese: 子, Pinyin: zǐ
	Chou                  // Chou 2nd, Chinese: 丑, Pinyin: chǒu
	Yin                   // Yin 3rd, Chinese: 寅, Pinyin: yín
	Mao                   // Mao 4th, Chinese: 卯, Pinyin: mǎo
	Chen                  // Chen 5th, Chinese: 辰, Pinyin: chén
	Si                    // Si 6th, Chinese: 巳, Pinyin: sì
	WU                    // Wu 7th, Chinese: 午（上声）, Pinyin: wǔ
	Wei                   // Wei 8th, Chinese: 未, Pinyin: wèi
	Shen                  // Shen 9th, Chinese: 申, Pinyin: shēn
	You                   // You 10th, Chinese: 酉, Pinyin: yǒu
	Xu                    // Xu 11th, Chinese: 戌, Pinyin: xū
	Hai                   // Hai 12th, Chinese: 亥, Pinyin: hài

	MinorInvalid Minor = 0  // 无效地支值
	MinorCycle         = 12 // 地支的周期

	子 = Zi   // Zi 1st, Chinese: 子, Pinyin: zǐ
	丑 = Chou // Chou 2nd, Chinese: 丑, Pinyin: chǒu
	寅 = Yin  // Yin 3rd, Chinese: 寅, Pinyin: yín
	卯 = Mao  // Mao 4th, Chinese: 卯, Pinyin: mǎo
	辰 = Chen // Chen 5th, Chinese: 辰, Pinyin: chén
	巳 = Si   // Si 6th, Chinese: 巳, Pinyin: sì
	午 = WU   // Wu 7th, Chinese: 午（上声）, Pinyin: wǔ
	未 = Wei  // Wei 8th, Chinese: 未, Pinyin: wèi
	申 = Shen // Shen 9th, Chinese: 申, Pinyin: shēn
	酉 = You  // You 10th, Chinese: 酉, Pinyin: yǒu
	戌 = Xu   // Xu 11th, Chinese: 戌, Pinyin: xū
	亥 = Hai  // Hai 12th, Chinese: 亥, Pinyin: hài
)

// Minors is the Chinese names of the 12 terrestrial branch
//
// 全部的地支名称
func Minors() [12]string {
	return minors
}

var minors = [...]string{
	"子",
	"丑",
	"寅",
	"卯",
	"辰",
	"巳",
	"午",
	"未",
	"申",
	"酉",
	"戌",
	"亥",
}

// Name returns the name of the minor
//
// 返回地支名称(子丑寅卯...)，无效地支返回空字符串
func (m Minor) Name() string {
	if !m.Valid() {
		return ""
	}
	return minors[(m-1)%MinorCycle]
}

// Sign returns the sign of the minor
//
// 返回地支对应的生肖(鼠牛虎兔龙蛇马羊猴鸡狗猪...)，无效地支返回 Invalid
func (m Minor) Sign() Sign {
	if !m.Valid() {
		return SignInvalid
	}
	return Sign(m)
}

func (m Minor) String() string {
	return m.Name()
}

// Value returns the value of the minor
//
// 返回地支序数(1234...)，无效地支返回 0
func (m Minor) Value() byte {
	if !m.Valid() {
		return 0
	}
	return byte(m)
}

func (m Minor) Valid() bool {
	return Zi <= m && m <= Hai
}

func (m Minor) YinYang() YinYang {
	return YinYang(m % 2)
}

func (m Minor) Wuxing() Wuxing {
	switch m {
	case Zi, Hai:
		return Water // 子，亥 > 水
	case Chou, Chen, Wei, Xu:
		return Earth // 丑，辰，未，戌 > 土
	case Yin, Mao:
		return Wood // 寅，卯 > 木
	case Si, WU:
		return Fire // 巳，午 > 火
	case Shen, You:
		return Metal // 酉，申 > 金
	default:
		return WuxingInvalid
	}
}

// MinorOf returns the minor(terrestrial branch) of the given value v, notation.MinorInvalid if 0.
//
// 根据任意序数返回地支，0 为无效值，返回 notation.MinorInvalid
func MinorOf(v int) Minor {
	return Minor(util.Cycle(v, MinorCycle))
}
