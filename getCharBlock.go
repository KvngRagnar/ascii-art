package main

func getCharBlock(char rune, lines []string) []string {

	ascii := int(char)

	if ascii < 32 || ascii > 126 {
		blank := make([]string, 8)

		for i := range blank {
			blank[i] = "        "
		}

		return blank
	}

	index := ascii - 32
	start := index*9 + 1

	return lines[start : start+8]
}
