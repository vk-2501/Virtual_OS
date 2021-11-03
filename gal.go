package main

import (
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"fmt"
    "io/ioutil"
     "log"
	 "strings"

)

func showGalleryApp(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Hello")
	// w.Resize(fyne.NewSize(1200,800));
	root_src:="C:\\Users\\TEMP\\Pictures";
	files, err := ioutil.ReadDir(root_src);

	if err != nil {
		        log.Fatal(err)
		    }
			
			tabs := container.NewAppTabs()
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extension:= strings.Split(file.Name(), ".")[1];
			if extension == "png" || extension == "jpg" || extension == "jpeg" {
				image := canvas.NewImageFromFile(root_src+"\\"+file.Name());
				image.FillMode = canvas.ImageFillContain
				tabs.Append(container.NewTabItem(file.Name(),image))

			}			
		}
    }		
	tabs.SetTabLocation(container.TabLocationBottom)
	w.SetContent(container.NewBorder(DeskBtn,nil,nil,nil,tabs),)
	w.Show()
}

