package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world, baineng!")
	info := fmt.Sprintln(r.Header.Get("Content-Type"), r.Host, r.URL)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, info, string(body))
}
func indexPost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "hello world, Jd!")
	// info := fmt.Sprintln("URL", r.URL, "HOST", r.Host, "Method", r.Method, "RequestURL", r.RequestURI, "RawQuery", r.URL.RawQuery)
	info := fmt.Sprintln(r.Header.Get("Content-Type"), r.Host, r.URL)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, info, string(body))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/post", indexPost)
	http.ListenAndServe(":8000", nil)

}
