package cron

import (
	"fmt"
	"net/http"
)

// SyncMapHandler returns handler of /cron/sync_map
func SyncMapHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
