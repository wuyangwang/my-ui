package display

import (
	"math"
	"strconv"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

// ProgressType is the shape a Progress is drawn in.
type ProgressType uint8

const (
	// ProgressLine draws a horizontal bar (the zero value).
	ProgressLine ProgressType = iota
	// ProgressCircle draws a round ring.
	ProgressCircle
)

// ProgressVariant is how a line Progress is painted.
type ProgressVariant uint8

const (
	// ProgressBar draws one continuous fill (the zero value).
	ProgressBar ProgressVariant = iota
	// ProgressSegments draws a row of discrete blocks.
	ProgressSegments
	// ProgressInset draws the fill as a pill beside the remaining track.
	ProgressInset
)

// ProgressStatus is the state a Progress reports.
type ProgressStatus uint8

const (
	// ProgressNormal is an ordinary bar, filled with the accent.
	ProgressNormal ProgressStatus = iota
	// ProgressActive adds a moving shine to the fill.
	ProgressActive
	// ProgressSuccess fills with the success color and shows a check.
	ProgressSuccess
	// ProgressException fills with the danger color and shows a cross.
	ProgressException
)

// progressMove is how long a change of value takes. The duration comes
// from the motion tokens, so every fill in the library moves at one pace.
var progressMove = ui.Animation{Duration: MotionNormal}

const progressShine = 1200 * time.Millisecond

// Progress shows how far along a task is. The constructor takes a value
// in percent, 0 to 100; for compatibility a value in (0, 1] is read as a
// fraction, so both 40 and 0.4 mean forty percent. A negative value shows
// activity of unknown length. Value(current, max) sets it from a pair.
//
//	display.NewProgress(c, 40).ShowInfo(true)
//	display.NewProgress(c, 0).Type(ProgressCircle).Status(ProgressSuccess)
//
// The fill grows to a new value through AnimateTo over MotionNormal; all
// colors are read from c.Theme() on every frame.
type Progress struct {
	ui.Element
	c *ui.Context

	percent  float32
	typ      ProgressType
	variant  ProgressVariant
	status   ProgressStatus
	size     Size
	showInfo bool
	label    string
	format   func(percent float32) string
	color    ui.Color
	height   float32
	segments int

	root      ui.Element
	info      ui.Element
	labelSlot ui.Element
	infoSlot  ui.Element
	lastLabel string
	lastInfo  string
	checkIcon ui.Element
	crossIcon ui.Element
	line      ui.Element
	track     ui.Element
	segs      ui.Element
	segBlocks []ui.Element
	segCount  int
	inset     ui.Element
	insetFill ui.Element
	insetRest ui.Element
	circle    ui.Element
	shown     float32
}

// NewProgress creates a progress bar. percent is 0..100; values in (0, 1]
// are read as a fraction for compatibility, and a negative value shows
// indeterminate activity.
func NewProgress(c *ui.Context, percent float64) *Progress {
	p := &Progress{c: c, height: 2, segments: 20}
	p.setPercent(percent)
	p.build()
	return p
}

func (p *Progress) setPercent(v float64) {
	switch {
	case v < 0:
		p.percent = -1
	case v <= 1:
		p.percent = float32(v * 100)
	default:
		p.percent = float32(v)
	}
	if p.percent > 100 {
		p.percent = 100
	}
}

// build constructs every form once; apply shows only the active one.
func (p *Progress) build() {
	t := p.c.Theme()
	root := ui.Column(p.c).Gap(t.Space(1.5)).FillWidth().Role(ui.RoleProgress)
	root.Children(func() {
		p.info = ui.Row(p.c).Gap(t.Space(3)).Justify(ui.SpaceBetween).
			FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted)
		p.info.Children(func() {
			p.labelSlot = ui.Box(p.c)
			p.infoSlot = ui.Box(p.c)
			p.checkIcon = ui.Icon(p.c, icons.Check()).FontSize(TextSize(t, TextSm)).Hidden(true)
			p.crossIcon = ui.Icon(p.c, icons.X()).FontSize(TextSize(t, TextSm)).Hidden(true)
		})

		p.line = ui.Box(p.c).FillWidth()
		p.line.Children(func() {
			p.track = ui.Box(p.c).FillWidth().Clip()
			p.track.Draw(func(pt *ui.Painter, r ui.Rect) { p.paintBar(pt, r) })
			p.segs = ui.Row(p.c).Gap(t.Space(0.5)).AlignItems(ui.Center)
			p.inset = ui.Row(p.c).Gap(t.Space(0.75)).AlignItems(ui.Center)
			p.inset.Children(func() {
				p.insetFill = ui.Box(p.c).Basis(0).Hidden(true)
				p.insetRest = ui.Box(p.c).Basis(0).Background(t.SurfaceHover)
			})
		})
		p.buildSegments()

		p.circle = ui.Box(p.c).Shrink(0)
		p.circle.Draw(func(pt *ui.Painter, r ui.Rect) { p.paintCircle(pt, r) })
	})
	p.root = root
	p.Element = root
	p.apply()
}

// buildSegments builds the blocks once for the current segment count.
func (p *Progress) buildSegments() {
	t := p.c.Theme()
	n := p.segments
	if n <= 0 {
		n = 20
	}
	p.segCount = n
	block := p.blockSize()
	p.segs.Height(block)
	p.segs.Children(func() {
		p.segBlocks = make([]ui.Element, n)
		for i := 0; i < n; i++ {
			p.segBlocks[i] = ui.Box(p.c).Grow(1).Basis(0).Height(block).
				Radius(1).Background(t.SurfaceHover)
		}
	})
}

// apply picks the active form and paints its colors on this frame's
// elements. The animated value lives on root, whose identity stays stable
// from frame to frame, so the fill moves instead of jumping.
func (p *Progress) apply() {
	t := p.c.Theme()
	fill := p.fillColor()
	p.root = p.root.Range(0, 100, float64(p.percent))
	p.shown = p.root.AnimateTo("progress-pct", p.percent, progressMove)
	shown := p.shown
	pct := clampPct(shown)
	frac := pct / 100

	p.info.Hidden(p.label == "" && !p.showInfo)
	setSlotText(p.c, p.labelSlot, &p.lastLabel, p.label, t.TextMuted)
	statusInfo := p.showInfo && (p.status == ProgressSuccess || p.status == ProgressException)
	p.infoSlot.Hidden(!p.showInfo || statusInfo)
	p.infoI(statusInfo, fill)
	if p.showInfo && !statusInfo {
		text := strconv.FormatFloat(float64(pct), 'f', -1, 32) + "%"
		if p.format != nil {
			text = p.format(pct)
		}
		setSlotText(p.c, p.infoSlot, &p.lastInfo, text, t.Text)
	}

	lineOn := p.typ == ProgressLine
	p.line.Hidden(!lineOn)
	p.circle.Hidden(lineOn)
	if !lineOn {
		d := p.circleSize()
		p.circle.Size(d, d)
		return
	}

	switch p.variant {
	case ProgressSegments:
		if p.segCount != effectiveSegments(p.segments) {
			p.buildSegments()
		}
		p.track.Hidden(true)
		p.inset.Hidden(true)
		p.segs.Hidden(false)
		filled := int(math.Round(float64(len(p.segBlocks)) * float64(pct) / 100))
		for i, b := range p.segBlocks {
			col := t.SurfaceHover
			if i < filled {
				col = fill
			}
			b.Background(col)
		}
	case ProgressInset:
		p.track.Hidden(true)
		p.segs.Hidden(true)
		p.inset.Hidden(false)
		block := p.blockSize()
		p.insetFill.Height(block).Radius(block / 2).Background(fill).Hidden(frac <= 0)
		if frac > 0 {
			p.insetFill = p.insetFill.Grow(frac).MinWidth(block)
		}
		p.insetRest.Height(block).Radius(block / 2).Hidden(frac >= 1)
		if frac < 1 {
			p.insetRest = p.insetRest.Grow(1 - frac)
		}
	default:
		p.track.Hidden(false)
		p.segs.Hidden(true)
		p.inset.Hidden(true)
		thick := p.lineThickness()
		p.track = p.track.Height(thick).Radius(thick / 2).Background(t.SurfaceHover)
	}
}

// infoI shows the status glyph (check or cross) and hides the percent.
func (p *Progress) infoI(on bool, fill ui.Color) {
	p.checkIcon.Hidden(!on || p.status != ProgressSuccess)
	p.crossIcon.Hidden(!on || p.status != ProgressException)
	if on {
		p.checkIcon = p.checkIcon.TextColor(fill)
		p.crossIcon = p.crossIcon.TextColor(fill)
	}
}

// paintBar draws the fill of a Bar, or a travelling segment when the
// value is indeterminate. It repaints each frame without rebuilding.
func (p *Progress) paintBar(pt *ui.Painter, r ui.Rect) {
	fill := p.fillColor()
	if p.percent >= 0 {
		at := clampPct(p.shown) / 100
		if at <= 0 {
			return
		}
		pt.Fill(ui.Rect{X: r.X, Y: r.Y, W: r.W * at, H: r.H}, fill, r.H/2)
		return
	}
	pt.AnimationFrame()
	phase := float32(pt.Now().UnixMilli()%int64(progressShine)) / float32(progressShine)
	w := r.W * 0.3
	x := r.X - w + (r.W+w)*phase
	pt.Clip(r, 0, func() {
		pt.Fill(ui.Rect{X: x, Y: r.Y, W: w, H: r.H}, fill, r.H/2)
	})
}

// paintCircle draws a track ring and an arc of the animated value.
func (p *Progress) paintCircle(pt *ui.Painter, r ui.Rect) {
	t := p.c.Theme()
	frac := float32(0)
	if p.shown > 0 {
		frac = p.shown / 100
	}
	stroke := p.circleStroke()
	cx, cy, rad := r.X+r.W/2, r.Y+r.H/2, r.W/2-stroke
	var ring ui.Path
	pt.StrokePath(ring.Circle(cx, cy, rad), stroke, t.SurfaceHover)
	if frac <= 0 {
		return
	}
	var arc ui.Path
	steps := max(2, int(96*frac))
	for i := 0; i <= steps; i++ {
		a := float64(frac) * 2 * math.Pi * float64(i) / float64(steps)
		x, y := cx+rad*float32(math.Sin(a)), cy-rad*float32(math.Cos(a))
		if i == 0 {
			arc.MoveTo(x, y)
		} else {
			arc.LineTo(x, y)
		}
	}
	pt.StrokePath(arc.Cap(ui.RoundCap), stroke, p.fillColor())
}

// fillColor is the hue of the fill on this frame's theme.
func (p *Progress) fillColor() ui.Color {
	if p.color != (ui.Color{}) {
		return p.color
	}
	t := p.c.Theme()
	switch p.status {
	case ProgressActive:
		return SemanticColor(t, SemanticInfo)
	case ProgressSuccess:
		return SemanticColor(t, SemanticSuccess)
	case ProgressException:
		return t.Danger
	}
	return t.Accent
}

// Type picks Line or Circle and returns the progress, for chaining.
func (p *Progress) Type(v ProgressType) *Progress {
	p.typ = v
	p.apply()
	return p
}

// Variant picks how a line is painted: Bar, Segments or Inset.
func (p *Progress) Variant(v ProgressVariant) *Progress {
	p.variant = v
	p.apply()
	return p
}

// Status colors the bar and, with ShowInfo, chooses the check or the
// cross instead of the percent.
func (p *Progress) Status(s ProgressStatus) *Progress {
	p.status = s
	p.apply()
	return p
}

// ShowInfo shows the percent (or the status icon) beside the bar.
func (p *Progress) ShowInfo(on bool) *Progress {
	p.showInfo = on
	p.apply()
	return p
}

// Format sets the function that writes the percent, such as "40 / 100".
// A nil Format restores "40%".
func (p *Progress) Format(fn func(percent float32) string) *Progress {
	p.format = fn
	p.apply()
	return p
}

// Color overrides the fill with c. A zero Color clears the override.
func (p *Progress) Color(c ui.Color) *Progress {
	p.color = c
	p.apply()
	return p
}

// Size sets the thickness of a line and the diameter of a circle.
func (p *Progress) Size(s Size) *Progress {
	p.size = s
	p.apply()
	return p
}

// Segments sets how many blocks the Segments variant draws; zero or less
// restores the default of twenty.
func (p *Progress) Segments(n int) *Progress {
	p.segments = n
	p.apply()
	return p
}

// Percent sets the value in 0..100 and returns the progress. Negative
// shows indeterminate activity.
func (p *Progress) Percent(v float32) *Progress {
	p.percent = v
	if p.percent > 100 {
		p.percent = 100
	}
	p.apply()
	return p
}

// Value sets the bar to current out of max and returns it, for chaining.
func (p *Progress) Value(current, max float64) *Progress {
	if max > 0 {
		p.percent = float32(current / max * 100)
	} else {
		p.percent = 0
	}
	if p.percent > 100 {
		p.percent = 100
	}
	p.apply()
	return p
}

// Indeterminate makes the bar show activity of unknown length.
func (p *Progress) Indeterminate() *Progress {
	p.percent = -1
	p.apply()
	return p
}

// Height sets the thickness of the bar in DIPs.
func (p *Progress) Height(h float32) *Progress {
	p.height = h
	p.apply()
	return p
}

// Label names the bar for assistive technology and returns it.
func (p *Progress) Label(label string) *Progress {
	p.label = label
	p.Element = p.Element.Label(label)
	p.apply()
	return p
}

// setSlotText puts s in slot, rebuilding only when it changed since the
// last call in this frame. The last value lives on the component, which
// is built anew each frame, so a new frame always rebuilds its children.
func setSlotText(c *ui.Context, slot ui.Element, last *string, s string, ink ui.Color) {
	if *last == s {
		return
	}
	slot.Children(func() {
		ui.Text(c, s).TextColor(ink).FontFeatures("tnum")
	})
	*last = s
}

func (p *Progress) lineThickness() float32 {
	if p.height > 0 {
		return p.height
	}
	switch p.size {
	case SizeXs:
		return 2
	case SizeLg:
		return 6
	default:
		return 4
	}
}

func (p *Progress) blockSize() float32 {
	b := p.lineThickness() * 2
	t := p.c.Theme()
	if b < t.Space(1) {
		b = t.Space(1)
	}
	return b
}

func (p *Progress) circleSize() float32 {
	switch p.size {
	case SizeXs:
		return 20
	case SizeLg:
		return 72
	default:
		return 40
	}
}

func (p *Progress) circleStroke() float32 {
	switch p.size {
	case SizeXs:
		return 2
	case SizeLg:
		return 4
	default:
		return 3
	}
}

func effectiveSegments(n int) int {
	if n <= 0 {
		return 20
	}
	return n
}

func clampPct(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}
