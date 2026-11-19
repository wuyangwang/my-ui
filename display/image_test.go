package display

import (
	"image"
	"image/color"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func imageTestSource() *ui.Bitmap {
	img := image.NewRGBA(image.Rect(0, 0, 8, 8))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x * 30), G: uint8(y * 30), B: 128, A: 255})
		}
	}
	return ui.NewBitmap(img)
}

func TestImageStates(t *testing.T) {
	src := imageTestSource()
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(3)).Children(func() {
			NewImage(c, src, ImageOptions{Size: 64, Alt: "picture"})
			NewImage(c, nil, ImageOptions{Size: 64, Loading: true, Alt: "loading"})
			NewImage(c, nil, ImageOptions{Size: 64, Failed: true, Alt: "failed"})
			NewImage(c, src, ImageOptions{Width: 64, Height: 48, Contain: true,
				Variant: Ghost, NoPreview: true, Alt: "contained"})
		})
	}
	tt := ui.NewTester(view, 440, 160)
	tt.Advance(time.Second)
	for _, label := range []string{"picture", "loading", "failed", "contained"} {
		if _, ok := tt.Access(label); !ok {
			t.Fatalf("image %s missing", label)
		}
	}
}

func TestImagePreview(t *testing.T) {
	src := imageTestSource()
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewImage(c, src, ImageOptions{Size: 64, Alt: "open me"})
	}
	tt := ui.NewTester(view, 200, 140)
	tt.Advance(time.Second)
	if err := tt.Click("open me"); err != nil {
		t.Fatalf("click: %v", err)
	}
	tt.Advance(time.Second)
	if tt.Image() == nil {
		t.Fatal("no frame drawn after preview")
	}
}
