package display

import (
	"testing"
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestCarouselOneClickMovesOneSlide(t *testing.T) {
	index := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		// The fluent calls used to rebuild the whole component in the
		// same frame, reading the frame's click more than once.
		NewCarousel(c, &index, 3, func(i int) {
			ui.Text(c, []string{"第一张", "第二张", "第三张"}[i])
		}).Arrows(true).Dots(true)
	}
	tt := ui.NewTester(view, 320, 240)
	adv := func() { tt.Advance(time.Second); tt.Advance(time.Second) }
	adv()
	if err := tt.Click("Next page"); err != nil {
		t.Fatal(err)
	}
	adv()
	if index != 1 {
		t.Fatalf("one click advanced to slide %d, want 1", index)
	}
	if _, ok := tt.Find("第二张"); !ok {
		t.Fatal("the second slide is missing")
	}
}

func TestCarouselPagesAndDots(t *testing.T) {
	index := 0
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewCarousel(c, &index, 3, func(i int) {
			ui.Text(c, []string{"第一张", "第二张", "第三张"}[i])
		})
	}
	tt := ui.NewTester(view, 320, 240)
	adv := func() { tt.Advance(time.Second); tt.Advance(time.Second) }
	if _, ok := tt.Find("第一张"); !ok {
		t.Fatal("the first slide is missing")
	}
	if err := tt.Click("Next page"); err != nil {
		t.Fatal(err)
	}
	adv()
	if index != 1 {
		t.Errorf("next gave %d, want 1", index)
	}
	if _, ok := tt.Find("第二张"); !ok {
		t.Fatal("the second slide is missing")
	}
	if err := tt.Click("第 3 张"); err != nil {
		t.Fatal(err)
	}
	adv()
	if index != 2 {
		t.Errorf("the dot gave %d, want 2", index)
	}
}
