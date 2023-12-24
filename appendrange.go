package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var arrA []int
	for i := min; i < max; i++ {
		arrA = append(arrA, i)
	}
	return arrA
}
