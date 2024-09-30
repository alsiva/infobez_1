package main

func GetAlphabet(data string) []rune {
	uniqueLetters := make(map[rune]bool)
	for _, r := range data {
		uniqueLetters[r] = true
	}
	result := make([]rune, 0, len(uniqueLetters))
	for letter := range uniqueLetters {
		result = append(result, letter)
	}
	return result
}
