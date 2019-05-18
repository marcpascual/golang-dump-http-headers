package main

import (
	"fmt"
	"net/http"
	"strings"
)

// https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
func handler(w http.ResponseWriter, r *http.Request) {
	var name string
	// Create return string
	var request []string

	found := r.URL.Query().Get("name")
	if found != "" {
		name = found
	} else {
		name = "world"
	}

	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	fmt.Fprintf(w, "Hello, %s!\n", name)
	fmt.Fprintf(w, strings.Join(request, "\n"))
}

func main() {
	fmt.Print("connect to localhost port 3000\n")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
