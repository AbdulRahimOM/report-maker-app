package screens

import (
	"errors"
	"strconv"
	"time"

	"github.com/AbdulRahimOM/report-maker-app/data"
	"github.com/AbdulRahimOM/report-maker-app/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type minWidthLayout struct {
	minWidth float32
}

func (m *minWidthLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	for _, obj := range objs {
		minSize := obj.MinSize()
		obj.Resize(fyne.NewSize(m.minWidth, minSize.Height))
		obj.Move(fyne.NewPos(0, 0))
	}
}

func (m *minWidthLayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	var minHeight float32 = 0.0
	for _, obj := range objs {
		minHeight = fyne.Max(minHeight, obj.MinSize().Height)
	}
	return fyne.NewSize(m.minWidth, minHeight)
}

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
	trainerEntry := widget.NewEntry()
	mainCordEntry := widget.NewEntry()
	asstCordEntry := widget.NewEntry()
	{ // Set the values/place-holders,etc of the input boxes
		nameEntry.SetPlaceHolder("enter batch name here")
		nameEntry.SetText(data.Batch.Name)
		nameEntry.Validator = func(s string) error {
			if len(s) < 3 {
				return errors.New("name should be atleast 4 characters long")
			}
			return nil
		}
		trainerEntry.SetPlaceHolder("enter trainer name here")
		trainerEntry.SetText(data.Batch.Trainer)

		mainCordEntry.SetPlaceHolder("enter main coordinator name here")
		mainCordEntry.SetText(data.Batch.MainCord)

		asstCordEntry.SetPlaceHolder("enter asst coordinator name here")
		asstCordEntry.SetText(data.Batch.AsstCord)
	}

	// Grouping Labels and Value Labels
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

	// Members
	membersLabel := widget.NewLabel("Members: ")
	membersLabel.TextStyle.Bold = true

	// Existing members
	members := data.Batch.Members
	var membersLabels []*widget.Label
	membersCtnrL := container.NewVBox()
	membersCtnrR := container.NewVBox()
	gap := widget.NewLabel("      ")
	membersCtnr := container.NewHBox(membersCtnrL, gap, membersCtnrR)
	strength := len(members)
	for i := 0; i <= strength/2; i++ {
		membersLabels = append(membersLabels, widget.NewLabel(strconv.Itoa(i+1)+". "+members[i]))
		membersCtnrL.Add(membersLabels[i])
	}
	for i := strength/2 + 1; i < strength; i++ {
		membersLabels = append(membersLabels, widget.NewLabel(strconv.Itoa(i+1)+". "+members[i]))
		membersCtnrR.Add(membersLabels[i])
	}

	// Add new members
	addMemberLabel := widget.NewLabel("Add Member:")
	addMemberLabel.TextStyle.Bold = true
	inputInfoLabel := widget.NewLabel("Invalid name")
	inputInfoLabel.Hide()
	addMemberEntry := widget.NewEntry()
	addMemberEntry.SetPlaceHolder("enter new name here")
	customLayout := &minWidthLayout{minWidth: 400}
	addMemberEntryCtnr := container.New(customLayout, addMemberEntry)

	// Members being added
	var newMemberNames []*widget.Label
	newMembersCtnr := container.NewVBox()

	//plus button
	plusButton := widget.NewButton("+", func() {
		if tools.ValidateName(addMemberEntry.Text) {
			newMemberNames = append(newMemberNames, widget.NewLabel(strconv.Itoa(len(membersLabels)+len(newMemberNames)+1)+". "+addMemberEntry.Text))
			newMembersCtnr.Add(newMemberNames[len(newMemberNames)-1])
			addMemberEntry.SetText("")
		} else {
			inputInfoLabel.Show()
			go func() {
				time.Sleep(3 * time.Second)
				inputInfoLabel.Hide()
			}()
		}
	})

	// Create a submit button
	submitButton := widget.NewButton("Submit", func() {
		// Get the values from the input boxes
		name := nameEntry.Text
		trainer := trainerEntry.Text
		mainCord := mainCordEntry.Text
		asstCord := asstCordEntry.Text
		for _, member := range newMemberNames {
			members = append(members, member.Text)
		}

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

		// Save data to data.Batch and to JSON file
		data.SaveData(data.BatchData{
			Name:     name,
			Trainer:  trainer,
			MainCord: mainCord,
			AsstCord: asstCord,
			Members:  members,
		})

		setHomeScreen()
	})

	seperator := widget.NewSeparator()

	changeDetailScreen := container.NewVBox(
		headingLabel,
		localSplit,
		seperator,

		membersLabel,
		membersCtnr,
		seperator,

		container.NewHBox(
			addMemberLabel,
			inputInfoLabel,
		),
		container.NewHBox(
			addMemberEntryCtnr,
			plusButton,
		),
		newMembersCtnr,
		seperator,

		submitButton,
	)

	loadContentScreen("Configure batch details", changeDetailScreen)
}
