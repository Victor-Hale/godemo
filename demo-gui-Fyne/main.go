package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//a := app.New()
	//w := a.NewWindow("客户端")
	//
	//hello := widget.NewLabel("Hello welcome to here ")
	//
	//w.SetContent(container.NewVBox(
	//	hello,
	//	widget.NewButton("enter!", func() {
	//		hello.SetText("Welcome :)")
	//
	//	}),
	//))

	myApp := app.New()
	myWin := myApp.NewWindow("Entry")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")
	nameEntry.OnChanged = func(content string) {
		fmt.Println("name:", nameEntry.Text, "entered")
	}

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("input password")

	nameBox := container.NewHBox(widget.NewLabel("Name"), layout.NewSpacer(), nameEntry)
	passwordBox := container.NewHBox(widget.NewLabel("Password"), layout.NewSpacer(), passEntry)

	loginBtn := widget.NewButton("Login", func() {
		fmt.Println("name:", nameEntry.Text, "password:", passEntry.Text, "login in")
	})

	multiEntry := widget.NewEntry()
	multiEntry.SetPlaceHolder("please enter\nyour description")
	multiEntry.MultiLine = true

	content := container.NewVBox(nameBox, passwordBox, loginBtn, multiEntry)
	myWin.SetContent(content)
	myWin.ShowAndRun()

	//input := widget.NewEntry()
	//input.SetPlaceHolder("Enter text...")
	//content := container.NewVBox(input, widget.NewButton("Save", func() {
	//	log.Println("Content was:", input.Text)
	//}))
	//w.SetContent(content)
	//w.ShowAndRun()
}