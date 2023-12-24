package piscine

func RecursivePower(nbt, power int) int {
	if power < 0 {
		return 0
	}

	return recursivePowerHelper(nbt, power)
}

func recursivePowerHelper(nbt, power int) int {
	if power == 0 {
		return 1
	}

	return nbt * recursivePowerHelper(nbt, power-1)
}
