package controls

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// TestCopyButtonCopiesAndShowsTheCheck checks the clipboard write, the
// Copied flag and the renamed icon while the framework animation runs.
func TestCopyButtonCopiesAndShowsTheCheck(t *testing.T) {
	var b *CopyButton
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		b = NewCopyButton(c, "npm install my-ui")
	}
	tt := ui.NewTester(view, 320, 200)
	tt.Frame()
	if err := tt.Click("Copy"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if got := tt.Clipboard(); got != "npm install my-ui" {
		t.Errorf("the clipboard holds %q", got)
	}
	if !b.Copied() {
		t.Error("Copied did not report the copied state")
	}
	if _, ok := tt.Access("Copied"); !ok {
		t.Error("the check's accessible name did not appear")
	}
	// Advance the tester's clock, never wall time: the check goes away
	// when the framework animation reaches rest.
	tt.Advance(3 * time.Second)
	tt.Frame()
	if b.Copied() {
		t.Error("Copied stayed true after the animation")
	}
	if _, ok := tt.Access("Copy"); !ok {
		t.Error("the copy icon did not come back")
	}
}

// TestCopyButtonLabels checks that an app can name both icon states in
// its own language.
func TestCopyButtonLabels(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewCopyButton(c, "text").Labels("复制", "已复制")
	}
	tt := ui.NewTester(view, 320, 200)
	tt.Frame()
	if _, ok := tt.Access("复制"); !ok {
		t.Error("the localized copy name is not on screen")
	}
}
