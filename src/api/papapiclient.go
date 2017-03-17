package main

import (
	"encoding/json"
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

func lookupArea(area string) (HoodInfo, error) {

	var data HoodInfo
	var areaCode int
	var err error
	if areaCode, err = strconv.Atoi(area); err != nil {
		return data, err
	}

	if areaCode > 0 {

		//lookup which area it is with the pap-api
		areaCodeString := strconv.Itoa(areaCode)
		first := areaCodeString[0:3]
		last := areaCodeString[3:len(areaCodeString)]
		url := "https://papapi.se/json/?z=" + first + "+" + last + "&token=15e0f9807b307da4ea6b7b19749498788459692d"
		resp, _ := http.Get(url)
		decoder := json.NewDecoder(resp.Body)
		if err = decoder.Decode(&data); err != nil {
			return data, err
		}
	}

	return data, nil

}
