package all

import "testing"

func TestTheSetHasItsNames(t *testing.T) {
	if len(Names) < 2000 {
		t.Errorf("the set has %d names, want the whole set", len(Names))
	}
	// The names are the ones of the set, sorted and once each.
	seen := map[string]bool{}
	for i, name := range Names {
		if seen[name] {
			t.Errorf("the name %q is twice in the list", name)
		}
		seen[name] = true
		if i > 0 && Names[i-1] > name {
			t.Errorf("the names are not in order: %q before %q", Names[i-1], name)
		}
	}
}

func TestByNameFindsTheIcon(t *testing.T) {
	svg := ByName("Save")
	if svg == nil {
		t.Fatal("the save icon is not in the set")
	}
	if ByName("Nothing") != nil {
		t.Error("the set has an icon of a name it does not have")
	}
}
