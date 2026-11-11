package myui

import (
	"github.com/egoist/mygo/ui"
)

// Kbd is a key cap: the key a shortcut needs.
//
//	myui.NewKbd(c, "Ctrl", "K")
//
// The key cap is styled by the library, as shadcn's: a small box with a
// border, in the monospaced face.
type Kbd struct {
	ui.Element
}

// NewKbd creates the key caps of a shortcut, one for each key.
func NewKbd(c *ui.Context, keys ...string) *Kbd {
	t := c.Theme()
	k := &Kbd{}
	row := ui.Row(c).Gap(t.Space(1)).AlignItems(ui.Center)
	row.Children(func() {
		for _, key := range keys {
			cap := ui.Text(c, key).
				FontSize(TextSize(t, TextXs)).
				TextColor(t.TextMuted).
				Background(t.SurfaceHover).
				Border(1, t.Border).
				Radius(t.Radius/2).
				Padding(t.Space(0.5), t.Space(1.5))
			cap.Font("monospace")
		}
	})
	k.Element = row
	return k
}
