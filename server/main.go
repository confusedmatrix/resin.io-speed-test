package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        cwd, _ := os.Getwd()
        tmpl, err := template.ParseFiles(filepath.Join(cwd, "templates/index.html"))
        if err != nil {
            log.Fatalf("template parse: %s", err)
        }

        err = tmpl.Execute(w, nil)
        if err != nil {
            log.Fatalf("template execution: %s", err)
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
