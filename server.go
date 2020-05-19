package main

import (
	"log"
	"net/http"
	"time"

	"./datasource"
	"github.com/gorilla/mux"
)

var (
	exp     datasource.Experience
	project datasource.ProjectTaken
	tech    datasource.Technology
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	exp = &datasource.AllExperience{}
	project = &datasource.AllProject{}
	tech = &datasource.AllTechnology{}
	exp.Initialize()
	project.ProjectInitialize()
	tech.TechnologyInitialize()
}

type server struct{}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/experience", showAllExperience)
	r.HandleFunc("/experience/{company_name}", SearchCompany)
	r.HandleFunc("/projects", showAllProject)
	r.HandleFunc("/technologies", showAllTechnologies)
	log.Println("Starting Server Now")
	log.Fatal(http.ListenAndServe(":8080", r))

}
