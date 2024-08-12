package refine

func Punctuations(text string) string {
	punctuations := []string{",", ".", "!", "?", ":", ";"}
	chars := []rune(text)

	// For Punctuations
	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == ' ' && contains(chars[i+1], punctuations) {
			// This handles when there is a space before the punctuation
			copy(chars[i:], chars[i+1:])
			chars = chars[:len(chars)-1]
			i--
		} else if contains(chars[i], punctuations) && chars[i+1] != ' ' && !contains(chars[i+1], punctuations) {
			// This handles the punctuation that do not have a space after them
			temp := chars[len(chars)-1]
			copy(chars[i+2:], chars[i+1:])
			chars[i+1] = ' '
			chars = append(chars, temp)
			continue
		}
	}

	// For Apostrophes
	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == '\'' && chars[i+1] == ' ' {
			copy(chars[i+1:], chars[i+2:])
			chars = chars[:len(chars)-1]
			i--
			for j := i+1; j < len(chars)-1; j++ {
				if chars[j] == ' ' && chars[j+1] == '\'' {
					copy(chars[j:], chars[j+1:])
					chars = chars[:len(chars)-1]
					j--
					continue
				}
			}
		}
	}

	// For a & an
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i := 0; i < len(chars)-1; i++ {
		if (i == 0 || (i > 0 && chars[i-1] == ' ')) && (chars[i] == 'a' || chars[i] == 'A') && chars[i+1] == ' ' && contains(chars[i+2], vowels) {
			temp := chars[len(chars)-1]
			copy(chars[i+2:], chars[i+1:])
			chars[i+1] = 'n'
			chars = append(chars, temp)
			continue
		}
	}

	return string(chars)
}

func contains(char rune, slice []string) bool {
	for _, item := range slice {
		if item == string(char) {
			return true
		}
	}
	return false
}
