package notation

// YinYang is a concept that originated in Chinese philosophy,
// describing an opposite but interconnected, self-perpetuating cycle.
//
// 阴阳是一个起源于中国哲学的概念，它描述了一种对立但相互联系的、自我延续的循环。
type YinYang int

const (
	YIN  YinYang = 0 + iota // YIN is negative 阴
	YANG                    // YANG is positive 阳

	Negative = YIN  // Negative is negative 阴
	Positive = YANG // Positive is positive 阳

	阴 = YIN
	阳 = YANG
)
