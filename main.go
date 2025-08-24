package main

import (
	"encoding/json"
	"net/http"

	"github.com/Michela-DC/review-finder-clients/mockclient"

	"github.com/Michela-DC/review-finder/domain"
)

func main() {
	//web server http
	mux := http.NewServeMux()
	// todo: generalizzare l'handleFunction per gestire tutti i tipi di richiesta (movie, anime, games)
	// todo: rendere i clients injectable
	// todo: resolve client based on endpoint
	// todo: move server to infrastracture folder
	// todo: move api in-out to domain
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			w.Write([]byte(`{"error": "missing name query parameter"}`))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		mc := &mockclient.Movies{}
		mcResp := mc.FindMovies(name)

		resp := make([]*domain.Movie, len(mcResp))
		for i, r := range mcResp {
			resp[i] = &domain.Movie{
				Title:       r.Title,
				Description: r.Description + " from review finder",
				Rating:      r.Rating,
			}
		}

		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	})
	http.ListenAndServe(":8080", mux)

}
