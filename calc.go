package main

import (
	"strconv"
	// "fmt"
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Calculator")
	output:=""
	input := widget.NewLabel(output)

	historyStr:=""
	history:=widget.NewLabel(historyStr);
	 var historyarr []string;
	 ifhis:=false;
	historybtn:=widget.NewButton("history",func(){
		 
			 if ifhis{
				historyStr="";
			 }else{
				 for i:=len(historyarr)-1;i>=0;i--{
            	historyStr=historyStr+historyarr[i];
				historyStr+="\n";
			 }
			
		 }
		 ifhis=!ifhis
		 history.SetText(historyStr);
	})

	Backbtn:=widget.NewButton("back",func(){
		if len(output)>0{
			output=output[:len(output)-1];
		input.SetText(output);
		}
		
	})

	Clearbtn:=widget.NewButton("clear",func(){
		output=""
		input.SetText(output);
	})
	Openbtn:=widget.NewButton("(",func(){
		output=output+"("
		input.SetText(output);
	})
    Closebtn:=widget.NewButton(")",func(){
		output=output+")"
		input.SetText(output);
	})
	Dividebtn:=widget.NewButton("/",func(){
		output=output+"/"
		input.SetText(output);
	})
	Sevenbtn:=widget.NewButton("7",func(){
		output=output+"7"
		input.SetText(output);
	})
	Eightbtn:=widget.NewButton("8",func(){
		output=output+"8"
		input.SetText(output);
	})
	Ninebtn:=widget.NewButton("9",func(){
		output=output+"9"
		input.SetText(output);
	})
	mulbtn:=widget.NewButton("*",func(){
		output=output+"*"
		input.SetText(output);
	})


	fourbtn:=widget.NewButton("4",func(){
		output=output+"4"
		input.SetText(output);
	})
	fivebtn:=widget.NewButton("5",func(){
		output=output+"5"
		input.SetText(output);
	})
	sixbtn:=widget.NewButton("6",func(){
		output=output+"6"
		input.SetText(output);
	})
	minusbtn:=widget.NewButton("-",func(){
		output=output+"-"
		input.SetText(output);
	})

	onebtn:=widget.NewButton("1",func(){
		output=output+"1"
		input.SetText(output);
	})
	twobtn:=widget.NewButton("2",func(){
		output=output+"2"
		input.SetText(output);
	})
	threebtn:=widget.NewButton("3",func(){
		output=output+"3"
		input.SetText(output);
	})
	plusbtn:=widget.NewButton("+",func(){
		output=output+"+"
		input.SetText(output);
	})

	zerobtn:=widget.NewButton("0",func(){
		output=output+"0"
		input.SetText(output);
	})
	dotbtn:=widget.NewButton(".",func(){
		output=output+"."
		input.SetText(output);
	})
	equalbtn:=widget.NewButton("=",func(){
		expression, err := govaluate.NewEvaluableExpression(output);
		if err==nil{
			result, err := expression.Evaluate(nil);
			if err==nil{
				ans:=strconv.FormatFloat(result.(float64),'f',-1,64);
				StringTobeAppended:=output+"="+ans;
				historyarr=append(historyarr,StringTobeAppended);
				output=ans;

			}else{
				output="error";
			}
		}else{
			output="error";
		}
		input.SetText(output);
		
	})

	


	
	 calcContainer:=container.NewVBox(
		 
		input,
		history,
		container.NewGridWithColumns(1,
		  container.NewGridWithColumns(2,
		  historybtn,
		  Backbtn,
		),
		container.NewGridWithColumns(4,
			Clearbtn,
			Openbtn,
			Closebtn,
			Dividebtn,
		),
		container.NewGridWithColumns(4,
			Sevenbtn,
			Eightbtn,
			Ninebtn,
			mulbtn,
		),

		container.NewGridWithColumns(4,
			fourbtn,
		    fivebtn,
			sixbtn,
			minusbtn,
		),
		container.NewGridWithColumns(4,
			onebtn,
			twobtn,
			threebtn,
			plusbtn,
		),
		container.NewGridWithColumns(2,
		container.NewGridWithColumns(2,
			zerobtn,
			dotbtn,),
			equalbtn,
		),

		

		),
		
	)

	// w=myApp.NewWindow("Calc");
	// w.Resize(fyne.NewSize(500,280));

	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,calcContainer),
		
	)
	w.Show()
}