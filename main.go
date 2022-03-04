package main

import (
	"fmt"
	"net/http"
)

func main() {
	//accepts a path and a function as arguments
	//handler function has to have appropriate signature (as described by "handler" func below)
	http.HandleFunc("/", handler)

	//after defining server, listen and serve on port 8080
	//second arg=handler--left as nil for now
	//handler defined above is used
	http.ListenAndServe(":8080", nil)

}

//handler= handler func. Follows func signature of ResponseWriter and Request type
// as the args.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
