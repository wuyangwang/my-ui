package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// DrawerSide is the edge of the window a drawer comes from.
type DrawerSide uint8

const (
	// DrawerRight is the right edge of the window, as a settings panel.
	DrawerRight DrawerSide = iota
	// DrawerLeft is the left edge of the window.
	DrawerLeft
	// DrawerBottom is the bottom edge of the window.
	DrawerBottom
)

// Drawer shows a panel over the edge of the window, while the app's flag
// is true.
//
//	myui.NewDrawer(c, &app.open, myui.DrawerRight, func() ui.Element {
//		return myui.NewButton(c, "设置").Element
//	}, func(panel ui.Element) {
//		ui.Text(c, "面板里的内容")
//	})
//
// The trigger is built by the app, and a click on it opens the panel; a
// click on the scrim closes the panel, and Escape closes it too.
//
// The scrim and the panel are styled by the library, as shadcn's sheet:
// the scrim dims the window, the panel is a card at the edge of it, with
// a close button in its corner.
type Drawer struct {
	ui.Element
	open  *bool
	build func(panel ui.Element)
}

// NewDrawer shows a panel over the edge of the window while open is
// true. The flag is the app's: it opens and closes the panel.
//
// The side says which edge the panel comes from: DrawerRight,
// DrawerLeft or DrawerBottom.
//
// The trigger is built where the drawer is, so it sits beside the panel
// it opens; a click on it opens the panel in the same frame.
//
// The build callback fills the panel: a title, a text and the controls.
func NewDrawer(c *ui.Context, open *bool, side DrawerSide, trigger func() ui.Element, build func(panel ui.Element)) *Drawer {
	d := &Drawer{open: open, build: build}
	if trigger != nil {
		t := trigger()
		// A click on the trigger opens the panel.
		if t.Clicked() && d.open != nil {
			*d.open = true
		}
	}
	if d.open == nil || !*d.open {
		return d
	}
	th := c.Theme()
	ui.Overlay(c, func() {
		// The scrim: it dims the window and takes the clicks outside
		// the panel, as a dialog's backdrop.
		back := ui.Box(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).Modal()
		back.Background(ui.RGBA(0, 0, 0, 0.4))
		if back.Clicked() && d.open != nil {
			*d.open = false
		}
		back.Children(func() {
			// The panel takes three quarters of the room, as shadcn's
			// sheet does, and no more than a card's width; the rest of
			// the scrim stays, so the window behind shows what is
			// happening.
			switch side {
			case DrawerLeft:
				back.Row().AlignItems(ui.Stretch)
				d.panel(c, ui.Box(c).Grow(3).MaxWidth(th.Space(90)), th)
				ui.Box(c).Grow(1).IgnoreInput(true)
			case DrawerBottom:
				back.Column().Justify(ui.End)
				d.panel(c, ui.Box(c).AlignSelf(ui.Stretch).MaxHeight(th.Space(80)), th)
			default:
				back.Row().AlignItems(ui.Stretch)
				ui.Box(c).Grow(1).IgnoreInput(true)
				d.panel(c, ui.Box(c).Grow(3).MaxWidth(th.Space(90)), th)
			}
		})
	})
	return d
}

// panel styles the panel of the drawer and fills it: a card at the edge
// of the window, with a close button in its corner.
func (d *Drawer) panel(c *ui.Context, panel ui.Element, th *ui.Theme) {
	panel.Background(th.Surface).Border(1, th.Border).Radius(th.Radius).
		Shadow(0, 0, 24, 0, ui.RGBA(0, 0, 0, 0.18)).
		Padding(th.Space(4)).
		Transition(stateTransition())
	// Escape closes the panel.
	if panel.OverlayShortcut(0, ui.KeyEscape) && d.open != nil {
		*d.open = false
	}
	panel.Children(func() {
		// The close button, in the corner of the panel, as shadcn's
		// sheet.
		close := NewButton(c, "").Ghost().Xs()
		close.Icon(icons.X()).Label(Message(c, "close"))
		close.Element = close.Element.Absolute().Top(th.Space(4)).Right(th.Space(4))
		if close.Clicked() && d.open != nil {
			*d.open = false
		}
		if d.build != nil {
			d.build(panel)
		}
	})
}

// Open reports whether the panel shows.
func (d *Drawer) Open() bool { return d.open != nil && *d.open }

// SetOpen opens or closes the panel and returns it, for chaining.
func (d *Drawer) SetOpen(open bool) *Drawer {
	if d.open != nil {
		*d.open = open
	}
	return d
}
