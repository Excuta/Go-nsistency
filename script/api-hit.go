package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {

	url := "http://192.168.1.108:8080/increment"
	method := "POST"

	client := &http.Client{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go request(method, url, client, &wg)
	}
	wg.Wait()
}

func request(method string, url string, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(string(body))
}
