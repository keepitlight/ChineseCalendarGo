package astrology

import (
	"testing"
	"time"
)

func loc() *time.Location {
	l, _ := time.LoadLocation("PRC")
	return l
}

func makeDate(m time.Month, d int) time.Time {
	y := time.Now().Year()
	return time.Date(y, m, d, 0, 0, 0, 0, loc())
}

func TestNewSign(t *testing.T) {
	cases := []struct {
		name string
		time time.Time
		want Sign
	}{
		{"Aries", makeDate(4, 1), Aries},
		{"Taurus", makeDate(5, 1), Taurus},
		{"Gemini", makeDate(6, 1), Gemini},
		{"Cancer", makeDate(7, 1), Cancer},
		{"Leo", makeDate(8, 1), Leo},
		{"Virgo", makeDate(9, 1), Virgo},
		{"Libra", makeDate(10, 1), Libra},
		{"Scorpio", makeDate(11, 1), Scorpio},
		{"Sagittarius", makeDate(12, 1), Sagittarius},
		{"Capricorn", makeDate(1, 1), Capricorn},
		{"Aquarius", makeDate(2, 1), Aquarius},
		{"Pisces", makeDate(3, 1), Pisces},
	}
	for _, c := range cases {
		if got := Get(c.time); got != c.want || got.String() != c.name {
			t.Errorf("zodiac sign %s, got %s", c.want.String(), got.String())
		} else {
			t.Logf("zodiac sign %q %s %s %s", got.String(), got.Emoji(), got.Symbol(), got.ChineseName())
		}
	}
}
