package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// groupAvatarPage shows the group layouts and corners.
func (s *showcase) groupAvatarPage(c *ui.Context, t *ui.Theme) {
	members := func(n int) []displaypkg.GroupAvatarItem {
		all := []displaypkg.GroupAvatarItem{
			{Avatar: "Ada", Fallback: displaypkg.AvatarText},
			{Avatar: "🤯"},
			{Avatar: "Grace", Fallback: displaypkg.AvatarText},
			{Avatar: "🚀"},
			{Avatar: "🧠"},
			{Avatar: "👻"},
		}
		return all[:min(n, len(all))]
	}
	s.sectionNote(c, t, "网格", "两列最多四人，三列最多九人", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				for _, cfg := range []struct {
					name   string
					grid   displaypkg.GroupGrid
					count  int
					corner displaypkg.GroupCorner
				}{
					{"one", displaypkg.GroupGridTwo, 1, displaypkg.GroupCornerSquare},
					{"two", displaypkg.GroupGridTwo, 2, displaypkg.GroupCornerSquare},
					{"four", displaypkg.GroupGridTwo, 4, displaypkg.GroupCornerSquare},
					{"auto", displaypkg.GroupGridAuto, 6, displaypkg.GroupCornerSquare},
					{"circle", displaypkg.GroupGridThree, 4, displaypkg.GroupCornerCircle},
				} {
					ui.Column(c).Gap(t.Space(1)).AlignItems(ui.Center).Children(func() {
						displaypkg.NewGroupAvatar(c, displaypkg.GroupAvatarOptions{
							Size: 48, Grid: cfg.grid, Corner: cfg.corner,
							Avatars: members(cfg.count),
						})
						ui.Text(c, cfg.name).FontSize(11).TextColor(t.TextMuted)
					})
				}
			})
		})
	})
}
