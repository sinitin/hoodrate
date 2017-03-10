package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HoodInfo struct {
	API struct {
		Name     string `json:"name"`
		URL      string `json:"url"`
		Version  string `json:"version"`
		Encoding string `json:"encoding"`
	} `json:"api"`
	Result []struct {
		Street       string `json:"street"`
		Number       string `json:"number"`
		Zipcode      string `json:"zipcode"`
		City         string `json:"city"`
		Municipality string `json:"municipality"`
		Code         string `json:"code"`
		State        string `json:"state"`
	} `json:"result"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func HoodShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered the hoodshow")
	vars := mux.Vars(r)
	var areaCode int
	var err error
	if areaCode, err = strconv.Atoi(vars["areaCode"]); err != nil {
		panic(err)
	}

	if areaCode > 0 {

		//lookup which area it is with the pap-api
		areaCodeString := strconv.Itoa(areaCode)
		first := areaCodeString[0:3]
		last := areaCodeString[3:len(areaCodeString)]
		url := "https://papapi.se/json/?z=" + first + "+" + last + "&token=15e0f9807b307da4ea6b7b19749498788459692d"
		resp, _ := http.Get(url)
		fmt.Printf("%+v\n", resp)

		decoder := json.NewDecoder(resp.Body)
		var data HoodInfo
		err = decoder.Decode(&data)

		fmt.Printf("%+v\n", data.Result[0].City)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data.Result[0].City); err != nil {
			panic(err)
		}

		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

}
