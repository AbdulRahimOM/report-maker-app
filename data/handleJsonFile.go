package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var batchDataFile string
var relativeFilePathAndName = "data/batch-details.json"

func init() {
	var baseDir string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %v", err)
	}

	switch runtime.GOOS {
	case "darwin":
		// macOS
		baseDir = filepath.Join(homeDir, "Library", "Application Support", "Report-Maker-App")
	case "linux":
		// Linux
		baseDir = filepath.Join(homeDir, ".Report-Maker-App")
	case "windows":
		// Windows
		appDataDir := os.Getenv("APPDATA")
		if appDataDir == "" {
			log.Fatalf("Error getting AppData directory")
		}
		baseDir = filepath.Join(appDataDir, "Report-Maker-App")
	default:
		// log.Fatalf("Unsupported OS: %s", runtime.GOOS)
		fmt.Println("Didn't recognize the OS. So, the current directory will be used to save config data.")
		baseDir = filepath.Join(".", "Report-Maker-App")
	}

	// Create the full path to the data file
	batchDataFile = filepath.Join(baseDir, relativeFilePathAndName)
	fmt.Println("Data file path:", batchDataFile)
}
func SaveData(data BatchData) {
	// Ensure the directory exists
	dir := filepath.Dir(batchDataFile)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}
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
	} else {
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

	return batchData
}
