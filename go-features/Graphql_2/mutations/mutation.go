package mutations

import (
	"git.garena.com/russell.chanxl/personal/Graphql_2/db"
	"git.garena.com/russell.chanxl/personal/Graphql_2/models"
	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create_tutorial": &graphql.Field{
			Type:        models.TutorialType,
			Description: "Create a new Tutorial",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				tutorial := models.Tutorial{
					Title: params.Args["title"].(string),
				}
				db.Tutorials = append(db.Tutorials, tutorial)
				return tutorial, nil
			},
		},
	},
})

