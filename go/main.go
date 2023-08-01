package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/garcia-s/protobuf-practice/proto/users"
	"google.golang.org/protobuf/proto"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplication/octet-stream")
		newUser := &users.NewUser{}
		body, _ := io.ReadAll(r.Body)
        fmt.Printf(" This should be the raw bytes: %x", body)
		_ = proto.Unmarshal(body, newUser)

		fmt.Printf("This is the string %s", newUser)
	})
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
