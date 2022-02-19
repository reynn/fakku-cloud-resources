package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		healthObjects, marshalErr := json.Marshal(map[string]string{
			"database": "ok",
			"cache": "ok",
		})
		if marshalErr != nil {
			http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
		}
		w.Write(healthObjects)
	})

	http.ListenAndServe(":8000", nil)
}
