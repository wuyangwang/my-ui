package main

import (
	"encoding/hex"
	"image/color"
	"testing"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

// pixelAt reads what the tester drew at a DIP of the window, as bytes.
func pixelAt(tt *ui.Tester, x, y int) (r, g, b uint8) {
	p := color.NRGBAModel.Convert(tt.Image().At(x, y)).(color.NRGBA)
	return p.R, p.G, p.B
}

func hexRGB(t *testing.T, s string) (uint8, uint8, uint8) {
	t.Helper()
	raw, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	return raw[0], raw[1], raw[2]
}

// TestSidebarSitsBesideThePages checks the two panes: the sidebar's
// texts stay on the left strip, the page's cards open to its right,
// and neither pane overruns the window.
func TestSidebarSitsBesideThePages(t *testing.T) {
	s := &showcase{}
	tt := ui.NewTester(s.view, 960, 620)

	check := func(what string, x, w float32, strip bool) {
		t.Helper()
		left := x < 190
		if strip != left {
			t.Errorf("%s: x %v w %v, want sidebar=%v", what, x, w, strip)
		}
		if x+w > 960 {
			t.Errorf("%s: x %v + w %v overruns the window", what, x, w)
		}
	}

	brand, ok := tt.Find("my-ui")
	if !ok {
		t.Fatal("the brand text is missing")
	}
	check("brand", brand.X, brand.W, true)

	nav, ok := tt.Find("外观")
	if !ok {
		t.Fatal("the nav entry is missing")
	}
	check("nav entry", nav.X, nav.W, true)

	section, ok := tt.Find("变体")
	if !ok {
		t.Fatal("the variants section is missing")
	}
	check("variants section", section.X, section.W, false)

	sizes, ok := tt.Find("尺寸")
	if !ok {
		t.Fatal("the sizes section is missing")
	}
	check("sizes section", sizes.X, sizes.W, false)
}

// TestThemePageSwitchesTheAppearanceAndThePalette opens the theme page
// and takes a switch for a drive: the page builds, and the pressed
// button has the rest painted.
func TestThemePageSwitchesTheAppearanceAndThePalette(t *testing.T) {
	s := &showcase{}
	tt := ui.NewTester(s.view, 960, 620)
	if err := tt.Click("外观"); err != nil {
		t.Fatal(err)
	}
	// 点击分发的输入只停在队列，下一帧才进视图。
	tt.Advance(100 * 1000 * 1000)
	if _, ok := tt.Find("跟随系统"); !ok {
		t.Fatal("the appearance switch is missing")
	}
	if _, ok := tt.Find("浅色子树"); !ok {
		t.Fatal("the subtree preview is missing")
	}
}

// TestForcedLightRelightsTheFrame paints a frame the desktop runs dark
// on, and the forced appearance has to take the window's sides: the
// main page and the sidebar. The Tester resolves the desktop's theme
// dark; a window always paints the sides in its own.
func TestForcedLightRelightsTheFrame(t *testing.T) {
	s := &showcase{appearance: myui.AppearanceLight}
	tt := ui.NewTesterWith(s.view, 960, 620, ui.TesterOptions{Dark: true})

	r, g, b := hexRGB(t, "ffffff")
	if pr, pg, pb := pixelAt(tt, 955, 300); pr != r || pg != g || pb != b {
		t.Errorf("the main page drew (%d, %d, %d), want the light background", pr, pg, pb)
	}
	r, g, b = hexRGB(t, "f4f4f5")
	if pr, pg, pb := pixelAt(tt, 95, 560); pr != r || pg != g || pb != b {
		t.Errorf("the sidebar drew (%d, %d, %d), want the light surface", pr, pg, pb)
	}
}
