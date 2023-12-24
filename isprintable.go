package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if char < 32 || char > 126 {
			// ASCII'de yazdırılabilir olmayan karakterler
			return false
		}
	}
	return true
}
