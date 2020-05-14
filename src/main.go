package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/compute", getResult).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe("127.0.0.1:5500", router))
}

func getResult(w http.ResponseWriter, r *http.Request) {
	typeInt, typeErr := strconv.Atoi(r.URL.Query().Get("type"))
	iFloat, iErr := strconv.ParseFloat(r.URL.Query().Get("i"), 64)
	nFloat, nErr := strconv.ParseFloat(r.URL.Query().Get("n"), 64)

	jsonMessage := map[string]string{
		"result":     "0",
		"typeStatus": "OK",
		"iStatus":    "OK",
		"jStatus":    "OK",
		"nStatus":    "OK",
	}

	if typeErr != nil {
		jsonMessage["typeStatus"] = "error"
	}
	if iErr != nil {
		jsonMessage["iStatus"] = "error"
	}
	if nErr != nil {
		jsonMessage["nStatus"] = "error"
	}

	var result float64

	if typeInt == 9 || typeInt == 10 {
		jFloat, jErr := strconv.ParseFloat(r.URL.Query().Get("j"), 64)
		if jErr != nil {
			jsonMessage["jStatus"] = "error"
		} else if typeErr == nil && iErr == nil && nErr == nil && jErr == nil {
			switch typeInt {
			case 9:
				result = PGivenAWithJ(iFloat, jFloat, nFloat)
			case 10:
				result = FGivenAWithJ(iFloat, jFloat, nFloat)
			}
			jsonMessage["result"] = strconv.FormatFloat(result, 'f', 4, 64)
		}

	} else if typeErr == nil && iErr == nil && nErr == nil {
		switch typeInt {
		case 1:
			result = PGivenF(iFloat, nFloat)
		case 2:
			result = FGivenP(iFloat, nFloat)
		case 3:
			result = PGivenA(iFloat, nFloat)
		case 4:
			result = AGivenP(iFloat, nFloat)
		case 5:
			result = FGivenA(iFloat, nFloat)
		case 6:
			result = AGivenF(iFloat, nFloat)
		case 7:
			result = PGivenG(iFloat, nFloat)
		case 8:
			result = AGivenG(iFloat, nFloat)
		}
		jsonMessage["result"] = strconv.FormatFloat(result, 'f', 4, 64)
	}

	// return computed result or error message for each variable
	json.NewEncoder(w).Encode(jsonMessage)
}
