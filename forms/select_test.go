package forms

import (
	. "github.com/wuyangwang/my-ui"
	testutil "github.com/wuyangwang/my-ui/internal/testutil"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestSelectShowsTheLabelOfTheChoice(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{
			{Label: "简体中文", Value: "zh-Hans"},
			{Label: "繁體中文", Value: "zh-Hant"},
		}).Outline().Label("语言")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	// Nothing is chosen yet, so the trigger shows no label.
	if _, ok := tt.Find("简体中文"); ok {
		t.Error("the trigger shows a label before a choice is made")
	}
	if _, ok := tt.Access("语言"); !ok {
		t.Fatal("the select is not on screen")
	}
}

func TestSelectOpensAndChooses(t *testing.T) {
	var value string
	var seen string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{
			{Label: "简体中文", Value: "zh-Hans"},
			{Label: "繁體中文", Value: "zh-Hant"},
		}).Outline().Label("语言").OnChanged(func(v string) { seen = v })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)

	if err := tt.Click("语言"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("繁體中文"); !ok {
		t.Fatal("the popup did not open")
	}
	if err := tt.Click("繁體中文"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if value != "zh-Hant" {
		t.Errorf("the select holds %q, want %q", value, "zh-Hant")
	}
	if seen != "zh-Hant" {
		t.Errorf("the change reported %q, want %q", seen, "zh-Hant")
	}
	// The trigger shows the label of the choice, not its value.
	if _, ok := tt.Find("繁體中文"); !ok {
		t.Error("the trigger does not show the label of the choice")
	}
}

func TestSelectPlaceholderShowsWhileNothingIsChosen(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{{Value: "小"}}).Outline().
			Placeholder("选择尺寸").Label("尺寸")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if _, ok := tt.Find("选择尺寸"); !ok {
		t.Error("the placeholder is not on screen")
	}
}

func TestSelectDisabledTakesNoClick(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewSelect(c, &value, []SelectOption{{Value: "小"}}).Outline().
			Label("尺寸").Disabled(true)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("尺寸")
	if !ok {
		t.Fatal("the select is not on screen")
	}
	if !info.Disabled {
		t.Error("the select does not read as disabled")
	}
	if err := tt.Click("尺寸"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("小"); ok {
		t.Error("a disabled select opened its popup")
	}
}

func TestSelectGroupsItsOptions(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{
			{Group: "最近", Label: "英语", Value: "en"},
			{Group: "最近", Label: "日语", Value: "ja"},
			{Group: "更多", Label: "韩语", Value: "ko"},
		}).Outline().Label("语言")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("语言"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	// The name of a group shows above its options.
	if _, ok := tt.Find("最近"); !ok {
		t.Error("the group is not on screen")
	}
	if _, ok := tt.Find("更多"); !ok {
		t.Error("the second group is not on screen")
	}
}

func TestSelectClearEmptiesTheChoice(t *testing.T) {
	var value string
	cleared := false
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{{Value: "小"}, {Value: "大"}}).
			Outline().Label("尺寸").OnCleared(func() { cleared = true })
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("尺寸"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if err := tt.Click("小"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if value != "小" {
		t.Fatalf("the select holds %q, want %q", value, "小")
	}

	// The clear icon takes the chevron's place at the right edge: the
	// trigger's role keeps nested controls out of the accessibility
	// tree, so press the edge of the trigger and read the value.
	trigger, ok := tt.Access("尺寸")
	if !ok {
		t.Fatal("the trigger is not on screen")
	}
	x := trigger.Bounds.X + trigger.Bounds.W - 20
	y := trigger.Bounds.Y + trigger.Bounds.H/2
	tt.Press(x, y)
	tt.Advance(time.Second)
	tt.Release(x, y)
	tt.Advance(time.Second)
	if value != "" {
		t.Errorf("the select holds %q after the clear, want it empty", value)
	}
	if !cleared {
		t.Error("the clear did not report")
	}
}

// TestSelectNotClearable checks that NewSelectClear leaves out the clear
// button when asked to.
func TestSelectNotClearable(t *testing.T) {
	value := "小"
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelectClear(c, &value, []SelectOption{{Value: "小"}, {Value: "大"}}, false).
			Outline().Label("尺寸")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	// Pressing the right edge of a non-clearable trigger must leave the
	// value alone: that spot is the chevron, not a clear icon.
	trigger, ok := tt.Access("尺寸")
	if !ok {
		t.Fatal("the trigger is not on screen")
	}
	tt.Press(trigger.Bounds.X+trigger.Bounds.W-20, trigger.Bounds.Y+trigger.Bounds.H/2)
	tt.Advance(time.Second)
	tt.Release(trigger.Bounds.X+trigger.Bounds.W-20, trigger.Bounds.Y+trigger.Bounds.H/2)
	tt.Advance(time.Second)
	if value != "小" {
		t.Errorf("a non-clearable select lost its value: %q", value)
	}
}

func TestSelectKeepsADefaultWidth(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewSelect(c, &value, []SelectOption{{Value: "小"}}).Outline().Label("尺寸")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("尺寸")
	if !ok {
		t.Fatal("the select is not on screen")
	}
	// The trigger keeps a width of its own: it does not collapse to the
	// width of a word.
	if info.Bounds.W < 120 {
		t.Errorf("the select is %v wide, want a width of its own", info.Bounds.W)
	}
}

func TestSelectSizesGiveRoomierBounds(t *testing.T) {
	height := func(tt *ui.Tester, name string) float32 {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		return info.Bounds.H
	}
	var a, b string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewSelect(c, &a, []SelectOption{{Value: "小"}}).Sm().Outline().Label("small")
			NewSelect(c, &b, []SelectOption{{Value: "中"}}).Md().Outline().Label("medium")
		})
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if h := height(tt, "small"); h >= height(tt, "medium") {
		t.Errorf("the small select is not shorter than the middle's: %v against %v", h, height(tt, "medium"))
	}
}

func TestSelectShowsOptionIconInTrigger(t *testing.T) {
	labelX := func(withIcon bool) float32 {
		value := "on"
		o := SelectOption{Label: "Online", Value: "on"}
		if withIcon {
			o.Icon = testIcon
		}
		tt := ui.NewTester(func(c *ui.Context) {
			Apply(c, Options{Appearance: AppearanceLight})
			NewSelect(c, &value, []SelectOption{o}).Outline().Label("状态")
		}, 320, 240)
		tt.Advance(time.Second)
		r, ok := tt.Find("Online")
		if !ok {
			t.Fatal("the chosen label is not on screen")
		}
		return r.X
	}
	if with, without := labelX(true), labelX(false); with <= without {
		t.Errorf("the icon did not push the label right: with %v, without %v", with, without)
	}
}

func TestSelectShowsOptionIconInPopup(t *testing.T) {
	labelX := func(withIcon bool) float32 {
		var value string
		o := SelectOption{Label: "Online", Value: "on"}
		if withIcon {
			o.Icon = testIcon
		}
		tt := ui.NewTester(func(c *ui.Context) {
			Apply(c, Options{Appearance: AppearanceLight})
			NewSelect(c, &value, []SelectOption{o}).Outline().Label("状态")
		}, 320, 240)
		tt.Advance(time.Second)
		if err := tt.Click("状态"); err != nil {
			t.Fatal(err)
		}
		tt.Advance(time.Second)
		r, ok := tt.Find("Online")
		if !ok {
			t.Fatal("the option is not on screen")
		}
		return r.X
	}
	if with, without := labelX(true), labelX(false); with <= without {
		t.Errorf("the icon did not push the label right: with %v, without %v", with, without)
	}
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestSelectArrowTurnsWhenOpen(t *testing.T) {
	var value string
	var bg [3]uint8
	view := func(c *ui.Context) {
		th := Apply(c, Options{Appearance: AppearanceLight})
		bg = [3]uint8{th.Background.R, th.Background.G, th.Background.B}
		NewSelect(c, &value, []SelectOption{{Value: "小"}}).Outline().Label("状态")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	tr, ok := tt.Access("状态")
	if !ok {
		t.Fatal("the select is not on screen")
	}
	cx := tr.Bounds.X + tr.Bounds.W - 21
	cy := tr.Bounds.Y + tr.Bounds.H/2
	// Read the chevron as a grid of ink dots. Open turns it 180 degrees,
	// so the open grid is the closed one flipped top to bottom: compare
	// the two against a flipped copy rather than against a colour guess.
	sample := func() [11][11]bool {
		var g [11][11]bool
		for dy := -5; dy <= 5; dy++ {
			for dx := -5; dx <= 5; dx++ {
				r, gg, b := testutil.Pixel(tt, cx+float32(dx), cy+float32(dy))
				d := absInt(int(r)-int(bg[0])) + absInt(int(gg)-int(bg[1])) + absInt(int(b)-int(bg[2]))
				g[dy+5][dx+5] = d > 40
			}
		}
		return g
	}
	before := sample()
	if err := tt.Click("状态"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	after := sample()
	same, flipped := 0, 0
	for dy := range before {
		for dx := range before[dy] {
			if before[dy][dx] == after[dy][dx] {
				same++
			}
			if before[dy][dx] == after[10-dy][dx] {
				flipped++
			}
		}
	}
	if flipped <= same {
		t.Errorf("the arrow did not turn: same %d, flipped %d", same, flipped)
	}
}
