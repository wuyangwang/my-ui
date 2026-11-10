package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestInputVariantsApplyAndSize(t *testing.T) {
	width := func(tt *ui.Tester, name string) float32 {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		return info.Bounds.W
	}
	var a, b, cval string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewInput(c, &a).Sm().Placeholder("Name").Label("Name")
			NewInput(c, &b).Md().Outline().Placeholder("Email").Label("Email")
			NewInput(c, &cval).Lg().Ghost().Placeholder("Tag").Label("Tag")
		})
	}
	tt := ui.NewTester(view, 480, 320)
	tt.Advance(time.Second)
	if w := width(tt, "Name"); w >= width(tt, "Email") {
		t.Errorf("small input not smaller than middle: %v >= %v", w, width(tt, "Email"))
	}
	if w := width(tt, "Email"); w >= width(tt, "Tag") {
		t.Errorf("middle input not smaller than large: %v >= %v", w, width(tt, "Tag"))
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
