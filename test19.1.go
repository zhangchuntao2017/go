package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	reponse, err := http.Get("https://ifconfig.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer reponse.Body.Close()
	body, err := ioutil.ReadAll(reponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}
