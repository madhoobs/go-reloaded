package refine

import (
	"strconv"
	"strings"
	// "regexp"
)

func Lowercase(text string) string {
	words := strings.Split(string(text), " ")

	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i-- // neglect the (low)
		}

		// capitalize with num
		if words[i] == "(low," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			if number >= len(words[i]) {
				number = len(words[i])
			}
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i-- // neglect the (low, <number>)
		}
	}

	return strings.Join(words, " ")
}

// func Lowercase(text string) string {
// 	if len(text) == 0 {
// 		return text
// 	}

// 	regex := `(\w+)\s*\(low\)`
// 	re := regexp.MustCompile(regex)

// 	result := re.ReplaceAllStringFunc(text, func(match string) string {
// 		word := re.FindStringSubmatch(match)[1]
// 		return strings.ToLower(word)
// 	})

// 	return result
// }
