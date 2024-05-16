package tools

import (
	"bufio"
	"os"
)

func GetAlternative(defaultString string) string {
	reader := bufio.NewReader(os.Stdin)
	new, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if new == "\n" {
		return defaultString
	}
	return new

}
