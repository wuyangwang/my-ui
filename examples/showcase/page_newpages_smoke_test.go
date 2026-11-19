package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
)

func TestNewPagesSmoke(t *testing.T) {
	for _, p := range []struct {
		name  string
		build func(*showcase, *ui.Context, *ui.Theme)
	}{
		{"alert dialog", (*showcase).alertDialogPage},
		{"list", (*showcase).listPage},
		{"outline", (*showcase).outlinePage},
	} {
		s := &showcase{}
		view := func(c *ui.Context) {
			th := myui.Apply(c, myui.Options{Appearance: myui.AppearanceLight})
			p.build(s, c, th)
		}
		tt := ui.NewTester(view, 640, 480)
		tt.Frame()
		if len(tt.Texts()) == 0 {
			t.Errorf("%s: the page drew no text", p.name)
		}
	}
}
