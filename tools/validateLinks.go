package tools

import "strings"

func ValidateTldvLink(link string) bool {
	return (strings.HasPrefix(link, "https://tldv.io") || strings.HasPrefix(link, "tldv.io") && len(link) < 70)
}

func ValidateGMeetLink(link string) bool {
	return (strings.HasPrefix(link, "https://meet.google.com/") || strings.HasPrefix(link, "meet.google.com/") && len(link) < 50)
}
