package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := flag.String("url", "http://127.0.0.1", "Address URL")
	data := flag.String("data", "", "Data to post")
	flag.Parse()
	reponse := doPost(url, data)
	fmt.Println(string(reponse))
}

func doPost(url, data *string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("POST", *url, strings.NewReader(*data))
	if err != nil {
		fmt.Println("Error: create request fail")
		return nil
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(request)
	if response.StatusCode != 200 {
		fmt.Println("Error: return status code is not 200")
		fmt.Printf("%+v\n", response)
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: read response body fail")
		return nil
	}
	return body
}
