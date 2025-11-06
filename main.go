package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

type Store struct {
	Date         string `json:"date"`
	UTCOffsetMin int    `json:"tzoffset"`
	TZ           string `json:"tz"`
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		formatString := "January 02, 2006 03:04:05 PM"
		datetime := time.Now().Local().Format(formatString)
		component := Index(datetime)
		component.Render(r.Context(), w)
	})

	http.HandleFunc("POST /time", func(w http.ResponseWriter, r *http.Request) {
		store := &Store{}
		_ = datastar.ReadSignals(r, store)
		fmt.Printf("store.TZ: %v\n", store.TZ)
		offsetSeconds := store.UTCOffsetMin * -60
		userTimeZone := time.FixedZone("user", offsetSeconds)

		t := time.Now().In(userTimeZone)
		fmt.Printf("t: %v\n", t)
		change := Change(t.Local().String())
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(change)
	})

	http.ListenAndServe(":8080", nil)
}
