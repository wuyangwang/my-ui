package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// avatarPage shows the avatar's fallbacks, sizes and shapes.
func (s *showcase) avatarPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "回退方式", "纹理按姓名生成，取自当前主题", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				for _, name := range []string{"Ada Lovelace", "Grace Hopper", "Alan Turing", "Lobe"} {
					displaypkg.NewAvatar(c.Key(name), name, displaypkg.AvatarOptions{Size: 44})
				}
			})
		})
	})
	s.sectionNote(c, t, "文字与 emoji", "首字母、完整文字与 emoji", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				displaypkg.NewAvatar(c, " Ada Lovelace", displaypkg.AvatarOptions{
					Size: 44, Fallback: displaypkg.AvatarText,
				})
				displaypkg.NewAvatar(c, "Grace Hopper", displaypkg.AvatarOptions{
					Size: 44, Fallback: displaypkg.AvatarText, Full: true,
				})
				displaypkg.NewAvatar(c, "🤯", displaypkg.AvatarOptions{Size: 44})
			})
		})
	})
	s.sectionNote(c, t, "尺寸与形状", "共享 Size 词表与圆形", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
				displaypkg.NewAvatar(c, "XS", displaypkg.AvatarOptions{Fallback: displaypkg.AvatarText}).Size(myui.SizeXs)
				displaypkg.NewAvatar(c, "SM", displaypkg.AvatarOptions{Fallback: displaypkg.AvatarText}).Size(myui.SizeSm)
				displaypkg.NewAvatar(c, "MD", displaypkg.AvatarOptions{Fallback: displaypkg.AvatarText}).Size(myui.SizeMd)
				displaypkg.NewAvatar(c, "LG", displaypkg.AvatarOptions{Fallback: displaypkg.AvatarText}).Size(myui.SizeLg)
				displaypkg.NewAvatar(c, "Circle", displaypkg.AvatarOptions{Size: 48}).Circle()
			})
		})
	})
}
