package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    cwd, _ := os.Getwd()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles(filepath.Join(cwd, "templates/index.html"))
        if err != nil {
            log.Fatalf("template parse: %s", err)
        }

        err = tmpl.Execute(w, nil)
        if err != nil {
            log.Fatalf("template execution: %s", err)
        }
    })

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(cwd, "static")))))

    fmt.Println("Starting webserver at http://localhost:80")
    log.Fatal(http.ListenAndServe(":80", nil))
}
