package left

import (
	"bibles/db"
	"bibles/models"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

type Leftside struct {
	//Zhang fyne.CanvasObject
}

var (
	data *gorm.DB
	//Zhang fyne.CanvasObject
)

func init() {
	data = db.DB

}
func Newleftside() Leftside {
	return Leftside{}
}

func (l *Leftside) Drawchapter_old() fyne.CanvasObject {
	var chap []models.Chapname

	data.Where("book<40").Find(&chap)
	var Icons []fyne.CanvasObject
	for _, val := range chap {
		//str_ch := strconv.Itoa(i + 1)

		oneicon := widget.NewButton(val.Jianchen, func() {
			clickId := val.Book
			l.Drawzhangjie(clickId)
			log.Println(clickId, val.Bookname)

		})
		Icons = append(Icons, oneicon)
	}
	cc := container.New(layout.NewGridLayout(4), Icons...)

	// l.Drawbookname(cc, zhang)
	return cc
}

func (l *Leftside) Drawzhang() fyne.CanvasObject {
	book := 1

	var chapn models.Chapname
	data.Where("book=?", book).Find(&chapn)

	var lab_str []string
	for i := 1; i <= chapn.Chnum; i++ {
		lab := strconv.Itoa(i)
		lab_str = append(lab_str, lab)
	}

	zhangjie.Set(lab_str)
	list := widget.NewListWithData(zhangjie, func() fyne.CanvasObject {
		return widget.NewButton("", func() {
			log.Println("good zhangjie")
		})
	}, func(di binding.DataItem, co fyne.CanvasObject) {
		txt, _ := di.(binding.String).Get()
		co.(*widget.Button).SetText(txt)
	})
	zhang_con := container.NewBorder(nil, nil, nil, nil, list)
	return zhang_con
}
func (l *Leftside) Drawchap_new() fyne.CanvasObject {
	var chap []models.Chapname

	data.Where("book>?", 39).Find(&chap)
	var Icons []fyne.CanvasObject
	for _, val := range chap {

		//str_ch := strconv.Itoa(i + 1)
		oneicon := widget.NewButton(val.Jianchen, func() {
			clickId := val.Book
			l.Drawzhangjie(clickId)
			log.Println(clickId, val.Bookname)

		})
		Icons = append(Icons, oneicon)
	}
	// content_zhang := l.Drawbookname(zhang)
	cc := container.New(layout.NewGridLayout(4), Icons...)

	return cc
}

func (l *Leftside) Drawbook() fyne.CanvasObject {

	title_old := widget.NewLabel("旧约")
	old_book := l.Drawchapter_old()
	title_new := widget.NewLabel("新约")
	new_book := l.Drawchap_new()
	cc := container.NewVBox(title_old, old_book, title_new, new_book)
	// var x fyne.CanvasObject
	// if Zhang != nil {
	// 	x = Zhang
	// }
	// c1 := container.NewVBox(x)
	//c1 := container.NewVBox(l.Zhang)
	zhj := l.Drawzhang()
	z1 := container.NewBorder(zhj, nil, nil, nil, nil)
	rs := l.Drawtabs(cc, z1)
	return rs
}

// func (l *Leftside) Drawbook_old() fyne.CanvasObject {

// 	title_old := widget.NewLabel("旧约")
// 	old_book := l.Drawchapter_old()
// 	// title_new := widget.NewLabel("新约")
// 	// new_book := l.Drawchap_new()
// 	cc := container.NewVBox(title_old, old_book)
// 	// var x fyne.CanvasObject
// 	// if Zhang != nil {
// 	// 	x = Zhang
// 	// }
// 	// c1 := container.NewVBox(x)
// 	// c1 := container.NewVBox(l.Zhang)
// 	z1 := container.NewBorder(l.Zhang, nil, nil, nil, nil)
// 	rs := l.Drawtabs(cc, z1)
// 	return rs
// }

// func (l *Leftside)Drawall()fyne.CanvasObject{
// 	container.NewVBox()
// }
