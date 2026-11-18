package myui

import (
	"runtime"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// The metrics of the title bar. The Windows look draws its controls
// 46 DIPs wide on a 32 DIP bar, as Windows 11's own title bars do; the
// Mac look draws the traffic lights 12 DIPs across, 8 apart, as
// macOS's do.
const (
	titleBarHeight       = float32(32)
	titleBarControlWidth = float32(46)
	titleBarLightSize    = float32(12)
	titleBarLightGap     = float32(8)
)

// TitleBarLook picks how the title bar's window controls look. The
// zero look takes the build's platform: the Mac's traffic lights on a
// darwin build, the Windows icon buttons on any other (Linux and
// Windows draw Windows-style controls; a browser build assumes the
// Windows look, as most desktops run it).
type TitleBarLook uint8

const (
	// TitleBarAuto lets the platform decide the look.
	TitleBarAuto TitleBarLook = iota
	// TitleBarWindows is the Windows look: the controls sit on the
	// right, as flat icon buttons on a 32 DIP bar, and close hovers in
	// the danger color.
	TitleBarWindows
	// TitleBarMac is the macOS look: the traffic lights — close,
	// minimize and maximize — sit on the left as colored dots that
	// show their glyphs on hover, and the bar keeps its left-to-right
	// order under a right-to-left layout.
	TitleBarMac
)

// resolveLook resolves one look from the option and the build's
// platform.
func resolveLook(l TitleBarLook) TitleBarLook {
	if l == TitleBarAuto {
		if runtime.GOOS == "darwin" {
			return TitleBarMac
		}
		return TitleBarWindows
	}
	return l
}

// TitleBarOptions configures NewTitleBar. A nil callback leaves its
// window control out: no OnClose, no close button.
type TitleBarOptions struct {
	// Icon is the logo drawn in a themed square mark ahead of the
	// name: pass icons.Component() and nothing else is needed. With
	// neither Icon nor Logo, a default component mark stands in.
	Icon *ui.SVG
	// Logo builds a custom mark ahead of the name, with the frame's
	// context. It takes the place of Icon when both are given — use it
	// for an app that draws its own brand.
	Logo func()
	// Name is the name of the app, drawn beside the logo.
	Name string
	// Maximized points at the window's maximized state, which the app
	// owns: with it non-nil the maximize control flips between the
	// maximize and the restore glyph, and names itself accordingly.
	Maximized *bool
	// OnMinimize minimizes the window (win.Minimize); nil leaves the
	// button out.
	OnMinimize func()
	// OnMaximize maximizes or restores the window (win.ToggleMaximize);
	// nil leaves the button out.
	OnMaximize func()
	// OnClose closes the window (win.Close); nil leaves the button out.
	OnClose func()
	// Look picks the style of the window controls, Windows's or the
	// Mac's. The zero look follows the platform the library builds on.
	Look TitleBarLook
}

// TitleBar is a custom window title bar:
//
//	myui.NewTitleBar(c, myui.TitleBarOptions{Name: "MyGo App", Icon: icons.Component()})
//
// Only a name and an icon are needed: the brand mark is drawn for the
// app. The window's actions are wired separately — most apps want
// window.New below, which does it for them:
//
//	myui.NewTitleBar(c, myui.TitleBarOptions{Name: "MyGo App",
//		OnMinimize: win.Minimize, OnMaximize: win.ToggleMaximize,
//		OnClose:    win.Close})
//
// It draws what a system title bar would, in the theme's own colors —
// the system's does not follow the app's light and dark switch. The
// whole bar drags the window, and a double click on the empty room of
// it maximizes, as the platform's does.
//
// The look of the controls follows the platform (TitleBarOptions.Look
// takes it over): a Windows build draws flat icon buttons on the
// right, close hovering in danger; a Mac build draws the traffic
// lights on the left, dots that show their glyphs on hover.
//
// It goes with mygo.WindowOptions{TitleBarStyle: TitleBarHidden,
// HideWindowButtons: true}, which hides the system's bar and controls
// and keeps the shadow, the rounded corners and the resizing edges.
// The actions live on mygo.Window, not in the ui package, so the app
// hands them over as callbacks and owns the Maximized state. With the
// system controls still on (TitleBarHidden without HideWindowButtons),
// c.TitleBar reserves their room, the bar keeps clear of them and
// builds none of its own.
//
// The bar follows the theme every frame: its colors move with the
// light and dark switch, the palette and the window's active state.
type TitleBar struct {
	ui.Element
	c *ui.Context
	o TitleBarOptions
}

// NewTitleBar builds the title bar row and returns it. The row is
// 32 DIPs high, or the window's own title bar room when that is
// taller, and embeds the ui.Element of it, so the app styles it as any
// element: its background is transparent, over the window's own.
func NewTitleBar(c *ui.Context, o TitleBarOptions) *TitleBar {
	t := c.Theme()
	bar := c.TitleBar()
	height := titleBarHeight
	if bar.Height > height {
		height = bar.Height
	}
	tb := &TitleBar{c: c, o: o}
	// The brand is one line for an app that passes an Icon; the
	// builder is there for one that draws its own.
	brand := o.Logo
	if brand == nil {
		icon := o.Icon
		if icon == nil {
			icon = icons.Component()
		}
		brand = brandMark(c, icon)
	}
	// The whole bar is a drag handle, and the controls sit inside it:
	// a press they take never reaches the drag.
	row := ui.Row(c).Height(height).AlignItems(ui.Center).
		DragWindow().Transition(stateTransition())
	// The brand keeps off the window's left edge, as the system's
	// title bars do: Windows insets its icon, the Mac its lights. The
	// system's own controls on the left take their room instead.
	if bar.Left > 0 {
		row = row.PaddingStart(bar.Left)
	} else if resolveLook(o.Look) == TitleBarMac {
		row = row.PaddingStart(t.Space(2))
	} else {
		row = row.PaddingStart(t.Space(3))
	}
	if bar.Right > 0 {
		row = row.PaddingEnd(bar.Right)
	}
	// The window still shows its own controls: keep clear of them and
	// draw none of ours.
	own := bar.Left == 0 && bar.Right == 0
	switch resolveLook(o.Look) {
	case TitleBarMac:
		// The traffic lights sit on the left as the system's do, under
		// any layout direction.
		row = row.Direction(ui.LeftToRight)
		row.Children(func() {
			if own {
				tb.lights(height)
			}
			mark := ui.Row(c).AlignItems(ui.Center).
				MarginStart(t.Space(3))
			mark.Children(brand)
			if o.Name != "" {
				name := ui.Text(c, o.Name).SingleLine().
					FontSize(TextSize(t, TextSm)).TextColor(t.TextMuted)
				name.MarginStart(t.Space(2))
			}
		})
	default:
		row.Children(func() {
			brand()
			if o.Name != "" {
				name := ui.Text(c, o.Name).SingleLine().
					FontSize(TextSize(t, TextSm)).TextColor(t.TextMuted)
				name.MarginStart(t.Space(2))
			}
			ui.Box(c).Grow(1)
			if !own {
				return
			}
			if o.OnMinimize != nil {
				tb.control(Message(c, "minimize"), height, icons.Minus(), false, false, o.OnMinimize)
			}
			if o.OnMaximize != nil {
				maximized := o.Maximized != nil && *o.Maximized
				glyph, name := icons.Square(), Message(c, "maximize")
				if maximized {
					glyph, name = icons.Copy(), Message(c, "restore")
				}
				tb.control(name, height, glyph, false, true, o.OnMaximize)
			}
			if o.OnClose != nil {
				tb.control(Message(c, "close"), height, icons.X(), true, false, o.OnClose)
			}
		})
	}
	tb.Element = row
	return tb
}

// brandMark draws the default logo: an SVG in a themed accent square,
// centered both ways. ui.Box stacks its children at the top left, so
// a Row centers the icon.
func brandMark(c *ui.Context, svg *ui.SVG) func() {
	return func() {
		t := c.Theme()
		mark := ui.Row(c).AlignItems(ui.Center).Justify(ui.Center).
			Size(t.Space(5), t.Space(5)).Radius(t.Space(1.5)).
			Background(t.Accent)
		mark.Children(func() {
			ui.Icon(c, svg).Size(t.Space(3.5), t.Space(3.5)).
				TextColor(t.AccentText)
		})
	}
}

// control builds one Windows-style window control: an icon button with
// no look of its own, 46 DIPs wide and the height of the bar, whose
// ink dims when the window is not active. Close paints its hover and
// its press in the danger colors, over the theme's inverse ink, as
// Windows' close does. mark sets the window's maximize button (the
// snap layouts hover).
func (tb *TitleBar) control(name string, height float32, svg *ui.SVG, danger, mark bool, onClick func()) {
	c := tb.c
	t := c.Theme()
	b := NewButton(c, "").Ghost()
	b.Element = b.Element.Radius(0).Size(titleBarControlWidth, height)
	b.Label(name)
	b.FontSize(TextSize(t, TextSm))
	// The ink and the hover follow the bar's theme every frame: the
	// inactive window's controls dim, and close hovers in danger.
	ink := t.Text
	if !c.WindowActive() {
		ink = t.TextMuted
	}
	if danger {
		hover, pressed, over := shades(t.Danger, t.Dark)
		if b.Pressed() {
			b.Background(pressed)
			ink = over
		} else if b.Hovered() {
			b.Background(hover)
			ink = over
		}
	} else {
		if b.Pressed() {
			b.Background(t.SurfacePressed)
		} else if b.Hovered() {
			b.Background(t.SurfaceHover)
		}
	}
	b.TextColor(ink)
	b.Icon(svg)
	if mark {
		b.Element = b.Element.MaximizeButton()
	}
	if b.Clicked() && onClick != nil {
		onClick()
	}
}

// The colors of the traffic lights, as macOS paints them.
var (
	lightClose    = ui.Hex("#ff5f57")
	lightMinimize = ui.Hex("#febc2e")
	lightMaximize = ui.Hex("#28c840")
)

// lights builds the Mac's traffic lights: close, minimize and
// maximize, each a colored dot that shows its glyph on hover, and
// reads gray when the window is not active.
func (tb *TitleBar) lights(height float32) {
	c := tb.c
	row := ui.Row(c).Gap(titleBarLightGap).AlignItems(ui.Center)
	row.Children(func() {
		tb.light(Message(c, "close"), lightClose, lightGlyphClose, false, height, tb.o.OnClose)
		tb.light(Message(c, "minimize"), lightMinimize, lightGlyphMinus, false, height, tb.o.OnMinimize)
		tb.light(Message(c, "maximize"), lightMaximize, lightGlyphFullscreen, true, height, tb.o.OnMaximize)
	})
}

// light builds one traffic light: a round button the size of the dot,
// painted by the dot's own color, darkened on hover. glyph draws the
// glyph the hover shows; mark sets the window's maximize button; the
// app's onClick minimizes, maximizes or closes the window.
func (tb *TitleBar) light(name string, base ui.Color, glyph func(p *ui.Painter, r ui.Rect), mark bool, height float32, onClick func()) {
	c := tb.c
	t := c.Theme()
	b := NewButton(c, "")
	b.Element = b.Element.Size(titleBarLightSize, titleBarLightSize).
		Radius(titleBarLightSize / 2).Padding(0)
	b.Label(name)
	b.FontSize(TextSize(t, TextXs))
	// The dot: its own color, a step toward black on hover, and gray
	// on an inactive window. The glyph shows while the pointer is on
	// it, in a dark ink, as macOS's do.
	fill := base
	if !c.WindowActive() {
		if t.Dark {
			fill = ui.Hex("#4f4f4f")
		} else {
			fill = ui.Hex("#d8d8d8")
		}
	} else if b.Hovered() || b.Pressed() {
		fill = base.MixOklab(ui.RGB(0, 0, 0), 0.12)
	}
	b.Background(fill)
	if (b.Hovered() || b.Pressed()) && c.WindowActive() {
		b.Draw(func(p *ui.Painter, r ui.Rect) {
			glyph(p, r)
		})
	}
	if mark {
		b.Element = b.Element.MaximizeButton()
	}
	if b.Clicked() && onClick != nil {
		onClick()
	}
}

// The glyphs a hover shows on a traffic light, in a dark ink of the
// light's own hue, as macOS's are.

// lightGlyphClose draws the × a close light shows on hover.
func lightGlyphClose(p *ui.Painter, r ui.Rect) {
	cx, cy, d := r.X+r.W/2, r.Y+r.H/2, r.W/5
	var path ui.Path
	path.MoveTo(cx-d, cy-d).LineTo(cx+d, cy+d)
	path.MoveTo(cx+d, cy-d).LineTo(cx-d, cy+d)
	p.StrokePath(&path, 1, ui.RGB(74, 20, 20))
}

// lightGlyphMinus draws the – a minimize light shows on hover.
func lightGlyphMinus(p *ui.Painter, r ui.Rect) {
	cx, cy, d := r.X+r.W/2, r.Y+r.H/2, r.W/4.5
	var path ui.Path
	path.MoveTo(cx-d, cy).LineTo(cx+d, cy)
	p.StrokePath(&path, 1, ui.RGB(96, 60, 4))
}

// lightGlyphFullscreen draws the two triangles a maximize light shows
// on hover: the fullscreen glyph, one triangle pointing out to each
// side.
func lightGlyphFullscreen(p *ui.Painter, r ui.Rect) {
	cx, cy, d := r.X+r.W/2, r.Y+r.H/2, r.W/5
	var path ui.Path
	path.MoveTo(cx-d*1.4, cy).LineTo(cx-d*1.4, cy-d*0.75).LineTo(cx-d*0.55, cy).LineTo(cx-d*1.4, cy)
	path.MoveTo(cx+d*1.4, cy).LineTo(cx+d*1.4, cy+d*0.75).LineTo(cx+d*0.55, cy).LineTo(cx+d*1.4, cy)
	p.FillPath(&path, ui.RGB(4, 72, 14))
}
