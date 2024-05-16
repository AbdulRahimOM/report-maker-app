package generateReport

import (
	"github.com/AbdulRahimOM/report-maker-app/data"
)

type AudioReportData struct {
	DateText    string
	Topic       string
	Submissions []bool
	ReportedBy  string
}

func CreateAudioReport(batchDetails data.BatchData, audioReport AudioReportData) *string {
	topicStr := ""
	if audioReport.Topic != "" {
		topicStr = "\n" + `🔖 Topic: ` + audioReport.Topic
	}

	submissionStr := ""
	for i, v := range data.DefaultData.Members {
		if audioReport.Submissions[i] {
			submissionStr += `✅ ` + v + "\n"
		} else {
			submissionStr += `❌ ` + v + "\n"
		}
	}

	report :=
		`*🎙 Audio task Submission Report*
	

🌸 ` + batchDetails.Name + `
📅 ` + audioReport.DateText + `
` + topicStr + `
	
*Submission status:*
` + submissionStr + `
	
✒️Report prepared by :
  ` + audioReport.ReportedBy

	return &report
}
