package screens

import (
	"time"

	"github.com/AbdulRahimOM/report-maker-app/data"
	"github.com/AbdulRahimOM/report-maker-app/generateReport"
	"github.com/AbdulRahimOM/report-maker-app/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var submissions []bool = make([]bool, len(data.DefaultData.Members))

func makeAudioReport() {
	var audioReportScreenContent *fyne.Container

	headingLabel := widget.NewLabel("Audio Report")
	headingLabel.TextStyle.Bold = true

	dateLabel := widget.NewLabel("Date:")
	dateEntry := widget.NewEntry()
	dateEntry.SetText(time.Now().Format("January 02, 2006"))

	topicLabel := widget.NewLabel("Topic:")
	topicEntry := widget.NewEntry()
	topicEntry.SetPlaceHolder("enter topic here (optional)")

	reportCreatedByLabel := widget.NewLabel("Reported by:")
	reportCreatedByEntry := widget.NewEntry()
	reportCreatedByEntry.SetPlaceHolder("enter name here")
	reportCreatedByEntry.SetText(data.DefaultData.MainCord)

	labelSet := container.NewVBox(
		dateLabel,
		topicLabel,
		reportCreatedByLabel,
	)

	inputSet := container.NewVBox(
		dateEntry,
		topicEntry,
		reportCreatedByEntry,
	)

	localSplit := container.NewHSplit(labelSet, inputSet)
	localSplit.SetOffset(0.3)

	submissionsLabel := widget.NewLabel("Submissions:")
	submissionsLabel.TextStyle.Bold = true
	checkboxContainerL := container.NewVBox()
	checkboxContainerR := container.NewVBox()
	strength := len(data.Batch.Members)
	checkBoxes := make([]*widget.Check, strength)
	for i := 0; i <= strength/2; i++ {
		checkBoxes[i] = widget.NewCheck(data.Batch.Members[i], func(checked bool) {
			submissions[i] = checked
		})
		checkboxContainerL.Add(checkBoxes[i])
	}
	for i := strength/2 + 1; i < strength; i++ {
		checkBoxes[i] = widget.NewCheck(data.Batch.Members[i], func(checked bool) {
			submissions[i] = checked
		})
		checkboxContainerR.Add(checkBoxes[i])
	}

	var toggleAllButton *widget.Button
	toggleAllButton = widget.NewButton("Select all", func() {
		checkMark := !selectionToggle
		selectionToggle = checkMark
		for i := 0; i < len(data.Batch.Members); i++ {
			checkBoxes[i].SetChecked(checkMark)
		}
		if checkMark {
			toggleAllButton.SetText("Deselect all")
		} else {
			toggleAllButton.SetText("Select all")
		}
	})

	checkboxContainer := container.NewHBox(checkboxContainerL, checkboxContainerR)

	submitStatusLabel := widget.NewLabel("")

	submitButton := widget.NewButton("Create report", func() {
		date := dateEntry.Text
		topic := topicEntry.Text
		reportedBy := reportCreatedByEntry.Text

		audioReport := generateReport.AudioReportData{
			DateText:    date,
			Topic:       topic,
			Submissions: submissions,
			ReportedBy:  reportedBy,
		}

		report := generateReport.CreateAudioReport(data.DefaultData, audioReport)
		// tools.LogReport("Audio Report", report)
		tools.CopyToClipboard(report)
		
		// Go back to the home screen
		submitStatusLabel.SetText("Report created successfully")
		go func() {
			time.Sleep(3 * time.Second)
			submitStatusLabel.SetText("")
			setHomeScreen()
		}()

	})

	audioReportScreenContent = container.NewVBox(
		headingLabel,
		localSplit,
		spacer,
		container.NewHBox(submissionsLabel, toggleAllButton),
		checkboxContainer,
		submitButton,
		submitStatusLabel,
	)

	loadContentScreen("Audio report", audioReportScreenContent)
}
