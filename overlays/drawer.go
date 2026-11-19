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

// DrawerOptions configures a Drawer at construction time. The open flag
// is the app's. Every field is read again each frame from the options the
// view passes.
type DrawerOptions struct {
	// Open is the app's flag; the panel belongs to the app.
	Open *bool
	// Side is the edge the panel comes from. The zero value, DrawerRight,
	// is the settings-panel side.
	Side DrawerSide
	// Trigger builds the control that opens the drawer. It runs before
	// the overlay so its click lands in the same frame.
	Trigger func() ui.Element
	// Title is the bold heading and the panel's accessible name.
	Title string
	// Description is the muted line under the title.
	Description string
	// Body fills the panel. The library builds it inside panel.Children.
	Body func()
	// Footer builds the right-aligned action row.
	Footer func()
	// OnOpenChange runs when the drawer opens or closes, from any source.
	// It is kept across frames, so attaching it late does not lose a
	// change.
	OnOpenChange func(open bool)
	// OnDismiss runs when Escape or a click on the scrim closes it.
	OnDismiss func()
	// ShowClose hides the corner close button when false. Nil keeps it on.
	ShowClose *bool
	// Grow is the panel's share of the main axis; MaxMain is its cap in
	// spacing units. Both default to shadcn's w-3/4 max-w-sm.
	Grow, MaxMain float32
}

// drawerState is what a Drawer keeps in its Local between frames.
type drawerState struct {
	open          openState
	parts         panelParts
	side          DrawerSide
	grow, maxMain float32
}

// Drawer shows a panel over the edge of the window while the app's flag
// is true.
//
//	overlays.NewDrawer(c, overlays.DrawerOptions{
//		Open:    &app.settingsOpen,
//		Side:    overlays.DrawerRight,
//		Trigger: func() ui.Element { return controls.NewButton(c, "设置").Element },
//		Title:   "设置",
//		Body:    func() { ui.Text(c, "在这里调整偏好") },
//	})
//
// The library builds the trigger first, then the overlay, and puts the
// Body inside the panel. A click on the scrim or Escape closes it.
type Drawer struct {
	ui.Element
	st        *drawerState
	open      *bool
	dismissed bool
}

// NewDrawer shows a panel over the edge of the window while o.Open is
// true. Unlike the old positional constructor, it never returns early
// while closed, so fluent methods such as Title and Footer always take
// effect when the drawer next opens.
func NewDrawer(c *ui.Context, o DrawerOptions) *Drawer {
	st := stateFor(c, o.Open, "drawer", func() drawerState {
		return drawerState{
			parts:   panelParts{ShowClose: true},
			side:    DrawerRight,
			grow:    3,
			maxMain: 90,
		}
	})
	if o.Title != "" {
		st.parts.Title = o.Title
	}
	if o.Description != "" {
		st.parts.Description = o.Description
	}
	if o.Body != nil {
		st.parts.Body = o.Body
	}
	if o.Footer != nil {
		st.parts.Footer = o.Footer
	}
	if o.ShowClose != nil {
		st.parts.ShowClose = *o.ShowClose
	}
	st.side = o.Side
	if o.Grow > 0 {
		st.grow = o.Grow
	}
	if o.MaxMain > 0 {
		st.maxMain = o.MaxMain
	}
	if o.OnOpenChange != nil {
		st.open.onChange = o.OnOpenChange
	}
	if o.OnDismiss != nil {
		st.open.onDismiss = o.OnDismiss
	}
	st.open.flush()

	d := &Drawer{st: st, open: st.open.openPtr(o.Open)}
	if o.Trigger != nil {
		t := o.Trigger()
		if t.Clicked() && !*d.open {
			*d.open = true
			st.open.record(true)
		}
	}
	// While closed there is still a trigger and a state to build; only
	// the overlay is skipped.
	if *d.open {
		d.build(c)
	}
	return d
}

// build lays out the scrim and the panel at the chosen edge.
func (d *Drawer) build(c *ui.Context) {
	st, th := d.st, c.Theme()
	ui.Overlay(c, func() {
		back := ui.Box(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).Modal()
		back.Background(ui.RGBA(0, 0, 0, 0.5))
		if back.Clicked() {
			*d.open = false
			d.dismissed = true
			st.open.record(false)
			st.open.recordDismiss()
		}
		back.Children(func() {
			switch st.side {
			case DrawerLeft:
				back.Row().AlignItems(ui.Stretch)
				d.panel(c, ui.Box(c).Grow(st.grow).MaxWidth(th.Space(st.maxMain)), th)
				ui.Box(c).Grow(1).IgnoreInput(true)
			case DrawerBottom, DrawerTop:
				col := back.Column()
				if st.side == DrawerBottom {
					col.Justify(ui.End)
				}
				d.panel(c, ui.Box(c).AlignSelf(ui.Stretch).MaxHeight(th.Space(st.maxMain)), th)
			default:
				back.Row().AlignItems(ui.Stretch)
				ui.Box(c).Grow(1).IgnoreInput(true)
				d.panel(c, ui.Box(c).Grow(st.grow).MaxWidth(th.Space(st.maxMain)), th)
			}
		})
	})
}

// panel styles the panel of the drawer and fills it.
func (d *Drawer) panel(c *ui.Context, panel ui.Element, th *ui.Theme) {
	st := d.st
	panel.Background(th.Background).Border(1, th.Border).Radius(th.Radius).
		Shadow(0, 0, 24, 0, ui.RGBA(0, 0, 0, 0.18)).
		Padding(th.Space(5)).
		Transition(StateTransition())
	OverlayEnter(c, panel, "drawer-enter")
	if panel.OverlayShortcut(0, ui.KeyEscape) {
		*d.open = false
		d.dismissed = true
		st.open.record(false)
		st.open.recordDismiss()
	}
	if st.parts.Title != "" {
		panel.Label(st.parts.Title)
	}
	buildPanelParts(c, panel, st.parts, func() {
		*d.open = false
		st.open.record(false)
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

// Dismissed reports whether the panel was dismissed — by Escape or a
// click on the scrim — in this frame. The close button is not a dismissal.
func (d *Drawer) Dismissed() bool { return d.dismissed }

// OnOpenChange registers fn for open changes and returns the drawer. A
// change that happened before this call is delivered at once.
func (d *Drawer) OnOpenChange(fn func(open bool)) *Drawer {
	d.st.open.onChange = fn
	d.st.open.flush()
	return d
}

// OnDismiss registers fn for dismissals and returns the drawer.
func (d *Drawer) OnDismiss(fn func()) *Drawer {
	d.st.open.onDismiss = fn
	d.st.open.flush()
	return d
}

// Side sets the edge the panel comes from, for chaining. It works while
// the drawer is closed: the edge is used when it next opens.
func (d *Drawer) Side(s DrawerSide) *Drawer {
	d.st.side = s
	return d
}

// Size sets how much of the window the panel takes: grow is the share of
// the main axis and maxMain is the cap in spacing units. It returns the
// drawer for chaining.
func (d *Drawer) Size(grow, maxMain float32) *Drawer {
	d.st.grow, d.st.maxMain = grow, maxMain
	return d
}

// Title sets the heading of the panel, for chaining.
func (d *Drawer) Title(s string) *Drawer {
	d.st.parts.Title = s
	return d
}

// Description sets the muted line under the title, for chaining.
func (d *Drawer) Description(s string) *Drawer {
	d.st.parts.Description = s
	return d
}

// Footer builds the right-aligned action row of the panel.
func (d *Drawer) Footer(build func()) *Drawer {
	d.st.parts.Footer = build
	return d
}

// ShowClose shows or hides the corner close button. On by default.
func (d *Drawer) ShowClose(on bool) *Drawer {
	d.st.parts.ShowClose = on
	return d
}
