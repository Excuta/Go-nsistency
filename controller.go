package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/counters", getAllHanlder)
	router.GET("/counters/:id", makeCountersHandler(getHanlder))
	router.POST("/counters/:id", makeCountersHandler(editHanlder))

	log.Fatal(http.ListenAndServe("", router))
}

func makeCountersHandler(fn func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, counterId string)) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		if id == "" {
			http.Error(w, "Counter id is required", http.StatusBadRequest)
			return
		}
		fn(w, r, ps, id)
	}

}

func getAllHanlder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("get all"))
}

func getHanlder(w http.ResponseWriter, r *http.Request, ps httprouter.Params, counterId string) {
	w.Write([]byte("get"))
}

func editHanlder(w http.ResponseWriter, r *http.Request, ps httprouter.Params, counterId string) {
	w.Write([]byte("edit"))
}
