package overlays

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestSplitButtonRunsItsAction(t *testing.T) {
	clicked := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSplitButton(c, "保存", func() { clicked = true }, []MenuItem{
			{Label: "另存为", OnClick: func() { clicked = true }},
		})
	}
	tt := ui.NewTester(view, 320, 120)
	tt.Advance(time.Second)
	if err := tt.Click("保存"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !clicked {
		t.Error("the split button's main action did not run")
	}
}
