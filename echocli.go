package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const entradas = 4
const neuronas = entradas + 1

// Response is a struct
type Response struct {
	LargoSepalo   float64 `json:"largoSepalo"`
	AnchoSepalo   float64 `json:"anchoSepalo"`
	LargoPetalo   float64 `json:"largoPetalo"`
	AnchoPetalo   float64 `json:"anchoPetalo"`
	Clasificacion string  `json:"clasificacion"`
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

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getClasification(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)["parameters"]
	msg := parameters + "\n"

	con1, _ := net.Dial("tcp", "localhost:8001")
	con2, _ := net.Dial("tcp", "localhost:8002")
	con3, _ := net.Dial("tcp", "localhost:8003")
	con4, _ := net.Dial("tcp", "localhost:8004")

	defer con1.Close()
	defer con2.Close()
	defer con3.Close()
	defer con4.Close()

	r1 := bufio.NewReader(con1)
	r2 := bufio.NewReader(con2)
	r3 := bufio.NewReader(con3)
	r4 := bufio.NewReader(con4)

	if cont == 1 {
		fmt.Fprint(con1, msg)
		cont++
		msg, _ = r1.ReadString('\n')
	} else if cont == 2 {
		fmt.Fprint(con2, msg)
		cont++
		msg, _ = r2.ReadString('\n')
	} else if cont == 3 {
		fmt.Fprint(con3, msg)
		cont++
		msg, _ = r3.ReadString('\n')
	} else if cont == 4 {
		fmt.Fprint(con4, msg)
		cont = 1
		msg, _ = r4.ReadString('\n')
	}

	fmt.Print("Respuesta: ", msg)
	fmt.Print("\n")

	w.Write([]byte(msg))
}

func postClasification(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var resp Response
	_ = json.Unmarshal(reqBody, &resp)
	fmt.Println(resp)
	
	x1 := fmt.Sprintf("%f", resp.LargoSepalo)
	x2 := fmt.Sprintf("%f", resp.AnchoSepalo)
	x3 := fmt.Sprintf("%f", resp.LargoPetalo)
	x4 := fmt.Sprintf("%f", resp.AnchoPetalo)
	msg := x1 + "," + x2 + "," + x3 + "," + x4 + "\n"

	con1, _ := net.Dial("tcp", "localhost:8001")
	con2, _ := net.Dial("tcp", "localhost:8002")
	con3, _ := net.Dial("tcp", "localhost:8003")
	con4, _ := net.Dial("tcp", "localhost:8004")

	defer con1.Close()
	defer con2.Close()
	defer con3.Close()
	defer con4.Close()

	r1 := bufio.NewReader(con1)
	r2 := bufio.NewReader(con2)
	r3 := bufio.NewReader(con3)
	r4 := bufio.NewReader(con4)

	if cont == 1 {
		fmt.Fprint(con1, msg)
		cont++
		msg, _ = r1.ReadString('\n')
	} else if cont == 2 {
		fmt.Fprint(con2, msg)
		cont++
		msg, _ = r2.ReadString('\n')
	} else if cont == 3 {
		fmt.Fprint(con3, msg)
		cont++
		msg, _ = r3.ReadString('\n')
	} else if cont == 4 {
		fmt.Fprint(con4, msg)
		cont = 1
		msg, _ = r4.ReadString('\n')
	}

	fmt.Print("Respuesta: ", msg)
	fmt.Print("\n")

	w.Write([]byte(msg))
}

var cont int = 1

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/iris", postClasification).Methods("POST")
	router.HandleFunc("/iris/{parameters}", getClasification).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
