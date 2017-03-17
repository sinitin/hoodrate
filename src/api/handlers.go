package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Rating struct {
	City      string `json:"city"`
	Area      string `json:"area"`
	Rating    int    `json:"rating"`
	Potential int    `json:"potential"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func HoodShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered the hoodshow")
	vars := mux.Vars(r)

	data, err := lookupArea(vars["areaCode"])
	if err != nil {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
	}

	//askBooli()
	//averagecurrent := 44000
	//averagestockholm := 90000

	var rating Rating

	ratingcurrent := 5
	potential := 7

	rating.City = data.Result[0].City
	rating.Rating = ratingcurrent
	rating.Potential = potential

	fmt.Printf("%+v\n", rating)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(rating); err != nil {
		panic(err)
	}

	return

}
