package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func readcsv() [itemsEntrenam]datosEntrenam {

	var entrenam [itemsEntrenam]datosEntrenam
	cont := 0
	var aux [entradas + 1]float64
	aux[0] = -1
	csvfile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		for i := 0; i < entradas; i++ {
			aux[i+1], err = strconv.ParseFloat(record[i], 64)
		}

		if record[entradas] == "setosa" {
			entrenam[cont] = datosEntrenam{aux, -1}
		}
		if record[entradas] == "versicolor" {
			entrenam[cont] = datosEntrenam{aux, 0}
		}
		if record[entradas] == "virginica" {
			entrenam[cont] = datosEntrenam{aux, 1}
		}
		cont++
	}
	return entrenam
}
