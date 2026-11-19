package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestGroupAvatarMembers(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Row(c).Gap(c.Theme().Space(3)).Children(func() {
			NewGroupAvatar(c, GroupAvatarOptions{Size: 48, Avatars: []GroupAvatarItem{
				{Avatar: "Ada", Fallback: AvatarText},
				{Avatar: "Grace", Fallback: AvatarText},
			}})
			NewGroupAvatar(c, GroupAvatarOptions{Size: 48, Corner: GroupCornerCircle,
				Grid: GroupGridThree, Avatars: []GroupAvatarItem{
					{Avatar: "🤯"}, {Avatar: "🚀"}, {Avatar: "🧠"},
					{Avatar: "👻"},
				}})
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	for _, s := range []string{"AD", "GR", "🤯", "🚀", "🧠", "👻"} {
		if !tt.HasText(s) {
			t.Fatalf("member %s missing", s)
		}
	}
}

func TestGroupAvatarSingle(t *testing.T) {
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		NewGroupAvatar(c, GroupAvatarOptions{Size: 40, Avatars: []GroupAvatarItem{
			{Avatar: "Lobe", Fallback: AvatarText},
		}})
	}
	tt := ui.NewTester(view, 160, 120)
	tt.Advance(time.Second)
	if !tt.HasText("LO") {
		t.Fatal("single member missing")
	}
}
