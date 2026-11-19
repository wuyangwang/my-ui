package markdown

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	md "github.com/egoist/mygo/ui/markdown"
	myui "github.com/wuyangwang/my-ui"
)

const sample = "# Title\n\nSome **bold** text with `code`.\n"

func TestNewMarkdownShowsSource(t *testing.T) {
	view := func(c *ui.Context) {
		myui.Apply(c, myui.Options{Appearance: myui.AppearanceLight})
		NewMarkdown(c, sample)
	}
	tt := ui.NewTester(view, 600, 400)
	for _, want := range []string{"Title", "Some bold text with code."} {
		if !tt.HasText(want) {
			t.Errorf("no %q in %q", want, tt.Texts())
		}
	}
}

func TestNewStreamReadsSourceEachFrame(t *testing.T) {
	full := []rune("第一段 **粗体**。\n\n第二段文本。")
	shown := 2
	calls := 0
	view := func(c *ui.Context) {
		myui.Apply(c, myui.Options{Appearance: myui.AppearanceLight})
		NewStream(c, func() string {
			calls++
			return string(full[:shown])
		})
	}
	tt := ui.NewTester(view, 600, 400)
	if calls == 0 {
		t.Fatal("the stream func was never called")
	}
	shown = len(full)
	tt.Frame()
	// The reveal is paced by the framework's clock, advanced here, never
	// slept against.
	for range 100 {
		tt.Advance(20 * time.Millisecond)
		if tt.HasText("第二段文本") {
			break
		}
	}
	if !tt.HasText("第二段文本") {
		t.Errorf("the new text never showed in %q", tt.Texts())
	}
	if calls < 2 {
		t.Errorf("the stream func was called %d times, want at least 2", calls)
	}
}

func TestNewStreamNilFunc(t *testing.T) {
	view := func(c *ui.Context) {
		myui.Apply(c, myui.Options{Appearance: myui.AppearanceLight})
		NewStream(c, nil)
	}
	tt := ui.NewTester(view, 200, 100)
	tt.Advance(20 * time.Millisecond)
	for _, s := range tt.Texts() {
		if s != "" {
			t.Errorf("a nil stream showed %q", s)
		}
	}
}

func TestThemeMapsFrameTheme(t *testing.T) {
	var light, dark *md.Theme
	view := func(c *ui.Context) {
		light = Theme(myui.NewTheme(c, myui.Options{Appearance: myui.AppearanceLight}))
		dark = Theme(myui.NewTheme(c, myui.Options{Appearance: myui.AppearanceDark}))
		NewMarkdown(c, "hi")
	}
	ui.NewTester(view, 200, 100)

	if light == nil || dark == nil {
		t.Fatal("the theme mapping returned nil")
	}
	if light.Text != myui.ShadcnLight().Text || dark.Text != myui.ShadcnDark().Text {
		t.Error("body text does not come from the frame's theme")
	}
	if light.Alerts[0] == dark.Alerts[0] {
		t.Error("alerts do not follow the appearance")
	}
	if light.Mark.A == 0 || light.CodeBlockRadius <= 0 {
		t.Error("marks and code blocks lost their themed metrics")
	}
}

func TestRevealed(t *testing.T) {
	var revealed, foreignRevealed bool
	tt := ui.NewTester(func(c *ui.Context) {
		myui.Apply(c, myui.Options{Appearance: myui.AppearanceLight})
		foreignRevealed = Revealed(ui.Box(c))
		revealed = Revealed(NewMarkdown(c, "hi"))
	}, 200, 100)
	if foreignRevealed {
		t.Error("an unrelated element should not be revealed")
	}
	for range 100 {
		if revealed {
			return
		}
		tt.Advance(20 * time.Millisecond)
	}
	t.Error("a finished Markdown view never revealed")
}
