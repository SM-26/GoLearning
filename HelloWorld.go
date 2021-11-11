package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "github.com/sm-26/GoLearning/translations"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	
	icon, iconErr := fyne.LoadResourceFromPath(".\\Resources\\RSLicon.png")
	if iconErr != nil {
		log.Fatalln("Error in icon loading")
	}
	myWindow.SetIcon(icon)
	myWindow.SetContent(widget.NewLabel("Hello, World!"))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(550, 550))
	
	// i18n default setup
	local := "en-gb"
	// Declare variable to hold the target language tag.
	var lang language.Tag

	switch local{
	case "en-gb":
		lang = language.MustParse("en-GB")
		lang = language.BritishEnglish
	case "en-us":
		//lang = language.MustParse("en-US")
		lang = language.AmericanEnglish
	case "he-il":
		//lang = language.MustParse("he-IL")
		lang = language.Hebrew
		
	default:
		log.Fatalln("Local not found")
		return
	}

	 // Initialize a message.Printer which uses the target language.
	 p := message.NewPrinter(lang)
	 // Print the welcome message translated into the target language.
	 p.Fprintf(os.Stdout,"The set local is %s\n",lang)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Sideloader", theme.HomeIcon(), widget.NewLabel("Hello")),
		container.NewTabItemWithIcon("Games", theme.ListIcon(), widget.NewLabel("World!")),
	)

	tabs.Append(container.NewTabItemWithIcon("Settings", theme.SettingsIcon(),	widget.NewButton("Set Hebrew", func() {lang = language.Hebrew})))
	tabs.Append(container.NewTabItemWithIcon("About", theme.HelpIcon(), widget.NewLabel("here is what we got so far")))

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()

	
}
