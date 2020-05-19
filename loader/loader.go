package loader

import (
	"encoding/csv"
	"io"
	"log"
)

type ExperienceRow struct {
	CompanyName  string `json:"company_name"`
	RoleTaken    string `json:"role_taken"`
	YearStarted  string `json:"year_started"`
	YearEnded    string `json:"year_ended"`
	MonthStarted string `json:"month_started"`
	MonthEnded   string `json:"month_ended"`
}
type ProjectTakenRow struct {
	CompanyName        string `json:"company_name"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
	ToolsUsed          string `json:"tools_used"`
	ProjectYear        string `json:"project_year"`
	ProjectRole        string `json:"project_role"`
}

type Technologies struct {
	Cloud         string `json:"cloud"`
	CloudServices string `json:"cloud_services"`
	Languages     string `json:"languages"`
	CICD          string `json:"cicd"`
	Monitoring    string `json:"monitoring"`
	Tech          string `json:"technologies"`
}

type All struct {
	experienceRow   ExperienceRow
	projecttakenRow ProjectTakenRow
	technologies    Technologies
}

func LoadTechnologies(r io.Reader) *[]*Technologies {
	reader := csv.NewReader(r)
	ret := make([]*Technologies, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of file Technologies")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		technologies_used := &Technologies{
			Cloud:         row[0],
			CloudServices: row[1],
			Languages:     row[2],
			CICD:          row[3],
			Monitoring:    row[4],
			Tech:          row[5],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, technologies_used)
	}

	return &ret
}

func LoadProjects(r io.Reader) *[]*ProjectTakenRow {

	reader := csv.NewReader(r)
	ret := make([]*ProjectTakenRow, 0, 0)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of file for Projects")
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		projectTaken := &ProjectTakenRow{
			CompanyName:        row[0],
			ProjectName:        row[1],
			ProjectDescription: row[2],
			ToolsUsed:          row[3],
			ProjectYear:        row[4],
			ProjectRole:        row[5],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, projectTaken)
	}

	return &ret
}
func LoadExperience(r io.Reader) *[]*ExperienceRow {
	reader := csv.NewReader(r)
	ret := make([]*ExperienceRow, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of file Experience")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		experience := &ExperienceRow{
			CompanyName:  row[0],
			RoleTaken:    row[1],
			YearStarted:  row[2],
			YearEnded:    row[3],
			MonthStarted: row[4],
			MonthEnded:   row[5],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, experience)
	}

	return &ret
}
