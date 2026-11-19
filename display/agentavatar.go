package display

import (
	"math"
	"unicode/utf16"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// AgentShape names the silhouette of an agent avatar.
type AgentShape int

const (
	// AgentShapeSeeded takes the shape its seed picks.
	AgentShapeSeeded AgentShape = iota
	AgentFace
	AgentBowl
	AgentCloud
	AgentTile
	AgentHex
	AgentShield
	AgentPudding
	AgentCookie
	AgentStar
	AgentCapsule
	AgentGhost
	AgentCat
)

// AgentColor names the body color of an agent avatar.
type AgentColor int

const (
	// AgentColorSeeded takes the color its seed picks.
	AgentColorSeeded AgentColor = iota
	AgentPurple
	AgentOrange
	AgentBlue
	AgentVolcano
	AgentMagenta
	AgentInk
	AgentGray
)

// AgentState is what the agent is doing, shown by its eyes and pose.
type AgentState int

const (
	AgentIdle AgentState = iota
	AgentWorking
	AgentWaiting
	AgentBlocked
	AgentThinking
	AgentDone
)

// AgentAvatarOptions configures NewAgentAvatar.
type AgentAvatarOptions struct {
	// Seed picks the shape, color and lean; Name when empty. The same seed
	// always gets the same agent.
	Seed string
	// Name names the avatar for assistive technology.
	Name string
	// Shape and Color replace what the seed picks.
	Shape AgentShape
	Color AgentColor
	// CustomColor replaces the color entirely.
	CustomColor ui.Color
	// Size is the width and height in DIPs; 40 when zero.
	Size float32
	// State is what the agent is doing.
	State AgentState
	// NoTilt stands the agent upright.
	NoTilt bool
}

// AgentAvatar draws an agent as a flat silhouette with a pair of big eyes,
// its shape, color and lean seeded by its name, and its state shown by its
// eyes and badge. The drawing is static: no timers or goroutines animate
// it, so it follows the desktop's motion preferences by standing still.
func NewAgentAvatar(c *ui.Context, o AgentAvatarOptions) ui.Element {
	t := c.Theme()
	size := o.Size
	if size <= 0 {
		size = 40
	}
	seed := o.Seed
	if seed == "" {
		seed = o.Name
	}
	h := agentHash("myui:" + seed)
	shape, colorKey := o.Shape, o.Color
	if shape == AgentShapeSeeded {
		shape = AgentShape(h%11) + AgentFace
	}
	if colorKey == AgentColorSeeded {
		colorKey = AgentColor((h>>3)%7) + AgentPurple
	}
	body := o.CustomColor
	if body == (ui.Color{}) {
		body = agentColor(t, colorKey)
	}
	tilt := float32(int((h>>11)%7)-3) * 1.5
	if o.NoTilt {
		tilt = 0
	}
	root := ui.Box(c).Size(size, size).Shrink(0)
	if o.Name != "" {
		root.Role(ui.RoleImage).Label(o.Name)
	}
	small := size <= 24
	state := o.State
	root.Draw(func(p *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2
		p.Transform(ui.Rotation(tilt), cx, cy, func() {
			agentBody(p, r, shape, body, t)
			if !small {
				agentEyes(p, r, state, t)
			}
		})
		if state == AgentBlocked {
			d := r.W * 0.22
			p.Fill(ui.Rect{X: r.X + r.W - d, Y: r.Y, W: d, H: d}, t.Danger, d/2)
		}
	})
	return root
}

// agentHash is a stable FNV-1a over the UTF-16 code units of s.
func agentHash(s string) uint32 {
	h := uint32(2166136261)
	for _, u := range utf16.Encode([]rune(s)) {
		h ^= uint32(u)
		h *= 16777619
	}
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

// agentColor derives the body hues from the frame's theme rather than a
// fixed palette, so a custom theme or dark mode changes the agent too.
func agentColor(t *ui.Theme, c AgentColor) ui.Color {
	switch c {
	case AgentOrange:
		return SemanticColor(t, SemanticWarning)
	case AgentBlue:
		return SemanticColor(t, SemanticInfo)
	case AgentVolcano:
		return t.Danger
	case AgentMagenta:
		return t.Accent.Mix(t.Danger, 0.45)
	case AgentInk:
		return t.Text
	case AgentGray:
		return t.TextMuted
	default:
		return t.Accent
	}
}

func agentBody(p *ui.Painter, r ui.Rect, shape AgentShape, body ui.Color, t *ui.Theme) {
	switch shape {
	case AgentHex, AgentShield, AgentStar, AgentCat:
		p.FillPath(agentPolygon(shape, r), body)
	case AgentCloud:
		p.Fill(ui.Rect{X: r.X + r.W*0.10, Y: r.Y + r.H*0.30, W: r.W * 0.80, H: r.H * 0.55}, body, r.W*0.20)
		p.Fill(ui.Rect{X: r.X + r.W*0.10, Y: r.Y + r.H*0.22, W: r.W * 0.42, H: r.W * 0.42}, body, r.W*0.21)
		p.Fill(ui.Rect{X: r.X + r.W*0.48, Y: r.Y + r.H*0.14, W: r.W * 0.44, H: r.W * 0.44}, body, r.W*0.22)
	default:
		p.Fill(r, body, agentRadius(shape, r.W))
	}
}

func agentRadius(shape AgentShape, side float32) float32 {
	switch shape {
	case AgentTile:
		return side * 0.12
	case AgentBowl, AgentPudding:
		return side * 0.40
	case AgentCapsule, AgentGhost:
		return side * 0.45
	case AgentCookie:
		return side * 0.34
	default:
		return side * 0.30
	}
}

func agentPolygon(shape AgentShape, r ui.Rect) *ui.Path {
	var pts [][2]float32
	cx, cy := r.X+r.W/2, r.Y+r.H/2
	rx, ry := r.W/2*0.95, r.H/2*0.95
	add := func(angle, scale float32) {
		a := float64(angle) * math.Pi / 180
		pts = append(pts, [2]float32{cx + rx*float32(math.Cos(a))*scale, cy + ry*float32(math.Sin(a))*scale})
	}
	switch shape {
	case AgentHex:
		for i := 0; i < 6; i++ {
			add(float32(i*60)+30, 1)
		}
	case AgentShield:
		pts = [][2]float32{
			{cx - rx, cy - ry}, {cx + rx, cy - ry}, {cx + rx, cy + ry*0.2},
			{cx, cy + ry}, {cx - rx, cy + ry*0.2},
		}
	case AgentStar:
		for i := 0; i < 10; i++ {
			s := float32(0.5)
			if i%2 == 0 {
				s = 1
			}
			add(float32(i*36)-90, s)
		}
	case AgentCat:
		pts = [][2]float32{
			{cx - rx, cy - ry*0.2}, {cx - rx*0.9, cy - ry}, {cx - rx*0.4, cy - ry*0.55},
			{cx + rx*0.4, cy - ry*0.55}, {cx + rx*0.9, cy - ry}, {cx + rx, cy - ry*0.2},
			{cx + rx, cy + ry}, {cx - rx, cy + ry},
		}
	}
	path := new(ui.Path)
	for i, pt := range pts {
		if i == 0 {
			path.MoveTo(pt[0], pt[1])
		} else {
			path.LineTo(pt[0], pt[1])
		}
	}
	path.Close()
	return path
}

func agentEyes(p *ui.Painter, r ui.Rect, state AgentState, t *ui.Theme) {
	ink := t.Background
	dx, dy := float32(0), r.H*0.04
	closed := false
	switch state {
	case AgentWorking:
		dy = r.H * 0.12
	case AgentWaiting:
		dx = r.W * 0.05
	case AgentThinking:
		dy = -r.H * 0.08
	case AgentDone:
		closed = true
	}
	ew, eh := r.W*0.13, r.H*0.17
	if state == AgentBlocked {
		ew, eh = r.W*0.16, r.H*0.21
	}
	for _, side := range []float32{-1, 1} {
		x := r.X + r.W*(0.5+side*0.19) - ew/2 + dx
		y := r.Y + r.H*0.40 + dy
		if closed {
			p.Line(x, y+eh/2, x+ew, y+eh/2, max(1, r.W*0.035), ink)
		} else {
			p.Fill(ui.Rect{X: x, Y: y, W: ew, H: eh}, ink, min(ew, eh)/2)
		}
	}
}
