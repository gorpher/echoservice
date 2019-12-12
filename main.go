package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	if err := http.ListenAndServe(":8901", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
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
