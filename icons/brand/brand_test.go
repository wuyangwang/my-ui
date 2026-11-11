package brand

import "testing"

func TestIconHasTheSizeOfTheSet(t *testing.T) {
	w, h := GitHub().Size()
	if w != 24 || h != 24 {
		t.Errorf("the GitHub icon is %v by %v, want 24 by 24", w, h)
	}
}

func TestIconIsParsedOnce(t *testing.T) {
	if GitHub() != GitHub() {
		t.Error("the GitHub icon was parsed again")
	}
}

// TestTheSetHasTheUsualBrands fails to build if the set loses a brand
// an app reaches for.
func TestTheSetHasTheUsualBrands(t *testing.T) {
	brands := []any{
		GitHub(), GitLab(), Google(), Apple(), X(),
		Anthropic(), Claude(), Figma(), Vercel(), Docker(), Reddit(),
	}
	if len(brands) != 11 {
		t.Errorf("the set has %d brands here, want 11", len(brands))
	}
}
