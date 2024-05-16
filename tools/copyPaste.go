package tools

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func GetClipboardText() string {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Error reading clipboard:", err)
		return ""
	}
	return text
}

//copyToClipboard copies the given text to the clipboard
func CopyToClipboard(text *string) {
	err := clipboard.WriteAll(*text)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
	}
}