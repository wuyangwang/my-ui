package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestInputSizesGiveRoomierBounds(t *testing.T) {
	height := func(tt *ui.Tester, name string) float32 {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		return info.Bounds.H
	}
	var a, b, cval string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewInput(c, &a).Sm().Outline().Label("small")
			NewInput(c, &b).Md().Outline().Label("medium")
			NewInput(c, &cval).Lg().Outline().Label("large")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if h := height(tt, "small"); h >= height(tt, "medium") {
		t.Errorf("the small input is not shorter than the middle's: %v against %v", h, height(tt, "medium"))
	}
	if h := height(tt, "medium"); h >= height(tt, "large") {
		t.Errorf("the middle input is not shorter than the large's: %v against %v", h, height(tt, "large"))
	}
}

func TestInputDisabledAndReadOnly(t *testing.T) {
	var locked, ro string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewInput(c, &locked).Outline().Label("Locked").Disabled(true)
			NewInput(c, &ro).Outline().Label("ReadOnly").ReadOnly(true)
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("Locked")
	if !ok {
		t.Fatal("disabled input missing")
	}
	if !info.Disabled {
		t.Error("disabled input not reported as disabled")
	}
	info2, ok := tt.Access("ReadOnly")
	if !ok {
		t.Fatal("readonly input missing")
	}
	if info2.Disabled {
		t.Error("readonly input reported as disabled")
	}
	if !info2.ReadOnly {
		t.Error("readonly input not reported as readonly")
	}
}

func TestInputDarkTheme(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewInput(c, &value).Outline().Label("Field")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Access("Field"); !ok {
		t.Fatal("input missing in dark theme")
	}
}

func TestInputClearableEmptiesTheValue(t *testing.T) {
	var value string
	cleared := false
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewInput(c, &value).Outline().Label("Name").Clearable().OnCleared(func() {
			cleared = true
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)

	if err := tt.Click("Name"); err != nil {
		t.Fatal(err)
	}
	tt.Type("hello")
	tt.Advance(time.Second)
	if value != "hello" {
		t.Fatalf("the input holds %q, want %q", value, "hello")
	}
	if _, ok := tt.Access("清除"); !ok {
		t.Fatal("the clear button is not on screen")
	}
	if err := tt.Click("清除"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if value != "" {
		t.Errorf("the input holds %q after the clear, want it empty", value)
	}
	if !cleared {
		t.Error("the clear did not report")
	}
}

func TestInputOnChangedReportsTheNewValue(t *testing.T) {
	var value string
	var seen string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewInput(c, &value).Outline().Label("Name").OnChanged(func(v string) {
			seen = v
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("Name"); err != nil {
		t.Fatal(err)
	}
	tt.Type("abc")
	tt.Advance(time.Second)
	if seen != "abc" {
		t.Errorf("the change reported %q, want %q", seen, "abc")
	}
}

func TestInputOnSubmitFires(t *testing.T) {
	var value string
	submitted := false
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewInput(c, &value).Outline().Label("Name").OnSubmit(func() {
			submitted = true
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("Name"); err != nil {
		t.Fatal(err)
	}
	tt.Type("go")
	tt.Advance(time.Second)
	tt.KeyDown(ui.KeyEnter)
	tt.Advance(time.Second)
	if !submitted {
		t.Error("the Enter key did not report")
	}
}

func TestInputErrorTextIsReadWithTheInput(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewInput(c, &value).Outline().Label("Email").Invalid(true).ErrorText("邮箱格式不对")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("Email")
	if !ok {
		t.Fatal("the input is not on screen")
	}
	if info.Description != "邮箱格式不对" {
		t.Errorf("the input reads %q, want the error message", info.Description)
	}
}

func TestInputPrefixAndSuffixShowBesideTheText(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewInput(c, &value).Outline().Label("Site").Prefix(func() {
			ui.Text(c, "https://")
		}).Suffix(func() {
			ui.Text(c, ".com")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("https://"); !ok {
		t.Error("the prefix is not on screen")
	}
	if _, ok := tt.Find(".com"); !ok {
		t.Error("the suffix is not on screen")
	}
}

func TestInputPrefixIconShows(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewInput(c, &value).Outline().Label("Plain")
			NewInput(c, &value).Outline().Label("Search").PrefixIcon(testIcon)
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	plain, ok := tt.Access("Plain")
	if !ok {
		t.Fatal("the plain input is not on screen")
	}
	withIcon, ok := tt.Access("Search")
	if !ok {
		t.Fatal("the input with the icon is not on screen")
	}
	// The icon takes room before the text, so the text starts later.
	if withIcon.Bounds.X <= plain.Bounds.X {
		t.Errorf("the icon is not before the text: %v against %v", withIcon.Bounds.X, plain.Bounds.X)
	}
}

func TestInputClearButtonIsTheIconAlone(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewInput(c, &value).Outline().Label("Name").Clearable()
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("Name"); err != nil {
		t.Fatal(err)
	}
	tt.Type("hello")
	tt.Advance(time.Second)
	info, ok := tt.Access("清除")
	if !ok {
		t.Fatal("the clear button is not on screen")
	}
	// The button is the icon alone: it takes no room of a button's.
	if info.Bounds.W > 24 || info.Bounds.H > 24 {
		t.Errorf("the clear button is %v by %v, want the icon alone", info.Bounds.W, info.Bounds.H)
	}
}

func TestInputFocusTakesTheAccentBorder(t *testing.T) {
	accent, _, _ := palettes[PaletteBlue].Triplet(false)
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight, Palette: PaletteBlue})
		NewInput(c, &value).Outline().Label("Name")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	input, ok := tt.Access("Name")
	if !ok {
		t.Fatal("the input is not on screen")
	}
	// Before the focus, the border is the theme's.
	if hasColor(tt, input.Bounds.X-16, input.Bounds.Y-8, 8, input.Bounds.H+16, accent) {
		t.Error("the border is the accent before the input has the focus")
	}
	if err := tt.Click("Name"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !hasColor(tt, input.Bounds.X-16, input.Bounds.Y-8, 8, input.Bounds.H+16, accent) {
		t.Error("the border did not take the accent while the input has the focus")
	}
}

// hasColor reports whether the tester drew the color somewhere in a
// box of the window.
func hasColor(tt *ui.Tester, x, y, w, h float32, want ui.Color) bool {
	for dy := 0; dy < int(h); dy++ {
		for dx := 0; dx < int(w); dx++ {
			r, g, b := pixel(tt, x+float32(dx), y+float32(dy))
			if want == (ui.Color{}) {
				// Any ink at all: the pixel is not the background.
				if r != 0xfa || g != 0xfa || b != 0xfa {
					return true
				}
				continue
			}
			if r == want.R && g == want.G && b == want.B {
				return true
			}
		}
	}
	return false
}
