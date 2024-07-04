package time

import (
	"testing"
)

func TestDays(t *testing.T) {
	for _, tt := range []struct {
		year int
		days int
	}{
		{2017, 354},
	} {
		y := Year(tt.year)
		if tt.days != y.Days() {
			t.Errorf("Year(%d).Days(), got %d, want %d", tt.year, y.Days(), tt.days)
		}
	}
}

func TestYearDaysOfYear(t *testing.T) {
	for i := 1900; i <= 2100; i++ {
		t.Logf("%d: %d days", i, DaysOfYear(i))
	}
}
