package main


import (
    "fmt"
    "log"
    "net/http"
);

func main(){
    http.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/octet-stream")
    })
    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
