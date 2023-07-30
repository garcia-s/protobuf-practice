package main

import (
	"fmt"
	"net/http"
    "github.com/garcia-s/protobuf-practice/client/fake"
)

func main() {
	url := "http://localhost:8080/client"
    user := pb.NewUser(

    )
    r, err := http.NewRequest("POST", url)
    if(err != nil) {
        panic(err)
    }
    fmt.Println("Body:" ,  r.Body);
}
