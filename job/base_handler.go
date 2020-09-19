package job

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	log.Printf("[ERROR] %+v", err)
	sentry.CaptureException(err)
	w.WriteHeader(500)
	fmt.Fprint(w, err)
}
