package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type IntPair struct {
	i int
	j int
}

func getGreenText(input string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", input)
}

func getRedText(input string) string {
	return fmt.Sprintf("\033[31m%s\033[0m\n", input)
}

func main() {

	//Считываем текст, который необходимо закодировать
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Printf(getRedText("Ошибка при чтении файла: %s"), err)
		return
	}
	toEncode := string(data)

	fmt.Printf("Данные для кодирования: \n%s\n", getGreenText(toEncode))

	//Формируем алфавит
	alphabet := GetAlphabet(toEncode)
	/*
		Проводим замешивание алфавита
		Необходимо для формирования
		Случайной матрицы Полибоя
	*/
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomizer.Shuffle(len(alphabet), func(i, j int) {
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
	})

	//Формируем матрицу полибоя
	polyboyMatrix := GetPolyBoy(alphabet)

	fmt.Printf("\nСгенерированная матрица полибоя \n%s\n", getGreenText(PolyBoyAsString(polyboyMatrix)))

	encodedData := make([]IntPair, 0)
	//Производим кодирование текста
	for _, letter := range toEncode {
		i, j, ok := GetPosition(letter, polyboyMatrix)
		if !ok {
			fmt.Println(fmt.Errorf("(%c) нет в матрице %s", letter, getRedText(PolyBoyAsString(polyboyMatrix))))
			return
		}
		encodedData = append(encodedData, IntPair{i, j})
	}

	fmt.Printf("Закодированные данные: \n%v\n", encodedData)

	//Производим декодирование текста
	decodedData := ""
	for _, pair := range encodedData {
		letter, ok := GetLetter(uint(pair.i), uint(pair.j), polyboyMatrix)
		if !ok {
			fmt.Println(fmt.Errorf("%c нет в матрице \n%v", letter, getRedText(PolyBoyAsString(polyboyMatrix))))
			return
		}
		decodedData += string(letter)
	}

	fmt.Printf("\nДекодированные данные: \n%s\n\n", getGreenText(decodedData))

	if toEncode == decodedData {
		fmt.Println("Исходные и выходные данные равны")
	} else {
		fmt.Println("Данные не сходятся")
	}

}
