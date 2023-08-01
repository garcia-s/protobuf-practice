package main

import (
    "bytes"
	"fmt"
	"github.com/garcia-s/protobuf-practice/client/proto/users"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func main() {
	url := "http://localhost:8080/client"
	user := users.NewUser{
		Username: "juanito",
		Password: "thisisapassword",
	}
	v, e := proto.Marshal(user)
	if e != nil {
		panic(e)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(v))

    req.Header.Set("Content-Type","application/octet-stream");
    client := &http.Client{}
    client.Do(req);

	fmt.Println("Body:", req.Body)
}
