package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:              ":8080",
		ReadTimeout:       time.Hour,
		ReadHeaderTimeout: time.Hour,
		WriteTimeout:      time.Hour,
		ErrorLog:          log.New(os.Stderr, "[http]", log.LstdFlags),
	}
	fmt.Println(srv.ListenAndServe())
}
