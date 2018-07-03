package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	cars := Cars{
		Car{
			Id: 1,
    	Fuel: "Gás",
    	Brand: "Fiat",
    	Model: "Uno",
    	LicensePlate: "ITP-1224"},
		Car{
			Id: 2,
    	Fuel: "Gás",
    	Brand: "Ford",
    	Model: "Ka",
    	LicensePlate: "ITP-1000"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		panic(err)
	}
}

func GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo Show:", todoId)
}