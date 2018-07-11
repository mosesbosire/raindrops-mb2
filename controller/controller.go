package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/armakuni/raindrops-mb2/raindrops"

	"github.com/gorilla/mux"
)

func Start() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", raindropsController).Methods("GET")

	return router
}

func raindropsController(w http.ResponseWriter, r *http.Request) {
	numberParamSlice := r.URL.Query()["number"]
	if len(numberParamSlice) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "To use this API please provide a 'number' param")
		return
	}
	number, err := strconv.Atoi(numberParamSlice[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "The 'number' param provided was '%s', it must be a valid integer", numberParamSlice[0])
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", raindrops.Process(number))
	return
}
