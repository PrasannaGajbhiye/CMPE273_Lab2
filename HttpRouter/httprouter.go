package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type test_struct struct {
	Name string
}
type test_structRespone struct {
	Greeting string
}

func getHello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func greet(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic("Error in reading body.")
	}

	var t test_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic("Error in unmarshalling.")
	}

	var res test_structRespone
	res.Greeting = "Hello, " + t.Name + "!"

	mapB, _ := json.Marshal(res)

	fmt.Fprintf(rw, string(mapB)+"\n")
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", getHello)
	mux.POST("/hello/:name", hello)
	mux.POST("/hello", greet)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
