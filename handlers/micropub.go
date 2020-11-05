package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CrowderSoup/gozette/config"
)

type Config struct {
	SyndicateTo   []config.SyndicationTarget `json:"syndicate-to,omitempty"`
	MediaEndpoint string                     `json:"media-endpoint,omitempty"`
}

// MicropubHandler handles micropub requests
func MicropubHandler(w http.ResponseWriter, r *http.Request) {
	appConfig := config.NewConfig()

	if r.Method == http.MethodGet {
		// TODO: implement get
	}

	fmt.Fprintf(w, "%s", appConfig)
}

func get(w http.ResponseWriter, r *http.Request) {
	if q := r.URL.Query().Get("q"); q == "config" || q == "syndicate-to" {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&Config{
			SyndicateTo:   config.SyndicationTargets,
			MediaEndpoint: config.MediaEndpointURL,
		})
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{}"))
	return
}
