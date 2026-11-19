package ai

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

func TestModelPickerShowsGroupsAndChipsWhenOpen(t *testing.T) {
	options := []ModelOption{
		{ID: "a", Name: "Model A", Group: "Vendor One", Context: "128K", Price: "$"},
		{ID: "b", Name: "Model B", Group: "Vendor Two", Speed: "Fast"},
	}
	var text string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewModelPicker(c, &text, options).Outline().Placeholder("搜索模型").Label("选择模型")
	}
	tt := ui.NewTester(view, 420, 300)
	tt.Advance(20 * time.Millisecond)
	if err := tt.Click("选择模型"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(20 * time.Millisecond)
	for _, s := range []string{"Vendor One", "Model A", "128K", "Vendor Two", "Model B", "Fast"} {
		if _, ok := tt.Find(s); !ok {
			t.Errorf("the popup is missing %q", s)
		}
	}
}

func TestModelPickerReportsTheChosenID(t *testing.T) {
	options := []ModelOption{
		{ID: "a", Name: "Alpha"},
		{ID: "b", Name: "Beta"},
	}
	var text string
	var p *ModelPicker
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		p = NewModelPicker(c, &text, options).Label("选择模型")
	}
	tt := ui.NewTester(view, 420, 300)
	tt.Advance(20 * time.Millisecond)
	if err := tt.Click("选择模型"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(20 * time.Millisecond)
	if err := tt.Click("Beta"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(20 * time.Millisecond)
	if got := p.Chosen(); got != "b" {
		t.Errorf("Chosen is %q, want %q", got, "b")
	}
	if sel := p.Selected(); sel == nil || sel.ID != "b" {
		t.Errorf("Selected is %v, want Beta", sel)
	}
	if text != "Beta" {
		t.Errorf("the search field holds %q, want %q", text, "Beta")
	}
}

func TestModelPickerEmptyText(t *testing.T) {
	options := []ModelOption{{ID: "a", Name: "Alpha"}}
	text := "zzz"
	var p *ModelPicker
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceDark})
		p = NewModelPicker(c, &text, options)
		p.EmptyText("没有模型")
		p.Label("选择模型")
	}
	tt := ui.NewTester(view, 420, 300)
	tt.Advance(20 * time.Millisecond)
	if err := tt.Click("选择模型"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("没有模型"); !ok {
		t.Errorf("the empty popup shows no hint; texts: %v", tt.Texts())
	}
}
