package main

//Board to display messages.
// API based with a VUE frontend
// Goal Create Teams
// Teams can have Members
// Only members can view the boards
// ACLs have to be used to implement this.

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
