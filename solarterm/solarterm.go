package solarterm

// SolarTerm is a 24-term of Chinese calendar.
// English name from Beijing 2022 Winter Olympics official translation
//
// 表示二十四节气之一。英文名来自 2022 年冬奥会官方翻译。
type SolarTerm byte

const (
	MinorCold          SolarTerm = 1 + iota // 小寒，英文名来自 2022 年冬奥会官方翻译
	MajorCold                               // 大寒，英文名来自 2022 年冬奥会官方翻译
	BeginningOfSpring                       // 立春，英文名来自 2022 年冬奥会官方翻译
	RainWater                               // 雨水，英文名来自 2022 年冬奥会官方翻译
	AwakeningOfInsects                      // 惊蛰，英文名来自 2022 年冬奥会官方翻译
	SpringEquinox                           // 春分，英文名来自 2022 年冬奥会官方翻译
	PureBrightness                          // 清明，英文名来自 2022 年冬奥会官方翻译
	GrainRain                               // 谷雨，英文名来自 2022 年冬奥会官方翻译
	BeginningOfSummer                       // 立夏，英文名来自 2022 年冬奥会官方翻译
	GrainBuds                               // 小满，英文名来自 2022 年冬奥会官方翻译
	GrainInEar                              // 芒种，英文名来自 2022 年冬奥会官方翻译
	SummerSolstice                          // 夏至，英文名来自 2022 年冬奥会官方翻译
	MinorHeat                               // 小暑，英文名来自 2022 年冬奥会官方翻译
	MajorHeat                               // 大暑，英文名来自 2022 年冬奥会官方翻译
	BeginningOfAutumn                       // 立秋，英文名来自 2022 年冬奥会官方翻译
	EndOfHeat                               // 处暑，英文名来自 2022 年冬奥会官方翻译
	WhiteDew                                // 白露，英文名来自 2022 年冬奥会官方翻译
	AutumnEquinox                           // 秋分，英文名来自 2022 年冬奥会官方翻译
	ColdDew                                 // 寒露，英文名来自 2022 年冬奥会官方翻译
	FrostsDescent                           // 霜降，英文名来自 2022 年冬奥会官方翻译
	BeginningOfWinter                       // 立冬，英文名来自 2022 年冬奥会官方翻译
	MinorSnow                               // 小雪，英文名来自 2022 年冬奥会官方翻译
	MajorSnow                               // 大雪，英文名来自 2022 年冬奥会官方翻译
	WinterSolstice                          // 冬至，英文名来自 2022 年冬奥会官方翻译
	Invalid            SolarTerm = 0        // 无效节气
)

const (
	小寒 = MinorCold
	大寒 = MajorCold
	立春 = BeginningOfSpring
	雨水 = RainWater
	惊蛰 = AwakeningOfInsects
	春分 = SpringEquinox
	清明 = PureBrightness
	谷雨 = GrainRain
	立夏 = BeginningOfSummer
	小满 = GrainBuds
	芒种 = GrainInEar
	夏至 = SummerSolstice
	小暑 = MinorHeat
	大暑 = MajorHeat
	立秋 = BeginningOfAutumn
	处暑 = EndOfHeat
	白露 = WhiteDew
	秋分 = AutumnEquinox
	寒露 = ColdDew
	霜降 = FrostsDescent
	立冬 = BeginningOfWinter
	小雪 = MinorSnow
	大雪 = MajorSnow
	冬至 = WinterSolstice

	ModerateCold    = MinorCold          // 别名，小寒，英文名来自香港天文台
	SevereCold      = MajorCold          // 别名，大寒，英文名来自香港天文台
	SpringCommences = BeginningOfSpring  // 别名，立春，英文名来自香港天文台
	SpringShowers   = RainWater          // 别名，雨水，英文名来自香港天文台
	InsectsWaken    = AwakeningOfInsects // 别名，惊蛰，英文名来自香港天文台
	VernalEquinox   = SpringEquinox      // 别名，春分，英文名来自香港天文台
	BrightAndClear  = PureBrightness     // 别名，清明，英文名来自香港天文台
	CornRain        = GrainRain          // 别名，谷雨，英文名来自香港天文台
	SummerCommences = BeginningOfSummer  // 别名，立夏，英文名来自香港天文台
	CornForms       = GrainBuds          // 别名，小满，英文名来自香港天文台
	CornOnEar       = GrainInEar         // 别名，芒种，英文名来自香港天文台
	ModerateHeat    = MinorHeat          // 别名，小暑，英文名来自香港天文台
	GreatHeat       = MajorHeat          // 别名，大暑，英文名来自香港天文台
	AutumnCommences = BeginningOfAutumn  // 别名，立秋，英文名来自香港天文台
	AutumnalEquinox = AutumnEquinox      // 别名，秋分，英文名来自香港天文台
	Frost           = FrostsDescent      // 别名，霜降，英文名来自香港天文台
	WinterCommences = BeginningOfWinter  // 别名，立冬，英文名来自香港天文台
	LightSnow       = MinorSnow          // 别名，小雪，英文名来自香港天文台
	HeavySnow       = MajorSnow          // 别名，大雪，英文名来自香港天文台
)

var _ChineseNames = [...]string{
	"小寒", "大寒", "立春", "雨水", "惊蛰", "春分",
	"清明", "谷雨", "立夏", "小满", "芒种", "夏至",
	"小暑", "大暑", "立秋", "处暑", "白露", "秋分",
	"寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
}

func ChineseNames() [24]string {
	return _ChineseNames
}

// Name returns the Chinese name of the solar term.
//
// 返回节气名称(立春雨水...)，无效节气返回空字符串
func (st SolarTerm) Name() string {
	if st <= Invalid || st > WinterSolstice {
		return ""
	}
	return _ChineseNames[(st-1)%24]
}

func (st SolarTerm) String() string {
	return st.Name()
}

// Value returns the value of the solar term.
//
// 返回节气的数值，序数(1,2,...,23,24)，无效节气返回 0
func (st SolarTerm) Value() byte {
	if !st.Valid() {
		return 0
	}
	return byte(st)
}

func (st SolarTerm) Valid() bool {
	return Invalid < st && st <= WinterSolstice
}

// IsMajorSolarTerm reports whether the solar term is a major solar term.
//
// 返回当前节气是否为中气
func (st SolarTerm) IsMajorSolarTerm() bool {
	return st%2 == 0
}
