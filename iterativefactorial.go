package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 25 {
		return 0
	}
	result := 1

	if nb == 0 {
		return 1
	}

	for i := 1; i <= nb; i++ {
		result *= i
	}

	return result
}
