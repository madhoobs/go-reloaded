package refine

import (
	"regexp"
	"strings"
)

func Capitalize(text string) string {
	if len(text) == 0 {
		return text
	}
	
	regex := `(\w+)\s*\(cap\)`
	re := regexp.MustCompile(regex)

	result := re.ReplaceAllStringFunc(text, func(match string) string {
		word := re.FindStringSubmatch(match)[1]
		return strings.ToUpper(word[:1]) + word[1:]
	})

	return result
}
