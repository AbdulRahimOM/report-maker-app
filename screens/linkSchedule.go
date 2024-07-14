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

func linkScheduleScreen() {
	var scheduleScreenContent *fyne.Container

	//heading
	headingLabel := widget.NewLabel("Create Meet Schedule")
	headingLabel.TextStyle.Bold = true

	//date
	dateLabel := widget.NewLabel("Date:")
	dateEntry := widget.NewEntry()
	dateEntry.SetText(time.Now().Format("January 02, 2006"))

	//time slot
	timeSlotLabel := widget.NewLabel("Time")
	timeSlotEntry := widget.NewEntry()
	timeSlotEntry.SetText(data.DefaultData.UsualTime)

	//meeting link
	linkLabel := widget.NewLabel("Link:")
	pasteInfoLabel := widget.NewLabel("")
	linkEntry := widget.NewEntry()
	linkEntry.SetPlaceHolder("enter meeting link here")
	{
		clipboardText := tools.GetClipboardText()
		if tools.ValidateGMeetLink(clipboardText) {
			linkEntry.SetText(clipboardText)
		}
	}

	//paste button
	pasteButton := widget.NewButton("Paste", func() {
		text := tools.GetClipboardText()
		if tools.ValidateGMeetLink(text) {
			linkEntry.SetText(text)
		} else {
			pasteInfoLabel.SetText("Invalid link in clipboard!!")
			go func() {
				time.Sleep(3 * time.Second)
				pasteInfoLabel.SetText("")
				loadContentScreen("Create Meet Schedule", scheduleScreenContent)
			}()
		}

	})

	//submit button
	statusLabel := widget.NewLabel("")
	submitButton := widget.NewButton("Make Link Schedule", func() {
		if tools.ValidateGMeetLink(linkEntry.Text) {
			report:=generateReport.CreateLinkScheduke(data.DefaultData, generateReport.LinkSchedule{
				DateText:    dateEntry.Text,
				MeetingLink: linkEntry.Text,
				MeetingSlot: timeSlotEntry.Text,
			})
			// tools.LogReport("Link Schedule",report)
			tools.CopyToClipboard(report)

			statusLabel.SetText("Schedule copied to clipboard")
			go func() {
				time.Sleep(4 * time.Second)
				statusLabel.SetText("")
				loadContentScreen("Create Meet Schedule", scheduleScreenContent)
				time.Sleep(2 * time.Second)
				setHomeScreen()
			}()
		} else {
			statusLabel.SetText("Invalid link!!")
			go func() {
				time.Sleep(3 * time.Second)
				statusLabel.SetText("")
				if len(linkEntry.Text) > 50 {
					linkEntry.SetText("")
				}
				loadContentScreen("Create Meet Schedule", scheduleScreenContent)
			}()
		}

	})

	scheduleScreenContent = container.NewVBox(
		headingLabel,

		dateLabel,
		dateEntry,

		timeSlotLabel,
		timeSlotEntry,
		container.NewHBox(linkLabel, pasteButton, pasteInfoLabel),
		linkEntry,

		spacer,
		submitButton,
		statusLabel,
	)

	loadContentScreen("Create Meet Schedule", scheduleScreenContent)
}
