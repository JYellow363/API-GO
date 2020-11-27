package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	//fmt.Println("Pesos")
	for scanner.Scan() {
		arr[i], err = strconv.ParseFloat(scanner.Text(), 64)
		//fmt.Println(arr[i])
		i++
	}
	fmt.Println(" ")
	return arr
}

func write(pesos [entradas*entradas + entradas]float64) {
	f, err := os.Create("pesos.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	for _, peso := range pesos {
		_, err := f.WriteString(strconv.FormatFloat(peso, 'f', -1, 64) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("done")
}
