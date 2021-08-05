package queries

import (
	"git.garena.com/russell.chanxl/personal/Graphql_2/db"
	"git.garena.com/russell.chanxl/personal/Graphql_2/models"
	"github.com/graphql-go/graphql"
)

// theses are the fields that can be used in the queries
var fields = graphql.Fields{
	"tutorial_by_id": &graphql.Field{
		Description: "Get tutorial by ID",

		// follows Tutorial struct
		Type: models.TutorialType,

		// takes in arg with type int
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},

		// what the queries returns
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {
				for _, tutorial := range db.Tutorials {
					if tutorial.Id == id {
						return tutorial, nil
					}
				}
			}
			return nil, nil
		},
	},

	"all_tutorials": &graphql.Field{
		Description: "Get full tutorial list",
		Type:        graphql.NewList(models.TutorialType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return db.Tutorials, nil
		},
	},
}

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: fields,
})

