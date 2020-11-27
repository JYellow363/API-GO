package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type datosEntrenam struct {
	x [entradas + 1]float64
	z float64
}

func aprox(x float64) float64 {
	return math.Round(x*1000) / 1000
}

func sigmoid(h float64) float64 {
	return 1 / (1 + math.Exp(-h))
}

func _sigmoid(h float64) float64 {
	return h * (1 - h)
}

func asig(l, i int, wg *sync.WaitGroup) {
	h[i] = 0
	for j := 1; j <= entradas; j++ {
		h[i] = h[i] + entrenam[l].x[j]*pesos[i][j]
	}
	f[i] = sigmoid(h[i])
	_f[i] = _sigmoid(f[i])
	defer wg.Done()
}

func hallarD(i int, wg *sync.WaitGroup) {
	dh[i] = _f[i] * pesos[neuronas][i] * dh[neuronas]
	defer wg.Done()
}

func ajustarPesos(l, i int, wg *sync.WaitGroup) {
	for j := 1; j <= entradas; j++ {
		pesos[i][j] = pesos[i][j] + a*dh[i]*entrenam[l].x[i]
	}
	defer wg.Done()
}

func imprimirAsig() {
	for i := 1; i <= neuronas; i++ {
		//fmt.Println("Neurona #", i, ":")
		//fmt.Println("h(", i, "): ", aprox(h[i]))
		//fmt.Println("f(", i, "): ", aprox(f[i]))
		//fmt.Println("f'(", i, "): ", aprox(_f[i]))
	}
}

func imprimirD() {
	for i := 1; i <= neuronas; i++ {
		fmt.Println("Neurona #", i, ":")
		fmt.Println("dh(", i, "): ", aprox(dh[i]))
	}
}

// Número de items de entrenamiento
const itemsEntrenam = 150

// Número de entradas
const entradas = 4

// Número de neuronas
const neuronas = entradas + 1

var entrenam [itemsEntrenam]datosEntrenam
var writerAux [entradas*entradas + entradas]float64
var pesos [neuronas + 1][entradas + 1]float64
var h [neuronas + 1]float64
var f [neuronas + 1]float64
var _f [neuronas + 1]float64
var dh [neuronas + 1]float64
var err float64
var a float64 = 0.25
var epoc int = 100

func main() {

	var wg sync.WaitGroup
	start := time.Now()

	var readerAux [entradas*entradas + entradas]float64 = read()
	cont := 0
	for i := 1; i <= neuronas; i++ {
		for j := 1; j <= entradas; j++ {
			pesos[i][j] = readerAux[cont]
			cont++
		}
	}
	entrenam = readcsv()

	for k := 0; k < epoc; k++ {
		for l := 0; l < itemsEntrenam; l++ {
			for i := 1; i < neuronas; i++ {
				wg.Add(1)
				asig(l, i, &wg)
			}
			wg.Wait()
			h[neuronas] = 0
			for j := 1; j <= entradas; j++ {
				h[neuronas] = h[neuronas] + f[j]*pesos[neuronas][j]
			}
			f[neuronas] = sigmoid(h[neuronas])
			_f[neuronas] = _sigmoid(f[neuronas])

			err = entrenam[l].z - f[neuronas]
			dh[neuronas] = _f[neuronas] * err

			for i := 1; i < neuronas; i++ {
				wg.Add(1)
				hallarD(i, &wg)
			}
			wg.Wait()
			for i := 1; i < neuronas; i++ {
				wg.Add(1)
				ajustarPesos(l, i, &wg)
			}
			wg.Wait()
			for j := 1; j <= entradas; j++ {
				pesos[neuronas][j] = pesos[neuronas][j] + a*dh[neuronas]*f[j]
			}
		}
	}
	cont = 0
	fmt.Println("Pesos Finales:\n")
	for i := 1; i <= neuronas; i++ {
		fmt.Println("Neurona #", i, ":")
		for j := 1; j <= entradas; j++ {
			fmt.Println("w", i, j, ": ", aprox(pesos[i][j]))
			writerAux[cont] = pesos[i][j]
			cont++
		}
	}
	write(writerAux)
	dur := time.Since(start)
	fmt.Println("\nDuración del programa: ", dur)

}
