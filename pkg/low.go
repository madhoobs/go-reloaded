package refine

import (
	"regexp"
	"strings"
)

func Lowercase(text string) string {
	if len(text) == 0 {
		return text
	}
	
	regex := `(\w+)\s*\(low\)`
	re := regexp.MustCompile(regex)

	result := re.ReplaceAllStringFunc(text, func(match string) string {
		word := re.FindStringSubmatch(match)[1]
		return strings.ToLower(word)
	})

	return result
}
