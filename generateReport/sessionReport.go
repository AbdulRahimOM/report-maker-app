package generateReport

import (
	"github.com/AbdulRahimOM/report-maker-app/data"
)

//	sessionReport := generateReport.SessionReportData{
//		DateText:    date,
//		TimeText:    time,
//		Activity:    activity,
//		TLDVLink:    tldvLink,
//		ReportedBy:  reportedBy,
//		Summary:     summary,
//	}
type SessionReportData struct {
	DateText   string
	TimeText   string
	Activity   string
	TLDVLink   string
	ReportedBy string
	Attendance []bool
	Summary    string
}

// CreateSessionReport creates a session report and copies it to clipboard
func CreateSessionReport(batch data.BatchData, details SessionReportData) *string {
	var asstCordStr, activityStr, attendance, tldvStr string
	//preparing optional strings
	{
		if batch.AsstCord != "" {
			asstCordStr = ` & ` + batch.AsstCord
		}

		if details.Activity != "" {
			activityStr = "\n" + `⛳ Activity: ` + details.Activity
		}

		for i, present := range details.Attendance {
			if present {
				attendance += `✅ ` + batch.Members[i] + "\n"
			} else {
				attendance += `❌ ` + batch.Members[i] + "\n"
			}
		}

		if details.TLDVLink != "" {
			tldvStr = "*📽️ TLDV link:*\n" + details.TLDVLink + "\n\n"
		}
	}

	report :=
		`*🔰 Session Report- ` + batch.Name + `*

🗓 Date : ` + details.DateText + `
🕜 Timing : ` + details.TimeText + `
👨🏽‍🏫 Trainer: ` + batch.Trainer + `
🕵🏽‍♂️ Coordinators:` + batch.MainCord + asstCordStr +
			activityStr + `

*📃 Session Summary:*
` + details.Summary + `

*Attendance:*
` + attendance + `

` + tldvStr + `
✒️Report prepared by :
  ` + details.ReportedBy

	// tools.CopyToClipboard(&report)
	// tools.LogReport("Session Report", report)
	return &report
}
