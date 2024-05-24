package left

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (l *Leftside) Drawtabs(one, two fyne.CanvasObject) fyne.CanvasObject {

	if two == nil {
		two = widget.NewLabel("无内容")
	}
	cc := container.NewAppTabs(
		container.NewTabItem("书卷", one),

		container.NewTabItem("章节", two),
	)
	return cc

}
