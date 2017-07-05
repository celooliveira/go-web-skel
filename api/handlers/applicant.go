package handlers

import (
    "encoding/json"
    "gitlab.com/thecodenation/thecodenation/datastore"
    "net/http"
)

func ApplicantIndex(ds datastore.ModelDatastore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        beers, err := ds.AllApplicants()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        if err := json.NewEncoder(w).Encode(beers); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}
