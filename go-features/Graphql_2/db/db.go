package db

import "git.garena.com/russell.chanxl/personal/Graphql_2/models"

var (
	Tutorials = populate()
)


func populate() []models.Tutorial {
	author := &models.Author{Name: "Jake", Tutorials: []int{1}}
	tutorial := models.Tutorial{
		Title:    "Tutorial 1",
		Id:       1,
		Author:   *author,
		Comments: []models.Comment{
			{"First comment"},
		},
	}
	tutorials := []models.Tutorial{tutorial}

	return tutorials
}

