package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

const highlighterSample = "package main\n\nfunc main() {\n\tprintln(\"hello\")\n}\n"

func TestHighlighterShowsCode(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewHighlighter(c, highlighterSample, HighlighterOptions{Language: "go"})
	}
	tt := ui.NewTester(view, 420, 260)
	tt.Advance(time.Second)
	if !tt.HasText("func") || !tt.HasText("hello") {
		t.Fatalf("code missing: %v", tt.Texts())
	}
}

func TestHighlighterFullFeatured(t *testing.T) {
	open := true
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewHighlighter(c, highlighterSample, HighlighterOptions{
			Language: "go", FullFeatured: true, FileName: "main.go",
			Collapsed: &open, Variant: CodeOutlined,
		})
	}
	tt := ui.NewTester(view, 420, 260)
	tt.Advance(time.Second)
	if !tt.HasText("main.go") || !tt.HasText("func") {
		t.Fatal("full-featured header or body missing")
	}
	if err := tt.Click("Toggle code"); err != nil {
		t.Fatalf("toggle: %v", err)
	}
	tt.Advance(time.Second)
	if open {
		t.Fatal("toggle did not collapse the code")
	}
	if tt.HasText("func") {
		t.Fatal("collapsed code still visible")
	}
}
