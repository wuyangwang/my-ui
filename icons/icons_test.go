package icons

import "testing"

func TestIconHasTheSizeOfTheSet(t *testing.T) {
	w, h := Save().Size()
	if w != 24 || h != 24 {
		t.Errorf("the save icon is %v by %v, want 24 by 24", w, h)
	}
}

func TestIconIsParsedOnce(t *testing.T) {
	if Save() != Save() {
		t.Error("the save icon was parsed again")
	}
}

func TestAliasPointsAtItsIcon(t *testing.T) {
	if Close() != X() {
		t.Error("the alias Close is not the icon X")
	}
	if Trash2() != Trash2() {
		t.Error("the alias Trash is not the icon Trash2")
	}
}

// TestTheSetHasTheUsualIcons fails to build if the set loses an icon an
// app reaches for.
func TestTheSetHasTheUsualIcons(t *testing.T) {
	icons := []any{
		Save(), X(), Check(), Trash2(), Search(), Settings(),
		ChevronDown(), LoaderCircle(), CircleAlert(), TriangleAlert(),
		Plus(), Minus(), Pencil(), House(), Funnel(), Ellipsis(),
	}
	if len(icons) != 16 {
		t.Errorf("the set has %d icons here, want 16", len(icons))
	}
}
