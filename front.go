// +build front

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(os.Getenv("SERVER_ADDR"))
		if err != nil {
			fmt.Printf("error getting respond from back server: %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, "Front Message : %v", string(body))
	})
	if os.Getenv("SERVER_ADDR") == "" {
		os.Setenv("SERVER_ADDR", "http://localhost:8080")
	}

	fmt.Printf("server_addr: %s \n", os.Getenv("SERVER_ADDR"))
	fmt.Printf("front server listening on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
