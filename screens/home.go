package screens

import (
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

	//Reseting the home screen
	setHomeScreen := container.NewVBox(
		batchDetailsHeading,
		batchDetailsSet,
	)

	loadContentScreen("", setHomeScreen)

}
