package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	sonuc := args[0]
	for i := 1; i < len(args); i++ {
		sonuc += "\n" + args[i]
	}
	return sonuc
}
