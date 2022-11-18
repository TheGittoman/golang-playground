package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const width float32 = 500
const height float32 = 500

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	size := fyne.Size{Width: width, Height: height}
	w.Resize(size)
	cont := container.NewVBox(makeUI())
	cont.Add(clockWidget())
	w.SetContent(cont)

	w.Show()
	a.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("exited")
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func clockWidget() *widget.Label {
	clock := widget.NewLabel("")
	updateTime(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	return clock
}

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello World!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}
