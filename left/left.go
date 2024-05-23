package left

import (
	"bibles/db"
	"bibles/left"
	"bibles/models"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

var clickId int = 0

type Leftside struct {
}

var data *gorm.DB

func init() {
	data = db.DB
}
func Newleftside() Leftside {
	return Leftside{}
}

func (l *Leftside) Drawchapter_old() fyne.CanvasObject {
	var chap []models.Chapname

	data.Find(&chap)
	var Icons []fyne.CanvasObject
	for i, val := range chap {
		//str_ch := strconv.Itoa(i + 1)
		if i >= 40 {
			break
		}
		oneicon := widget.NewButton(val.Bookname, func() {
			log.Println(val.Jianchen, val.Bookname)
			clickId = val.Id
		})
		Icons = append(Icons, oneicon)
	}

	cc := container.New(layout.NewGridLayout(4), Icons...)
	return cc
}
func (l *Leftside) Drawchap_new() fyne.CanvasObject {
	var chap []models.Chapname

	data.Find(&chap)
	var Icons []fyne.CanvasObject
	for i, val := range chap {
		if i < 40 {
			continue
		}
		//str_ch := strconv.Itoa(i + 1)
		oneicon := widget.NewButton(val.Bookname, func() {
			clickId = val.Id
			log.Println(val.Jianchen, val.Bookname)
			lzj := left.Left_zhangjie
			lzj.Drawzhangjie()
		})
		Icons = append(Icons, oneicon)
	}

	cc := container.New(layout.NewGridLayout(4), Icons...)
	return cc
}
func (l *Leftside) Drawbookname() fyne.CanvasObject {

	title_old := widget.NewLabel("旧约")
	old_book := l.Drawchapter_old()
	title_new := widget.NewLabel("新约")
	new_book := l.Drawchap_new()
	cc := container.NewVBox(title_old, old_book, title_new, new_book)
	return cc
}
