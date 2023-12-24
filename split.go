package piscine

func Split(s, sep string) []string {
	var words []string
	var currentWord string
	a := []rune(s)
	b := []rune(sep)

	for i := 0; i < len(a); i++ {
		match := true
		for j := 0; j < len(b); j++ {
			if i+j >= len(a) || a[i+j] != b[j] {
				match = false
				break
			}
		}

		if match {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
			i += len(b) - 1
		} else {
			currentWord += string(a[i])
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}
