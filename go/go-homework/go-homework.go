package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/calc/sum/{valA}/{valB}", sum)
	router.HandleFunc("/calc/subtract/{valA}/{valB}", subtract)
	router.HandleFunc("/calc/multiply/{valA}/{valB}", multiply)
	router.HandleFunc("/calc/divide/{valA}/{valB}", divide)

	http.ListenAndServe(":8080", router)
}

func subtract(w http.ResponseWriter, r *http.Request) {
	intValA, intValB, err := parseAndCheckValues(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "The subtraction result is: %d", intValA-intValB)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	intValA, intValB, err := parseAndCheckValues(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "The multiplication result is: %d", intValA*intValB)
}

func divide(w http.ResponseWriter, r *http.Request) {
	intValA, intValB, err := parseAndCheckValues(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "The division result is: %d", intValA/intValB)
}

func sum(w http.ResponseWriter, r *http.Request) {
	intValA, intValB, err := parseAndCheckValues(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "The sum result is: %d", intValA+intValB)
}

func parseAndCheckValues(w http.ResponseWriter, r *http.Request) (int64, int64, error) {
	vars := mux.Vars(r)

	intValA, err := parseAndCheckIntValue(vars["valA"], w)
	if err != nil {
		return 0, 0, err
	}
	intValB, err := parseAndCheckIntValue(vars["valB"], w)
	if err != nil {
		return 0, 0, err
	}
	return intValA, intValB, nil
}

func parseAndCheckIntValue(val string, w http.ResponseWriter) (int64, error) {
	varA, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Parameter must '%s' be an Integer", val)
		return 0, err
	}
	return varA, err
}
