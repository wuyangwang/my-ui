package main

import (
	"github.com/egoist/mygo/ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
)

// editableTextPage demonstrates the EditableText component.
func (s *showcase) editableTextPage(c *ui.Context, t *ui.Theme) {

	s.sectionNote(c, t, "就地编辑", "像 Finder 里重命名文件名一样", func() {
		s.card(c, t, func() { formspkg.NewEditableText(c, &s.editableText) })
	})
}
