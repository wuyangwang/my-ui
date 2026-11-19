package forms

import (
	"strings"
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestUploadFiltersAndReportsDrops(t *testing.T) {
	got := []string{}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewUpload(c, UploadOptions{
			Accept:  func(path string) bool { return len(path) > 3 && path[len(path)-3:] == ".go" },
			OnFiles: func(paths []string) { got = append(got, paths...) },
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.SetLocale("zh-Hans")
	tt.Advance(time.Second)
	if _, ok := tt.Find("把文件拖到这里"); !ok {
		t.Fatal("the drop zone's label is missing")
	}
	if !tt.DropFiles(160, 80, "a.txt", "b.go") {
		t.Fatal("the drop was not taken")
	}
	tt.Advance(time.Second)
	if len(got) != 1 || got[0] != "b.go" {
		t.Errorf("the accepted files are %v, want [b.go]", got)
	}
}

// TestUploadPickForwardsTheChosenPath checks that setting OnPick adds a
// button, and that a click on it calls the opener and sends the path
// through the same Accept/OnFiles route a drop uses.
func TestUploadPickForwardsTheChosenPath(t *testing.T) {
	got := []string{}
	opens := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewUpload(c, UploadOptions{
			Accept:  func(path string) bool { return strings.HasSuffix(path, ".go") },
			OnFiles: func(paths []string) { got = append(got, paths...) },
			OnPick: func() (string, bool) {
				opens++
				return "picked.go", true
			},
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.SetLocale("zh-Hans")
	tt.Advance(time.Second)

	if _, ok := tt.Find("打开"); !ok {
		t.Fatal("the pick button is missing")
	}
	if err := tt.Click("打开"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if opens != 1 {
		t.Errorf("the opener was called %d times, want 1", opens)
	}
	if len(got) != 1 || got[0] != "picked.go" {
		t.Errorf("the picked file is %v, want [picked.go]", got)
	}
}

// TestUploadPickWithoutAChoiceReportsNothing covers the cancel path:
// OnPick returning ok == false must not call OnFiles.
func TestUploadPickWithoutAChoiceReportsNothing(t *testing.T) {
	got := []string{}
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewUpload(c, UploadOptions{
			OnFiles: func(paths []string) { got = append(got, paths...) },
			OnPick:  func() (string, bool) { return "", false },
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.SetLocale("zh-Hans")
	tt.Advance(time.Second)
	if err := tt.Click("打开"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if len(got) != 0 {
		t.Errorf("a cancelled pick reported %v, want nothing", got)
	}
}
