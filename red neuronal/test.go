package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
)

// Número de entradas
const entradas = 3

// Número de neuronas
const neuronas = entradas + 1

type datosEntrenam struct {
	x [entradas + 1]float64
	z float64
}

func sigmoid(h float64) float64 {
	return 1 / (1 + math.Exp(-h))
}

func read() [entradas*entradas + entradas]float64 {

	f, err := os.Open("pesos.txt")
	var arr [entradas*entradas + entradas]float64
	var i int32 = 0
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		arr[i], err = strconv.ParseFloat(scanner.Text(), 64)
		i++
	}
	fmt.Println(" ")
	return arr
}


var pesos [neuronas + 1][entradas + 1]float64

func hallaF(x1, x2, x3, x4 float64) float64 {
	f1 := sigmoid(x1*pesos[1][1] + x2*pesos[1][2] + x3*pesos[1][3] + x4*pesos[1][4])
	f2 := sigmoid(x1*pesos[2][1] + x2*pesos[2][2] + x3*pesos[2][3] + x4*pesos[2][4])
	f3 := sigmoid(x1*pesos[3][1] + x2*pesos[3][2] + x3*pesos[3][3] + x4*pesos[3][4])
	f4 := sigmoid(x1*pesos[4][1] + x2*pesos[4][2] + x3*pesos[4][3] + x4*pesos[4][4])
	fy := sigmoid(f1*pesos[5][1] + f2*pesos[5][2] + f3*pesos[5][3] + x4*pesos[5][4])

	if fy < 0.8 {
		fmt.Println("Respuesta:", fy)
	}
	if fy < 0.935 && fy >= 0.8 {
		fmt.Println("Respuesta:", fy)
	}
	if fy >= 0.935 {
		fmt.Println("Respuesta:", fy)
	}
	return fy
}

func main() {

	// Leer pesos
	var readerAux [entradas*entradas + entradas]float64 = read()
	cont := 0
	for i := 1; i <= neuronas; i++ {
		for j := 1; j <= entradas; j++ {
			pesos[i][j] = readerAux[cont]
			cont++
		}
	}

	n := 12
	cont = 0

	// 1
	if hallaF(42, 0, 1) == 1 {
		cont++
	}
	if hallaF(51, 4, 1) == 1 {
		cont++
	}
	if hallaF(59, 8, 1) == 1 {
		cont++
	}
	if hallaF(64, 12, 1) == 1 {
		cont++
	}

	fmt.Println("\nExactitud con datos de prueba: ", (cont/n)*100, "%")

}
