package main

import {
    "log"
    "net/http"
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)) {
        fmt.Printf("Hello World 2!")
    }

    log.Fatal(http.ListenAndServe(":80", nil))
}
