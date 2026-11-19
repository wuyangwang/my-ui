package display

import (
	"strings"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// SnippetOptions configures NewSnippet.
type SnippetOptions struct {
	// Language highlights the line with the same small tokenizer as the
	// highlighter.
	Language string
	// Prefix shows before the code and is left out of what Copy copies,
	// as "$" before a command.
	Prefix string
	// Variant picks the surface.
	Variant CodeVariant
	// NoCopy leaves out the copy button.
	NoCopy bool
	// Shadow lifts the snippet.
	Shadow bool
	// FontSize is the code's size in DIPs; the theme's size when zero.
	FontSize float32
	// Label names the snippet for assistive technology.
	Label string
	// OnCopy runs with the copied source (without the prefix).
	OnCopy func(src string)
}

// Snippet shows a single line of code, as a command to run, with a prefix
// and a copy button. Long lines scroll sideways instead of wrapping.
func NewSnippet(c *ui.Context, src string, o SnippetOptions) ui.Element {
	t := c.Theme()
	src = strings.TrimSpace(src)
	root := ui.Row(c).AlignItems(ui.Center).Gap(t.Space(2)).
		Height(t.Space(10)).Padding(0, t.Space(3), 0, t.Space(3)).
		Radius(Radius(t, RadiusMd)).Clip().Role(ui.RoleGroup)
	label := o.Label
	if label == "" {
		label = "Code snippet"
	}
	root.Label(label)
	switch o.Variant {
	case CodeFilled:
		root.Background(t.SurfaceHover)
	case CodeOutlined:
		root.Border(1, t.Border).Background(t.Surface)
	}
	if o.Shadow {
		root.Shadow(0, 1, 2, -1, t.Border.Alpha(0.5))
	}
	shown := src
	if o.Prefix != "" {
		shown = o.Prefix + " " + src
	}
	spans := codeSpans(shown, t)
	root.Children(func() {
		scroll := ui.ScrollHorizontal(c).Grow(1).MinWidth(0).
			Font("monospace").TabSize(4)
		if o.FontSize > 0 {
			scroll.FontSize(o.FontSize)
		}
		scroll.Children(func() {
			text := ui.RichText(c, spans...)
			if o.FontSize > 0 {
				text.FontSize(o.FontSize)
			}
		})
		if !o.NoCopy {
			ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
				SVG: icons.Copy(), IconSize: IconSizeXs, Label: "Copy",
				OnClick: func() {
					c.WriteClipboard(src)
					if o.OnCopy != nil {
						o.OnCopy(src)
					}
				},
			}).Shrink(0)
		}
	})
	return root
}
