package solarterm

import "testing"

func TestSolarTerm_String(t *testing.T) {
	for i := 0; i < len(_ChineseNames); i++ {
		if _ChineseNames[i] != SolarTerm(i+1).String() {
			t.Errorf("ChineseNames[%d]!= SolarTerm(%d).String()", i, i+1)
		}
	}
	if SolarTerm(0).String() != "" {
		t.Errorf("SolarTerm(0).String()!= \"\"")
	}
}
func TestSolarTerm_Value(t *testing.T) {
	for i := 0; i < len(_ChineseNames); i++ {
		if i+1 != int(SolarTerm(i+1).Value()) {
			t.Errorf("ChineseNames[%d]!= SolarTerm(%d).String()", i, i+1)
		}
	}
	if SolarTerm(0).Value() != 0 {
		t.Errorf("SolarTerm(0).Value()!= 0")
	}
}
func TestSolarTerm_IsMajorSolarTerm(t *testing.T) {
	for i := 0; i < len(_ChineseNames); i++ {
		if SolarTerm(i+1).IsMajorSolarTerm() != (i%2 == 1) {
			t.Errorf("ChineseNames[%d]!= SolarTerm(%d).IsMajorSolarTerm()", i, i+1)
		}
	}
}
