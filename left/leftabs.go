package left

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Lefttabs struct {
}

func Newtabs() Lefttabs {
	return Lefttabs{}
}

var lss Leftside

func (lt *Lefttabs) Drawtabs() fyne.CanvasObject {
	shujuan := lss.Drawbookname()
	if clickId != 0 {
		log.Println(clickId)
	}
	cc := container.NewAppTabs(
		container.NewTabItem("书卷", shujuan),
		container.NewTabItem("章节", canvas.NewText("test", color.Black)),
	)
	return cc

}
