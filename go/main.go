package main

import (
	"fmt"
	"github.com/garcia-s/protobuf-practice/proto/users"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
	"os"
)

func check(e error) {
    if(e != nil) {
        panic(e)
    } 
}

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplication/octet-stream")
		newUser := &users.NewUser{}
		body, _ := io.ReadAll(r.Body)
		err := os.WriteFile("./userbinary", body, 0777)
        check(err)
		_ = proto.Unmarshal(body, newUser)

	})

	http.HandleFunc("/getUsers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplication/octet-stream")
		file, err := os.Open("./userbinary")
        check(err)
		bits := make([]byte, 5)
		_, er := file.Read(bits)
        check(er)
		user := &users.NewUser{}
		_ = proto.Unmarshal(bits, user)

	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}


