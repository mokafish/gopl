//go:build ignore

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/view", viewer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Printf("Header[%q] = %q\n", k, v)
	}
	fmt.Printf("Host = %q\n", r.Host)
	fmt.Printf("RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Printf("Form[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func viewer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
