package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "80", "port on localhost to check")
	flag.Parse()

	resp, err := http.Get("http://127.0.0.1:" + *port + "/health")
	if err != nil || resp.StatusCode != 200 {
		log.Println("down")
		os.Exit(1)
	}
	os.Exit(0)
}
