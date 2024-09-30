package main

import (
	"fmt"
	"math"
)

func GetPolyBoy(alphabet []rune) [][]rune {
	//Формируем матрицу полибоя
	matrixSize := int(math.Ceil(math.Sqrt(float64(len(alphabet)))))
	polyboyMatrix := make([][]rune, 0)
	alphaIndex := 0

	for i := 0; i < matrixSize; i++ {

		if len(alphabet)-alphaIndex > matrixSize {
			polyboyMatrix = append(polyboyMatrix, make([]rune, matrixSize))
		} else {
			polyboyMatrix = append(polyboyMatrix, make([]rune, len(alphabet)-alphaIndex))
		}

		for j := 0; j < len(polyboyMatrix[i]); j++ {
			polyboyMatrix[i][j] = alphabet[alphaIndex]
			alphaIndex++
		}
	}
	return polyboyMatrix
}

func GetLetter(i uint, j uint, polyBoyMatrix [][]rune) (rune, bool) {
	if int(i) < len(polyBoyMatrix) {
		if int(j) < len(polyBoyMatrix[i]) {
			return polyBoyMatrix[i][j], true
		}
	}
	return rune(0), false
}

func GetPosition(letter rune, polyBoyMatrix [][]rune) (int, int, bool) {
	for i := 0; i < len(polyBoyMatrix); i++ {
		for j := 0; j < len(polyBoyMatrix[i]); j++ {
			if polyBoyMatrix[i][j] == letter {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

func PolyBoyAsString(polyBoyMatrix [][]rune) string {
	result := ""
	for i := 0; i < len(polyBoyMatrix); i++ {
		for j := 0; j < len(polyBoyMatrix[i]); j++ {
			result += fmt.Sprintf("(%d,%d)=%q ", i, j, polyBoyMatrix[i][j])
		}
		result += "\n"
	}
	return result
}
