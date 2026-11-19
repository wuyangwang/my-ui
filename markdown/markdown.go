// Package markdown wraps mygo's Markdown view with my-ui's theme. It shows
// finished documents and documents still being written — a language model's
// reply — without holding any state of its own: the source and the theme are
// read from the frame, so a subtree theme or the desktop's appearance takes
// effect on the next frame.
package markdown

import (
	"github.com/egoist/mygo/ui"
	md "github.com/egoist/mygo/ui/markdown"
	. "github.com/wuyangwang/my-ui"
)

// Theme maps a my-ui theme onto the Markdown view's theme. Fields it does
// not style are left at zero, so mygo fills them from its own default for
// the same theme; the syntax colors of code blocks keep coming from mygo's
// editor theme. It builds a new value on every call and holds nothing, so
// call it with the frame's theme.
func Theme(t *ui.Theme) *md.Theme {
	if t == nil {
		return nil
	}
	info := SemanticColor(t, SemanticInfo)
	success := SemanticColor(t, SemanticSuccess)
	warning := SemanticColor(t, SemanticWarning)
	th := &md.Theme{
		FontSize:            t.FontSize,
		Font:                t.Font,
		Text:                t.Text,
		Muted:               t.TextMuted,
		Link:                t.Accent,
		Strong:              t.Text,
		Code:                t.Text,
		CodeBackground:      t.Text.Alpha(0.07),
		QuoteBar:            t.Border,
		QuoteText:           t.TextMuted,
		Rule:                t.Border,
		CodeBlockBackground: t.Surface,
		CodeBlockBorder:     t.Border,
		CodeBlockRadius:     Radius(t, RadiusLg),
		TableBorder:         t.Border,
		TableHeader:         t.Surface,
		Mark:                warning.Alpha(0.35),
		Checkbox:            t.Accent,
	}
	// An alert takes my-ui's semantic colors only when the theme supplies
	// them all; otherwise the view uses its own light/dark alert colors.
	if info.A != 0 && success.A != 0 && warning.A != 0 && t.Danger.A != 0 {
		th.Alerts = [5]md.Color{info, success, t.Accent, warning, t.Danger}
	}
	return th
}

// Revealed reports whether an element NewMarkdown or NewStream returned
// this frame shows all of its source, done fading in. A chat uses it to
// show a reply's time and actions once the reply has really arrived, not
// as its last bytes come in. Call it in the same frame, right after the
// view was built; it is false for any other element.
var Revealed = md.Revealed

// NewMarkdown shows source as finished Markdown: CommonMark and mygo's
// default extensions. The text is selectable and copies as text, HTML or
// Markdown. The view is built with the frame's theme; no source or theme
// is kept between frames.
func NewMarkdown(c *ui.Context, source string) ui.Element {
	return md.View(c, source, options(c))
}

// NewStream shows Markdown still being written. stream is read once per
// frame and answers the source accumulated so far; it belongs to the app,
// which replaces it with NewMarkdown when the reply is complete. A nil
// stream shows an empty document. mygo parses only what was appended and
// fades new text in under the desktop's motion rules.
func NewStream(c *ui.Context, stream func() string) ui.Element {
	source := ""
	if stream != nil {
		source = stream()
	}
	return md.Stream(c, source, options(c))
}

// options builds the view options of one frame. The upstream view keeps
// the options it is given, so a fresh value here is what makes the theme
// follow SetTheme and the desktop on the next frame.
func options(c *ui.Context) *md.Options {
	return &md.Options{Theme: Theme(c.Theme())}
}
