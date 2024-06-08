package notation

import (
	"time"
)

// Time is notation(sexagesimal cycle) time
//
// 干支计时
type Time struct {
	t *time.Time

	year  Notation
	month Notation
	day   Notation
	clock Notation
}

// New creates a new notation(sexagesimal cycle) time, t can be nil
//
// 创建干支时间，不需要精确时间时 t 可以为 nil
func New(year, month, day, clock Notation, t *time.Time) *Time {
	return &Time{
		year:  year,
		month: month,
		day:   day,
		clock: clock,
		t:     t,
	}
}

// Sign returns year's sign (symbol/animal)
//
// 返回年份生肖
func (t *Time) Sign() Sign {
	if !t.year.Valid() {
		return SignInvalid
	}
	_, i := t.year.Pair()
	return i.Sign()
}

func (t *Time) String() string {
	if !t.Valid() {
		return ""
	}
	// 返回八字
	y, m, d, c := t.year, t.month, t.day, t.clock
	return y.String() + m.String() + d.String() + c.String()
}

func (t *Time) Valid() bool {
	return t.year.Valid() && t.month.Valid() && t.day.Valid() && t.clock.Valid()
}

func (t *Time) Equal(b *Time) bool {
	if b == nil {
		return t == b
	}
	return t.year == b.year &&
		t.month == b.month &&
		t.day == b.day &&
		t.clock == b.clock
}

func (t *Time) Year() Notation {
	return t.year
}
func (t *Time) Month() Notation {
	return t.month
}
func (t *Time) Day() Notation {
	return t.day
}
func (t *Time) Clock() Notation {
	return t.clock
}

func (t *Time) Date() (year, month, day Notation) {
	return t.year, t.month, t.day
}

func (t *Time) Time() *time.Time {
	return t.t
}
