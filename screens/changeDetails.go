package screens

import (
	"errors"
	"strconv"
	"time"

	"github.com/AbdulRahimOM/report-maker-app/data"
	"github.com/AbdulRahimOM/report-maker-app/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
	usualTimeLabel := widget.NewLabel("Usual Time: ")

	// Create input boxes
	nameEntry := widget.NewEntry()
	trainerEntry := widget.NewEntry()
	mainCordEntry := widget.NewEntry()
	asstCordEntry := widget.NewEntry()
	usualTimeEntry := widget.NewEntry()
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

		usualTimeEntry.SetPlaceHolder("enter usual time here")
		usualTimeEntry.SetText(data.Batch.UsualTime)
	}

	// Grouping Labels and Value Labels
	labelSet := container.NewVBox(
		nameLabel,
		trainerLabel,
		mainCordLabel,
		asstCordLabel,
		usualTimeLabel,
	)
	inputSet := container.NewVBox(
		nameEntry,
		trainerEntry,
		mainCordEntry,
		asstCordEntry,
		usualTimeEntry,
	)
	localSplit := container.NewHSplit(labelSet, inputSet)

	// Members
	membersLabel := widget.NewLabel("Members: ")
	membersLabel.TextStyle.Bold = true

	// Existing members
	members := data.Batch.Members
	// var membersLabels []*widget.Label
	membersCtnrL := container.NewVBox()
	membersCtnrR := container.NewVBox()
	gap := widget.NewLabel("      ")
	membersCtnr := container.NewHBox(membersCtnrL, gap, membersCtnrR)
	strength := len(members)

	//Members being removed
	removedMembersLabel := widget.NewLabel("Members being removed:")
	removedMembersLabel.TextStyle.Bold = true
	removedMembersLabel.Hide()
	// var removedMembersLabels []*widget.Label
	// var removedMembers []string
	removedMembersCtnr := container.NewVBox()

	var membersRemovalStatus []bool = make([]bool, strength)
	var removedMembersLabels []*widget.Label = make([]*widget.Label, strength)

	removalCount := 0	//for toggling show/hide of 'removedMembersLabel'
	for i := 0; i < strength; i++ {
		index := i
		membersLabel := widget.NewLabel(strconv.Itoa(i+1) + ". " + members[i])
		var removeButton *widget.Button
		removedMembersLabels[i] = widget.NewLabel(strconv.Itoa(i+1) + ". " + members[i])
		removeButton = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			if !membersRemovalStatus[index] {
				removedMembersCtnr.Add(removedMembersLabels[index])
				removeButton.SetIcon(theme.ViewRefreshIcon())
				membersRemovalStatus[index] = true
				removalCount++
				removedMembersLabel.Show()
			} else {
				removedMembersCtnr.Remove(removedMembersLabels[index])
				removeButton.SetIcon(theme.DeleteIcon())
				membersRemovalStatus[index] = false

				removalCount--
				if removalCount == 0 {
					removedMembersLabel.Hide()
				}
			}
		})

		// Arrange label and button in a horizontal box
		labelAndButton := container.NewHBox(
			// label,
			removeButton,
			membersLabel,
		)
		if i <= strength/2 {
			membersCtnrL.Add(labelAndButton)
		} else {
			membersCtnrR.Add(labelAndButton)
		}
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
	var newMemberNamesLabels []*widget.Label
	var newMembers []string
	newMembersCtnr := container.NewVBox()
	addButton := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		if tools.ValidateName(addMemberEntry.Text) {
			newMembers = append(newMembers, addMemberEntry.Text)
			newMemberNamesLabels = append(newMemberNamesLabels, widget.NewLabel(strconv.Itoa(len(members)+len(newMemberNamesLabels)+1)+". "+addMemberEntry.Text))
			newMembersCtnr.Add(newMemberNamesLabels[len(newMemberNamesLabels)-1])
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

		updatedMembers := []string{}
		//remove members in delete list
		for i := 0; i < strength; i++ {
			if !membersRemovalStatus[i] {
				updatedMembers = append(updatedMembers, members[i])
			}
		}
		// adding new members to the existing members
		updatedMembers = append(updatedMembers, newMembers...)

		// Check if the values are empty
		if nameEntry.Text == "" {
			nameEntry.Text = data.Batch.Name
		}

		// Save data to data.Batch and to JSON file
		data.SaveData(data.BatchData{
			Name:      nameEntry.Text,
			Trainer:   trainerEntry.Text,
			MainCord:  mainCordEntry.Text,
			AsstCord:  asstCordEntry.Text,
			UsualTime: usualTimeEntry.Text,
			Members:   updatedMembers,
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

		removedMembersLabel,
		removedMembersCtnr,

		container.NewHBox(
			addMemberLabel,
			inputInfoLabel,
		),
		container.NewHBox(
			addMemberEntryCtnr,
			addButton,
		),
		newMembersCtnr,
		seperator,

		submitButton,
	)

	loadContentScreen("Configure batch details", changeDetailScreen)
}
