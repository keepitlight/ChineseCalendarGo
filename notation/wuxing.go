package notation

import "github.com/keepitlight/ChineseCalendarGo/util"

// Wuxing Chinese: 五行; pinyin: wǔxíng, usually translated as Five Phases or Five Agents,
// is a fivefold conceptual scheme used in many traditional Chinese fields of study to explain
// a wide array of phenomena, including cosmic cycles, the interactions between internal organs,
// the succession of political regimes, and the properties of herbal medicines.
type Wuxing int

// Phases is alias of Wuxing
type Phases = Wuxing

const (
	Wood  Wuxing = 1 + iota // Wood 木
	Fire                    // Fire 火
	Earth                   // Earth 土
	Metal                   // Metal 金
	Water                   // Water 水

	WuxingInvalid = 0
	WuxingCycle   = 5

	木 = Wood
	火 = Fire
	土 = Earth
	金 = Metal
	水 = Water
)

func (w Wuxing) Valid() bool {
	return w > 0 && w < 6
}

// IsInterPromoting checks whether the current wuxing w is inter-promoting with the given wuxing x.
//
// 检测 w 与 x 是否相生，即 w 生 x，即 x 是五行循环中 w 的下一个元素。
func (w Wuxing) IsInterPromoting(x Wuxing) bool {
	if !w.Valid() || !x.Valid() {
		return false
	}
	return int(x) == util.Cycle(int(w+1), WuxingCycle)
}

// IsInterRegulating checks whether the current wuxing w is inter-regulating with the given wuxing x.
//
// 检测 w 与 x 是否相克，即 w 克 x
func (w Wuxing) IsInterRegulating(x Wuxing) bool {
	if !w.Valid() || !x.Valid() {
		return false
	}
	return int(x) == util.Cycle(int(w+2), WuxingCycle)
}
