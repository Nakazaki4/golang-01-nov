package piscine

func PodiumPosition(podium [][]string) [][]string {
	p1 := len(podium) - 1

	for i := 0; i < len(podium)/2; i++ {
		podium[i], podium[p1] = podium[p1], podium[i]
		p1--
	}
	return podium
}
