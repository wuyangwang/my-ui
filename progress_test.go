package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestProgressIsOnScreen(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewProgress(c, 0.4).Label("上传进度")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Access("上传进度"); !ok {
		t.Error("the bar is not on screen")
	}
}
