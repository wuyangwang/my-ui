package myui

import (
	"github.com/egoist/mygo/ui"
)

// Appearance chooses which appearance a framed view shows in. The zero
// value follows the system, which MyGo resolves from the desktop and
// keeps up to date: the light or dark appearance, the accent color,
// higher contrast and the text scale, which controls follow too, all
// with no code of the app's.
type Appearance uint8

const (
	// AppearanceSystem follows the desktop's appearance: the theme is the
	// one MyGo picks for the system, in dark or light as the user runs
	// other apps. An Options of only this and no palette does nothing:
	// the frame keeps MyGo's own default theme and every preference it
	// follows.
	AppearanceSystem Appearance = 0
	// AppearanceLight forces the light appearance.
	AppearanceLight Appearance = 1
	// AppearanceDark forces the dark appearance.
	AppearanceDark Appearance = 2
)

// Options choose the theme a framed view builds with.
type Options struct {
	// Appearance says light or dark. AppearanceSystem follows the
	// desktop; a preview pane shows the opposite appearance by building
	// inside ui.Themed(c, myui.NewTheme(c, Options{Appearance:
	// AppearanceDark, ...})).
	Appearance Appearance
	// Palette chooses the accent colors, whose hues the theme presets
	// (Palette) describe. PaletteNone keeps what the appearance source
	// provides: the system's accent color, or the accent the built-in
	// theme ships with where the desktop gives none.
	Palette PaletteID
}

// Apply returns nil, or the theme the frame uses from here on. Call it
// again each frame, as often as the view runs, so the theme keeps
// following the desktop's preferences; MyGo resets the frame's theme to
// the system's each frame, and this sets it back after that.
//
// The view calls it before what the theme colors:
//
//	myui.Apply(c, myui.Options{Appearance: myui.AppearanceDark, Palette: myui.PaletteBlue})
//
// A nil return leaves the frame on MyGo's default theme, which follows
// the system; any other value replaces it for this frame. Where an app
// toggles a switch to change the appearance, the app turns the Options,
// and the next frame takes it.
//
// Apply is a one-call theme switcher, not an app-wide state keeper: it
// reads the desktop's appearance and preferences and derives the theme
// from them. It holds no state across frames.
func Apply(c *ui.Context, o Options) *ui.Theme {
	if o.Appearance == AppearanceSystem && o.Palette == PaletteNone {
		return nil
	}
	t := NewTheme(c, o)
	c.SetTheme(t)
	return t
}

// NewTheme builds the theme o asks for on the frame's preferences, and
// leaves the frame alone: it answers the ask without setting anything,
// for a subtree that shows another theme as the children of a
// ui.Themed:
//
//	ui.Themed(c, myui.NewTheme(c, myui.Options{Appearance: myui.AppearanceDark}))
//
// Texts outside that Themed inherit the window's ink, which only the
// frame's own Apply sets; a subtree theme must never reach them.
// An Options that asks for nothing (the system's appearance, no
// palette) hands back the theme the frame already has. NewTheme reads
// the desktop's preferences and derives from them; it holds no state.
func NewTheme(c *ui.Context, o Options) *ui.Theme {
	if o.Appearance == AppearanceSystem && o.Palette == PaletteNone {
		return c.Theme()
	}
	prefs := c.Preferences()
	dark := false
	switch o.Appearance {
	case AppearanceSystem:
		// The frame's theme is MyGo's default this early in the frame
		// (reset above the view), so its Dark flag is the system's.
		dark = c.Theme().Dark
	case AppearanceDark:
		dark = true
	}
	base := ui.LightTheme()
	if dark {
		base = ui.DarkTheme()
	}
	t := *base
	if scale := prefs.TextScale; scale > 0 && scale != 1 {
		t.FontSize = t.FontSize * scale
	}
	if pal, ok := palette(o.Palette); ok {
		accent, hover, pressed := pal.Triplet(dark)
		t.Accent, t.AccentHover, t.AccentPressed, t.AccentText = accent, hover, pressed, OnAccent(accent)
	} else if a := prefs.Accent; a.A != 0 {
		// The system's accent colors the theme, and the derived shades
		// hover and pressed over it, as a palette's authored ones sit
		// beside their base.
		t.AccentHover, t.AccentPressed, t.AccentText = shades(a, dark)
		t.Accent = a
	}
	if prefs.HighContrast {
		// Draw borders and secondary text darker (lighter in the dark)
		// and make the focus ring opaque, as the default theme does.
		b, w := ui.RGB(0, 0, 0), ui.RGB(255, 255, 255)
		strong := b
		if dark {
			strong = w
		}
		t.Border = t.Border.MixOklab(strong, 0.3)
		t.TextMuted = t.TextMuted.MixOklab(strong, 0.25)
		t.Focus = ui.RGBA(t.Focus.R, t.Focus.G, t.Focus.B, 1)
	}
	return &t
}

// OnAccent returns the color text shows in on the accent: dark ink on a
// light accent, white on a dark one, at half the luma. The built-in
// themes sets AccentText by hand; a palette's derives its own.
func OnAccent(c ui.Color) ui.Color {
	if luma(c) > 0.5 {
		return ui.RGB(24, 24, 27)
	}
	return ui.RGB(255, 255, 255)
}

// InverseColors are the colors a tip shows in: the theme turned over, so
// that the tip stands out from what it is over. A theme that leaves
// Inverse and InverseText alone gets the ink and the background of the
// window, as MyGo does for its own tips and toasts.
func InverseColors(t *ui.Theme) (fill, text ui.Color) {
	fill, text = t.Inverse, t.InverseText
	if fill == ui.Transparent {
		fill = t.Text
	}
	if text == ui.Transparent {
		text = t.Background
	}
	return fill, text
}

// luma is the color's lightness, from 0, black, to 1, white, by the
// weights the built-in themes read intensities by.
func luma(c ui.Color) float32 {
	return (0.2126*float32(c.R) + 0.7152*float32(c.G) + 0.0722*float32(c.B)) / 255
}
