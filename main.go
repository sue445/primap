package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sue445/primap/cron"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cron/sync_map", cron.SyncMapHandler)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
