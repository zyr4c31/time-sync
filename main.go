package main

import (
	"net/http"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		formatString := "January 02, 2006 03:04:05 PM"
		datetime := time.Now().Local().Format(formatString)
		component := Index(datetime)
		component.Render(r.Context(), w)
	})

	http.HandleFunc("GET /time", func(w http.ResponseWriter, r *http.Request) {
		formatString := "January 02, 2006 03:04:05 PM"
		datetime := time.Now().Local().Format(formatString)
		change := Change(datetime)
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(change)
	})

	http.ListenAndServe(":8080", nil)
}
