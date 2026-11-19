package bind

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	chartpkg "github.com/wuyangwang/my-ui/chart"
	markdownpkg "github.com/wuyangwang/my-ui/markdown"
)

// Bars builds a themed bar chart of values with numbered categories.
func (u UI) Bars(values []float64) ui.Element {
	return chartpkg.NewBars(u.c, values)
}

// Line builds a themed line chart of values with a soft fill and dots.
func (u UI) Line(values []float64) ui.Element {
	return chartpkg.NewLine(u.c, values)
}

// Area builds a themed area chart of values.
func (u UI) Area(values []float64) ui.Element {
	return chartpkg.NewArea(u.c, values)
}

// Pie builds a themed pie chart of values with a swatch and label each.
func (u UI) Pie(values []float64) ui.Element {
	return chartpkg.NewPie(u.c, values)
}

// Palette builds a chart palette from the frame's theme. It is chart.NewPalette
// with u.Theme(), the constructor's only argument coming from the frame.
func (u UI) Palette() chartpkg.Palette {
	return chartpkg.NewPalette(u.Theme())
}

// Markdown shows source as finished Markdown styled with the frame's theme.
func (u UI) Markdown(source string) ui.Element {
	return markdownpkg.NewMarkdown(u.c, source)
}

// Stream shows Markdown still being written; stream answers the source
// accumulated so far and is read once per frame.
func (u UI) Stream(stream func() string) ui.Element {
	return markdownpkg.NewStream(u.c, stream)
}

// NewTheme builds the theme o asks for on the frame's preferences, leaving
// the frame's own theme alone; it is NewTheme with the frame's context.
// Apply is the one-call form that also sets the theme for the frame.
func (u UI) NewTheme(o Options) *ui.Theme {
	return NewTheme(u.c, o)
}
