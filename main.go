package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		result := map[string]interface{}{
			"host":        r.Host,
			"hostname":    hostname,
			"method":      r.Method,
			"url":         r.URL.String(),
			"remote_addr": r.RemoteAddr,
		}
		js, _ := json.MarshalIndent(result, "", "  ")
		log.Println(string(js))
		w.Write(js)
	})

	addr := "0.0.0.0:80"
	log.Printf("Starting listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
