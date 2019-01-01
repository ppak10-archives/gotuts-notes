package main

import ("fmt"
"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whoa, Go is neat!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whoa, Go is neat 2!")
}

func main() {
  // creates server localhost:8000
  http.HandleFunc("/", index_handler)
  http.HandleFunc("/about/", about_handler)
  http.ListenAndServe(":8000", nil)
}
