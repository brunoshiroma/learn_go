package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	for i := 0; i <= 10000000000; i++ {
		log.Print(fmt.Sprintf("request %d", i))
		getRequest("http://localhost:8080")
	}
}

func getRequest(url string) {
	var netClient = &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := netClient.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if robots != nil {

	}
	//fmt.Printf("%s", robots)
}
