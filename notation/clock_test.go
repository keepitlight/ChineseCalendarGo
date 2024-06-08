package notation

import (
	"testing"
)

func TestDayPeriod(t *testing.T) {
	for _, tt := range []struct {
		hour uint8
		want string
	}{
		{0, "凌晨"},
		{1, "凌晨"},
		{2, "凌晨"},
		{3, "凌晨"},
		{4, "凌晨"},
		{5, "凌晨"},
		{6, "早上"},
		{7, "早上"},
		{8, "上午"},
		{9, "上午"},
		{10, "上午"},
		{11, "上午"},
		{12, "中午"},
		{13, "中午"},
		{14, "下午"},
		{15, "下午"},
		{16, "下午"},
		{17, "下午"},
		{18, "下午"},
		{19, "晚上"},
		{20, "晚上"},
		{21, "晚上"},
		{22, "晚上"},
		{23, "晚上"},
		{24, "凌晨"},
		{25, "凌晨"},
		{26, "凌晨"},
		{27, "凌晨"},
		{28, "凌晨"},
	} {
		if got := Period(int(tt.hour)); got != tt.want {
			t.Errorf("DayPeriod(%d) = %q, want %q", tt.hour, got, tt.want)
		} else {
			t.Logf("DayPeriod(%d) > %q", tt.hour, got)
		}
	}
}

func TestClock(t *testing.T) {
	for _, tt := range []struct {
		hour, minute, second int
		major                Major
		want, mm             string
	}{
		{0, 20, 20, 甲, "凌晨零点二十分", "甲子"},
		{1, 45, 0, 乙, "凌晨一点四十五分", "乙丑"},
		{2, 30, 0, 丙, "凌晨二点三十分", "丙丑"},
		{3, 5, 0, 丁, "凌晨三点零五分", "丁寅"},
		{4, 0, 0, 戊, "凌晨四点", "戊寅"},
		{5, 15, 0, 己, "凌晨五点十五分", "己卯"},
		{6, 20, 10, 庚, "早上六点二十分", "庚卯"},
		{7, 45, 0, 辛, "早上七点四十五分", "辛辰"},
		{8, 45, 0, 壬, "上午八点四十五分", "壬辰"},
		{9, 45, 0, 癸, "上午九点四十五分", "癸巳"},
		{10, 45, 0, 甲, "上午十点四十五分", "甲巳"},
		{11, 45, 0, 乙, "上午十一点四十五分", "乙午"},
		{12, 45, 0, 丙, "中午十二点四十五分", "丙午"},
		{13, 45, 0, 丁, "中午一点四十五分", "丁未"},
		{14, 45, 0, 戊, "下午二点四十五分", "戊未"},
		{15, 45, 0, 己, "下午三点四十五分", "己申"},
		{16, 45, 0, 庚, "下午四点四十五分", "庚申"},
		{17, 45, 0, 辛, "下午五点四十五分", "辛酉"},
		{18, 45, 0, 壬, "下午六点四十五分", "壬酉"},
		{19, 45, 0, 癸, "晚上七点四十五分", "癸戌"},
	} {
		i := MinorOfHour(tt.hour)
		n, ok := Pair(tt.major, i)
		if !ok {
			t.Fatalf("Pair(%d, %d) = %q, %t, want %q, %t", tt.major, i, n, ok, n, true)
		}
		got := &Clock{
			hour:       tt.hour,
			minute:     tt.minute,
			second:     tt.second,
			nanosecond: 39984021,
			notation:   n,
		}
		if got.String() != tt.want {
			t.Errorf("NewClock(%d, %d, %d, 0, notation.Major(1)).String() = %q, want %q", tt.hour, tt.minute, tt.second, got, tt.want)
		}
		if h, m, s := got.Clock(); h != tt.hour || m != tt.minute || s != tt.second {
			t.Errorf("NewClock(%d, %d, %d, 0, notation.Major(1)).Clock() = %d, %d, %d, want %d, %d, %d", tt.hour, tt.minute, tt.second, h, m, s, tt.hour, tt.minute, tt.second)
		}
		if n := got.Nanosecond(); n != 39984021 {
			t.Errorf("NewClock(%d, %d, %d, 0, notation.Major(1)).Nanosecond() = %d, want 39984021", tt.hour, tt.minute, tt.second, n)
		}
		n = got.Notation()
		if n.String() != tt.mm {
			t.Errorf("NewClock(%d, %d, %d, 0, notation.Major(1)).Notation() = %q, want %q", tt.hour, tt.minute, tt.second, n.String(), tt.mm)
		}
		if !got.Equal(got) {
			t.Errorf("NewClock(%d, %d, %d, 0, notation.Major(1)).Equal(NewClock(%d, %d, %d, 0, notation.Major(1))) = false, want true", tt.hour, tt.minute, tt.second, tt.hour, tt.minute, tt.second)
		}
		if !(*Clock)(nil).Equal((*Clock)(nil)) {
			t.Errorf("Equal(nil) error")
		}
	}
	for _, tt := range []struct {
		hour, minute, second int
	}{
		{25, 20, 0},
		{10, 60, 1},
		{10, 30, 61},
		{-10, 30, 61},
		{1, 30, 30},
	} {
		x := &Clock{
			hour:       tt.hour,
			minute:     tt.minute,
			second:     tt.second,
			nanosecond: 10000,
			notation:   0,
		}
		if x != nil {
			t.Errorf("ClockOf(%d, %d, %d, 0, notation.Major(10)) error", tt.hour, tt.minute, tt.second)
		}
	}
}
