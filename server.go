// +build server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var i int
	i = 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		i++
		fmt.Fprintf(w, "Server Message %d", i)

	})

	fmt.Printf("bacl server listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
