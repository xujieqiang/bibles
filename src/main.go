package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	//"fyne.io/fyne/container"
	"bibles/left"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/golang/freetype/truetype"
)

func init() {
	// fontPath, err := findfont.Find("SIMYOU.TTF")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Found 'arial.ttf' in '%s'\n", fontPath)
	fontPath := "./Fontlib/fonts/Yahei.ttf"

	// load the font with the freetype library
	// 原作者使用的ioutil.ReadFile已经弃用
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}

func main() {
	fmt.Println("hello bibles!")
	bible := app.New()
	//bible.Settings().SetTheme()
	//w := a.NewWindow("圣经")
	mainWindow := bible.NewWindow("圣经")
	mainWindow.SetTitle("圣经")

	text2 := canvas.NewText("password", color.Black)
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
		//container.NewAppTabs()

		container.NewTabItem("cccc", text2),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)
	// var sp container.Split
	// custom_sp := sp.CreateRenderer()

	// custom_sp.Layout()
	nleft := left.Newleftside()
	cc := nleft.Drawbook()

	// cc := container.NewVBox(old, nn)
	content := container.NewBorder(nil, nil, cc, nil, tabs)
	//content := container.NewHBox(lab1, p, lab2, i1, btn, tabs)
	mainWindow.SetContent(content)

	// 窗口设置
	mainWindow.Resize(fyne.NewSize(1200, 1000))

	// mainWindow.SetFixedSize(true)

	mainWindow.Show()

	bible.Run()
	endprogram()

}
func endprogram() {
	fmt.Println("Exit the program!")
}
func updatetime(clock *widget.Label) {
	format := time.Now().Format("Time:03:04:05")
	clock.SetText(format)
}

func Left_side() *fyne.Container {
	t1 := canvas.NewText("1", color.Black)
	t2 := canvas.NewText("2", color.Black)
	b1 := widget.NewButton("3", func() {
		log.Println("good 3")
	})
	c := container.NewHBox(t1, t2, b1)
	c.Resize(fyne.NewSize(200, 800))
	dd := container.NewBorder(c, nil, nil, nil, nil)
	return dd
}

func Widg_label(x, y float32, tt string) *widget.Label {
	c := widget.NewLabel(tt)
	c.Move(fyne.NewPos(x, y))

	return c
}

func Serach_btn(x, y float32, str string, i1 *widget.Entry) *widget.Button {
	b := widget.NewButton(str, func() {
		log.Println("good btn")
		log.Println("Content was:", i1.Text)
	})
	b.Move(fyne.NewPos(x, y))
	b.Resize(fyne.NewSize(70, 35))

	return b
}

func Inputentry(x, y float32) *widget.Entry {
	i := widget.NewEntry()
	i.PlaceHolder = "请输入查询的信息..."
	i.Move(fyne.NewPos(x, y))
	i.Resize(fyne.NewSize(200, 30))
	return i
}
func Passentry(x, y float32) *widget.Entry {
	i := widget.NewPasswordEntry()
	i.PlaceHolder = ""
	i.Move(fyne.NewPos(x, y))
	i.Resize(fyne.NewSize(200, 30))
	return i
}
