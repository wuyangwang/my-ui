package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestSnippetPrefixAndCopy(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewSnippet(c, "go get github.com/wuyangwang/my-ui", SnippetOptions{Language: "sh", Prefix: "$"})
			NewSnippet(c, "bun run dev", SnippetOptions{Language: "sh", Variant: CodeOutlined, NoCopy: true})
		})
	}
	tt := ui.NewTester(view, 440, 180)
	tt.Advance(time.Second)
	if !tt.HasText("$") || !tt.HasText("go get") {
		t.Fatalf("prefixed snippet missing: %v", tt.Texts())
	}
	if !tt.HasText("bun run dev") {
		t.Fatal("outlined snippet missing")
	}
}

func TestSnippetCopyFlag(t *testing.T) {
	copied := ""
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewSnippet(c, "npm test", SnippetOptions{OnCopy: func(src string) { copied = src }})
	}
	tt := ui.NewTester(view, 300, 120)
	tt.Advance(time.Second)
	if err := tt.Click("Copy"); err != nil {
		t.Fatalf("copy: %v", err)
	}
	tt.Advance(time.Second)
	if copied != "npm test" {
		t.Fatalf("copied = %q", copied)
	}
}
