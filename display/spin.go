package display

import (
	"math"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// SpinVariant is the shape of the loading indicator.
type SpinVariant uint8

const (
	// SpinLoader is a turning ring (the zero value).
	SpinLoader SpinVariant = iota
	// SpinNetwork is a small network of pulsing nodes, for work over a
	// network such as an agent's.
	SpinNetwork
	// SpinPercentKind is a ring filled to a percent. SpinPercent selects
	// it; the constant is exported for callers that want the value.
	SpinPercentKind
)

// Spin is a loading indicator: a turning loader, a percent ring, or a
// node network, with an optional tip and an optional dimmed overlay over
// content. All motion uses Element.Loop, so the desktop's Reduce Motion
// comes free — there are no timers in the component. It is a RoleStatus,
// so a screen reader announces it as advisory information.
type Spin struct {
	ui.Element
	c *ui.Context

	size         Size
	variant      Variant
	spinKind     SpinVariant
	percent      float32
	tip          string
	root         ui.Element
	glyphRow     ui.Element
	appArea      ui.Element
	overlay      ui.Element
	loaders      [2]ui.Element
	rings        [2]ui.Element
	networks     [2]ui.Element
	tips         [2]ui.Element
	contentOn    bool
	networkPhase float64
}

// spinGlyphs are the elements of one indicator set: the standalone set at
// index 0 and the set over content at index 1.
type spinGlyphs struct {
	loader  ui.Element
	ring    ui.Element
	network ui.Element
	tip     ui.Element
}

// NewSpin creates a spinner at the text's own size and the accent's
// color.
func NewSpin(c *ui.Context) *Spin {
	s := &Spin{c: c}
	t := c.Theme()
	root := ui.Box(c).AlignSelf(ui.Start).Role(ui.RoleStatus)
	root.Children(func() {
		s.appArea = ui.Box(c)
		s.appArea.Children(func() {
			s.glyphRow = ui.Column(c).Gap(t.Space(2)).AlignItems(ui.Center)
			g := s.buildGlyphs(c, 0)
			s.loaders[0], s.rings[0], s.networks[0], s.tips[0] = g.loader, g.ring, g.network, g.tip
		})
		s.overlay = ui.Column(c).Absolute().Fill().Gap(t.Space(2)).Center().
			Background(t.Background.Alpha(0.7)).Hidden(true)
		s.overlay.Children(func() {
			g := s.buildGlyphs(c, 1)
			s.loaders[1], s.rings[1], s.networks[1], s.tips[1] = g.loader, g.ring, g.network, g.tip
		})
	})
	s.Element = root
	s.root = root
	s.apply()
	return s
}

// buildGlyphs builds one loader, ring, network and tip slot. idx keeps
// each Loop's key distinct between the two sets.
func (s *Spin) buildGlyphs(c *ui.Context, idx int) spinGlyphs {
	d := s.glyphSize()
	loader := ctrlpkg.NewIcon(c, icons.LoaderCircle()).Spin().Element.Size(d, d)
	ring := ui.Box(c).Size(d, d).Shrink(0)
	ring.Draw(func(pt *ui.Painter, r ui.Rect) { s.paintRing(pt, r) })
	network := ui.Box(c).Size(d, d).Shrink(0)
	network.Draw(func(pt *ui.Painter, r ui.Rect) { s.paintNetwork(pt, r, idx) })
	tip := ui.Box(c)
	return spinGlyphs{loader: loader, ring: ring, network: network, tip: tip}
}

// Size picks the spinner's DIP size.
func (s *Spin) Size(sz Size) *Spin {
	s.size = sz
	s.apply()
	return s
}

// Variant colors the spinner with the shared semantic colors.
func (s *Spin) Variant(v Variant) *Spin {
	s.variant = v
	s.apply()
	return s
}

// SpinPercent shows a ring filled to p (0..100) instead of the turning
// loader. A value outside the range is clamped.
func (s *Spin) SpinPercent(p float32) *Spin {
	s.spinKind = SpinPercentKind
	s.percent = p
	if s.percent < 0 {
		s.percent = 0
	}
	if s.percent > 100 {
		s.percent = 100
	}
	s.apply()
	return s
}

// SpinNetwork switches the indicator to a network of pulsing nodes.
func (s *Spin) SpinNetwork() *Spin {
	s.spinKind = SpinNetwork
	s.apply()
	return s
}

// SpinTip shows s under the indicator (or over content). An empty string
// clears the tip.
func (sp *Spin) SpinTip(text string) *Spin {
	sp.tip = text
	for _, slot := range sp.tips {
		slot.Children(func() {
			ui.Text(sp.c, text).TextColor(sp.c.Theme().TextMuted)
		})
	}
	return sp
}

// SpinContent lays content out and, while *spinning is true, dims it and
// shows the indicator with the tip over it. A nil spinning means always
// spinning. Call it each frame with the same arguments.
func (s *Spin) SpinContent(content func(), spinning *bool) *Spin {
	s.contentOn = true
	if content != nil {
		s.appArea.Children(content)
	}
	on := spinning == nil || *spinning
	s.glyphRow.Hidden(on)
	s.overlay.Hidden(!on)
	return s
}

func (s *Spin) apply() {
	d := s.glyphSize()
	// Read the animation in the build phase, not from inside Draw.
	s.networkPhase = float64(s.networks[0].Loop("spin-network", 20*time.Second, ui.Linear))
	for i := range s.loaders {
		s.loaders[i].Size(d, d)
		s.rings[i].Size(d, d)
		s.networks[i].Size(d, d)
		switch s.spinKind {
		case SpinPercentKind:
			s.loaders[i].Hidden(true)
			s.rings[i].Hidden(false)
			s.networks[i].Hidden(true)
		case SpinNetwork:
			s.loaders[i].Hidden(true)
			s.rings[i].Hidden(true)
			s.networks[i].Hidden(false)
		default:
			s.loaders[i].Hidden(false)
			s.rings[i].Hidden(true)
			s.networks[i].Hidden(true)
		}
	}
}

func (s *Spin) glyphSize() float32 {
	switch s.size {
	case SizeXs:
		return IconSizeXs
	case SizeSm:
		return IconSizeSm
	case SizeLg:
		return IconSizeLg
	default:
		return IconSizeMd
	}
}

// ink is the indicator's color on this frame's theme.
func (s *Spin) ink() ui.Color {
	t := s.c.Theme()
	switch s.variant {
	case Primary, Link:
		return t.Accent
	case Secondary:
		return t.TextMuted
	case Destructive:
		return t.Danger
	case Success, Info, Warning:
		role, _ := SemanticOf(s.variant)
		return SemanticColor(t, role)
	}
	return t.TextMuted
}

// paintRing draws a track and an arc filled to the percent.
func (s *Spin) paintRing(pt *ui.Painter, r ui.Rect) {
	t := s.c.Theme()
	frac := s.percent / 100
	stroke := max(float32(1), r.W/8)
	cx, cy, rad := r.X+r.W/2, r.Y+r.H/2, r.W/2-stroke
	var track ui.Path
	pt.StrokePath(track.Circle(cx, cy, rad), stroke, t.SurfaceHover)
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
	pt.StrokePath(arc.Cap(ui.RoundCap), stroke, s.ink())
}

// paintNetwork draws three columns of nodes joined by lines. Nodes pulse
// in sequence; the period comes from Loop, not from a timer.
func (s *Spin) paintNetwork(pt *ui.Painter, r ui.Rect, idx int) {
	phase := s.networkPhase
	col := s.ink()
	k := r.W / 100
	at := func(x, y float32) (float32, float32) { return r.X + x*k, r.Y + y*k }

	for layer := 0; layer < 2; layer++ {
		for from := 0; from < 3; from++ {
			for to := 0; to < 3; to++ {
				x1, y1 := at(25+float32(layer)*25, 25+float32(from)*25)
				x2, y2 := at(50+float32(layer)*25, 25+float32(to)*25)
				var line ui.Path
				pt.StrokePath(line.MoveTo(x1, y1).LineTo(x2, y2), 0.5*k, col.Alpha(0.3))
			}
		}
	}
	i := 0
	for layer := 0; layer < 3; layer++ {
		for row := 0; row < 3; row++ {
			a := spinKeyframe(math.Mod(phase-0.2*float64(i), 1), 0.3, 1)
			x, y := at(25+float32(layer)*25, 25+float32(row)*25)
			var dot ui.Path
			pt.FillPath(dot.Circle(x, y, max(6, 2/k)*k/2), col.Alpha(a))
			i++
		}
	}
}

// spinKeyframe goes from rest to peak and back along an easing.
func spinKeyframe(phase float64, rest, peak float32) float32 {
	if phase < 0 {
		phase += 1
	}
	if phase < 0.5 {
		return rest + (peak-rest)*ui.EaseInOut(float32(phase*2))
	}
	return peak + (rest-peak)*ui.EaseInOut(float32(phase*2-1))
}
