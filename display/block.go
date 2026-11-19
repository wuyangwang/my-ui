package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// BlockOptions configures NewBlock.
type BlockOptions struct {
	// Variant picks the surface: Default is a tinted fill, Outline a
	// bordered face, Ghost no face at all.
	Variant Variant
	// Clickable makes the block a button; ask the returned element about
	// Clicked or set OnChanged.
	Clickable bool
	// Shadow lifts the block with a soft shadow.
	Shadow bool
	// Label names the block for assistive technology.
	Label string
	// Disabled dims a clickable block and stops its clicks.
	Disabled bool
	// OnChanged runs when a clickable block is clicked.
	OnChanged func()
}

// Block is a rounded surface that holds its own children, as a card-like
// tile in a grid. A clickable block is a button whose colors fade through
// the shared state transition.
//
//	b := display.NewBlock(c, display.BlockOptions{Clickable: true}).
//		Label("Open").Outline()
//	b.Children(func() { ... })
//	if b.Clicked() { ... }
type Block struct {
	ui.Element
	c *ui.Context
	o BlockOptions
}

// NewBlock creates a block with the options o.
func NewBlock(c *ui.Context, o BlockOptions) *Block {
	b := &Block{c: c, o: o}
	var root ui.Element
	if o.Clickable {
		root = ui.ButtonBase(c).Column().AlignItems(ui.Stretch).
			Justify(ui.Start).Cursor(ui.CursorPointer)
	} else {
		root = ui.Box(c)
	}
	root.Radius(Radius(c.Theme(), RadiusLg)).Transition(StateTransition())
	b.Element = root
	b.applyLook()
	if o.Label != "" {
		root.Label(o.Label)
	}
	root.Disabled(o.Disabled)
	if o.Clickable && root.Clicked() && o.OnChanged != nil {
		o.OnChanged()
	}
	return b
}

// applyLook paints the current variant, hover and shadow on the existing
// element, so a fluent call never builds a second element in the frame.
func (b *Block) applyLook() {
	t := b.c.Theme()
	root := b.Element
	root.NoShadow()
	hover := b.o.Clickable && root.Hovered()
	switch b.o.Variant {
	case Outline:
		root.Border(1, t.Border)
		if hover {
			root.Background(t.SurfaceHover)
		} else {
			root.Background(t.Surface)
		}
	case Ghost:
		root.Border(0, ui.Transparent)
		if hover {
			root.Background(t.SurfaceHover)
		} else {
			root.Background(ui.Transparent)
		}
	case Secondary:
		root.Border(0, ui.Transparent)
		root.Background(t.Surface)
	default:
		root.Border(0, ui.Transparent)
		// shadcn's muted surface at rest, its hover shade under the
		// pointer: never a hover color while idle.
		if hover {
			root.Background(t.SurfaceHover)
		} else {
			root.Background(t.Surface)
		}
	}
	if b.o.Shadow {
		// The same soft two-layer shadow as display.NewCard.
		root.Shadow(0, 1, 2, 0, t.Text.Alpha(0.06)).
			Shadow(0, 8, 24, -8, t.Text.Alpha(0.10))
	}
}

// Variant applies the shared variant v and returns the block.
func (b *Block) Variant(v Variant) *Block {
	b.o.Variant = v
	b.applyLook()
	return b
}

// Outline gives the block a border and returns it.
func (b *Block) Outline() *Block { return b.Variant(Outline) }

// Ghost removes the block's face and returns it.
func (b *Block) Ghost() *Block { return b.Variant(Ghost) }

// Shadow lifts the block and returns it.
func (b *Block) Shadow(on bool) *Block {
	b.o.Shadow = on
	b.applyLook()
	return b
}

// Label names the block for assistive technology.
func (b *Block) Label(s string) *Block {
	b.Element.Label(s)
	return b
}

// Disabled dims a clickable block and stops its clicks.
func (b *Block) Disabled(off bool) *Block {
	b.Element.Disabled(off)
	return b
}

// OnChanged runs fn when a clickable block is clicked.
func (b *Block) OnChanged(fn func()) *Block {
	b.o.OnChanged = fn
	if b.o.Clickable && b.Element.Clicked() && fn != nil {
		fn()
	}
	return b
}
