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

	// The clear button sits inside the trigger, and the trigger's role
	// keeps it out of the accessibility tree, so the icon is found by the
	// ink it draws: the muted gray of the X, against the dark arrow.
	trigger, ok := tt.Access("尺寸")
	if !ok {
		t.Fatal("the trigger is not on screen")
	}
	muted := ui.Hex("#71717a")
	at := clearIconAt(t, tt, trigger.Bounds, muted)
	if at < 0 {
		t.Fatal("the clear icon is not drawn")
	}
	x := float32(at)
	y := trigger.Bounds.Y + trigger.Bounds.H/2
	tt.Move(x, y)
	tt.Advance(time.Second)
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
	trigger, ok := tt.Access("尺寸")
	if !ok {
		t.Fatal("the trigger is not on screen")
	}
	if at := clearIconAt(t, tt, trigger.Bounds, ui.Hex("#71717a")); at >= 0 {
		t.Fatalf("the clear icon is drawn though the select is not clearable (at x=%v)", at)
	}
}

// clearIconAt is where the muted icon is drawn in a row of the window,
// or -1 when it is nowhere.
func clearIconAt(t *testing.T, tt *ui.Tester, bounds ui.Rect, want ui.Color) int {
	t.Helper()
	for x := int(bounds.X+bounds.W) - 60; x < int(bounds.X+bounds.W)-30; x++ {
		for y := int(bounds.Y); y < int(bounds.Y+bounds.H); y++ {
			r, g, b := testutil.Pixel(tt, float32(x), float32(y))
			if abs(int(r)-int(want.R)) < 12 && abs(int(g)-int(want.G)) < 12 && abs(int(b)-int(want.B)) < 12 {
				return x
			}
		}
	}
	return -1
}

// abs is the size of a difference.
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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
