package tools

import "strings"

func ValidateTldvLink(link string) bool {
	return (strings.HasPrefix(link, "https://tldv.io") || strings.HasPrefix(link, "tldv.io") && len(link) < 70)
}

func ValidateGMeetLink(link string) bool {
	return (strings.HasPrefix(link, "https://meet.google.com/") || strings.HasPrefix(link, "meet.google.com/") && len(link) < 50)
}

func ValidateName(name string) bool {
	//check if name contains only alphabets
	for _, char := range name {
		switch {
		case char >= 'A' && char <= 'Z',
			char >= 'a' && char <= 'z',
			char == ' ',
			char == '.':
			// do nothing
		default:
			return false
		}
		// if (char < 'A' || (char > 'Z' && char < 'a') || char > 'z' )||char == ' '{
		// 	return false
		// }
	}

	// check if name is atleast 2 characters long
	return len(name) >= 2
}
