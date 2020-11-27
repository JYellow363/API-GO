package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

// Response is a struct
type Response struct {
	Localhost     string  `json:"localhost"`
	LargoSepalo   float64 `json:"largoSepalo"`
	AnchoSepalo   float64 `json:"anchoSepalo"`
	LargoPetalo   float64 `json:"largoPetalo"`
	AnchoPetalo   float64 `json:"anchoPetalo"`
	Clasificacion string  `json:"clasificacion"`
}

const entradas = 4
const neuronas = entradas + 1

var pesos [neuronas + 1][entradas + 1]float64

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

func hallaF(x1, x2, x3, x4 float64) string {
	f1 := sigmoid(x1*pesos[1][1] + x2*pesos[1][2] + x3*pesos[1][3] + x4*pesos[1][4])
	f2 := sigmoid(x1*pesos[2][1] + x2*pesos[2][2] + x3*pesos[2][3] + x4*pesos[2][4])
	f3 := sigmoid(x1*pesos[3][1] + x2*pesos[3][2] + x3*pesos[3][3] + x4*pesos[3][4])
	f4 := sigmoid(x1*pesos[4][1] + x2*pesos[4][2] + x3*pesos[4][3] + x4*pesos[4][4])
	fy := sigmoid(f1*pesos[5][1] + f2*pesos[5][2] + f3*pesos[5][3] + f4*pesos[5][4])

	if fy < 0.45 {
		return "setosa"
	}
	if fy < 0.575 && fy >= 0.45 {
		return "versicolor"
	}
	if fy >= 0.575 {
		return "virginica"
	}
	return ""
}

func main() {
	ln, _ := net.Listen("tcp", "localhost:8003")
	defer ln.Close()

	var readerAux [entradas*entradas + entradas]float64 = read()
	cont := 0
	for i := 1; i <= neuronas; i++ {
		for j := 1; j <= entradas; j++ {
			pesos[i][j] = readerAux[cont]
			cont++
		}
	}

	for {
		con, _ := ln.Accept()
		go handle(con)
	}
}

func leerRow(msg string) []string {
	msgReader := strings.NewReader(msg)
	r := csv.NewReader(msgReader)
	record, err := r.Read()
	if err == io.EOF {
		fmt.Println("Error")
	}
	return record
}

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	var aux [entradas + 1]float64
	for {
		msg, _ := r.ReadString('\n')
		if len(msg) > 0 {
			fmt.Printf("%s: %s", con.RemoteAddr(), msg)
			record := leerRow(msg)
			if len(record) > 0 {
				for i := 0; i < entradas; i++ {
					aux[i+1], _ = strconv.ParseFloat(record[i], 64)
				}
				y := hallaF(aux[1], aux[2], aux[3], aux[4])

				data := &Response{}
				data.LargoSepalo = aux[1]
				data.AnchoSepalo = aux[2]
				data.LargoPetalo = aux[3]
				data.AnchoPetalo = aux[4]
				data.Localhost = "8003"
				data.Clasificacion = y

				//response := "localhost:8003 -> " + record[0] + " " + record[1] + " " + record[2] + " " + record[3] + " -> " + y + "\n"
				responseJSON, _ := json.Marshal(data)
				fmt.Fprint(con, string(responseJSON)+"\n")
			}
		}
	}
}
