package solarterm

import (
	"testing"
	"time"
)

var solarTermsArray = [...]SolarTerm{
	MinorCold,
	MajorCold,
	BeginningOfSpring,
	RainWater,
	AwakeningOfInsects,
	SpringEquinox,
	PureBrightness,
	GrainRain,
	BeginningOfSummer,
	GrainBuds,
	GrainInEar,
	SummerSolstice,
	MinorHeat,
	MajorHeat,
	BeginningOfAutumn,
	EndOfHeat,
	WhiteDew,
	AutumnEquinox,
	ColdDew,
	FrostsDescent,
	BeginningOfWinter,
	MinorSnow,
	MajorSnow,
	WinterSolstice,
}

func TestSolarTerm_Get(t *testing.T) {
	for y := 1904; y <= 3000; y++ {
		for j, st := range solarTermsArray {
			i := (y-1904)*24 + j
			gt := getTime(i)
			tt := Get(y, st)
			if !tt.Valid() {
				t.Errorf("Get(%d, %s) invalid", y, st.Name())
			} else if gt.Unix() != tt.Unix() {
				t.Errorf("Get(%d, %s).Time() != %s", y, st.Name(), gt.Format("2006-01-02 15:04:05"))
			} else if tt.SolarTerm() != st {
				t.Errorf("Get(%d, %s).SolarTerm() != %s", y, st.Name(), st.Name())
			} else if !gt.Equal(tt.Time()) {
				t.Errorf("Get(%d, %s).Time() != %s", y, st.Name(), gt.Format("2006-01-02 15:04:05"))
			} else if !tt.Is(st) {
				t.Errorf("Get(%d, %s).SolarTerm() != %s", y, st.Name(), st.Name())
			}
		}
	}
	if Get(1900, MinorCold).Valid() {
		t.Error("Get(1900, 小寒) error")
	}
	if Get(2025, SolarTerm(0)).Valid() {
		t.Error("Get(2025, 无效节气) error")
	}
	if Get(2022, MinorCold).IsToday() {
		t.Error("Get(2022, 小寒).IsToday() error")
	}
	if !Get(2023, BeginningOfSummer).Prev().Is(GrainRain) {
		t.Error("Get(2023, 立夏).Prev().Is() error")
	}
	if !Get(2099, BeginningOfWinter).Next().Is(MinorSnow) {
		t.Error("Get(2099, 立夏).Next().Is() error")
	}
}

func TestAnnual(t *testing.T) {
	for i, index := range Annual(2023) {
		if index.Unix() != getTimestamp((2023-1904)*24+i) {
			t.Error("Annual(2023) error")
		} else if index.SolarTerm().Name() != ChineseNames[i] {
			t.Error("Annual(2023) error")
		}
	}
	if len(Annual(1900)) > 0 {
		t.Error("Annual(1900) error")
	}
}

func TestCast(t *testing.T) {
	c := time.Date(2023, 8, 1, 0, 0, 0, 0, time.Local)
	pw := getTime((2023-1904)*24 + 14 - 1) // 大暑 2023.7.23
	nw := getTime((2023-1904)*24 + 15 - 1) // 立秋 2023.8.8
	p, n := Lookup(c)
	if !p.Is(MajorHeat) || p.Unix() != pw.Unix() {
		t.Errorf("Lookup(%s) error", c.Format("2006-01-02 15:04:05"))
	}
	if !n.Is(BeginningOfAutumn) || n.Unix() != nw.Unix() {
		t.Errorf("Lookup(%s) error", c.Format("2006-01-02 15:04:05"))
	}
}
func TestCurrent(t *testing.T) {
	now := time.Now()
	p, n := Current()
	pw, nw := Lookup(now)
	if p != pw || n != nw {
		t.Error("Current() error")
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		jd float64
	}
	tests := []struct {
		name   string
		args   args
		wantY  int
		wantM  int
		wantD  int
		wantH  int
		wantM2 int
		wantS  int
	}{
		{"2018-02-04 05:28:29", args{2458153.72812377}, 2018, 2, 4, 5, 28, 29},
		{"50046-01-13 12:00:00", args{19999999}, 50046, 1, 13, 12, 00, 00},
		{"-4712-01-02 12:00:00", args{1}, -4712, 1, 2, 12, 00, 00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotY, gotM, gotD, gotH, gotM2, gotS := GetDateTime(JulianDay(tt.args.jd))
			if gotY != tt.wantY {
				t.Errorf("Convert() got_Y = %v, want %v", gotY, tt.wantY)
			}
			if int(gotM) != tt.wantM {
				t.Errorf("Convert() got_M = %v, want %v", gotM, tt.wantM)
			}
			if gotD != tt.wantD {
				t.Errorf("Convert() got_D = %v, want %v", gotD, tt.wantD)
			}
			if gotH != tt.wantH {
				t.Errorf("Convert() got_h = %v, want %v", gotH, tt.wantH)
			}
			if gotM2 != tt.wantM2 {
				t.Errorf("Convert() got_m = %v, want %v", gotM2, tt.wantM2)
			}
			if gotS != tt.wantS {
				t.Errorf("Convert() got_s = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
