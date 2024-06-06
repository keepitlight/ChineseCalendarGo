package notation

import "testing"

func TestMajor(t *testing.T) {
	for i, major := range majors {
		m := Major(i + 1)
		if m.String() != major {
			t.Errorf("%d: %s error", i+1, major)
		}
		if !m.Valid() {
			t.Errorf("%d: %s Valid() error", i+1, major)
		}
		if int(m.Value()) != i+1 {
			t.Errorf("%d: %s Value() error", i+1, major)
		}
	}
	if Major(0).Valid() {
		t.Error("0: Major(0) Valid() error")
	}
	if Major(0).String() != "" {
		t.Error("0: Major(0) String() error")
	}
	if Major(0).Value() != 0 {
		t.Error("0: Major(0) Value() error")
	}
}
