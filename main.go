package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var delaySeconds int

func sayHello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(delaySeconds) * time.Second)
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	if delay := os.Getenv("DELAY"); delay != "" {
		i, err := strconv.Atoi(delay)
		if err != nil {
			panic(err)
		}
		delaySeconds = i
	}

	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
