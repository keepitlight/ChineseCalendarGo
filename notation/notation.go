package notation

import "github.com/keepitlight/ChineseCalendarGo/util"

// Notation Sexagesimal cycle
//
// 干支，表示一个计元单位，由天干和地支构成，常用于表示年、月、日、时，也可用于其它计数用途
type Notation byte

const (
	JiaZi    Notation = 1 + iota // 1, 甲子 jiǎ zǐ
	YiChou                       // 2, 乙丑 yǐc hǒu
	BingYin                      // 3, 丙寅 bǐng yín
	DingMao                      // 4, 丁卯 dīng mǎo
	WuChen                       // 5, 戊辰 wù chén
	JiSi                         // 6, 己巳 jǐ sì
	GengWu                       // 7, 庚午 gēng wǔ
	XinWei                       // 8, 辛未 xīn wèi
	RenShen                      // 9, 壬申 rén shēn
	GuiYou                       // 10, 癸酉 guǐ yǒu
	JiaXu                        // 11, 甲戌 jiǎ xū
	YiHai                        // 12, 乙亥 yǐ hài
	BingZi                       // 13, 丙子 bǐng zǐ
	DingChou                     // 14, 丁丑 dīng chǒu
	WuYin                        // 15, 戊寅 wù yín
	JiMao                        // 16, 己卯 jǐ mǎo
	GengChen                     // 17, 庚辰 gēng chén
	XinSi                        // 18, 辛巳 xīn sì
	RenWu                        // 19, 壬午 rén wǔ
	GuiWei                       // 20, 癸未 guǐ wèi
	JiaShen                      // 21, 甲申 jiǎ shēn
	YiYou                        // 22, 乙酉 yǐ yǒu
	BingXu                       // 23, 丙戌 bǐng xū
	DingHai                      // 24, 丁亥 dīng hài
	WuZi                         // 25, 戊子 wù zǐ
	JiChou                       // 26, 己丑 jǐ chǒu
	GengYin                      // 27, 庚寅 gēng yín
	XinMao                       // 28, 辛卯 xīn mǎo
	RenChen                      // 29, 壬辰 rén chén
	GuiSi                        // 30, 癸巳 guǐ sì
	JiaWu                        // 31, 甲午 jiǎ wǔ
	YiWei                        // 32, 乙未 yǐ wèi
	BingShen                     // 33, 丙申 bǐng shēn
	DingYou                      // 34, 丁酉 dīng yǒu
	WuXu                         // 35, 戊戌 wù xū
	JiHai                        // 36, 己亥 jǐ hài
	GengZi                       // 37, 庚子 gēng zǐ
	XinChou                      // 38, 辛丑 xīn chǒu
	RenYin                       // 39, 壬寅 rén yín
	GuiMao                       // 40, 癸卯 guǐ mǎo
	JiaChen                      // 41, 甲辰 jiǎ chén
	YiSi                         // 42, 乙巳 yǐ sì
	BingWu                       // 43, 丙午 bǐng wǔ
	DingWei                      // 44, 丁未 dīng wèi
	WuShen                       // 45, 戊申 wù shēn
	JiYou                        // 46, 己酉 jǐ yǒu
	GengXu                       // 47, 庚戌 gēng xū
	XinHai                       // 48, 辛亥 xīn hài
	RenZi                        // 49, 壬子 rén zǐ
	GuiChou                      // 50, 癸丑 guǐ chǒu
	JiaYin                       // 51, 甲寅 jiǎ yín
	YiMao                        // 52, 乙卯 yǐ mǎo
	BingChen                     // 53, 丙辰 bǐng chén
	DingSi                       // 54, 丁巳 dīng sì
	WuWu                         // 55, 戊午 wù wǔ
	JiWei                        // 56, 己未 jǐ wèi
	GengShen                     // 57, 庚申 gēng shēn
	XinYou                       // 58, 辛酉 xīn yǒu
	RenXu                        // 59, 壬戌 rén xū
	GuiHai                       // 60, 癸亥 guǐ hài

	Invalid = 0 // 无效值

	N1  = JiaZi    // 甲子 jiǎ zǐ
	N2  = YiChou   // 乙丑 yǐ chǒu
	N3  = BingYin  // 丙寅 bǐng yín
	N4  = DingMao  // 丁卯 dīng mǎo
	N5  = WuChen   // 戊辰 wù chén
	N6  = JiSi     // 己巳 jǐ sì
	N7  = GengWu   // 庚午 gēng wǔ
	N8  = XinWei   // 辛未 xīn wèi
	N9  = RenShen  // 壬申 rén shēn
	N10 = GuiYou   // 癸酉 guǐ yǒu
	N11 = JiaXu    // 甲戌 jiǎ xū
	N12 = YiHai    // 乙亥 yǐ hài
	N13 = BingZi   // 丙子 bǐng zǐ
	N14 = DingChou // 丁丑 dīng chǒu
	N15 = WuYin    // 戊寅 wù yín
	N16 = JiMao    // 己卯 jǐ mǎo
	N17 = GengChen // 庚辰 gēng chén
	N18 = XinSi    // 辛巳 xīn sì
	N19 = RenWu    // 壬午 rén wǔ
	N20 = GuiWei   // 癸未 guǐ wèi
	N21 = JiaShen  // 甲申 jiǎ shēn
	N22 = YiYou    // 乙酉 yǐ yǒu
	N23 = BingXu   // 丙戌 bǐng xū
	N24 = DingHai  // 丁亥 dīng hài
	N25 = WuZi     // 戊子 wù zǐ
	N26 = JiChou   // 己丑 jǐ chǒu
	N27 = GengYin  // 庚寅 gēng yín
	N28 = XinMao   // 辛卯 xīn mǎo
	N29 = RenChen  // 壬辰 rén chén
	N30 = GuiSi    // 癸巳 guǐ sì
	N31 = JiaWu    // 甲午 jiǎ wǔ
	N32 = YiWei    // 乙未 yǐ wèi
	N33 = BingShen // 丙申 bǐng shēn
	N34 = DingYou  // 丁酉 dīng yǒu
	N35 = WuXu     // 戊戌 wù xū
	N36 = JiHai    // 己亥 jǐ hài
	N37 = GengZi   // 庚子 gēng zǐ
	N38 = XinChou  // 辛丑 xīn chǒu
	N39 = RenYin   // 壬寅 rén yín
	N40 = GuiMao   // 癸卯 guǐ mǎo
	N41 = JiaChen  // 甲辰 jiǎ chén
	N42 = YiSi     // 乙巳 yǐ sì
	N43 = BingWu   // 丙午 bǐng wǔ
	N44 = DingWei  // 丁未 dīng wèi
	N45 = WuShen   // 戊申 wù shēn
	N46 = JiYou    // 己酉 jǐ yǒu
	N47 = GengXu   // 庚戌 gēng xū
	N48 = XinHai   // 辛亥 xīn hài
	N49 = RenZi    // 壬子 rén zǐ
	N50 = GuiChou  // 癸丑 guǐ chǒu
	N51 = JiaYin   // 甲寅 jiǎ yín
	N52 = YiMao    // 乙卯 yǐ mǎo
	N53 = BingChen // 丙辰 bǐng chén
	N54 = DingSi   // 丁巳 dīng sì
	N55 = WuWu     // 戊午 wù wǔ
	N56 = JiWei    // 己未 jǐ wèi
	N57 = GengShen // 庚申 gēng shēn
	N58 = XinYou   // 辛酉 xīn yǒu
	N59 = RenXu    // 壬戌 rén xū
	N60 = GuiHai   // 癸亥 guǐ hài

	甲子 = JiaZi
	乙丑 = YiChou
	丙寅 = BingYin
	丁卯 = DingMao
	戊辰 = WuChen
	己巳 = JiSi
	庚午 = GengWu
	辛未 = XinWei
	壬申 = RenShen
	癸酉 = GuiYou
	甲戌 = JiaXu
	乙亥 = YiHai
	丙子 = BingZi
	丁丑 = DingChou
	戊寅 = WuYin
	己卯 = JiMao
	庚辰 = GengChen
	辛巳 = XinSi
	壬午 = RenWu
	癸未 = GuiWei
	甲申 = JiaShen
	乙酉 = YiYou
	丙戌 = BingXu
	丁亥 = DingHai
	戊子 = WuZi
	己丑 = JiChou
	庚寅 = GengYin
	辛卯 = XinMao
	壬辰 = RenChen
	癸巳 = GuiSi
	甲午 = JiaWu
	乙未 = YiWei
	丙申 = BingShen
	丁酉 = DingYou
	戊戌 = WuXu
	己亥 = JiHai
	庚子 = GengZi
	辛丑 = XinChou
	壬寅 = RenYin
	癸卯 = GuiMao
	甲辰 = JiaChen
	乙巳 = YiSi
	丙午 = BingWu
	丁未 = DingWei
	戊申 = WuShen
	己酉 = JiYou
	庚戌 = GengXu
	辛亥 = XinHai
	壬子 = RenZi
	癸丑 = GuiChou
	甲寅 = JiaYin
	乙卯 = YiMao
	丙辰 = BingChen
	丁巳 = DingSi
	戊午 = WuWu
	己未 = JiWei
	庚申 = GengShen
	辛酉 = XinYou
	壬戌 = RenXu
	癸亥 = GuiHai
)

// SexagesimalCycle is alias of Notation
//
// 别名，干支，表示一个计元单位，由天干和地支构成，常用于表示年、月、日、时，也可用于其它计数用途
type SexagesimalCycle = Notation

// Pair to pairs major and minor to create a notation, nil if major invalid, or minor invalid, or pair invalid.
//
// 天干地支配对组合成有效的干支值，参数无效或配对无效返回 nil
func Pair(major Major, minor Minor) (Notation, bool) {
	if major.Valid() && minor.Valid() {
		for r := 0; r <= 5; r++ {
			i := int(major) * r
			if util.Cycle(i, MinorCycle) == int(minor) {
				return Notation(i), true
			}
		}
	}
	return Invalid, false
}

// String returns notation string
//
// 返回干支字符串，无效值返回空字符串
func (n Notation) String() string {
	if !n.Valid() {
		return ""
	}
	return names[n-1]
}

// Index returns the index of notations
//
// 返回计时单位的序号（即六十干支顺序值 0~59），无效的干支值返回 -1
func (n Notation) Index() int {
	if !n.Valid() {
		return -1
	}
	return int(n - 1)
}

// Valid detects whether the notation is valid
//
// 验证天干和地支是否有效
func (n Notation) Valid() bool {
	return n > Invalid && n <= Cycle
}

func (n Notation) Pinyin() string {
	return Pinyin(n)
}

func (n Notation) Equal(v Notation) bool {
	if !n.Valid() || !v.Valid() {
		return false
	}
	return n == v
}

func (n Notation) Pair() (major Major, minor Minor) {
	if !n.Valid() {
		return MajorInvalid, MinorInvalid
	}
	return Major(util.Cycle(int(n), MajorCycle)), Minor(util.Cycle(int(n), MinorCycle))
}

func (n Notation) Prev() Notation {
	if n == 1 {
		return Notation(Cycle)
	}
	return n - 1
}
func (n Notation) Next() Notation {
	if n == Cycle {
		return 1
	}
	return n + 1
}

var (
	names = [...]string{
		"甲子",
		"乙丑",
		"丙寅",
		"丁卯",
		"戊辰",
		"己巳",
		"庚午",
		"辛未",
		"壬申",
		"癸酉",
		"甲戌",
		"乙亥",
		"丙子",
		"丁丑",
		"戊寅",
		"己卯",
		"庚辰",
		"辛巳",
		"壬午",
		"癸未",
		"甲申",
		"乙酉",
		"丙戌",
		"丁亥",
		"戊子",
		"己丑",
		"庚寅",
		"辛卯",
		"壬辰",
		"癸巳",
		"甲午",
		"乙未",
		"丙申",
		"丁酉",
		"戊戌",
		"己亥",
		"庚子",
		"辛丑",
		"壬寅",
		"癸卯",
		"甲辰",
		"乙巳",
		"丙午",
		"丁未",
		"戊申",
		"己酉",
		"庚戌",
		"辛亥",
		"壬子",
		"癸丑",
		"甲寅",
		"乙卯",
		"丙辰",
		"丁巳",
		"戊午",
		"己未",
		"庚申",
		"辛酉",
		"壬戌",
		"癸亥",
	}
	pinyin = [...]string{
		"jiǎ zǐ",
		"yǐ chǒu",
		"bǐng yín",
		"dīng mǎo",
		"wù chén",
		"jǐ sì",
		"gēng wǔ",
		"xīn wèi",
		"rén shēn",
		"guǐ yǒu",
		"jiǎ xū",
		"yǐ hài",
		"bǐng zǐ",
		"dīng chǒu",
		"wù yín",
		"jǐ mǎo",
		"gēng chén",
		"xīn sì",
		"rén wǔ",
		"guǐ wèi",
		"jiǎ shēn",
		"yǐ yǒu",
		"bǐng xū",
		"dīng hài",
		"wù zǐ",
		"jǐc hǒu",
		"gēng yín",
		"xīn mǎo",
		"rén chén",
		"guǐ sì",
		"jiǎ wǔ",
		"yǐ wèi",
		"bǐng shēn",
		"dīng yǒu",
		"wù xū",
		"jǐ hài",
		"gēng zǐ",
		"xīn chǒu",
		"rén yín",
		"guǐ mǎo",
		"jiǎ chén",
		"yǐ sì",
		"bǐng wǔ",
		"dīng wèi",
		"wù shēn",
		"jǐ yǒu",
		"gēng xū",
		"xīn hài",
		"rén zǐ",
		"guǐ chǒu",
		"jiǎ yín",
		"yǐ mǎo",
		"bǐng chén",
		"dīng sì",
		"wù wǔ",
		"jǐ wèi",
		"gēng shēn",
		"xīn yǒu",
		"rén xū",
		"guǐ hài",
	}
)

// Names return the Chinese names of sexagesimal cycle
//
// 六十甲子
func Names() [60]string {
	return names
}

// Pinyin return the pinyin of the given notation(sexagesimal cycle) v
//
// 返回指定干支的拼音
func Pinyin(v Notation) string {
	if !v.Valid() {
		return ""
	}
	return pinyin[v-1]
}

// TryParse to parse the notation(sexagesimal cycle) by Chinese name
//
// 根据中文名解析干支符号
func TryParse(name string) (Notation, bool) {
	for i := 0; i < len(names); i++ {
		if names[i] == name {
			return Notation(byte(i + 1)), true
		}
	}
	return Invalid, false
}

const (
	Cycle = 60 // 60 cycles
)

// Of returns the notation(sexagesimal cycle) by index, notation.Invalid if index is 0
//
// 根据序号返回干支（1 为甲子），注意，没有 0
func Of(index int) Notation {
	return Notation(util.Cycle(index, Cycle))
}
