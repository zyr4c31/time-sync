package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

type Store struct {
	Date string `json:"date"`
	TZ   string `json:"tz"`
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		date := time.Now().Local()
		component := Index(date.String())
		component.Render(r.Context(), w)
	})

	http.HandleFunc("POST /time", func(w http.ResponseWriter, r *http.Request) {
		store := &Store{}
		if err := datastar.ReadSignals(r, store); err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Printf("store.Date: %v\n", store.Date)
		fmt.Printf("store.TZ: %v\n", store.TZ)

		location, err := time.LoadLocation(store.TZ)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		sse := datastar.NewSSE(w, r)
		if err := sse.PatchElementTempl(Change(location.String())); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	http.HandleFunc("GET /alarms", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r: %v\n", r)
	})

	addr := "http://localhost:8080"
	fmt.Printf("addr: %v\n", addr)
	http.ListenAndServe(":8080", nil)
}
