package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/counters", headersHandler(getAllHanlder))
	router.GET("/counters/:id", makeCountersHandler(getHanlder))
	router.POST("/counters/:id", makeCountersHandler(incrementHanlder))
	router.HandleOPTIONS = true

	log.Fatal(http.ListenAndServe("", router))
}

func makeCountersHandler(fn func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, counterId string)) httprouter.Handle {
	return headersHandler(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		id := ps.ByName("id")
		if id == "" {
			http.Error(w, "Counter id is required", http.StatusBadRequest)
			return
		}
		fn(w, r, ps, id)
	})

}

func getAllHanlder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	resp, err := GetAll()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	w.Write(resp)
}

func getHanlder(w http.ResponseWriter, r *http.Request, ps httprouter.Params, counterId string) {
	w.Write([]byte("{\"res\":\"get\"}"))
}

func incrementHanlder(w http.ResponseWriter, r *http.Request, ps httprouter.Params, counterId string) {
	w.Write([]byte("{\"res\":\"increment\"}"))
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
