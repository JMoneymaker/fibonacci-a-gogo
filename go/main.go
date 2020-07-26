package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"strconv"
)

type UserInput struct {
	Input int `validate:"min=1,max=300"`
}

func Fib() func() int {
	a, b, c := 0, 1, 0
	return func() int {
		c, a, b = a, b, a+b
		return c
	}
}

func GetNumbers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	maxFib, _ := strconv.Atoi(ps.ByName("number"))
	query := UserInput{Input: maxFib}
	
	if errs := validator.Validate(query); errs != nil {
		fmt.Fprint(w, "Please enter a number between 1 and 300.")
	} else {
		response := make([]int, 0)
		f := Fib()
		for i := 0; i < maxFib; i++ {
			response = append(response, f())
		}
		jsonArray, _ := json.Marshal(response)
		fmt.Fprint(w, string(jsonArray))
	}

}

func main() {
	router := httprouter.New()
	router.GET("/api/fibonacci/:number", GetNumbers)
	log.Fatal(http.ListenAndServe(":8080", router))
}
