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

var selectionToggle bool = false

func makeSessionReportScreen() {
	var sessionReportContent *fyne.Container

	//heading
	headingLabel := widget.NewLabel("Create Session Report")
	headingLabel.TextStyle.Bold = true

	//date
	dateLabel := widget.NewLabel("Date:")
	dateEntry := widget.NewEntry()
	dateEntry.SetText(time.Now().Format("January 02, 2006"))
	dateEntry.SetPlaceHolder("enter date here")

	//time
	timeLabel := widget.NewLabel("Time:")
	timeEntry := widget.NewEntry()
	timeEntry.SetPlaceHolder("enter time here")
	timeEntry.SetText(data.Batch.UsualTime)

	//activity
	activityLabel := widget.NewLabel("Activity:")
	activityEntry := widget.NewEntry()
	activityEntry.SetPlaceHolder("enter activity here(optional)")

	//tldv link
	tldvLinkLabel := widget.NewLabel("TLDV Link:")
	tldvLinkEntry := widget.NewEntry()
	tldvLinkEntry.SetPlaceHolder("enter TLDV link here(optional)")
	tldvPasteInfoLabel := widget.NewLabel("")
	tldvPasteButton := widget.NewButton("Paste", func() {
		text := tools.GetClipboardText()
		if tools.ValidateTldvLink(text) {
			tldvLinkEntry.SetText(text)
		} else {
			tldvPasteInfoLabel.SetText("Invalid link in clipboard!!")
			go func() {
				time.Sleep(3 * time.Second)
				tldvPasteInfoLabel.SetText("")
				loadContentScreen("Create Meet Schedule", sessionReportContent)
			}()
		}
	})

	//reported by
	reportedByLabel := widget.NewLabel("Reported by:")
	reportedByEntry := widget.NewEntry()
	reportedByEntry.SetPlaceHolder("enter name here")
	reportedByEntry.SetText(data.Batch.MainCord)

	//label and input sets (layout)
	labelSet := container.NewVBox(
		dateLabel,
		timeLabel,
		activityLabel,
		reportedByLabel,
	)
	inputSet := container.NewVBox(
		dateEntry,
		timeEntry,
		activityEntry,
		reportedByEntry,
	)
	localSplit := container.NewHSplit(labelSet, inputSet)
	localSplit.SetOffset(0.3)

	//summary
	summaryLabel := widget.NewLabel("Summary:")
	summaryEntry := widget.NewMultiLineEntry()
	summaryEntry.Wrapping = fyne.TextWrapBreak
	summaryEntry.SetPlaceHolder("enter summary here")

	//attendance
	attendanceLabel := widget.NewLabel("Attendance:")
	attendanceLabel.TextStyle.Bold = true
	checkboxContainerL := container.NewVBox()
	checkboxContainerR := container.NewVBox()
	strength := len(data.Batch.Members)
	checkBoxes := make([]*widget.Check, strength)
	{
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
	}
	checkboxContainer := container.NewHBox(checkboxContainerL, checkboxContainerR)

	//select/deselect all button
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

	//preview button
	previewButton := widget.NewButton("Preview", func() {
		//data for preview
		sessionReportData := generateReport.SessionReportData{
			DateText:   dateEntry.Text,
			TimeText:   timeEntry.Text,
			Activity:   activityEntry.Text,
			TLDVLink:   tldvLinkEntry.Text,
			Attendance: submissions,
			ReportedBy: reportedByEntry.Text,
			Summary:    summaryEntry.Text,
		}
		report := generateReport.CreateSessionReport(data.DefaultData, sessionReportData)
		showReportPreview(report)

	})

	//submit button
	submitStatusLabel := widget.NewLabel("")
	submitButton := widget.NewButton("Create report", func() {
		//data for report
		sessionReportData := generateReport.SessionReportData{
			DateText:   dateEntry.Text,
			TimeText:   timeEntry.Text,
			Activity:   activityEntry.Text,
			TLDVLink:   tldvLinkEntry.Text,
			Attendance: submissions,
			ReportedBy: reportedByEntry.Text,
			Summary:    summaryEntry.Text,
		}
		report := generateReport.CreateSessionReport(data.DefaultData, sessionReportData)
		tools.LogReport("Session Report", report)
		tools.CopyToClipboard(report)
		submitStatusLabel.SetText("Report created successfully")
		//resetting the screen
		go func() {
			time.Sleep(4 * time.Second)
			submitStatusLabel.SetText("")
			setHomeScreen()
		}()
	})

	sessionReportContent = container.NewVBox(
		headingLabel,
		localSplit,
		spacer,
		container.NewHBox(attendanceLabel, toggleAllButton),
		checkboxContainer,
		summaryLabel,
		summaryEntry,
		container.NewHBox(tldvLinkLabel, tldvPasteButton, tldvPasteInfoLabel),
		tldvLinkEntry,
		// previewButton,
		// submitButton,
		container.NewHBox(previewButton, submitButton),
		submitStatusLabel,
	)
	loadContentScreen("Create Session Report", sessionReportContent)

}
