package main

import (
"fmt"
"github.com/julienschmidt/httprouter"
"net/http"
"log"
"strconv"
"encoding/json"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getIntegerFromParams(str string) int {
	strVar := str
	intVar, _ := strconv.Atoi(strVar)
	return intVar
}

func fib() func() int {
	a, b, c := 0, 1, 0
	return func() int {
	c, a, b = a, b, a+b
	return c
	}
}

func GetNumbers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enableCors(&w)
	var maxFib = getIntegerFromParams(ps.ByName("number"))
	
	output := make([]int, 0)
	f := fib()
		for i := 0; i < maxFib; i++ {
			output = append(output, f())
		}
		jsonArray, _ := json.Marshal(output)
		fmt.Fprint(w, string(jsonArray))
}

func main() {
router := httprouter.New()
router.GET("/api/fibonacci/:number", GetNumbers)
log.Fatal(http.ListenAndServe(":8080", router))
}
