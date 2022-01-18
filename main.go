package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	rand2 "math/rand"
	"strconv"
	"time"
)

//go:embed img/fu.jpeg
var fuImg []byte

type Theme int

const (
	Do = iota
	Num
)

var TBtn Theme

var musicSingle = make(map[int]string)
var musicSingleDo = make(map[string]string)
var singleCardWget = &widget.Card{}
var btn = make([]*widget.Button, 7)
var randSingleNums = 0
var RandSeed = &rand2.Rand{}

func init() {
	musicSingle[1] = "Do"
	musicSingle[2] = "Ri"
	musicSingle[3] = "Mi"
	musicSingle[4] = "Fa"
	musicSingle[5] = "Suo"
	musicSingle[6] = "La"
	musicSingle[7] = "Xi"
	musicSingleDo["Do"] = "1"
	musicSingleDo["Ri"] = "2"
	musicSingleDo["Mi"] = "3"
	musicSingleDo["Fa"] = "4"
	musicSingleDo["Suo"] = "5"
	musicSingleDo["La"] = "6"
	musicSingleDo["Xi"] = "7"
	TBtn = Do
	RandSeed = rand2.New(rand2.NewSource(time.Now().UnixNano()))
}

func main() {
	myApp := app.New()
	//resource := fyne.NewStaticResource("time_icon", icon)
	//myApp.SetIcon(resource)
	myWindow := myApp.NewWindow("Music Practice")
	grid := container.New(layout.NewGridLayout(1), singleSyllable(), numsBtn())
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(500, 100))
	myWindow.ShowAndRun()
}

func singleSyllable() *fyne.Container {
	singleCardWget = widget.NewCard("", "------------------️", nil)
	singleCardWget.SetImage(&canvas.Image{
		Resource: &fyne.StaticResource{
			StaticName:    "name",
			StaticContent: fuImg,
		},
	})
	refreshCard()
	leftBtn := widget.NewButton("Do", func() {
		if TBtn == Do {
			refreshCard()
			return
		}
		TBtn = Do
		refreshCard()
		for _, obj := range btn {
			obj.SetText(musicSingleDo[obj.Text])
		}
	})
	rightBtn := widget.NewButton("Num", func() {
		if TBtn == Num {
			refreshCard()
			return
		}
		TBtn = Num
		refreshCard()
		for _, obj := range btn {
			n, _ := strconv.Atoi(obj.Text)
			obj.SetText(musicSingle[n])
		}
	})
	grid := container.New(layout.NewGridLayout(3), leftBtn, singleCardWget, rightBtn)
	return grid
}

func numsBtn() *fyne.Container {
	wgets := make([]fyne.CanvasObject, 7)
	for i := 0; i < 7; i++ {
		nums := i + 1
		btn[i] = widget.NewButton(strconv.Itoa(nums), func() {
			fmt.Println(nums, randSingleNums)
			if nums == randSingleNums {
				refreshCard()
			} else {
				singleCardWget.SetSubTitle("fail! try again")
			}
		})
		wgets[i] = btn[i]
	}
	rand2.Seed(time.Now().UnixNano())
	rand2.Shuffle(len(wgets), func(i, j int) {
		wgets[i], wgets[j] = wgets[j], wgets[i]
	})
	return container.New(layout.NewGridLayout(7), wgets...)
}

func refreshCard() {
	singleCardWget.SetSubTitle("-")
	randSingleNums = getMusicSingle()
	if TBtn == Do {
		singleCardWget.SetTitle(musicSingle[randSingleNums])
	} else {
		singleCardWget.SetTitle(strconv.Itoa(randSingleNums))
	}

}

func getMusicSingle() int {
	return RandSeed.Intn(7) + 1
}
