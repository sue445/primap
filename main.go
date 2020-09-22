package main

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/sue445/gcp-secretmanagerenv"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/job"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//var (
//	indexTmpl = readTemplate("index.html")
//)

func main() {
	projectID := os.Getenv("GCP_PROJECT")

	sentryDebug := os.Getenv("SENTRY_DEBUG") != ""
	secretmanager, err := secretmanagerenv.NewClient(context.Background(), projectID)
	if err != nil {
		panic(err)
	}

	sentryDsn, err := secretmanager.GetValueFromEnvOrSecretManager("SENTRY_DSN", false)
	if err != nil {
		panic(err)
	}

	googleMapsAPIKey, err := secretmanager.GetValueFromEnvOrSecretManager("GOOGLE_MAPS_API_KEY", false)
	if err != nil {
		panic(err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
		AttachStacktrace: true,
		Debug:            sentryDebug,
	})
	if err != nil {
		panic(err)
	}

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	config.Init(&config.InitParams{
		ProjectID:        projectID,
		GoogleMapsAPIKey: googleMapsAPIKey,
	})

	r := mux.NewRouter()
	r.HandleFunc("/job/cron/update_shops", job.CronUpdateShopsHandler).Methods("POST")
	r.HandleFunc("/job/queue/save_shop", job.QueueSaveShopHandler).Methods("POST")
	// r.HandleFunc("/", indexHandler)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		http.NotFound(w, r)
//		return
//	}
//
//	if os.Getenv("GCP_PROJECT") == "" {
//		// Hot reloading for local
//		indexTmpl = readTemplate("index.html")
//	}
//
//	vars := map[string]string{}
//
//	if err := indexTmpl.Execute(w, vars); err != nil {
//		sentry.CaptureException(err)
//		log.Printf("Error executing template: %v", err)
//		http.Error(w, "Internal server error", http.StatusInternalServerError)
//	}
//}

func readTemplate(name string) *template.Template {
	return template.Must(template.ParseFiles(filepath.Join("templates", name)))
}
