package notation

import (
	"testing"
)

func TestMinor(t *testing.T) {
	for i, minor := range minors {
		m := Minor(i + 1)
		if m.String() != minor {
			t.Errorf("%d: %s error", i+1, minor)
		}
		if !m.Valid() {
			t.Errorf("%d: %s Valid() error", i+1, minor)
		}
		if int(m.Value()) != i+1 {
			t.Errorf("%d: %s Value() error", i+1, minor)
		}
	}
	if Minor(0).Valid() {
		t.Error("0: Minor(0).Valid() error")
	}
	if Minor(13).Valid() {
		t.Error("13: Minor(13).Valid() error")
	}
	if Minor(0).Value() != 0 {
		t.Error("0: Minor(0).Value() error")
	}
	if Minor(0).String() != "" {
		t.Error("0: Minor(0).String() error")
	}
	if Minor(0).Sign() != SignInvalid {
		t.Error("0: Minor(0).Sign() error")
	}
	if Minor(1).Sign() != Shu {
		t.Error("1: Minor(1).Sign() error")
	}
}
