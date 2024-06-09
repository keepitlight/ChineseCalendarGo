package notation

import (
	"testing"
)

func _ss[S [4]T, T any](v1, v2, v3, v4 T) S {
	return S{v1, v2, v3, v4}
}

func TestTime(t *testing.T) {
	m := func(ns [4]Notation, ts [4]int) *Time {
		// x := time.Date(ts[0], time.Month(ts[1]), ts[2], ts[3], 0, 0, 0, time.Local)
		return New(ns[0], ns[1], ns[2], ns[3])
	}
	for _, c := range []struct {
		want string
		ns   [4]Notation
		ts   [4]int
	}{
		{"甲辰庚午壬寅甲子", _ss(JiaChen, GengWu, RenYin, JiaZi), _ss(2024, 6, 7, 23)},
	} {
		if x := m(c.ns, c.ts).String(); x != c.want {
			t.Errorf("want %s, got %s", c.want, x)
		}
	}
}
