package display

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// NewOutline wraps mygo's generic ui.Outline[K] in a shadcn container and
// gives each row a light, single-line look: the theme's surface, a
// hairline border, a rounded face, and text at the shared small size.
// Opening and closing branches, the keyboard, the selection and the tree
// semantics in assistive technology all remain mygo's.
//
// mygo's Outline takes its root items separately. This constructor keeps
// the shorter signature by using the zero value of K as a virtual root:
// children(zero) returns the root items, and children(item) returns a
// node's children (nil for a leaf). With string keys, "" is the virtual
// root. The zero value must not also be a real item; title is only called
// for real items.
//
// The tree state is the app's: set state.List.Selected (and optionally
// state.List.Key), fill state.Open to show branches, and bump
// state.Generation when children would return different items.
//
//	var state ui.OutlineState[string]
//	display.NewOutline(c, &state, func(path string) []string {
//		if path == "" {
//			return []string{"src", "README.md"}
//		}
//		return files[path] // nil for files
//	}, func(path string) string {
//		return path[strings.LastIndexByte(path, '/')+1:]
//	})
//
// K must be comparable, as mygo's OutlineState requires. Go 1.27 handles
// the generic call directly: there is no reflection and no separate
// non-generic variant. Colors and metrics come from c.Theme() on each
// build and changes fade through StateTransition; no timers or goroutines
// are used.
func NewOutline[K comparable](c *ui.Context, state *ui.OutlineState[K], children func(K) []K, title func(K) string) ui.Element {
	if c == nil || state == nil {
		return ui.Element{}
	}
	if children == nil {
		children = func(K) []K { return nil }
	}
	var root K
	roots := children(root)
	t := c.Theme()
	return ui.Outline(c, state, roots, children, func(item K) {
		label := ""
		if title != nil {
			label = title(item)
		}
		if label == "" {
			label = fmt.Sprint(item)
		}
		ui.Row(c).AlignItems(ui.Center).Children(func() {
			ui.Text(c, label).FontSize(TextSize(t, TextSm))
		})
	}).Background(t.Surface).Border(1, t.Border).
		Radius(Radius(t, RadiusLg)).Padding(t.Space(2)).
		Transition(StateTransition())
}
