package display

import (
	"strings"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// CodeVariant names the surface of a code block.
type CodeVariant int

const (
	// CodeFilled tints the block with the theme's surface.
	CodeFilled CodeVariant = iota
	// CodeOutlined draws a border instead.
	CodeOutlined
	// CodeBorderless shows only the code.
	CodeBorderless
)

// HighlighterOptions configures NewHighlighter.
type HighlighterOptions struct {
	// Language is the code's language, by name or extension. It is used by
	// the simplified tokenizer; unknown languages show plain text.
	Language string
	// Variant picks the surface.
	Variant CodeVariant
	// NoCopy leaves out the copy button.
	NoCopy bool
	// HideLanguage leaves out the language tag.
	HideLanguage bool
	// Wrap wraps long lines instead of scrolling them sideways.
	Wrap bool
	// Shadow lifts the block.
	Shadow bool
	// FullFeatured adds a header with the language or file name and a
	// collapse toggle.
	FullFeatured bool
	// FileName replaces the language name in the header.
	FileName string
	// Collapsed is the app's flag for the full-featured header. When nil,
	// the component keeps its own state.
	Collapsed *bool
	// Actions builds extra buttons before Copy.
	Actions func()
	// FontSize is the code's size in DIPs; the theme's size when zero.
	FontSize float32
	// OnCopy runs with the copied source when Copy is pressed.
	OnCopy func(src string)
}

// Highlighter shows a block of code with theme-derived colors, a copy
// button and the language tag. The coloring is a small token pass, not a
// full grammar: see docs/components/highlighter.md for the limits.
func NewHighlighter(c *ui.Context, src string, o HighlighterOptions) ui.Element {
	t := c.Theme()
	src = strings.TrimRight(strings.TrimLeft(src, "\r\n"), " \t\r\n")
	spans := codeSpans(src, t)
	root := ui.Column(c).AlignSelf(ui.Stretch).Radius(Radius(t, RadiusLg)).
		Clip().Role(ui.RoleGroup)
	if o.Language != "" {
		root.Label("Code: " + o.Language)
	}
	switch o.Variant {
	case CodeFilled:
		root.Background(t.SurfaceHover)
	case CodeOutlined:
		root.Border(1, t.Border).Background(t.Surface)
	}
	if o.Shadow {
		root.Shadow(0, 1, 2, -1, t.Border.Alpha(0.5))
	}
	if o.FullFeatured {
		open := o.Collapsed
		if open == nil {
			open = ui.Local(root, "code-open", func() bool { return true })
		}
		name := o.FileName
		if name == "" {
			name = o.Language
		}
		root.Children(func() {
			head := ui.Row(c).AlignItems(ui.Center).Gap(t.Space(1)).
				Padding(t.Space(1)).AlignSelf(ui.Stretch)
			chevron := icons.ChevronDown()
			if !*open {
				chevron = icons.ChevronRight()
			}
			head.Children(func() {
				ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
					SVG: chevron, IconSize: IconSizeXs,
					Label: "Toggle code",
					OnClick: func() {
						*open = !*open
					},
				})
				if name != "" {
					ui.Text(c, name).FontSize(TextSize(t, TextSm)).TextColor(t.Text)
				}
				codeActions(c, src, o)
			})
			if *open {
				codeBody(c, spans, o.Wrap, o.FontSize)
			}
		})
		return root
	}
	root.Children(func() {
		codeBody(c, spans, o.Wrap, o.FontSize)
		if (!o.NoCopy || o.Actions != nil) && (root.Hovered() || root.FocusWithin()) {
			ui.Row(c).Absolute().Top(t.Space(1)).Right(t.Space(1)).
				Gap(t.Space(1)).AlignItems(ui.Center).Children(func() {
				codeActions(c, src, o)
			})
		}
		if !o.HideLanguage && o.Language != "" {
			ui.Text(c, o.Language).Font("monospace").
				FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted).
				Absolute().Bottom(t.Space(1)).Right(t.Space(2))
		}
	})
	return root
}

func codeBody(c *ui.Context, spans []ui.Span, wrap bool, size float32) {
	t := c.Theme()
	var outer ui.Element
	if wrap {
		outer = ui.Column(c)
	} else {
		outer = ui.ScrollHorizontal(c)
	}
	outer.AlignSelf(ui.Stretch).Font("monospace").TabSize(4)
	if size > 0 {
		outer.FontSize(size)
	} else {
		outer.FontSize(TextSize(t, TextSm))
	}
	outer.Children(func() {
		inner := ui.Column(c).Padding(t.Space(3))
		if !wrap {
			inner.Shrink(0)
		}
		inner.Children(func() {
			ui.RichText(c, spans...)
		})
	})
}

func codeActions(c *ui.Context, src string, o HighlighterOptions) {
	if o.Actions != nil {
		o.Actions()
	}
	if o.NoCopy {
		return
	}
	ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
		SVG: icons.Copy(), IconSize: IconSizeXs, Label: "Copy",
		OnClick: func() {
			c.WriteClipboard(src)
			if o.OnCopy != nil {
				o.OnCopy(src)
			}
		},
	})
}

// codeSpans tokenizes src into themed spans. It is intentionally small:
// strings, comments, numbers, keywords and punctuation. Unknown tokens
// take the text's own color, so any language still reads.
func codeSpans(src string, t *ui.Theme) []ui.Span {
	keyword := SemanticColor(t, SemanticInfo)
	str := SemanticColor(t, SemanticSuccess)
	comment := t.TextMuted
	number := SemanticColor(t, SemanticWarning)
	plain := t.Text
	punct := t.TextMuted
	words := codeKeywords(src)
	spans := make([]ui.Span, 0, 32)
	add := func(text string, col ui.Color) {
		if text == "" {
			return
		}
		spans = append(spans, ui.Span{Text: text, Color: col})
	}
	for i := 0; i < len(src); {
		ch := src[i]
		switch {
		case ch == '/' && i+1 < len(src) && src[i+1] == '/':
			j := strings.IndexByte(src[i:], '\n')
			if j < 0 {
				j = len(src)
			} else {
				j += i
			}
			add(src[i:j], comment)
			i = j
		case ch == '#' || (ch == '-' && i+1 < len(src) && src[i+1] == '-'):
			j := strings.IndexByte(src[i:], '\n')
			if j < 0 {
				j = len(src)
			} else {
				j += i
			}
			add(src[i:j], comment)
			i = j
		case ch == '/' && i+1 < len(src) && src[i+1] == '*':
			j := strings.Index(src[i+2:], "*/")
			end := len(src)
			if j >= 0 {
				end = i + 2 + j + 2
			}
			add(src[i:end], comment)
			i = end
		case ch == '"' || ch == '\'' || ch == '`':
			j := i + 1
			for j < len(src) && src[j] != ch {
				if src[j] == '\\' && j+1 < len(src) {
					j++
				}
				j++
			}
			if j < len(src) {
				j++
			}
			add(src[i:j], str)
			i = j
		case ch >= '0' && ch <= '9':
			j := i + 1
			for j < len(src) && (isCodeWord(src[j]) || src[j] == '.') {
				j++
			}
			add(src[i:j], number)
			i = j
		case isCodeWord(ch):
			j := i + 1
			for j < len(src) && isCodeWord(src[j]) {
				j++
			}
			word := src[i:j]
			col := plain
			if words[word] {
				col = keyword
			}
			add(word, col)
			i = j
		default:
			j := i + 1
			add(src[i:j], punct)
			i = j
		}
	}
	return spans
}

func isCodeWord(ch byte) bool {
	return ch == '_' || ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9'
}

// codeKeywords is a shared set for the common languages; language-specific
// grammars are beyond this component's small tokenizer.
func codeKeywords(src string) map[string]bool {
	common := []string{
		"func", "return", "if", "else", "for", "range", "var", "const", "type", "struct",
		"import", "package", "go", "defer", "switch", "case", "break", "continue", "nil",
		"true", "false", "let", "const", "function", "class", "new", "async", "await",
		"export", "default", "from", "interface", "enum", "this", "def", "lambda",
		"pass", "with", "as", "in", "not", "and", "or", "print", "SELECT", "FROM",
		"WHERE", "INSERT", "UPDATE", "DELETE", "CREATE", "TABLE", "true", "false", "null",
	}
	set := make(map[string]bool, len(common))
	for _, w := range common {
		set[w] = true
	}
	return set
}
