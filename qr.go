package main

import (
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fmt"
	"strconv"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"

) 

func showqr(w fyne.Window){
	// a:=app.New()

	// w:=a.NewWindow("Qr Code Generator")

	// w.Resize(fyne.NewSize(400,400))

	url:= widget.NewEntry()
	url.SetPlaceHolder("Enter url ....")

	size:= widget.NewEntry()
	size.SetPlaceHolder("Enter file size ....")
	size_1 ,_:=strconv.Atoi(size.Text)

	file_name:= widget.NewEntry()
	file_name.SetPlaceHolder("Enter file name ....")

	btn:=widget.NewButton("Create",func(){
		err1:=qrcode.WriteFile(url.Text,qrcode.Highest,size_1,fmt.Sprintf("%s.png",file_name.Text,),
	)
if err1!=nil{
	fmt.Println(err1)
}

	})

	qrcont:=container.NewVBox(
		url,
		size,
		file_name,
		btn,
	)

	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,qrcont),
		
	)
	w.Show()
}