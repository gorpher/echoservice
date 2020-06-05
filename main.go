package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	if err := http.ListenAndServe(":8901", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.URL.Path == "/health" {
			m := map[string]interface{}{}
			hostname, err := os.Hostname()
			if err != nil {
				panic(err)
			}
			home, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}
			m["hostname"] = hostname
			m["home"] = home
			m["status"] = "up"
			m["start_time"] = startTime
			res, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}
		}
		if r.Method == http.MethodGet {
			_, err := w.Write([]byte(r.URL.RawQuery))
			if err != nil {
				panic(err)
			}
		}

		if r.Method == http.MethodPost && strings.Contains(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}
			m := map[string]interface{}{}
			for k := range r.PostForm {
				m[k] = r.PostForm.Get(k)
			}
			res, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}
		}

		if r.Method == http.MethodPost && r.Header.Get("Content-Type") == "" {
			_, err := w.Write([]byte(r.URL.RawQuery))
			if err != nil {
				panic(err)
			}
		}

		if r.Method == http.MethodPost && r.Header.Get("Content-Type") == "application/json" {
			for k := range r.Header {
				w.Header().Add(k, r.Header.Get(k))
				_, err := io.Copy(w, r.Body)
				if err != nil {
					panic(err)
				}
			}
		}
	})); err != nil {
		log.Fatal(err)
	}
}
