package refine

import (
	"strconv"
	"strings"
	// "regexp"
)

func Capitalize(text string) string {
	words := strings.Split(string(text), " ")

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" {
			words[i-1] = strings.ToUpper(words[i-1][:1]) + words[i-1][1:]
			words = append(words[:i], words[i+1:]...)
			i-- // neglect the (cap)
		}

		// capitalize with num
		if words[i] == "(cap," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			if number >= len(words[i]) {
				number = len(words[i])
			}
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToUpper(words[i-j][:1]) + words[i-j][1:]
			}
			words = append(words[:i], words[i+2:]...)
			i-- // neglect the (cap, <number>)
		}
	}

	return strings.Join(words, " ")
}

// func Capitalize(text string) string {
// 	if len(text) == 0 {
// 		return text
// 	}

// 	regex := `([\w\s]+?)\s*\(cap(?:,\s*(\d+))?\)`
// 	re := regexp.MustCompile(regex)
// 	var results [][]string

// 	result := re.ReplaceAllStringFunc(text, func(match string) string {
// 		submatches := re.FindStringSubmatch(match)

// 		for index, match := range submatches {
// 			// This will check if a number of words is specified
// 			if (index+1)%3 == 0 {
// 				num, err := strconv.Atoi(submatches[2])
// 				if err != nil {
// 					// This is when there is no specific number of words
// 					continue
// 				}

// 				matches := re.FindAllStringSubmatch(text, -1)
// 				for _, match := range matches {
// 					if len(match) > 1 {
// 						// Split the matched words into individual words
// 						words := strings.Fields(match[1])
// 						results = append(results, words)
// 					}
// 				}

// 				if len(submatches) > 2 && submatches[2] != "" {
// 					results[1] = CapitalizeListInitials(results[1], num)
// 				}

// 				// The bug is happening here !

// 				return strings.Join(results[1], " ")
// 			} else if index%2 == 0 {
// 				// Here, no specific number of words entered
// 				_, err := strconv.Atoi(submatches[2])
// 				if err == nil {
// 					// This is when there is a specific number of words
// 					continue
// 				}
// 				regex := `(\w+)\s*\(cap\)`
// 				re := regexp.MustCompile(regex)

// 				result := re.ReplaceAllStringFunc(match, func(match string) string {
// 					word := re.FindStringSubmatch(match)[1]
// 					return strings.ToUpper(word[:1])+word[1:] + " "
// 				})
// 				return result
// 			}
// 		}

// 		return match

// 	})

// 	return result
// }

// func CapitalizeListInitials(words []string, start int) []string {
// 	// This function capitalizes the initials for an array of strings starting from n position
// 	var capWords []string

// 	for i := 0; i < len(words); i++ {
// 		if len(words[i]) > 0 {
// 			if i >= len(words)-start {
// 				capWords = append(capWords, strings.ToUpper(words[i][:1])+words[i][1:])
// 			} else {
// 				capWords = append(capWords, words[i])
// 			}
// 		}
// 	}

// 	return capWords
// }

// func CapitalizeInitial(words []string) []string {
// 	// This function capitalizes the initials of a word
// 	var capWords []string

// 	for i := 0; i < len(words); i++ {
// 		if len(words[i]) > 0 {
// 			if i == len(words)-1 {
// 				capWords = append(capWords, strings.ToUpper(words[i][:1])+words[i][1:])
// 			} else {
// 				capWords = append(capWords, words[i])
// 			}
// 		}
// 	}

// 	return capWords
// }
