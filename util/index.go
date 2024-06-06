package util

import (
	"strconv"
	"strings"
)

// Mod to get the modulus of v % mod, and return v % mod + mod if v <= 0.
//
// 模运算 (modulo)，返回 v % mod，v <= 0 时返回 v % mod + mod
func Mod(v, mod int) (result int) {
	result = v % mod
	if result < 0 {
		result += mod
	}
	return
}

// Cycle to calculate the cycle of v without zero, where length denotes the period length and
// must be greater than 1, negative values are interpreted as cycling in the opposite direction,
//
// 无 0 的循环周期运算，length 为周期长度，必须大于 1，负数视为反方向的循环。
func Cycle(v int, length uint) int {
	if v == 0 {
		return 0
	}
	if v < 0 {
		v++ // offset 1
	}
	v %= int(length)
	if v <= 0 {
		v += int(length)
	}
	return v
}

var digits = []string{
	"〇", "一", "二", "三", "四", "五", "六", "七", "八", "九",
}

// Digits to convert a number to Chinese digits.
//
// 返回汉字的数字（一般用作编号），例如 2018 返回 “二〇一八”
func Digits(n uint64) string {
	s := strconv.FormatUint(n, 10)
	for i, replace := range digits {
		s = strings.Replace(s, strconv.Itoa(i), replace, -1)
	}
	return s
}

var (
	// 超四位的大数单位
	majors = [...]string{
		"", "万", "亿", "兆", "京", "垓", "秭",
		// "穰", "沟", "涧", "正", "载", "极", "恒河沙", "阿僧祇", "那由他", "不可思议", "无量大数"
	}
	// 四位以内小数单位
	minors = [...]string{
		"", "十", "百", "千",
	}
	numbers = [...]string{
		"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十",
	}
)

// 42_9496_7295
// 12_3456_7890
// 十亿千百十万千百十个

// Small returns the Chinese numerical representation of a decimal number within 10,000
// (excluding 10,000). For numbers exceeding 10,000, it takes the modulo and
// returns the corresponding Chinese string. For example, 12018 would return
// "二千零一十八" (which translates to 2018 in Chinese).
//
//   Small(12018) // 二千零一十八
//   Small(2018) // 二千零一十八
//
// 返回一万以内小数（较小的数）的十进制中文数值（不含 10000，超出 10000 取模），例如 12018 返回 “二千零一十八”
func Small(n uint16) string {
	v, _, _, _ := small(n)
	return v
}

func small(n uint16) (num, omit string, zero bool, length int) {
	if n >= w16 {
		n = n % w16
	}
	if n < 10 {
		return numbers[n], "", n == 0, 1
	}
	if n == 10 {
		return numbers[n], numbers[1], false, 1
	}
	if n < 20 {
		return numbers[10] + numbers[n-10], numbers[1], false, 2 // 例如，15 返回 “十五”
	}
	s := strconv.FormatUint(uint64(n), 10)
	length = len(s)
	var z bool
	for i := 0; i < length; i++ {
		v := s[i]
		if v == '0' {
			z = true
		} else {
			if z {
				num += numbers[0]
				z = false
			}
			u := length - i - 1
			num += numbers[int(s[i]-'0')]
			num += minors[u]
		}
	}
	return
}

const (
	w   uint64 = 10000 // 一万
	w16 uint16 = 10000
)

// Numbers returns the Chinese representation of the ordinal number n, with digits.
// For example, 2018 returns "二千零一十八".
//
// 返回汉字的数值（有数位），例如 2018 返回 “二千零一十八”
func Numbers(n uint64) (num string) {
	if n <= 10 {
		return numbers[n]
	}
	if n < 20 {
		return "十" + numbers[n-10] // 遇十忽略 1 ，例如 11 返回 “十一”
	}
	type field struct {
		num, omit, unit string
		zero            bool
		length          int
	}
	var fields []*field
	shift := func(v *field) {
		fields = append([]*field{v}, fields...)
	}
	for i := 0; i < len(majors); i++ {
		s, o, z, l := small(uint16(n % w))
		shift(&field{s, o, majors[i], z, l})
		if n < w {
			break
		}
		n /= w
	}
	var nums []string
	var zero bool
	for i, f := range fields {
		if f.zero {
			zero = true
			continue
		}
		var s []string
		if i > 0 {
			if f.length != 4 || zero {
				s = append(s, numbers[0])
				zero = false
			}
			if len(f.omit) > 0 {
				s = append(s, f.omit)
			}
		}
		s = append(s, f.num, f.unit)
		nums = append(nums, strings.Join(s, ""))
	}
	return strings.Join(nums, "")
}
