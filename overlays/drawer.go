package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
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
	// DrawerTop is the top edge of the window.
	DrawerTop
)

// Drawer shows a panel over the edge of the window, while the app's flag
// is true.
//
//	myui.NewDrawer(c, &app.open, myui.DrawerRight, trigger, func(panel ui.Element) {
//		ui.Text(c, "面板里的内容")
//	}).Title("设置").Description("在这里调整偏好").
//		Footer(func() { myui.NewButton(c, "完成").Primary() })
//
// The trigger is built by the app, and a click on it opens the panel; a
// click on the scrim closes the panel, and Escape closes it too.
//
// The scrim and the panel are styled by the library, as shadcn's sheet:
// the scrim dims the window, the panel is a card at the edge of it, with a
// close button in its corner. Title, Description and Footer add the
// structured parts of a shadcn sheet.
type Drawer struct {
	ui.Element
	open      *bool
	side      DrawerSide
	grow      float32
	maxMain   float32
	dismissed bool
	*Parts
}

// NewDrawer shows a panel over the edge of the window while open is
// true. The flag is the app's: it opens and closes the panel.
func NewDrawer(c *ui.Context, open *bool, side DrawerSide, trigger func() ui.Element, build func(panel ui.Element)) *Drawer {
	d := &Drawer{open: open, side: side, grow: 3, maxMain: 90}
	if trigger != nil {
		t := trigger()
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
		back.Background(ui.RGBA(0, 0, 0, 0.5))
		if back.Clicked() && d.open != nil {
			*d.open = false
			d.dismissed = true
		}
		back.Children(func() {
			switch side {
			case DrawerLeft:
				back.Row().AlignItems(ui.Stretch)
				d.panel(c, ui.Box(c).Grow(d.grow).MaxWidth(th.Space(d.maxMain)), th, build, open)
				ui.Box(c).Grow(1).IgnoreInput(true)
			case DrawerBottom, DrawerTop:
				col := back.Column()
				if side == DrawerBottom {
					col.Justify(ui.End)
				}
				d.panel(c, ui.Box(c).AlignSelf(ui.Stretch).MaxHeight(th.Space(d.maxMain)), th, build, open)
			default:
				back.Row().AlignItems(ui.Stretch)
				ui.Box(c).Grow(1).IgnoreInput(true)
				d.panel(c, ui.Box(c).Grow(d.grow).MaxWidth(th.Space(d.maxMain)), th, build, open)
			}
		})
	})
	return d
}

// panel styles the panel of the drawer and fills it.
func (d *Drawer) panel(c *ui.Context, panel ui.Element, th *ui.Theme, build func(panel ui.Element), open *bool) {
	panel.Background(th.Background).Border(1, th.Border).Radius(th.Radius).
		Shadow(0, 0, 24, 0, ui.RGBA(0, 0, 0, 0.18)).
		Padding(th.Space(5)).
		Transition(StateTransition())
	OverlayEnter(c, panel, "drawer-enter")
	if panel.OverlayShortcut(0, ui.KeyEscape) && d.open != nil {
		*d.open = false
		d.dismissed = true
	}
	d.Parts = BuildSlots(c, build, panel, open)
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

// Dismissed reports whether the panel was dismissed — by Escape or a
// click on the scrim — in this frame. The close button is not a dismissal,
// as in mygo's dialogs.
func (d *Drawer) Dismissed() bool { return d.dismissed }

// OnDismiss calls fn when the panel was dismissed in this frame.
func (d *Drawer) OnDismiss(fn func()) *Drawer {
	if fn != nil && d.Dismissed() {
		fn()
	}
	return d
}

// Size sets how much of the window the panel takes: grow is the share of
// the main axis and maxMain is the cap in spacing units, as shadcn's
// w-3/4 max-w-sm. It returns the drawer for chaining.
func (d *Drawer) Size(grow, maxMain float32) *Drawer {
	d.grow, d.maxMain = grow, maxMain
	return d
}

// Title sets the heading of the panel, for chaining.
func (d *Drawer) Title(s string) *Drawer {
	if d.Parts != nil {
		d.Parts.Title(s)
	}
	return d
}

// Description sets the muted line under the title, for chaining.
func (d *Drawer) Description(s string) *Drawer {
	if d.Parts != nil {
		d.Parts.Description(s)
	}
	return d
}

// Footer builds the right-aligned action row of the panel.
func (d *Drawer) Footer(build func()) *Drawer {
	if d.Parts != nil {
		d.Parts.Footer(build)
	}
	return d
}

// ShowClose shows or hides the corner close button. On by default.
func (d *Drawer) ShowClose(on bool) *Drawer {
	if d.Parts != nil {
		d.Parts.ShowClose(on)
	}
	return d
}
