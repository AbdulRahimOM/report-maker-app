package data

import (
	"encoding/json"
	"fmt"
	"os"
)

var batchDataFile = "./data/batch-details.json"

func SaveData(data BatchData) {
	Batch = data
	// Marshal data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	// Write JSON data to file
	err = os.WriteFile(batchDataFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
