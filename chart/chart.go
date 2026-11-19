// Package chart wraps mygo's headless chart primitives with my-ui's palette
// and ready-to-use presets. The plots are built from the frame's theme on
// every frame: colors, radius and spacing follow SetTheme and the desktop's
// appearance, accent and text scale on the next frame. Values, labels and
// an optional selection belong to the app; this package holds no state.
package chart

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	mchart "github.com/egoist/mygo/ui/chart"
	myui "github.com/wuyangwang/my-ui"
)

// Palette is the set of colors a chart paints with, taken from one frame's
// theme. Series holds the data colors in order; Grid colors grid lines,
// Axis the tick labels and Surface the plot's quiet background.
type Palette struct {
	Series     []ui.Color
	Grid, Axis ui.Color
	Surface    ui.Color
}

// Color returns the series color of index i, cycling through Series. It is
// safe for any i, including a negative one, and for an empty palette.
func (p Palette) Color(i int) ui.Color {
	if len(p.Series) == 0 {
		return ui.Transparent
	}
	n := len(p.Series)
	return p.Series[((i%n)+n)%n]
}

// NewPalette builds a chart palette from a my-ui theme. The primary color
// leads, followed by info, success, warning and danger; a theme that
// leaves a slot empty simply contributes fewer series colors. The result
// is built fresh on each call and never cached in package state.
func NewPalette(t *ui.Theme) Palette {
	if t == nil {
		t = ui.LightTheme()
	}
	var series []ui.Color
	add := func(c ui.Color) {
		if c.A != 0 {
			series = append(series, c)
		}
	}
	add(t.Accent)
	add(myui.SemanticColor(t, myui.SemanticInfo))
	add(myui.SemanticColor(t, myui.SemanticSuccess))
	add(myui.SemanticColor(t, myui.SemanticWarning))
	add(t.Danger)
	if len(series) == 0 {
		series = []ui.Color{t.Text}
	}
	return Palette{
		Series:  series,
		Grid:    t.Border,
		Axis:    t.TextMuted,
		Surface: t.Surface,
	}
}

// Options configures a chart preset. The zero value paints a themed chart
// with numbered categories, a default height and straight line segments.
type Options struct {
	// Labels names the values along the category axis. Extra names are
	// ignored; missing ones fall back to 1, 2, 3, …
	Labels []string
	// Selected, when set, is the value under the pointer or the keyboard.
	// The plot makes itself focusable and writes -1 when the pointer
	// leaves, as mygo's chart.Select does.
	Selected *int
	// Height is the plot's height in DIPs. Zero uses 45 theme spacing
	// units (180 DIPs at the default scale).
	Height float32
	// Curve is how a line joins its points; zero is a straight line.
	Curve mchart.Curve
	// Label names the chart for assistive technology; empty means "Chart".
	Label string
	// Description says what the chart shows, for assistive technology.
	Description string
}

// NewBars builds a themed bar chart of values, with numbered categories
// and round tick labels on both axes. It is Bars with zero Options.
func NewBars(c *ui.Context, values []float64) ui.Element {
	return Bars(c, values, Options{})
}

// NewLine builds a themed line chart of values, with a soft fill under the
// line and a dot at each value. It is Line with zero Options.
func NewLine(c *ui.Context, values []float64) ui.Element {
	return Line(c, values, Options{})
}

// NewArea builds a themed area chart of values. It is Area with zero
// Options.
func NewArea(c *ui.Context, values []float64) ui.Element {
	return Area(c, values, Options{})
}

// NewPie builds a themed pie chart of values, with a color swatch and a
// label for each value. It is Pie with zero Options.
func NewPie(c *ui.Context, values []float64) ui.Element {
	return Pie(c, values, Options{})
}

// Bars builds a bar chart with the labels, height, selection and
// accessible name of o. Bars are placed by mygo's band scale, so they
// follow the plot as it resizes.
func Bars(c *ui.Context, values []float64, o Options) ui.Element {
	t := c.Theme()
	p := NewPalette(t)
	labels := labelsOf(o, len(values))
	x := mchart.NewBand(labels...).Padding(0.2, 0.1)
	lo, hi := extent(values)
	ticks := tickCount(len(values))
	y := mchart.NewLinear(lo, hi).Nice(ticks)
	yt := y.Ticks(ticks)
	xs := make([]float32, len(values))
	for i := range xs {
		xs[i] = x.At(i)
	}
	plot := mchart.Plot(c, mchart.Axes{Left: yt, Bottom: x.Ticks()}, func(area ui.Element) {
		area.Background(p.Surface)
		drawGrid(area, p, yt)
		for i, v := range values {
			mchart.Rect(c, x.Band(i), y.Span(lo, v)).
				Background(p.Color(i)).
				Radius(myui.Radius(t, myui.RadiusSm)).
				Transition(myui.StateTransition()).
				Tooltip(labels[i] + ": " + format(v))
		}
		if o.Selected != nil {
			mchart.Select(c, area, o.Selected, xs, describe(labels, values))
			if i := *o.Selected; i >= 0 && i < len(values) {
				mchart.RuleX(c, xs[i], 1).Background(p.Grid)
				mchart.Point(c, xs[i], y.At(values[i]), 8, 8).
					Radius(4).Background(p.Color(i))
			}
		}
	})
	return finish(plot, o, t)
}

// Line builds a line chart with the labels, height, selection, curve and
// accessible name of o. A NaN value breaks the line where it occurs.
func Line(c *ui.Context, values []float64, o Options) ui.Element {
	return xyPlot(c, values, o, 0.18, true)
}

// Area builds a filled area chart with the labels, height, selection,
// curve and accessible name of o. The fill runs down to the value zero.
func Area(c *ui.Context, values []float64, o Options) ui.Element {
	return xyPlot(c, values, o, 0.35, false)
}

// Pie builds a pie chart with the labels, height and accessible name of
// o. Values below zero and NaN take no room, as in mygo's chart.Pie.
func Pie(c *ui.Context, values []float64, o Options) ui.Element {
	t := c.Theme()
	p := NewPalette(t)
	labels := labelsOf(o, len(values))
	slices := mchart.Pie(values, 0.02)
	size := heightOf(t, o)
	root := ui.Column(c).Gap(t.Space(3)).Label(labelOf(o)).Description(o.Description)
	root.Children(func() {
		ui.Box(c).Size(size, size).Draw(func(pa *ui.Painter, r ui.Rect) {
			cx, cy := r.X+r.W/2, r.Y+r.H/2
			outer := min(r.W, r.H)/2 - 2
			if outer <= 0 {
				return
			}
			for i, s := range slices {
				pa.FillPath(mchart.Sector(cx, cy, 0, outer, s), p.Color(i))
			}
		})
		ui.Column(c).Gap(t.Space(1)).Children(func() {
			for i, v := range values {
				row := ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center)
				row.Children(func() {
					ui.Box(c).Size(t.Space(3), t.Space(3)).
						Radius(myui.Radius(t, myui.RadiusSm)).Background(p.Color(i))
					ui.Text(c, labels[i]+": "+format(v)).
						FontSize(myui.TextSize(t, myui.TextSm))
				})
			}
		})
	})
	return root
}

// xyPlot is the shared builder of the line and area presets.
func xyPlot(c *ui.Context, values []float64, o Options, alpha float32, stroke bool) ui.Element {
	t := c.Theme()
	p := NewPalette(t)
	labels := labelsOf(o, len(values))
	x := mchart.NewBand(labels...).Padding(0, 0)
	lo, hi := extent(values)
	ticks := tickCount(len(values))
	y := mchart.NewLinear(lo, hi).Nice(ticks)
	yt := y.Ticks(ticks)
	xs, ys, base := make([]float32, len(values)), make([]float32, len(values)), make([]float32, len(values))
	for i, v := range values {
		xs[i], ys[i], base[i] = x.At(i), y.At(v), y.At(0)
	}
	fill, curve := p.Color(0).Alpha(alpha), o.Curve
	plot := mchart.Plot(c, mchart.Axes{Left: yt, Bottom: x.Ticks()}, func(area ui.Element) {
		area.Background(p.Surface)
		area.Draw(func(pa *ui.Painter, r ui.Rect) {
			for _, tk := range yt {
				yy := mchart.Y(r, tk.At)
				pa.Line(r.X, yy, r.X+r.W, yy, 1, p.Grid)
			}
			pa.FillPath(mchart.Area(r, xs, base, ys, curve), fill)
			if stroke {
				pa.StrokePath(mchart.Line(r, xs, ys, curve), 2, p.Color(0))
			}
		})
		for i := range values {
			mchart.Point(c, xs[i], ys[i], 6, 6).
				Radius(3).Background(p.Color(0)).
				Transition(myui.StateTransition()).
				Tooltip(labels[i] + ": " + format(values[i]))
		}
		if o.Selected != nil {
			mchart.Select(c, area, o.Selected, xs, describe(labels, values))
			if i := *o.Selected; i >= 0 && i < len(values) {
				mchart.RuleX(c, xs[i], 1).Background(p.Grid)
				mchart.Point(c, xs[i], ys[i], 8, 8).
					Radius(4).Background(p.Color(0))
			}
		}
	})
	return finish(plot, o, t)
}

// drawGrid paints the horizontal grid lines of a plot, in the palette's
// grid color. It runs in one Draw with any series paths, since a second
// Draw on the same element would overwrite the first.
func drawGrid(area ui.Element, p Palette, ticks []mchart.Tick) {
	area.Draw(func(pa *ui.Painter, r ui.Rect) {
		for _, tk := range ticks {
			yy := mchart.Y(r, tk.At)
			pa.Line(r.X, yy, r.X+r.W, yy, 1, p.Grid)
		}
	})
}

// describe answers what a selected value says to assistive technology.
func describe(labels []string, values []float64) func(int) string {
	return func(i int) string { return labels[i] + ": " + format(values[i]) }
}

// labelsOf returns exactly n labels, numbering the ones o did not name.
func labelsOf(o Options, n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = strconv.Itoa(i + 1)
	}
	copy(out, o.Labels)
	return out
}

// extent is a value domain that always includes zero and has room to
// draw, even when every value is zero or the input is empty.
func extent(values []float64) (float64, float64) {
	lo, hi := mchart.Extent(values...)
	lo = min(lo, 0)
	if !(hi > lo) {
		hi = lo + 1
	}
	return lo, hi
}

func tickCount(n int) int {
	if n < 2 {
		return 2
	}
	return min(n, 6)
}

func heightOf(t *ui.Theme, o Options) float32 {
	if o.Height > 0 {
		return o.Height
	}
	return t.Space(45)
}

func labelOf(o Options) string {
	if o.Label != "" {
		return o.Label
	}
	return "Chart"
}

func format(v float64) string { return strconv.FormatFloat(v, 'g', -1, 64) }

func finish(e ui.Element, o Options, t *ui.Theme) ui.Element {
	return e.Height(heightOf(t, o)).Gap(t.Space(2)).
		Label(labelOf(o)).Description(o.Description)
}
