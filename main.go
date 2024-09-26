package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
)

type IntPair struct {
	First  int
	Second int
}

type JsonInput struct {
	ToEncodeString string `json:"toEncode"`
}

func main() {
	var engAlphabet [26]rune

	for i := 0; i < len(engAlphabet); i++ {
		engAlphabet[i] = 'a' + rune(i)
	}

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomizer.Shuffle(len(engAlphabet), func(i, j int) {
		engAlphabet[i], engAlphabet[j] = engAlphabet[j], engAlphabet[i]
	})

	polyboyMatrix := make(map[rune]IntPair)
	for i := 0; i < 6; i++ {
		fmt.Println("")
		for j := 0; j < 6; j++ {
			if 6*i+j < len(engAlphabet) {
				polyboyMatrix[engAlphabet[6*i+j]] = IntPair{
					First:  i,
					Second: j,
				}
				fmt.Printf("(%d,%d)=%c ", i, j, engAlphabet[6*i+j])
			} else {
				break
			}
		}
	}

	jsonFile, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var jsonInput JsonInput
	err = json.Unmarshal(jsonFile, &jsonInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	toEncode := jsonInput.ToEncodeString
	encodedData := make([]string, len(toEncode))

	fmt.Printf("%s\n", toEncode)
	for i, c := range toEncode {
		pair, isFound := polyboyMatrix[unicode.ToLower(c)]
		if isFound {
			if unicode.IsUpper(c) {
				encodedData[i] = fmt.Sprintf("%d%d%c", pair.First, pair.Second, 'u')
			} else {
				encodedData[i] = fmt.Sprintf("%d%d%c", pair.First, pair.Second, 'l')
			}
		} else {
			encodedData[i] = string(c)
		}
	}

	fmt.Printf("[")
	for _, c := range encodedData {
		fmt.Printf("%s,", c)
	}
	fmt.Printf("]\n")

	decodedData := ""
	for _, s := range encodedData {
		if len(s) == 3 {
			pairToFind := IntPair{
				First:  int(s[0] - '0'),
				Second: int(s[1] - '0'),
			}

			for letter := range polyboyMatrix {
				pair := polyboyMatrix[letter]
				if pair == pairToFind {
					if s[2] == 'u' {
						decodedData += string(unicode.ToUpper(letter))
					} else {
						decodedData += string(unicode.ToLower(letter))
					}

				}
			}

		} else {
			decodedData += s
		}
	}

	fmt.Println(decodedData)

}
