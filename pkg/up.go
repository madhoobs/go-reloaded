package refine

import (
	"regexp"
	"strings"
)

func Uppercase(text string) string {
	if len(text) == 0 {
		return text
	}
	
	regex := `(\w+)\s*\(up\)`
	re := regexp.MustCompile(regex)

	result := re.ReplaceAllStringFunc(text, func(match string) string {
		word := re.FindStringSubmatch(match)[1]
		return strings.ToUpper(word)
	})

	return result
}
