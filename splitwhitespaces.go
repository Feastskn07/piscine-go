package piscine

func SplitWhiteSpaces(s string) []string {
	var kelimeler []string
	var mevcutKelime string

	for _, char := range s {
		if char == ' ' {
			if mevcutKelime != "" {
				kelimeler = append(kelimeler, mevcutKelime)
				mevcutKelime = ""
			}
		} else {
			mevcutKelime += string(char)
		}
	}

	if mevcutKelime != "" {
		kelimeler = append(kelimeler, mevcutKelime)
	}

	return kelimeler
}
