package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	//"path"

	"github.com/gorilla/mux"
	//"github.com/labstack/echo/v4"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al api reto calcular!")
}

func calcular(w http.ResponseWriter, r *http.Request) {
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
	// router := mux.NewRouter().StrictSlash(true)

	// router.HandleFunc("/", indexRoute)
	// router.HandleFunc("/calcular", createTask).Methods("POST")
	// log.Fatal(http.ListenAndServe(":3000", router))

	r := mux.NewRouter()
	r.HandleFunc("/", indexRoute)
	r.HandleFunc("/calcular", calcular)

	srv := &http.Server{
		Handler: r,
		Addr:    "https://reto-go-api.herokuapp.com",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	// fmt.Println("Go program")

	// server := echo.New()
	// server.GET(path.Join("/"), indexRoute)
}
