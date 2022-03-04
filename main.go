package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//new router func creates router and returns it.
//can use this func to instantiate and test router outside main func
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	//declare  new router
	//	r := mux.NewRouter()
	//accepts a path and a function as arguments
	//handler function has to have appropriate signature (as described by "handler" func below)
	//	r.HandleFunc("/hello", handler).Methods("GET")

	//router now formed by calling `newRouter` constructor func defined above
	r := newRouter()

	//after defining server, listen and serve on port 8080
	//second arg=handler--left as nil for now
	//handler defined above is used
	http.ListenAndServe(":8080", r)

}

//handler= handler func. Follows func signature of ResponseWriter and Request type
// as the args.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
