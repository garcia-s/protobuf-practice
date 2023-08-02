package main

import (
	"bytes"
	"fmt"

	"net/http"

	"github.com/garcia-s/protobuf-practice/client/proto/users"
	"google.golang.org/protobuf/proto"
)

func main() {
	url := "http://localhost:8080"
	createRoute := "/hello"

	user := users.NewUser{
		Username: "juanito",
		Password: "thisisapassword",
	}
	v, e := proto.Marshal(&user)
	if e != nil {
		panic(e)
	}

	req, _ := http.NewRequest("POST", url+createRoute, bytes.NewBuffer(v))
	req.Header.Set("Content-Type", "application/octet-stream")
	client := &http.Client{}
	client.Do(req)

	req2, _ := http.NewRequest("GET", url+"/getUsers", nil)

	req2.Header.Set("Content-Type", "application/octet-stream")

	client.Do(req2)


	fmt.Printf("Este es el responsecode %d", req2.Response.StatusCode )


    fmt.Printf("ESTE ES EL USUARIO %s",req2.Body)
}
