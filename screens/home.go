package screens

import (
	"strconv"

	"github.com/AbdulRahimOM/report-maker-app/data"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func setHomeScreen() {
	//heading
	batchDetailsHeading := widget.NewLabel("Batch Details")
	batchDetailsHeading.TextStyle.Bold = true

	//Labels
	batchNameLabel := widget.NewLabel("Batch Name: ")
	trainerLabel := widget.NewLabel("Trainer: ")
	mainCordLabel := widget.NewLabel("Main Coordinator: ")
	asstCordLabel := widget.NewLabel("Asst Coordinator: ")

	//Value Labels
	batchNameValueLabel := widget.NewLabel(data.Batch.Name)
	trainerValueLabel := widget.NewLabel(data.Batch.Trainer)
	mainCordValueLabel := widget.NewLabel(data.Batch.MainCord)
	asstCordValueLabel := widget.NewLabel(data.Batch.AsstCord)

	//Grouping Labels and Value Labels
	labelSet := container.NewVBox(
		batchNameLabel,
		trainerLabel,
		mainCordLabel,
		asstCordLabel,
	)
	valueSet := container.NewVBox(
		batchNameValueLabel,
		trainerValueLabel,
		mainCordValueLabel,
		asstCordValueLabel,
	)
	batchDetailsSet := container.NewHBox(labelSet, valueSet)

	//Members List
	membersListLabel := widget.NewLabel("Members: ")
	membersListLabel.TextStyle.Bold = true
	membersCtnrL := container.NewVBox()
	gap := widget.NewLabel("      ")
	membersCtnrR := container.NewVBox()
	membersCtnr := container.NewHBox(membersCtnrL,gap, membersCtnrR)
	strength := len(data.Batch.Members)
	for i := 0; i <= strength/2; i++ {
		membersCtnrL.Add(widget.NewLabel(strconv.Itoa(i+1) + ". " + data.Batch.Members[i]))
	}
	for i := strength/2 + 1; i < strength; i++ {
		membersCtnrR.Add(widget.NewLabel(strconv.Itoa(i+1) + ". " + data.Batch.Members[i]))
	}

	newHomeScreen := container.NewVBox(
		batchDetailsHeading,
		batchDetailsSet,
		widget.NewSeparator(),
		membersListLabel,
		membersCtnr,
	)

	loadContentScreen("", newHomeScreen)

}
