package main
import (
"fmt"
"github.com/julienschmidt/httprouter"
"net/http"
"log"
"strconv"
)

func fib() func() int {
	a, b, c := 0, 1, 0
	return func() int {
	c, a, b = a, b, a+b
	return c
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)
	fmt.Fprint(w, "Welcome!\n")
}

func getIntegerFromParams(str string) int {
	strVar := str
	intVar, _ := strconv.Atoi(strVar)
	return intVar
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enableCors(&w)
	w.Header().Set("content-type", "text/plain")
	var userInput = ps.ByName("number")
	var maxFib = getIntegerFromParams(userInput)
	f := fib()
		for i := 0; i < maxFib; i++ {
		fmt.Fprint(w, f())
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
router := httprouter.New()
router.GET("/api", Index)
router.GET("/api/fibonacci/:number", Hello)
log.Fatal(http.ListenAndServe(":8080", router))
}