package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// QRStatus is the overlay a QR code can show.
type QRStatus int

const (
	// QRActive shows the code.
	QRActive QRStatus = iota
	// QRLoading covers the code with a spinner icon.
	QRLoading
	// QRExpired covers the code with a message and a refresh button.
	QRExpired
)

// QRCodeOptions configures NewQRCode.
type QRCodeOptions struct {
	// Size is the code's width and height in DIPs, without its padding;
	// 160 when zero.
	Size float32
	// Color colors the modules; the theme's ink when zero.
	Color ui.Color
	// Background fills the card behind them; the theme's background when
	// zero.
	Background ui.Color
	// NoBorder leaves out the hairline around the card.
	NoBorder bool
	// Icon builds a logo over the middle of the code. The modules under
	// it are left out; see the spec for the trade-off.
	Icon func(c *ui.Context)
	// Status covers an expired code or one being fetched.
	Status QRStatus
	// OnRefresh adds a Refresh button to an expired code.
	OnRefresh func()
	// Label names the code for assistive technology; the value when empty.
	Label string
}

type qrState struct {
	value string
	icon  bool
	grid  [][]bool
	hole  int
	drawn struct {
		r     ui.Rect
		scale float32
		path  *ui.Path
	}
}

// NewQRCode shows value as a QR code on a card. The encoder handles byte
// mode, error level L and short values (up to version 5); see
// docs/components/qr-code.md for the limit.
func NewQRCode(c *ui.Context, value string, o QRCodeOptions) ui.Element {
	t := c.Theme()
	size := o.Size
	if size <= 0 {
		size = 160
	}
	ink, paper := o.Color, o.Background
	if ink == (ui.Color{}) {
		ink = t.Text
	}
	if paper == (ui.Color{}) {
		paper = t.Background
	}
	root := ui.Box(c).Padding(t.Space(4)).Radius(Radius(t, RadiusXl)).
		Background(paper).Shrink(0)
	if !o.NoBorder {
		root.Border(1, t.Border)
	}
	st := ui.Local(root, "qr", func() qrState { return qrState{} })
	icon := o.Icon != nil
	if st.grid == nil || st.value != value || st.icon != icon {
		grid := qrEncode(value)
		hole := 0
		if icon && len(grid) > 0 {
			hole = int(float32(len(grid))*0.22+0.999) | 1
		}
		*st = qrState{value: value, icon: icon, grid: grid, hole: hole}
	}
	label := o.Label
	if label == "" {
		label = value
	}
	return root.Children(func() {
		code := ui.Box(c).Size(size, size).Role(ui.RoleImage).Label(label)
		if len(st.grid) > 0 {
			code.Draw(func(p *ui.Painter, r ui.Rect) {
				d := &st.drawn
				if d.path == nil || d.r != r || d.scale != p.Scale() {
					s, unit := p.Scale(), r.W/float32(len(st.grid))
					snap := func(v float32) float32 { return float32(int(v*s+0.5)) / s }
					d.r, d.scale = r, s
					d.path = qrPath(st.grid, st.hole,
						func(k int) float32 { return snap(r.X + float32(k)*unit) },
						func(k int) float32 { return snap(r.Y + float32(k)*unit) })
				}
				p.FillPath(d.path, ink)
			})
		}
		if icon {
			d := size * 0.22
			ui.Box(c).Size(d, d).Absolute().
				Left(t.Space(4) + (size-d)/2).Top(t.Space(4) + (size-d)/2).
				Radius(d / 2).Clip().Background(paper).Children(func() { o.Icon(c) })
		}
		switch o.Status {
		case QRLoading:
			qrOverlay(c, paper, t.Space(4)).Children(func() {
				ui.Icon(c, icons.LoaderCircle()).Size(IconSizeLg, IconSizeLg).TextColor(t.TextMuted)
			})
		case QRExpired:
			qrOverlay(c, paper, t.Space(4)).Children(func() {
				ui.Text(c, "Expired").FontWeight(700).TextColor(t.Text)
				if o.OnRefresh != nil {
					ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
						SVG: icons.RefreshCw(), Label: "Refresh",
						IconSize: IconSizeSm, OnClick: o.OnRefresh,
					})
				}
			})
		}
	})
}

func qrPath(grid [][]bool, hole int, atX, atY func(int) float32) *ui.Path {
	n := len(grid)
	from := (n - hole) / 2
	inHole := func(x, y int) bool {
		return hole > 0 && x >= from && x < from+hole && y >= from && y < from+hole
	}
	path := new(ui.Path)
	for y, row := range grid {
		for x := 0; x < n; {
			if !row[x] || inHole(x, y) {
				x++
				continue
			}
			end := x
			for end < n && row[end] && !inHole(end, y) {
				end++
			}
			path.MoveTo(atX(x), atY(y)).LineTo(atX(end), atY(y)).
				LineTo(atX(end), atY(y+1)).LineTo(atX(x), atY(y+1)).Close()
			x = end
		}
	}
	return path
}

func qrOverlay(c *ui.Context, paper ui.Color, pad float32) ui.Element {
	return ui.Column(c).Absolute().Left(pad).Right(pad).Top(pad).Bottom(pad).
		Gap(c.Theme().Space(2)).AlignItems(ui.Center).Justify(ui.Center).
		Background(paper).TextColor(c.Theme().Text)
}
