package generateReport

import (
	"time"

	"github.com/AbdulRahimOM/report-maker-app/data"
)

type LinkSchedule struct {
	DateText    string
	MeetingLink string
	MeetingSlot string
}

func CreateLinkScheduke(batchDetails data.BatchData, scheduleData LinkSchedule) *string {

	//multi line string - do not misundertand
	//====================================================

	// `✨Good morning All✨
	report :=
		`🎙Communication Session

Batch: ` + batchDetails.Name + `🌺

🖇Meeting link:` + scheduleData.MeetingLink + `

📆Date:- ` + time.Now().Format("January 02, 2006") + `

⏰Time:- ` + scheduleData.MeetingSlot

	return &report
}
