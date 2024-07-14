package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	MyApp    = app.New()
	MyWindow = MyApp.NewWindow("Report maker")

	splitScreen *container.Split
	sidebar     *fyne.Container
	spacer      = widget.NewLabel("")
)

func init() {
	// Window settings
	MyWindow.Resize(fyne.NewSize(700, 500))

	// Sidebar
	sidebar = container.NewVBox(
		widget.NewButton("Session report", func() {
			makeSessionReportScreen()
		}),
		widget.NewButton("Audio report", func() {
			makeAudioReport()
		}),
		widget.NewButton("Meet schedule", func() {
			linkScheduleScreen()
		}),

		widget.NewButton("Change details", func() {
			changeParticularsScreen()
		}),
	)

	// Load home screen
	setHomeScreen()
}

// load different screens
func loadContentScreen(titleAdOn string, mainContent *fyne.Container) {
	scrollContainer := container.NewVScroll(mainContent)
	splitScreen = container.NewHSplit(sidebar, scrollContainer)
	splitScreen.SetOffset(0.3)

	MyWindow.SetContent(splitScreen)
	newTitle := "Report maker"
	if titleAdOn != "" {
		newTitle += " - " + titleAdOn
	}
	MyWindow.SetTitle(newTitle)

}
