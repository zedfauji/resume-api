package datasource

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"../loader"
	"github.com/landoop/tableprinter"
)

type AllExperience struct {
	TotalExperience *[]*loader.ExperienceRow `json:"all_experience"`
}

type AllProject struct {
	TotalProjects *[]*loader.ProjectTakenRow `json:"all_project"`
}

type AllTechnology struct {
	TotalTechnologies *[]*loader.Technologies `json:"all_technologies"`
}

func (AT *AllTechnology) TechnologyInitialize() {
	fmt.Println("Starting in Tech Initialize")

	filename := "/home/girish/tempproject/go-projects/training/resume-api/go/resume/datasource/technologies.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Can't open file", err)
	}

	defer file.Close()
	AT.TotalTechnologies = loader.LoadTechnologies(file)
}

func (AT *AllTechnology) ShowTechnologiesAll() *[]*loader.Technologies {
	fmt.Println("Here in Show All Technologies")
	marshData, _ := json.MarshalIndent(AT, "", "	")
	fmt.Println(string(marshData))
	return AT.TotalTechnologies
}

func (AP *AllProject) ProjectInitialize() {
	fmt.Println("Starting in Project Initialize")

	filename := "/home/girish/tempproject/go-projects/training/resume-api/go/resume/datasource/projects.csv"
	file, err := os.Open(filename)
	fmt.Println("Project File opened")
	if err != nil {
		log.Fatalln("Can't open the file", err)
	}

	defer file.Close()
	AP.TotalProjects = loader.LoadProjects(file)
}

func (AP *AllProject) ShowProjectsAll() *[]*loader.ProjectTakenRow {
	fmt.Println("Here in Show All Projectts")
	marshData, _ := json.MarshalIndent(AP, "", "	")
	fmt.Println(string(marshData))
	return AP.TotalProjects
}

func (AE *AllExperience) Initialize() {
	fmt.Println("Starting in Initialize")
	filename := "/home/girish/tempproject/go-projects/training/resume-api/go/resume/datasource/experience.csv"
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln("Can't open the file", err)
	}

	defer file.Close()
	AE.TotalExperience = loader.LoadExperience(file)
}

func (AE *AllExperience) ShowExperienceAll() *[]*loader.ExperienceRow {

	fmt.Println("Here in SHowExperienceAll")
	marshData, _ := json.Marshal(AE)
	fmt.Println(string(marshData))
	tableprinter.Print(os.Stdout, AE.TotalExperience)
	return AE.TotalExperience
}

func (AE *AllExperience) SearchExperience(company_name string) *loader.ExperienceRow {

	ret := Filter(AE.TotalExperience, func(v *loader.ExperienceRow) bool {
		return strings.ToLower(v.CompanyName) == strings.ToLower(company_name)
	})

	if len(*ret) > 0 {
		return (*ret)[0]
	}
	return nil

}

func Filter(vs *[]*loader.ExperienceRow, f func(*loader.ExperienceRow) bool) *[]*loader.ExperienceRow {
	vsf := make([]*loader.ExperienceRow, 0, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return &vsf
}
