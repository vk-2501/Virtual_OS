package main


import(
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"fmt"
	"math/rand"
)
func dice(w fyne.Window){
	// a:=app.New()

	// w:=a.NewWindow("Roll The Dice")

	//Resize Window
	// w.Resize(fyne.NewSize(800,700))
	
	img:=canvas.NewImageFromFile("E:\\Go_Tut\\MAC OS\\dice.png")
	img.FillMode=canvas.ImageFillOriginal

	btn1:=widget.NewButton("Play",func(){
		//Machine starts from 0 ,so tackle this issue with adding one
		rand:= rand.Intn(6)+1
        img.File=fmt.Sprintf("E:\\Go_Tut\\MAC OS\\dice%d.png",rand)
		img.Refresh()
	})


		dicecont:=container.NewVBox(
			img,
            btn1,
		)
	
		w.SetContent(
			container.NewBorder(DeskBtn,nil,nil,nil,dicecont),
			
		)
	
	//run the app
	w.Show()
}