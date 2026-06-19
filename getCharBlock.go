package main

func getCharBlock(char rune, lines []string) []string {

	ascii := int(char)

	if ascii < 32 || ascii > 126 {
		return []string{
			"        ",
			"        ",
			"        ",
			"        ",
			"        ",
			"        ",
			"        ",
			"        ",
		}
	}

	index := ascii - 32
	start := index*9 + 1

	return lines[start : start+8]
}
