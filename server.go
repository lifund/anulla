package main

import (
		"fmt"
		"log"
		"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })
	
	fmt.Println("[anulla] go server at port 3004")
	if err := http.ListenAndServe(":3004", nil); err != nil {
        log.Fatal(err)
    } 
}
