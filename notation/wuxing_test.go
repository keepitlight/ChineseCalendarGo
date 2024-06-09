package notation

import "testing"

func TestWuxing(t *testing.T) {
	if !木.IsInterPromoting(火) ||
		!火.IsInterPromoting(土) ||
		!土.IsInterPromoting(金) ||
		!金.IsInterPromoting(水) ||
		!水.IsInterPromoting(木) {
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
