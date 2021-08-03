package graph

import (
	"encoding/json"
	"fmt"
	"graphgl-demo/graph/model"
	"net/http"
)

type ServerInterfaceCustom interface {

}

type RootHandler func(http.ResponseWriter, *http.Request) error

func (re *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	println(r.Body)

}

func CustomHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("loan request", r.Body)

		var newLoan model.NewLoan
		if err := json.NewDecoder(r.Body).Decode(&newLoan); err != nil {
			println("error in parsing" , r.Body)
		}

		fmt.Printf("loan parsed", newLoan)
	})
}



