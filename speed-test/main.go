package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
)

func main() {
    url := os.Getenv("DOWNLOAD_URL")
    if url == "" {
        log.Fatalln("DOWNLOAD_URL environment variable not set")
    }

    filename := os.Getenv("SAVE_FILE")
    if filename == "" {
        log.Fatalln("SAVE_FILE environment variable not set")
    }

    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }
    
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    
    end := time.Now()
    
    // time in ms
    elapsed_second_time := float64(end.Sub(start) / 1000000) / 1000
    filesize := float64(len(body))

    // speed in Mbps
    speed := filesize*8/(1024*1024)/elapsed_second_time

    fmt.Println(time.Now())
    fmt.Println(speed)
    
    _, err = os.Stat(filename)
    if os.IsNotExist(err) {
      var file, err = os.Create(filename)
      if err != nil {
        log.Fatalln(err)
      }
      defer file.Close()
    }
    
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalln(err)
    }

    defer f.Close()
    if _, err = f.WriteString(start.Format(time.RFC3339) + "," + strconv.FormatFloat(speed, 'E', -1, 64) + "\n"); err != nil {
        log.Fatalln(err)
    }
}