package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

func main() {
	l := httplog.NewLogger("app", httplog.Options{
		JSON: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(l))
	checkEnv()

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Get("/*", processRequest)

	http.ListenAndServe(":"+port, r)
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	archiveUrl, status, err := RequestArchiveUrl(r.RequestURI)
	if err != nil {
		http.Error(w, http.StatusText(status), status)
		return
	}

	resp, err := http.Get(archiveUrl)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}

func checkEnv() {
	envKey := []string{"FAKE_DOMAIN"}
	for _, v := range envKey {
		env := os.Getenv(v)
		if env == "" {
			log.Fatal("Environment " + v + " not declared")
			os.Exit(1)
		}
	}
}
