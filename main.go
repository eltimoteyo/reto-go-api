package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	//"path"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/paytm/grace.v1"
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
	//router := mux.NewRouter()

	// router.HandleFunc("/", indexRoute)
	// router.HandleFunc("/calcular", calcular).Methods("POST")
	// log.Fatal(http.ListenAndServe(":80", router))

	// log.Fatal(srv.ListenAndServe())
	// fmt.Println("Go program")
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", indexRoute)
	muxRouter.HandleFunc("/calcular", calcular).Methods("POST")
	http.Handle("/", muxRouter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	grace.Serve(":"+port, context.ClearHandler(http.DefaultServeMux))

	// server := echo.New()
	// server.GET(path.Join("/"), indexRoute)
}
