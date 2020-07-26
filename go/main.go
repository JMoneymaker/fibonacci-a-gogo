package main

import (
"fmt"
"github.com/julienschmidt/httprouter"
"net/http"
"log"
"strconv"
"encoding/json"
)

func fib() func() int {
	a, b, c := 0, 1, 0
	return func() int {
	c, a, b = a, b, a+b
	return c
	}
}

func GetNumbers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	maxFib, _ := strconv.Atoi(ps.ByName("number"))
	response := make([]int, 0)
	f := fib()
		for i := 0; i < maxFib; i++ {
			response = append(response, f())
		}
		jsonArray, _ := json.Marshal(response)
		fmt.Fprint(w, string(jsonArray))
}

func main() {
router := httprouter.New()
router.GET("/api/fibonacci/:number", GetNumbers)
log.Fatal(http.ListenAndServe(":8080", router))
}
