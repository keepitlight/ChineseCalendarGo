package notation

import "testing"

func TestWuxing(t *testing.T) {
	if !Wood.IsInterPromoting(Fire) ||
		!Fire.IsInterPromoting(Earth) ||
		!Earth.IsInterPromoting(Metal) ||
		!Metal.IsInterPromoting(Water) ||
		!Water.IsInterPromoting(Wood) {
		t.Fatal("wuxing inter-promoting error")
	}
	if !Wood.IsInterRegulating(Earth) ||
		!Fire.IsInterRegulating(Metal) ||
		!Earth.IsInterRegulating(Water) ||
		!Metal.IsInterRegulating(Wood) ||
		!Water.IsInterRegulating(Fire) {
		t.Fatal("wuxing inter-regulating error")
	}
}
