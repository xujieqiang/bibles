package left

import (
	"bibles/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/binding"

	"fyne.io/fyne/v2/data/binding"
)

var zhangjie = binding.NewStringList()

func (l *Leftside) Drawzhangjie(id int) fyne.CanvasObject {
	book := id
	if book == 0 {
		return nil
	}
	var chapn models.Chapname
	data.Where("book=?", book).Find(&chapn)

	var lab_str []string
	for i := 1; i <= chapn.Chnum; i++ {
		lab := strconv.Itoa(i)
		lab_str = append(lab_str, lab)
	}
	zhangjie.Set(lab_str)
	// zhangjie.Set(lab_str)
	// list := widget.NewListWithData(zhangjie, func() fyne.CanvasObject {
	// 	return widget.NewButton("", func() {
	// 		log.Println("good zhangjie")
	// 	})
	// }, func(di binding.DataItem, co fyne.CanvasObject) {
	// 	txt, _ := di.(binding.String).Get()
	// 	co.(*widget.Button).SetText(txt)
	// })
	// cc := container.NewBorder(nil, nil, nil, nil, list)
	// return cc
}
