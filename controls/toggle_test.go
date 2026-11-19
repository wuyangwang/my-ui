package controls

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

func TestToggleFlipsAndReportsThroughHandle(t *testing.T) {
	on := false
	changedCount := 0
	handled := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		tg := NewToggle(c, &on, "粗体").Primary()
		if tg.Changed() {
			changedCount++
		}
		tg.Handle(func(v bool) {
			handled++
			if !v {
				t.Error("Handle got the old value")
			}
		})
	}
	tt := ui.NewTester(view, 320, 120)
	if _, ok := tt.Find("粗体"); !ok {
		t.Fatal("the toggle label is missing")
	}
	if err := tt.Click("粗体"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(20 * time.Millisecond)
	if !on {
		t.Error("the click did not set the bound bool")
	}
	if changedCount == 0 {
		t.Error("Changed never reported the flip")
	}
	if handled == 0 {
		t.Error("the registered Handle callback never fired")
	}
}

func TestToggleOnChangedFiresInTheClickFrame(t *testing.T) {
	on := false
	calls := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewToggle(c, &on, "开关").Sm().OnChanged(func(v bool) {
			calls++
			if !v {
				t.Error("OnChanged got the old value")
			}
		})
	}
	tt := ui.NewTester(view, 320, 120)
	if err := tt.Click("开关"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(20 * time.Millisecond)
	if calls == 0 {
		t.Error("OnChanged did not fire in the click frame")
	}
}

func TestToggleDisabledDoesNotChangeTheValue(t *testing.T) {
	on := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewToggle(c, &on, "禁用").Disabled(true)
	}
	tt := ui.NewTester(view, 320, 120)
	if err := tt.Click("禁用"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(20 * time.Millisecond)
	if on {
		t.Error("a disabled toggle changed its value")
	}
}

func TestToggleWithIconRendersInBothThemes(t *testing.T) {
	for _, a := range []Appearance{AppearanceLight, AppearanceDark} {
		on := true
		view := func(c *ui.Context) {
			Apply(c, Options{Appearance: a})
			NewToggle(c, &on, "通知").Icon(icons.Bell()).Label("通知开关")
		}
		tt := ui.NewTester(view, 320, 120)
		if _, ok := tt.Find("通知"); !ok {
			t.Errorf("%v: the toggle label is missing", a)
		}
		if info, ok := tt.Access("通知开关"); !ok || info.Disabled {
			t.Errorf("%v: the toggle is not an enabled named control", a)
		}
	}
}
