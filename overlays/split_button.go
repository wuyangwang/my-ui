package overlays

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// SplitButton is a button plus a drop-down of related actions, drawn
// in the window with mygo's PopoverMenuButton and my-ui's menu look.
type SplitButton struct {
	ui.Element
}

// NewSplitButton builds a split button. onClick runs when the main
// section is clicked; items fill the arrow's menu.
func NewSplitButton(c *ui.Context, label string, onClick func(), items []MenuItem) *SplitButton {
	t := c.Theme()
	parts := ui.SplitButton(c, label)
	parts.Group.Radius(Radius(t, RadiusMd)).Transition(StateTransition())
	parts.Arrow.PopoverMenuButton(func(m *ui.Menu) {
		for _, item := range items {
			if item.Separator {
				m.Separator()
				continue
			}
			if m.Item(item.Label).Chosen() && item.OnClick != nil {
				item.OnClick()
			}
		}
	})
	if parts.Button.Clicked() && onClick != nil {
		onClick()
	}
	return &SplitButton{Element: parts.Group}
}

// NewPrimarySplitButton is the accent-colored split button.
func NewPrimarySplitButton(c *ui.Context, label string, onClick func(), items []MenuItem) *SplitButton {
	b := NewSplitButton(c, label, onClick, items)
	b.Element = b.Element.Background(c.Theme().Accent).TextColor(c.Theme().AccentText)
	return b
}
