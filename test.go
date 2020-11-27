package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Número de entradas
const entradas = 4

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

func hallaF(x1, x2, x3, x4 float64) string {
	f1 := sigmoid(x1*pesos[1][1] + x2*pesos[1][2] + x3*pesos[1][3] + x4*pesos[1][4])
	f2 := sigmoid(x1*pesos[2][1] + x2*pesos[2][2] + x3*pesos[2][3] + x4*pesos[2][4])
	f3 := sigmoid(x1*pesos[3][1] + x2*pesos[3][2] + x3*pesos[3][3] + x4*pesos[3][4])
	f4 := sigmoid(x1*pesos[4][1] + x2*pesos[4][2] + x3*pesos[4][3] + x4*pesos[4][4])
	fy := sigmoid(f1*pesos[5][1] + f2*pesos[5][2] + f3*pesos[5][3] + f4*pesos[5][4])

	if fy < 0.8 {
		fmt.Println("Respuesta: setosa", fy)
		return "setosa"
	}
	if fy < 0.935 && fy >= 0.8 {
		fmt.Println("Respuesta: versicolor", fy)
		return "versicolor"
	}
	if fy >= 0.935 {
		fmt.Println("Respuesta: virginica", fy)
		return "virginica"
	}
	return ""
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

	// setosa
	if hallaF(5.2, 3.6, 1.5, 0.3) == "setosa" {
		cont++
	}
	if hallaF(5.1, 3.2, 1.6, 0.3) == "setosa" {
		cont++
	}
	if hallaF(4.8, 3.3, 1.4, 0.3) == "setosa" {
		cont++
	}
	if hallaF(4.7, 3.2, 1.6, 0.3) == "setosa" {
		cont++
	}

	// versicolor
	if hallaF(7.1, 3.3, 4.8, 1.5) == "versicolor" {
		cont++
	}
	if hallaF(6.5, 3.3, 4.6, 1.6) == "versicolor" {
		cont++
	}
	if hallaF(7.0, 3.2, 5.0, 1.6) == "versicolor" {
		cont++
	}
	if hallaF(5.6, 2.4, 4.1, 1.4) == "versicolor" {
		cont++
	}

	// virginica
	if hallaF(6.4, 3.4, 6.1, 2.6) == "virginica" {
		cont++
	}
	if hallaF(5.9, 2.8, 5.2, 2.0) == "virginica" {
		cont++
	}
	if hallaF(7.2, 3.1, 6.0, 2.2) == "virginica" {
		cont++
	}
	if hallaF(6.4, 3.0, 5.7, 1.9) == "virginica" {
		cont++
	}

	fmt.Println("\nExactitud con datos de prueba: ", (cont/n)*100, "%")

}
