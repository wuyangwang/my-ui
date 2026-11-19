package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

// ImageOptions configures NewImage.
type ImageOptions struct {
	// Width and Height size the box; Size sets both. The picture shows at
	// its own size, at most its parent's width, when none is set.
	Width, Height, Size float32
	// Contain fits the whole picture instead of cropping it to cover.
	Contain bool
	// Variant picks the frame: Default tints the box, Outline borders it,
	// Ghost has neither.
	Variant Variant
	// Loading shows a placeholder while there is no picture yet.
	Loading bool
	// Failed shows a placeholder when the picture could not be loaded.
	Failed bool
	// NoPreview keeps a click from opening the picture in a preview.
	NoPreview bool
	// Alt names the picture for assistive technology.
	Alt string
	// Actions builds buttons at the top right, shown while the pointer is
	// over the picture (or always when AlwaysShowActions).
	Actions           func()
	AlwaysShowActions bool
	// OnOpen runs when the preview opens.
	OnOpen func()
	// OnClose runs when the preview closes.
	OnClose func()
}

// Image shows a decoded ui.ImageSource in a rounded frame. A click opens
// a preview over the window; the zoom and pan tools of a full viewer are
// not part of this component (see docs/components/image.md).
func NewImage(c *ui.Context, src ui.ImageSource, o ImageOptions) ui.Element {
	t := c.Theme()
	w, h := o.Width, o.Height
	if w == 0 {
		w = o.Size
	}
	if h == 0 {
		h = o.Size
	}
	if o.Loading {
		return imagePlaceholder(c, w, h, false, t).Label(o.Alt)
	}
	previewable := !o.NoPreview && src != nil && !o.Failed
	var root ui.Element
	if previewable {
		root = ui.ButtonBase(c).Column().Cursor(ui.CursorZoomIn)
	} else {
		root = ui.Box(c)
	}
	root.Radius(Radius(t, RadiusLg)).Clip().Shrink(0).Unselectable()
	if o.Alt != "" {
		root.Label(o.Alt)
	}
	if w > 0 {
		root.Width(w)
	}
	if h > 0 {
		root.Height(h)
	}
	if w == 0 && h == 0 {
		root.MaxWidthPercent(100)
	}
	switch o.Variant {
	case Outline:
		root.Border(1, t.Border).Background(t.Surface)
	case Ghost:
	default:
		root.Border(1, t.Border).Background(t.SurfaceHover)
	}
	open := ui.Local(root, "image-preview", func() bool { return false })
	wasOpen := ui.Local(root, "image-was-open", func() bool { return false })
	preview := previewable && *open
	if previewable && root.Clicked() && !*open {
		*open = true
		if o.OnOpen != nil {
			o.OnOpen()
		}
	}
	root.Children(func() {
		if o.Failed || src == nil {
			imagePlaceholder(c, w, h, true, t)
		} else {
			img := ui.Image(c, src).Fit(ui.Cover)
			if o.Contain {
				img.Fit(ui.Contain)
			}
			if w > 0 || h > 0 {
				img.Fill()
			} else {
				img.MaxWidthPercent(100)
			}
			if preview {
				img.Opacity(0)
			}
		}
		if o.Actions != nil && (o.AlwaysShowActions || root.Hovered() || root.FocusWithin()) {
			ui.Row(c).Absolute().Top(0).Right(0).ZIndex(1).Children(o.Actions)
		}
	})
	if previewable {
		ui.DialogBase(c, open, func(backdrop ui.Element, panel ui.Element) {
			pt := c.Theme()
			backdrop.Background(ui.RGBA(0, 0, 0, 0.8))
			panel.Background(pt.Background).Radius(Radius(pt, RadiusXl)).
				Padding(pt.Space(3)).MaxWidth(pt.Space(160))
			ui.Image(c, src).Fit(ui.Contain).MaxWidthPercent(100).MaxHeight(pt.Space(140))
		})
		// DialogBase clears the flag on Escape or a backdrop click; the
		// callback runs on the frame the flag turns false.
		if *wasOpen && !*open && o.OnClose != nil {
			o.OnClose()
		}
		*wasOpen = *open
	}
	return root
}

func imagePlaceholder(c *ui.Context, w, h float32, failed bool, t *ui.Theme) ui.Element {
	box := ui.Row(c).Justify(ui.Center).AlignItems(ui.Center).
		Background(t.SurfaceHover).TextColor(t.TextMuted)
	if w > 0 || h > 0 {
		box.Fill()
	} else {
		box.Size(128, 128)
	}
	side := min(w, h)
	if side <= 0 {
		side = max(w, h, 128)
	}
	glyph := icons.Image()
	if failed {
		glyph = icons.TriangleAlert()
	}
	box.Children(func() {
		ui.Icon(c, glyph).Size(max(IconSizeLg, side*0.28), max(IconSizeLg, side*0.28))
	})
	return box
}
