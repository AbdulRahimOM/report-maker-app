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

func LoadData() BatchData {
    // Check if the file exists
    if _, err := os.Stat(batchDataFile); os.IsNotExist(err) {
		fmt.Println("batch-details.json does not exist. Using default data.")
        return DefaultData
    }else{
		if err != nil {
			fmt.Println("Error checking file:", err)
			return DefaultData
		}
	}

    // Read data from file
    jsonData, err := os.ReadFile(batchDataFile)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return DefaultData
    }

    // Unmarshal JSON data into struct
    var batchData BatchData
    err = json.Unmarshal(jsonData, &batchData)
    if err != nil {
        fmt.Println("Error unmarshalling data:", err)
        return DefaultData
    }

    // Batch = batchData
    // fmt.Println("Loaded Batch:", Batch)

    return batchData
}

