package myui

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestApplyDoesNothingForTheSystemAndNoPalette(t *testing.T) {
	var th *ui.Theme
	view := func(c *ui.Context) {
		th = Apply(c, Options{})
		ui.Text(c, "Plain")
	}
	tt := ui.NewTesterWith(view, 200, 100, ui.TesterOptions{Dark: true})
	if th != nil {
		t.Error("Apply in the system appearance with no palette replaced the theme, want nil")
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
	accent, hover, pressed := palettes[PaletteBlue].Triplet(true)
	if got.Accent != accent || got.AccentHover != hover || got.AccentPressed != pressed {
		t.Errorf("accent %v, hover %v, pressed %v; want the dark blue ones %v, %v, %v",
			got.Accent, got.AccentHover, got.AccentPressed, accent, hover, pressed)
	}
	if got.Background != ui.DarkTheme().Background {
		t.Errorf("background %v, want the dark theme's %v", got.Background, ui.DarkTheme().Background)
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
	accent, hover, pressed := palettes[PaletteRose].Triplet(true)
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
	if want := ui.DarkTheme().FontSize * 2; got.FontSize != want {
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
	if got.Border == ui.DarkTheme().Border {
		t.Error("the border stayed as the dark theme's, want it strengthened")
	}
	if got.TextMuted == ui.DarkTheme().TextMuted {
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
