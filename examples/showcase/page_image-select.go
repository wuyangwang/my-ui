package main

import (
	"image"
	"image/color"

	"github.com/egoist/mygo/ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
	"github.com/wuyangwang/my-ui/icons"
)

// imageSelectPage shows a theme picker of picture tiles. The chosen value
// lives in a ui.Local of a hidden holder.
func imagePreview(dark bool) ui.ImageSource {
	img := image.NewRGBA(image.Rect(0, 0, 48, 28))
	c := color.RGBA{R: 0xf8, G: 0xf8, B: 0xf8, A: 0xff}
	if dark {
		c = color.RGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xff}
	}
	for y := 0; y < 28; y++ {
		for x := 0; x < 48; x++ {
			img.SetRGBA(x, y, c)
		}
	}
	return ui.NewBitmap(img)
}

func (s *showcase) imageSelectPage(c *ui.Context, t *ui.Theme) {
	holder := ui.Box(c.Key("image-select-demo")).Hidden(true)
	selected := ui.Local(holder, "image-select-demo", func() string { return "auto" })
	if selected == nil {
		return
	}
	s.sectionNote(c, t, "ImageSelect", "图片单选：点选或方向键切换，选中的图片有随主题变色的圈", func() {
		s.card(c, t, func() {
			formspkg.NewImageSelect(c, selected, formspkg.ImageSelectProps[string]{
				MinTileWidth: 160,
				Options: []formspkg.ImageSelectItem[string]{
					{Value: "light", Label: "浅色", Icon: icons.Sun, Image: imagePreview(false)},
					{Value: "dark", Label: "深色", Icon: icons.Moon, Image: imagePreview(true)},
					{Value: "auto", Label: "跟随系统", Icon: icons.Monitor, Image: imagePreview(false)},
				},
			})
		})
	})
}
