package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestSpinSizesAndVariants(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(4)).Children(func() {
			NewSpin(c)
			NewSpin(c).Size(SizeSm).Variant(Success)
			NewSpin(c).Size(SizeLg).Variant(Destructive)
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	// A spinner draws under the framework's motion, without timers in
	// the component: this is a smoke test of the three presets.
	painted := false
	for y := 0; y < 40 && !painted; y++ {
		for x := 0; x < 120; x++ {
			r, g, b := pixelOfDisplay(tt, x, y)
			if !(r == 255 && g == 255 && b == 255) {
				painted = true
				break
			}
		}
	}
	if !painted {
		t.Error("no spinner was drawn")
	}
}

func TestSpinPercentNetworkAndTip(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(4)).Children(func() {
			NewSpin(c).SpinPercent(40).SpinTip("四成")
			NewSpin(c).SpinNetwork().SpinTip("节点")
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	for _, name := range []string{"四成", "节点"} {
		if _, ok := tt.Find(name); !ok {
			t.Errorf("spin tip %q missing", name)
		}
	}
}

func TestSpinContent(t *testing.T) {
	spinning := true
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSpin(c).SpinContent(func() {
			ui.Text(c, "下面的内容")
		}, &spinning).SpinTip("加载中")
	}
	tt := ui.NewTester(view, 320, 160)
	tt.Advance(time.Second)
	if _, ok := tt.Find("下面的内容"); !ok {
		t.Error("the content under the spinner is missing")
	}
	if _, ok := tt.Find("加载中"); !ok {
		t.Error("the tip over the content is missing")
	}
}
