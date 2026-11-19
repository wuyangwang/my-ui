package bind

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// TestHandleBuildsNewlyCoveredPackages checks that the handle can build one
// component from each package whose constructors were added to it: layout,
// chart and markdown, plus the root theme and chart palette helpers.
func TestHandleBuildsNewlyCoveredPackages(t *testing.T) {
	gotTheme := false
	gotPalette := false
	view := func(c *ui.Context) {
		h := Bind(c)
		h.Apply(Options{Appearance: AppearanceLight})
		gotTheme = h.NewTheme(Options{Appearance: AppearanceDark}) != nil
		gotPalette = len(h.Palette().Series) > 0
		ui.Column(c).Children(func() {
			h.Grid()
			h.Scroll()
			h.Bars([]float64{3, 7, 5, 9})
			h.Line([]float64{3, 7, 5, 9})
			h.Area([]float64{3, 7, 5, 9})
			h.Pie([]float64{3, 7, 5, 9})
			h.Markdown("hello from bind")
			h.Stream(func() string { return "stream from bind" })
		})
	}
	tt := ui.NewTester(view, 500, 900)
	tt.Advance(time.Second)
	if !gotTheme {
		t.Error("the handle did not build a theme")
	}
	if !gotPalette {
		t.Error("the handle did not build a chart palette")
	}
	if _, ok := tt.Find("hello from bind"); !ok {
		t.Error("the handle did not build the Markdown source")
	}
}
