package notation

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
	Zi   Minor = 1 + iota // Zi 1st 子
	Chou                  // Chou 2nd 丑
	Yin                   // Yin 3rd 寅
	Mao                   // Mao 4th 卯
	Chen                  // Chen 5th 辰
	Si                    // Si 6th 巳
	WU                    // Wu 7th 午（上声）
	Wei                   // Wei 8th 未
	Shen                  // Shen 9th 申
	You                   // You 10th 酉
	Xu                    // Xu 11th 戌
	Hai                   // Hai 12th 亥

	MinorInvalid Minor = 0  // 无效地支值
	MinorCycle         = 12 // 地支的周期
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
