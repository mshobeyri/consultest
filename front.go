// +build front

package main

import (
	"consultest/pb"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	if os.Getenv("SERVER_ADDR") == "" {
		os.Setenv("SERVER_ADDR", "localhost:8080")
	}

	var i int32
	i = 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := grpc.Dial(os.Getenv("SERVER_ADDR"), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewCountingClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := c.GetMessage(ctx, &pb.Message{
			Text:   "Front",
			Number: i,
		})

		i++

		if err != nil {
			fmt.Fprintf(w, "could not get message: %v", err)
			return
		}
		fmt.Fprintf(w, "Front Message : %v %v\n", res.GetText(), res.GetNumber())

	})

	fmt.Printf("server_addr: %s \n", os.Getenv("SERVER_ADDR"))
	fmt.Printf("front server is listening on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
