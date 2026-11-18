package myui

import (
	"image/color"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestApplyUsesShadcnThemeForTheSystemAndNoPalette(t *testing.T) {
	var th *ui.Theme
	view := func(c *ui.Context) {
		th = Apply(c, Options{})
		ui.Text(c, "Plain")
	}
	tt := ui.NewTesterWith(view, 200, 100, ui.TesterOptions{Dark: true})
	if th == nil {
		t.Fatal("Apply in the system appearance with no palette left the theme alone")
	}
	if !th.Dark || th.Background != ShadcnDark().Background {
		t.Errorf("the default theme is not shadcn's dark one: dark %v, background %v", th.Dark, th.Background)
	}
	tt.SetDark(false)
}

func TestApplyForcesTheAppearanceAndThePalette(t *testing.T) {
	var got *ui.Theme
	view := func(c *ui.Context) {
		got = Apply(c, Options{Appearance: AppearanceDark, Palette: PaletteBlue})
	}
	ui.NewTester(view, 200, 100)
	if !got.Dark {
		t.Error("the theme is not dark after Apply, want the forced dark")
	}
	accent, hover, pressed := Palettes[PaletteBlue].Triplet(true)
	if got.Accent != accent || got.AccentHover != hover || got.AccentPressed != pressed {
		t.Errorf("accent %v, hover %v, pressed %v; want the dark blue ones %v, %v, %v",
			got.Accent, got.AccentHover, got.AccentPressed, accent, hover, pressed)
	}
	if got.Background != ShadcnDark().Background {
		t.Errorf("background %v, want shadcn's dark one's %v", got.Background, ShadcnDark().Background)
	}
}

func TestApplyFollowsTheSystemOfTheFrameItRunsIn(t *testing.T) {
	var got *ui.Theme
	view := func(c *ui.Context) {
		got = Apply(c, Options{Palette: PaletteRose})
	}
	tt := ui.NewTesterWith(view, 200, 100, ui.TesterOptions{Dark: true})
	if !got.Dark {
		t.Error("the theme is not dark, want the system's")
	}
	accent, hover, pressed := Palettes[PaletteRose].Triplet(true)
	if got.Accent != accent || got.AccentHover != hover || got.AccentPressed != pressed {
		t.Errorf("accent %v, hover %v, pressed %v; want the rose dark ones", got.Accent, got.AccentHover, got.AccentPressed)
	}
	tt.SetDark(false)
}

func TestApplyTakesTheSystemAccent(t *testing.T) {
	accent := ui.Hex("#dc2626")
	var got *ui.Theme
	view := func(c *ui.Context) {
		got = Apply(c, Options{Appearance: AppearanceLight})
	}
	ui.NewTesterWith(view, 200, 100, ui.TesterOptions{
		Preferences: ui.Preferences{Accent: accent},
	})
	if got.Accent != accent {
		t.Errorf("accent %v, want the system's %v", got.Accent, accent)
	}
	if got.AccentText != ui.RGB(255, 255, 255) {
		t.Errorf("AccentText %v, want white on a dark accent", got.AccentText)
	}
	if got.AccentHover == accent || got.AccentPressed == accent {
		t.Error("hover and pressed do not move from the accent")
	}
}

func TestApplyScalesText(t *testing.T) {
	var got *ui.Theme
	view := func(c *ui.Context) {
		got = Apply(c, Options{Appearance: AppearanceDark})
	}
	ui.NewTesterWith(view, 300, 200, ui.TesterOptions{
		Preferences: ui.Preferences{TextScale: 2},
	})
	if want := ShadcnDark().FontSize * 2; got.FontSize != want {
		t.Errorf("FontSize %v, want twice the dark theme's %v", got.FontSize, want)
	}
}

func TestApplyHighContrast(t *testing.T) {
	var got *ui.Theme
	view := func(c *ui.Context) {
		got = Apply(c, Options{Appearance: AppearanceDark})
	}
	ui.NewTesterWith(view, 300, 200, ui.TesterOptions{
		Preferences: ui.Preferences{HighContrast: true},
	})
	if got.Focus.A != 255 {
		t.Errorf("focus ring alpha %d, want an opaque one", got.Focus.A)
	}
	if got.Border == ShadcnDark().Border {
		t.Error("the border stayed as the dark theme's, want it strengthened")
	}
	if got.TextMuted == ShadcnDark().TextMuted {
		t.Error("the muted text stayed as the dark theme's, want it strengthened")
	}
}

func TestScalesResolveOnTheTheme(t *testing.T) {
	th := ui.LightTheme()
	if got := TextSize(th, TextMd); got != th.FontSize {
		t.Errorf("TextSize at scale 1 is %v, want the theme's %v", got, th.FontSize)
	}
	if got := Space(th, SpacingXs); got != th.Space(1) {
		t.Errorf("Space at 1 is %v, want the theme's", got)
	}
	if got := Radius(th, RadiusLg); got != th.Radius*1.5 {
		t.Errorf("Radius at 1.5 is %v, want the theme's raised one", got)
	}
}

func TestOnAccent(t *testing.T) {
	if got := OnAccent(ui.Hex("#09090b")); got != ui.RGB(255, 255, 255) {
		t.Errorf("OnAccent on near black is %v, want white", got)
	}
	if got := OnAccent(ui.Hex("#fef08a")); got != ui.RGB(24, 24, 27) {
		t.Errorf("OnAccent on near white is %v, want ink", got)
	}
}

// TestSubtreeThemeLeavesTheWindowInk builds a light frame with a dark
// Themed subtree inside: the texts before and after the subtree keep
// the window's ink, which only the frame's Apply sets. A subtree's
// theme must never reach them — a pane's theme switch used to run over
// the root's text style, and dark panes turned light pages' inherited
// ink near-white, invisible on the light background.
func TestSubtreeThemeLeavesTheWindowInk(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Children(func() {
			ui.Text(c, "前")
			ui.Themed(c, NewTheme(c, Options{Appearance: AppearanceDark})).
				Children(func() { ui.Text(c, "pane") })
			ui.Text(c, "后")
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)

	for _, what := range []string{"前", "后"} {
		where, ok := tt.Find(what)
		if !ok {
			t.Fatalf("%s is not on screen", what)
		}
		darkest, _, ok := ink(tt, where)
		if !ok {
			t.Fatalf("%s drew no pixels", what)
		}
		if darkest > 0.35 {
			t.Errorf("the %s text draws luma %.2f, want the window's near-black ink", what, darkest)
		}
	}
	// The pane's own text shows the dark ink: the subtree really runs
	// the other theme, the window's ink just stays out.
	where, ok := tt.Find("pane")
	if !ok {
		t.Fatal("pane is not on screen")
	}
	_, brightest, ok := ink(tt, where)
	if !ok || brightest < 0.8 {
		t.Errorf("the pane's ink luma is %.2f, want the subtree's near-white ones", brightest)
	}
}

// ink walks a found text's bounds and reads the extremes the glyphs
// drew: the darkest and the brightest pixel's luma, from 0, black, to
// 1, white. A glyph that drew nothing reports ok false.
func ink(tt *ui.Tester, where ui.Rect) (darkest, brightest float32, ok bool) {
	darkest, brightest, ok = 1, 0, false
	for y := int(where.Y); y < int(where.Y+where.H); y++ {
		for x := int(where.X); x < int(where.X+where.W); x++ {
			p := color.NRGBAModel.Convert(tt.Image().At(x, y)).(color.NRGBA)
			l := (0.2126*float32(p.R) + 0.7152*float32(p.G) + 0.0722*float32(p.B)) / 255
			if l < darkest {
				darkest = l
			}
			if l > brightest {
				brightest = l
			}
			if p.A != 0 && (p.R != p.G || p.G != p.B) {
				// Any colored pixel is a drawn one over a plain face.
				ok = true
			}
		}
	}
	return darkest, brightest, ok
}
