package piscine

func ConcatParams(args []string) string {
	t := ""
	for index, c := range args {
		if index < len(args)-1 {
			t = t + c + "\n"
		} else {
			t = t + c
		}
	}
	return t
}
