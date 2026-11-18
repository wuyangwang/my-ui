package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/forms"
)

// uploadPage demonstrates a drop zone: mygo hands the paths to the
// app, the library only draws the target.
func (s *showcase) uploadPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "拖放上传", "只接收拖入的文件；mygo 给出路径，交给 app 处理", func() {
		s.card(c, t, func() {
			forms.NewUpload(c, forms.UploadOptions{
				OnFiles: func(paths []string) {
					if len(paths) > 0 {
						s.lastAction = paths[0]
					}
				},
			})
			ui.Text(c, "最近拖入："+s.lastAction).TextColor(t.TextMuted)
		})
	})
}
