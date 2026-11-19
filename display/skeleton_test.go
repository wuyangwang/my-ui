package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestSkeletonDrawsThemedBars(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSkeleton(c, 3).Width(200).Size(SizeSm).Animation(SkeletonStatic)
	}
	tt := ui.NewTester(view, 240, 160)
	tt.Advance(time.Second)
	// The bars are the theme's Surface: a pixel inside the first row
	// must differ from the page's white background.
	r, g, b := pixelOfDisplay(tt, 20, 10)
	if r == 255 && g == 255 && b == 255 {
		t.Error("no skeleton bar was painted")
	}
}

func TestSkeletonCircleAndDefaults(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(4)).Children(func() {
			NewSkeleton(c, 0).Circle()
			NewSkeleton(c, 2).Size(SizeLg)
		})
	}
	tt := ui.NewTester(view, 240, 160)
	tt.Advance(time.Second)
	painted := false
	for y := 0; y < 60 && !painted; y++ {
		for x := 0; x < 60; x++ {
			r, g, b := pixelOfDisplay(tt, x, y)
			if !(r == 255 && g == 255 && b == 255) {
				painted = true
				break
			}
		}
	}
	if !painted {
		t.Error("the avatar skeleton was not painted")
	}
}

func TestSkeletonAnimations(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(2)).Children(func() {
			NewSkeleton(c, 3).Animation(SkeletonFade)
			NewSkeleton(c, 3).Animation(SkeletonSweep)
			NewSkeleton(c, 3).Sweep()
			NewSkeleton(c, 1).Static()
		})
	}
	tt := ui.NewTester(view, 480, 160)
	tt.Advance(time.Second)
	// Each mode draws without a timer of its own: the framework's clock
	// advances the Loop, and Tester stays CI-safe.
	painted := false
	for y := 0; y < 60 && !painted; y++ {
		for x := 0; x < 160; x++ {
			r, g, b := pixelOfDisplay(tt, x, y)
			if !(r == 255 && g == 255 && b == 255) {
				painted = true
				break
			}
		}
	}
	if !painted {
		t.Error("no animated skeleton was painted")
	}
}

func pixelOfDisplay(tt *ui.Tester, x, y int) (uint8, uint8, uint8) {
	p := tt.Image().At(x, y)
	r, g, b, _ := p.RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}
