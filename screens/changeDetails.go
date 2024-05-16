package screens

import (
	"errors"

	"github.com/AbdulRahimOM/report-maker-app/data"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func changeParticularsScreen() {
	// Heading
	headingLabel := widget.NewLabel("Change Batch Details:")
	headingLabel.TextStyle.Bold = true

	// Labels
	nameLabel := widget.NewLabel("Batch Name")
	trainerLabel := widget.NewLabel("Trainer: ")
	mainCordLabel := widget.NewLabel("Main Coordinator: ")
	asstCordLabel := widget.NewLabel("Asst Coordinator: ")

	// Create input boxes
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("enter batch name here")
	nameEntry.SetText(data.Batch.Name)
	nameEntry.Validator = func(s string) error {
		if len(s) < 3 {
			return errors.New("name should be atleast 4 characters long")
		}
		return nil
	}
	trainerEntry := widget.NewEntry()
	trainerEntry.SetPlaceHolder("enter trainer name here")
	trainerEntry.SetText(data.Batch.Trainer)

	mainCordEntry := widget.NewEntry()
	mainCordEntry.SetPlaceHolder("enter main coordinator name here")
	mainCordEntry.SetText(data.Batch.MainCord)

	asstCordEntry := widget.NewEntry()
	asstCordEntry.SetPlaceHolder("enter asst coordinator name here")
	asstCordEntry.SetText(data.Batch.AsstCord)
	labelSet := container.NewVBox(
		nameLabel,
		trainerLabel,
		mainCordLabel,
		asstCordLabel,
	)
	inputSet := container.NewVBox(
		nameEntry,
		trainerEntry,
		mainCordEntry,
		asstCordEntry,
	)
	localSplit := container.NewHSplit(labelSet, inputSet)

	// Create a submit button
	submitButton := widget.NewButton("Submit", func() {
		// Get the values from the input boxes
		name := nameEntry.Text
		trainer := trainerEntry.Text
		mainCord := mainCordEntry.Text
		asstCord := asstCordEntry.Text

		// Check if the values are empty
		if name == "" {
			name = data.Batch.Name
		}
		if trainer == "" {
			trainer = data.Batch.Trainer
		}
		if mainCord == "" {
			mainCord = data.Batch.MainCord
		}

		//save the values to the data
		data.Batch= data.BatchData{
			Name:     name,
			Trainer:  trainer,
			MainCord: mainCord,
			AsstCord: asstCord,
		}

		setHomeScreen()
	})

	changeDetailScreen := container.NewVBox(
		headingLabel,
		localSplit,
		submitButton,
	)

	loadContentScreen("Configure batch details", changeDetailScreen)
}
