package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al api reto calcular!")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var requestNums [][]int
	var responseNums [][]int
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}
	json.Unmarshal(reqBody, &requestNums)

	for i, item := range requestNums {
		var itemArr []int
		for j := range item {
			itemArr = append(itemArr, requestNums[j][i])
		}
		responseNums = append([][]int{itemArr}, responseNums...)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseNums)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/calcular", createTask).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
