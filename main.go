package main

import (
	// "fyne.io/fyne"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"	
	"fyne.io/fyne/v2/theme"
)

var myApp fyne.App=app.New()

var myWindow fyne.Window=myApp.NewWindow("Mac Os")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var btn7 fyne.Widget



var img fyne.CanvasObject;
var DeskBtn fyne.Widget

var panelContent *fyne.Container


func main(){
	// myApp.Settings().SetTheme(theme.LightTheme())
	img:=canvas.NewImageFromFile("E:\\Go_Tut\\MAC OS\\image.jpg")

	

	btn1=widget.NewButtonWithIcon("Calculator",theme.ContentAddIcon(),func(){
		showCalc(myWindow)
	}) 


	btn2=widget.NewButtonWithIcon("GalleryApp",theme.StorageIcon(),func(){
		showGalleryApp(myWindow)
	}) 

	btn3=widget.NewButtonWithIcon("Text Editor",theme.DocumentIcon(),func(){
		showTextEditor(myWindow)
	})

	btn4=widget.NewButtonWithIcon("PassWord Generator",theme.HelpIcon(),func(){
		genpass(myWindow)
	})
	
	btn5=widget.NewButtonWithIcon("Qr",theme.DocumentIcon(),func(){
		showqr(myWindow)
	})

	btn6=widget.NewButtonWithIcon("Play Game",theme.ComputerIcon(),func(){
		dice(myWindow)
	})
	
	btn7=widget.NewButtonWithIcon("Weather App",theme.InfoIcon(),func(){
		showWeatherApp(myWindow)
	}) 

	DeskBtn=widget.NewButtonWithIcon("Home",theme.HomeIcon(),func(){
		myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	})

	panelContent=container.NewVBox((container.NewGridWithColumns(5,DeskBtn,btn1,btn2,btn3,btn4,btn5,btn6)))

	myWindow.Resize(fyne.NewSize(1280,720))
	myWindow.CenterOnScreen();

	myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img),)
	

	 myWindow.ShowAndRun()
}