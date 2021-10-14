package main

import (
	//"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello, World!"))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(550,550))

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Sideloader",theme.HomeIcon(), widget.NewLabel("Hello")),
		container.NewTabItemWithIcon("Games",theme.ListIcon(), widget.NewLabel("World!")),
		)

	tabs.Append(container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("Settings Text")))
	tabs.Append(container.NewTabItemWithIcon("About", theme.HelpIcon(), widget.NewLabel("here is what we got so far")))

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)

	myWindow.ShowAndRun()
}
