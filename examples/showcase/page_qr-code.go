package main

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// qrCodePage shows the QR code's colors, icon and status overlays.
func (s *showcase) qrCodePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "基本", "库内编码，字节模式，纠错级别 L", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				displaypkg.NewQRCode(c, "https://lobehub.com", displaypkg.QRCodeOptions{
					Size: 128,
				})
				displaypkg.NewQRCode(c, "https://mygo.lobehub.com", displaypkg.QRCodeOptions{
					Size: 128, NoBorder: true,
				})
			})
		})
	})
	s.sectionNote(c, t, "状态", "加载与过期覆盖层", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				displaypkg.NewQRCode(c, "https://lobehub.com", displaypkg.QRCodeOptions{
					Size: 120, Status: displaypkg.QRLoading,
				})
				displaypkg.NewQRCode(c, "https://lobehub.com", displaypkg.QRCodeOptions{
					Size: 120, Status: displaypkg.QRExpired,
					OnRefresh: func() { s.lastAction = "QR refreshed" },
				})
			})
		})
	})
}
