package myui

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestScaffoldRendersRegions(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		Scaffold(c, ScaffoldOptions{
			Top: func() { ui.Text(c, "Header") },
			Bottom: func() {
				ui.Row(c).Height(40).Children(func() { ui.Text(c, "Footer") })
			},
			BottomInset: 48,
			Content: func() {
				ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
					ui.Text(c, "Body")
					ui.Text(c, "End")
				})
			},
		})
	}
	tt := ui.NewTester(view, 320, 240)
	if _, ok := tt.Find("Header"); !ok {
		t.Fatal("header missing")
	}
	if _, ok := tt.Find("Body"); !ok {
		t.Fatal("body missing")
	}
	if _, ok := tt.Find("Footer"); !ok {
		t.Fatal("footer missing")
	}
	if _, ok := tt.Find("End"); !ok {
		t.Fatal("end content missing")
	}
}

func TestScaffoldUsesTitleBarSafeArea(t *testing.T) {
	var topY float32
	view := func(c *ui.Context) {
		Apply(c, Options{})
		Scaffold(c, ScaffoldOptions{
			UseTitleBar: true,
			Top:         func() { ui.Text(c, "Title") },
			Content:     func() { ui.Text(c, "Content") },
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.SetTitleBar(ui.TitleBar{Height: 28, Left: 70, Right: 0})
	tt.Advance(0)
	if r, ok := tt.Find("Title"); ok {
		topY = r.Y
	} else {
		t.Fatal("title missing")
	}
	if topY < 20 {
		t.Errorf("title not pushed below title bar area: y=%v", topY)
	}
}
