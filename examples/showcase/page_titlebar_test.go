package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestTitleBarPageShowsBothLooks checks the page lays both looks out:
// the Windows icon-button bar and the Mac traffic-light bar. (Find
// matches a label's full text.)
func TestTitleBarPageShowsBothLooks(t *testing.T) {
	s := &showcase{page: "titlebar"}
	tt := ui.NewTester(s.view, 960, 620)
	for _, what := range []string{
		"自绘标题栏：左侧品牌，右侧窗口按钮，随明暗换色。",
		"Windows 风格（Look: TitleBarWindows）",
		"macOS 风格（Look: TitleBarMac，红绿灯在左）",
	} {
		if _, ok := tt.Find(what); !ok {
			t.Fatalf("the block %q is missing", what)
		}
	}
	// The demo bars carry their own controls; the feedback line waits.
	if _, ok := tt.Access("Close"); !ok {
		t.Fatal("the demo bars drew no close control")
	}
	if _, ok := tt.Find("刚发生："); ok {
		t.Error("the feedback line already speaks before a demo button was pressed")
	}
}
