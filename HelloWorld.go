package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello, World!"))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(550, 550))

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Sideloader", theme.HomeIcon(), widget.NewLabel("Hello")),
		container.NewTabItemWithIcon("Games", theme.ListIcon(), widget.NewLabel("World!")),
	)

	tabs.Append(container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("Settings Text")))
	tabs.Append(container.NewTabItemWithIcon("About", theme.HelpIcon(), widget.NewLabel("here is what we got so far")))

	tabs.SetTabLocation(container.TabLocationTop)
	icon, iconErr := fyne.LoadResourceFromPath(".\\Resources\\RSLicon.png")

	if iconErr != nil {
		log.Fatalln("Error in icon loading")
	}
	myWindow.SetContent(tabs)
	myWindow.SetIcon(icon)

	drv := fyne.CurrentApp().Driver()
	w := drv.CreateSplashWindow()
				w.SetContent(widget.NewLabelWithStyle("Super early alpha\n\nYou have been warned!",
					fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
				w.Show()

				go func() {
					time.Sleep(time.Second * 3)
					w.Close()
				}()

	myWindow.ShowAndRun()
}
