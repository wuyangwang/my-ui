package forms

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestFieldsFillByDefaultAndWidthOverrides checks the default width of
// the input-like controls: they take the room of their column, and
// Width(w) pins them to a fixed size.
func TestFieldsFillByDefaultAndWidthOverrides(t *testing.T) {
	for _, tc := range []struct {
		name  string
		build func(c *ui.Context, fixed bool) ui.Element
	}{
		{"input", func(c *ui.Context, fixed bool) ui.Element {
			f := NewInput(c, new(string))
			if fixed {
				f.Width(170)
			}
			return f.Element
		}},
		{"textarea", func(c *ui.Context, fixed bool) ui.Element {
			f := NewTextArea(c, new(string))
			if fixed {
				f.Width(170)
			}
			return f.Element
		}},
		{"password", func(c *ui.Context, fixed bool) ui.Element {
			f := NewPasswordInput(c, new(string))
			if fixed {
				f.Width(170)
			}
			return f.Element
		}},
		{"select", func(c *ui.Context, fixed bool) ui.Element {
			f := NewSelect(c, new(string), []SelectOption{{Value: "a", Label: "A"}})
			if fixed {
				f.Width(170)
			}
			return f.Element
		}},
		{"combobox", func(c *ui.Context, fixed bool) ui.Element {
			f := NewComboBox(c, ComboOptions{Options: []SelectOption{{Value: "a", Label: "A"}}})
			if fixed {
				f.Width(170)
			}
			return f.Element
		}},
		{"multiselect", func(c *ui.Context, fixed bool) ui.Element {
			f := NewMultiSelect(c, new([]string), []SelectOption{{Value: "a", Label: "A"}})
			if fixed {
				f.Width(170)
			}
			return f.Element
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			measure := func(fixed bool) float32 {
				var w float32
				view := func(c *ui.Context) {
					ui.Column(c).Width(300).Children(func() {
						tc.build(c, fixed).Label(tc.name)
					})
				}
				tt := ui.NewTester(view, 320, 240)
				if info, ok := tt.Access(tc.name); ok {
					w = info.Bounds.W
				} else {
					t.Fatalf("%s is not on screen", tc.name)
				}
				return w
			}
			if got := measure(false); got < 250 {
				t.Errorf("the default width is %v, want the column's room", got)
			}
			if got := measure(true); got < 140 || got > 180 {
				// TokenField reports its content box, so its accessible width
				// is a little narrower than the field's 170 DIPs.
				t.Errorf("Width(170) gave %v, want about 170", got)
			}
		})
	}
}
