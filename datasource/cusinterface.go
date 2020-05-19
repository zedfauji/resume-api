package datasource

import (
	"github.com/zedfauji/resume-api/loader"
)

type Experience interface {
	Initialize()
	ShowExperienceAll() *[]*loader.ExperienceRow
	SearchExperience(company_name string) *loader.ExperienceRow
}
type ProjectTaken interface {
	ProjectInitialize()
	ShowProjectsAll() *[]*loader.ProjectTakenRow
	//SearchProject()
}

type Technology interface {
	TechnologyInitialize()
	ShowTechnologiesAll() *[]*loader.Technologies
}
