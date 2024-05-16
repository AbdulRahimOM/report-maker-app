package data

type BatchData struct {
	Name      string   `json:"name"`
	Trainer   string   `json:"trainer"`
	MainCord  string   `json:"mainCord"`
	AsstCord  string   `json:"asstCord"`
	UsualTime string   `json:"usualTime"`
	TimeSlot  string   `json:"timeSlot"`
	Members   []string `json:"members"`
}

var (
	//when the app is started for the first time, this data will be copied to the batchData variable and local file
	DefaultData = BatchData{
		Name:      "BCR39/40",
		Trainer:   "Siva Shakthi Sir",
		MainCord:  "Abdul Rahim O M",
		AsstCord:  "Afsal KT",
		UsualTime: "3:00 PM",
		TimeSlot:  "3:00 PM to 4:00 PM",
		Members: []string{
			"Shruthi Kiron",
			"Abdul Rahim O M",
			"Afsal KT",
			"Amal",
			"Anjali",
			"Anusha",
			"Gadha",
			"Sreedevan",
			"Arjun",
			"Aswin",
			"Alan",
			"Mishab",
			"Ajay",
		},
	}

	Batch BatchData
)

func init() {
	Batch =LoadData()
}