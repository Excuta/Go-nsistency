package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"excuta/go-nsistency/service"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/getcounter", headersHandler(getCounterHandler))
	router.POST("/increment", headersHandler(incrementHanlder))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("GON_PORT")), router))
}

func getCounterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	resp, err := service.GetCounter()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	w.Write(resp)
}
func incrementHanlder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := service.Increment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func headersHandler(fn func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Content-Type", "application/json")
		fn(w, r, ps)
	}
}
