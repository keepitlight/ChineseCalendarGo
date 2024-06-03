package util

import "testing"

func TestMod(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"test_1", args{0, 12}, 12},
		{"test_2", args{1, 12}, 1},
		{"test_3", args{12, 12}, 12},
		{"test_4", args{13, 12}, 1},
		{"test_5", args{2, 12}, 2},
		{"test_6", args{14, 12}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Mod(tt.args.a, tt.args.b); gotResult != tt.wantResult {
				t.Errorf("Mod() = %v, want %v", gotResult, tt.wantResult)
			}
		})
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
			gotY, gotM, gotD, gotH, gotM2, gotS := Convert(tt.args.jd)
			if gotY != tt.wantY {
				t.Errorf("Convert() got_Y = %v, want %v", gotY, tt.wantY)
			}
			if gotM != tt.wantM {
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

func TestSmallNumbers(t *testing.T) {
	for _, tt := range []struct {
		name string
		n    uint16
		want string
	}{
		{
			name: "0",
			n:    0,
			want: "零",
		},
		{
			name: "1",
			n:    1,
			want: "一",
		},
		{
			name: "2",
			n:    2,
			want: "二",
		},
		{
			name: "3",
			n:    3,
			want: "三",
		},
		{
			name: "4",
			n:    4,
			want: "四",
		},
		{
			name: "5",
			n:    5,
			want: "五",
		},
		{
			name: "6",
			n:    6,
			want: "六",
		},
		{
			name: "7",
			n:    7,
			want: "七",
		},
		{
			name: "8",
			n:    8,
			want: "八",
		},
		{
			name: "9",
			n:    9,
			want: "九",
		},
		{
			name: "10",
			n:    10,
			want: "十",
		},
		{
			name: "11",
			n:    11,
			want: "十一",
		},
		{
			name: "12",
			n:    12,
			want: "十二",
		},
		{
			name: "50",
			n:    50,
			want: "五十",
		},
		{
			name: "89",
			n:    89,
			want: "八十九",
		},
		{
			name: "90",
			n:    90,
			want: "九十",
		},
		{
			name: "100",
			n:    100,
			want: "一百",
		},
		{
			name: "101",
			n:    101,
			want: "一百零一",
		},
		{
			name: "102",
			n:    102,
			want: "一百零二",
		},
		{
			name: "113",
			n:    113,
			want: "一百一十三",
		},
		{
			name: "1000",
			n:    1000,
			want: "一千",
		},
		{
			name: "1003",
			n:    1003,
			want: "一千零三",
		},
		{
			name: "1004",
			n:    1004,
			want: "一千零四",
		},
		{
			name: "1056",
			n:    1056,
			want: "一千零五十六",
		},
		{
			name: "1105",
			n:    1105,
			want: "一千一百零五",
		},
		{
			name: "1157",
			n:    1157,
			want: "一千一百五十七",
		},
		{
			name: "2008",
			n:    2008,
			want: "二千零八",
		},
		{
			name: "5713",
			n:    5713,
			want: "五千七百一十三",
		},
		{
			name: "10009",
			n:    10009,
			want: "九",
		},
		{
			name: "10010",
			n:    10010,
			want: "十",
		},
		{
			name: "10011",
			n:    10011,
			want: "十一",
		},
		{
			name: "10000",
			n:    10000,
			want: "零",
		},
		{
			name: "9000",
			n:    9000,
			want: "九千",
		},
		{
			name: "9876",
			n:    9876,
			want: "九千八百七十六",
		},
	} {
		v := Small(tt.n)
		if tt.want != v {
			t.Errorf("%s: want %s, got %s", tt.name, tt.want, v)
		} else {
			t.Logf("%s -> %s", tt.name, v)
		}
	}
}

func TestNumbers(t *testing.T) {
	for _, tt := range []struct {
		name string
		n    uint64
		want string
	}{
		{
			name: "0",
			n:    0,
			want: "零",
		},
		{
			name: "1",
			n:    1,
			want: "一",
		},
		{
			name: "2",
			n:    2,
			want: "二",
		},
		{
			name: "3",
			n:    3,
			want: "三",
		},
		{
			name: "4",
			n:    4,
			want: "四",
		},
		{
			name: "5",
			n:    5,
			want: "五",
		},
		{
			name: "6",
			n:    6,
			want: "六",
		},
		{
			name: "7",
			n:    7,
			want: "七",
		},
		{
			name: "8",
			n:    8,
			want: "八",
		},
		{
			name: "9",
			n:    9,
			want: "九",
		},
		{
			name: "10",
			n:    10,
			want: "十",
		},
		{
			name: "11",
			n:    11,
			want: "十一",
		},
		{
			name: "12",
			n:    12,
			want: "十二",
		},
		{
			name: "50",
			n:    50,
			want: "五十",
		},
		{
			name: "89",
			n:    89,
			want: "八十九",
		},
		{
			name: "90",
			n:    90,
			want: "九十",
		},
		{
			name: "100",
			n:    100,
			want: "一百",
		},
		{
			name: "101",
			n:    101,
			want: "一百零一",
		},
		{
			name: "102",
			n:    102,
			want: "一百零二",
		},
		{
			name: "113",
			n:    113,
			want: "一百一十三",
		},
		{
			name: "1000",
			n:    1000,
			want: "一千",
		},
		{
			name: "1003",
			n:    1003,
			want: "一千零三",
		},
		{
			name: "1004",
			n:    1004,
			want: "一千零四",
		},
		{
			name: "1056",
			n:    1056,
			want: "一千零五十六",
		},
		{
			name: "1105",
			n:    1105,
			want: "一千一百零五",
		},
		{
			name: "1157",
			n:    1157,
			want: "一千一百五十七",
		},
		{
			name: "2008",
			n:    2008,
			want: "二千零八",
		},
		{
			name: "10009",
			n:    10009,
			want: "一万零九",
		},
		{
			name: "10010",
			n:    10010,
			want: "一万零一十",
		},
		{
			name: "10011",
			n:    10011,
			want: "一万零一十一",
		},
		{
			name: "1 0012",
			n:    10012,
			want: "一万零一十二",
		},
		{
			name: "10 1010",
			n:    101010,
			want: "十万一千零一十",
		},
		{
			name: "20 0110",
			n:    200110,
			want: "二十万零一百一十",
		},
		{
			name: "10 0020 0110",
			n:    10_0020_0110,
			want: "十亿零二十万零一百一十",
		},
		{
			name: "30 0020 0110",
			n:    30_0020_0110,
			want: "三十亿零二十万零一百一十",
		},
		{
			name: "42 3192 8765",
			n:    42_3192_8765,
			want: "四十二亿三千一百九十二万八千七百六十五",
		},
		{
			name: "4188 5042 3192 8765",
			n:    4188_5042_3192_8765,
			want: "四千一百八十八兆五千零四十二亿三千一百九十二万八千七百六十五",
		},
		{
			name: "1844 4188 5042 3192 8765",
			n:    1844_4188_5042_3192_8765,
			want: "一千八百四十四京四千一百八十八兆五千零四十二亿三千一百九十二万八千七百六十五",
		},
		{
			name: "1844 4000 5002 3092 0765",
			n:    1844_4000_5002_3092_0765,
			want: "一千八百四十四京四千兆五千零二亿三千零九十二万零七百六十五",
		},
		{
			name: "1044 0000 0000 3092 0765",
			n:    1044_0000_0000_3092_0765,
			want: "一千零四十四京零三千零九十二万零七百六十五",
		},
		{
			name: "1000 0000 0000 3092 0765",
			n:    1000_0000_0000_3092_0765,
			want: "一千京零三千零九十二万零七百六十五",
		},
		{
			name: "1000 0000 0000 0000 0765",
			n:    1000_0000_0000_0000_0765,
			want: "一千京零七百六十五",
		},
	} {
		v := Numbers(tt.n)
		if tt.want != v {
			t.Errorf("%s: want %s, got %s", tt.name, tt.want, v)
		} else {
			t.Logf("%s -> %s", tt.name, v)
		}
	}
}

func TestDigits(t *testing.T) {
	for _, tt := range []struct {
		name string
		n    int64
		want string
	}{
		{
			name: "0",
			n:    0,
			want: "〇",
		},
		{
			name: "1",
			n:    1,
			want: "一",
		},
		{
			name: "2",
			n:    2,
			want: "二",
		},
		{
			name: "3",
			n:    3,
			want: "三",
		},
		{
			name: "4",
			n:    4,
			want: "四",
		},
		{
			name: "5",
			n:    5,
			want: "五",
		},
		{
			name: "6",
			n:    6,
			want: "六",
		},
		{
			name: "7",
			n:    7,
			want: "七",
		},
		{
			name: "8",
			n:    8,
			want: "八",
		},
		{
			name: "9",
			n:    9,
			want: "九",
		},
		{
			name: "10",
			n:    10,
			want: "一〇",
		},
		{
			name: "100",
			n:    100,
			want: "一〇〇",
		},
		{
			name: "101",
			n:    101,
			want: "一〇一",
		},
		{
			name: "102",
			n:    102,
			want: "一〇二",
		},
		{
			name: "110",
			n:    110,
			want: "一一〇",
		},
		{
			name: "111",
			n:    111,
			want: "一一一",
		},
		{
			name: "230",
			n:    230,
			want: "二三〇",
		},
		{
			name: "1000",
			n:    1000,
			want: "一〇〇〇",
		},
		{
			name: "1000",
			n:    1000,
			want: "一〇〇〇",
		},
		{
			name: "1050",
			n:    1050,
			want: "一〇五〇",
		},
		{
			name: "2187",
			n:    2187,
			want: "二一八七",
		},
		{
			name: "1234567890",
			n:    1234567890,
			want: "一二三四五六七八九〇",
		},
	} {
		if got := Digits(uint64(tt.n)); got != tt.want {
			t.Errorf("%s: Digits(%d) = %q, want %q", tt.name, tt.n, got, tt.want)
		}
	}
}
