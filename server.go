package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		delay := r.URL.Query().Get("delay")
		if delay == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		delayDuration, err := time.ParseDuration(delay)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		time.Sleep(delayDuration)
		fmt.Fprintf(w, "Response was delayed by %s", delayDuration)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
