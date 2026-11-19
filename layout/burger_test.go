package layout

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

func TestBurgerOpensAndChooses(t *testing.T) {
	open, selected := false, "chat"
	view := func(c *ui.Context) {
		NewBurger(c, &open, BurgerProps{
			Selected: &selected,
			Items: []BurgerItem{
				{Key: "chat", Label: "聊天", Icon: icons.MessageSquare()},
				{Key: "files", Label: "文件", Icon: icons.Folder()},
			},
		})
	}
	tt := ui.NewTester(view, 400, 300)
	tt.Advance(time.Second)
	if err := tt.Click("menu"); err != nil {
		t.Fatalf("the menu trigger did not open: %v", err)
	}
	tt.Advance(time.Second)
	if !tt.HasText("聊天") || !tt.HasText("文件") {
		t.Fatal("the menu rows are not on screen")
	}
	if err := tt.Click("文件"); err != nil {
		t.Fatalf("the row did not choose: %v", err)
	}
	tt.Advance(time.Second)
	if selected != "files" {
		t.Errorf("the choice was %q, want files", selected)
	}
	if open {
		t.Error("the menu did not close after a choice")
	}
}
