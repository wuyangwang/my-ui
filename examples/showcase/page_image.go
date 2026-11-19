package main

import (
	"image"
	"image/color"

	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

func showcaseImage() *ui.Bitmap {
	img := image.NewRGBA(image.Rect(0, 0, 64, 44))
	for y := 0; y < 44; y++ {
		for x := 0; x < 64; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x * 4), G: uint8(y * 5), B: 180, A: 255})
		}
	}
	return ui.NewBitmap(img)
}

// imagePage shows the image frame, fit and states.
func (s *showcase) imagePage(c *ui.Context, t *ui.Theme) {
	src := showcaseImage()
	s.sectionNote(c, t, "预览", "点击图片打开覆盖窗口的预览", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).Children(func() {
				displaypkg.NewImage(c, src, displaypkg.ImageOptions{Size: 120, Alt: "cover"})
				displaypkg.NewImage(c, src, displaypkg.ImageOptions{
					Width: 120, Height: 80, Contain: true, Alt: "contained",
				})
			})
		})
	})
	s.sectionNote(c, t, "状态与边框", "加载、失败与变体", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).Children(func() {
				displaypkg.NewImage(c, nil, displaypkg.ImageOptions{Size: 96, Loading: true, Alt: "loading"})
				displaypkg.NewImage(c, nil, displaypkg.ImageOptions{Size: 96, Failed: true, Alt: "failed"})
				displaypkg.NewImage(c, src, displaypkg.ImageOptions{
					Size: 96, Variant: myui.Outline, NoPreview: true, Alt: "outlined",
				})
			})
		})
	})
}
