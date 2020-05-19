package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SearchCompany(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "SearchCompany")
	queries := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if val, ok := queries["company_name"]; ok {
		data := exp.SearchExperience(val)
		if data != nil {
			b, err := json.MarshalIndent(data, "", "		")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return

			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return

		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

func showAllExperience(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "showAllExperience")
	w.Header().Set("Content-Type", "application/json")

	allExp := exp.ShowExperienceAll()
	if allExp != nil {
		//respData, err := json.Marshal(allExp)
		respData, err := json.MarshalIndent(allExp, "", "		")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respData)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}
func showAllTechnologies(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "showAllTechnologies")
	w.Header().Set("Content-Type", "application/json")

	allTech := tech.ShowTechnologiesAll()
	if allTech != nil {
		//respData, err := json.Marshal(allExp)
		respData, err := json.MarshalIndent(allTech, "", "		")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respData)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

func showAllProject(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "ShowALlProject")
	w.Header().Set("Content-Type", "application/json")

	allProject := project.ShowProjectsAll()

	if allProject != nil {
		respData, err := json.MarshalIndent(allProject, "", "		")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respData)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))

}
