package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"
)

func askBooli() {

	timestamp := time.Now().Unix()
	fmt.Println("")
	fmt.Println("The unix timestamp")
	fmt.Println("")
	fmt.Println(timestamp)
	fmt.Println("")

	r := randString(16)

	fmt.Println("")
	fmt.Println("The random string")
	fmt.Println("")
	fmt.Println(r)
	fmt.Println("")

	fmt.Println("")
	fmt.Println("The concatenated string")
	fmt.Println("")
	fmt.Printf("reversed_guide%d6b1MHSquuAAWBpRDqfFhldrdIVdlvQLYfCzhAR1v%s", timestamp, r)
	fmt.Println("")
	ba := []byte(fmt.Sprintf("reversed_guide%d6b1MHSquuAAWBpRDqfFhldrdIVdlvQLYfCzhAR1v%s", timestamp, r))

	sh := sha1.Sum(ba)
	s := fmt.Sprintf("%x", sh)

	fmt.Println("")
	fmt.Println("The sha1 string")
	fmt.Println("")
	fmt.Printf("%s", s)
	fmt.Println("")

	//url := fmt.Sprintf("https://www.booli.se/api/sold?q=solna&minPublished=20150313&maxPublished=20160313&minRooms=2&maxRooms=2&callerId=reversed_guide&time=%d&unique=%s&hash=%s", t, r, s)
	url := fmt.Sprintf("https://www.booli.se/api/sold?q=solna&callerId=reversed_guide&time=%d&unique=%s&hash=%s", timestamp, r, s)
	//url := fmt.Sprintf("https://www.booli.se/api/sold?q=solna&callerId=reversed_guide&time=%d&&hash=%s", t, s)

	//url := "https://www.booli.se/api/sold?q=solna"

	fmt.Println("")
	fmt.Println("The url")
	fmt.Println("")
	fmt.Println(url)
	fmt.Println("")

	//resp, _ := http.Get(url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.booli-v2+json")
	//r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	fmt.Println("")
	fmt.Println("The request")
	fmt.Println("")
	//fmt.Printf("%+v\n", req)

	// Save a copy of this request for debugging.
	requestDump, _ := httputil.DumpRequest(req, false)
	fmt.Println(string(requestDump))
	fmt.Println("")

	resp, _ := client.Do(req)

	respDump, _ := httputil.DumpResponse(resp, false)

	fmt.Println("")
	fmt.Println("The response")
	fmt.Println("")
	//fmt.Printf("%+v\n", resp)
	fmt.Println(string(respDump))
	fmt.Println("")
}

func randString(n int) string {

	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
