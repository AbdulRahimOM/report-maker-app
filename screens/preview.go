package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showReportPreview(report *string) {
	prevWindow := MyApp.NewWindow("Report Preview")
	prevWindow.Resize(fyne.NewSize(400, 500))

	reportAsLabel := widget.NewLabel(*report)
	backButton := widget.NewButton("Back", func() {
		prevWindow.Close()
	})

	// Create a container for the label and button
	reportContainer := container.NewVBox(
		reportAsLabel,
		backButton,
	)

	// Set the content of the window to the container
	prevWindow.SetContent(reportContainer)
	prevWindow.Show()

}
