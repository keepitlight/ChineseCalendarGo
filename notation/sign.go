package notation

// Sign is the animal signs of year.
// In Chinese astrology, the zodiac of twelve animal signs represents
// twelve different types of personality.
//
// 表示中国传统文化中的十二生肖
type Sign byte

// ChineseZodiac is an alias of Sign
//
// 别名，表示中国传统文化中的十二生肖
type ChineseZodiac = Sign

// Symbol is an alias of Sign
//
// 别名，表示中国传统文化中的十二生肖
type Symbol = Sign

// Animal is an alias of Sign
type Animal = Sign

const (
	Rat     Sign = 1 + iota // Rat 1st 鼠
	Ox                      // Ox 2nd 牛
	Tiger                   // Tiger 3rd 虎
	Rabbit                  // Rabbit 4th 兔
	Dragon                  // Dragon 5th 龙
	Snake                   // Snake 6th 蛇
	Horse                   // Horse 7th 马
	Sheep                   // Sheep 8th 羊
	Monkey                  // Monkey 9th 猴
	Rooster                 // Rooster 10th 鸡
	Dog                     // Dog 11th 狗
	Pig                     // Pig 12th 猪

	SignInvalid Sign = 0 // 无效生肖

	Goat  = Sheep  // 羊，别名
	Loong = Dragon // 龙，别名

	Shu  = Rat     // 鼠，别名
	Niu  = Ox      // 牛，别名
	Hu   = Tiger   // 虎，别名
	Tu   = Rabbit  // 兔，别名
	Long = Dragon  // 龙，别名
	Se   = Snake   // 蛇，别名
	Ma   = Horse   // 马，别名
	Yang = Sheep   // 羊，别名
	HOU  = Monkey  // 猴，别名
	JI   = Rooster // 鸡，别名
	Gou  = Dog     // 狗，别名
	Zhu  = Pig     // 猪，别名

	SignCycle = 12

	鼠 = Rat
	牛 = Ox
	虎 = Tiger
	兔 = Rabbit
	龙 = Dragon
	蛇 = Snake
	马 = Horse
	羊 = Sheep
	猴 = Monkey
	鸡 = Rooster
	狗 = Dog
	猪 = Pig
)

// Signs is the names of Chinese zodiac.
//
// 生肖名称
func Signs() [12]string {
	return signs
}

var signs = [...]string{
	"鼠",
	"牛",
	"虎",
	"兔",
	"龙",
	"蛇",
	"马",
	"羊",
	"猴",
	"鸡",
	"狗",
	"猪",
}

// Name returns the name of Chinese zodiac.
//
// 返回生肖名称(鼠牛虎...)，无效生肖返回空字符串
func (s Sign) Name() string {
	if s == SignInvalid {
		return ""
	}
	return signs[(s-1)%12]
}

func (s Sign) String() string {
	return s.Name()
}

// Minor returns the corresponding minor of Chinese zodiac.
//
// 返回对应的地支
func (s Sign) Minor() Minor {
	return Minor(s)
}

// Value returns the corresponding index of Chinese zodiac.
//
// 返回生肖的索引值，无效生肖返回 0
func (s Sign) Value() byte {
	if !s.Valid() {
		return byte(SignInvalid)
	}
	return byte(s)
}

// Valid returns whether the sign is valid.
//
// 指示生肖是否有效
func (s Sign) Valid() bool {
	return Rat <= s && s <= Pig
}
