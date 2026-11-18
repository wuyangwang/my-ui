package forms

import (
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
