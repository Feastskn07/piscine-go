package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arrayA := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		arrayA[i] = min + i
	}
	return arrayA
}
