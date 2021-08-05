package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

func main() {

	// schema
	fields := graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "zuu", nil
			},
		},
		"age": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 22, nil
			},
		},
	}

	// ObjectConfig -(is used by)-> NewObject --> SchemaConfig --> NewSchema --> Params --> Do
	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery", Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal(err)
	}

	// queries
	query := `
		{
			name,
			age
		}
	`

	params := graphql.Params{
		Schema: schema,
		RequestString: query,
	}

	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatal(r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Println(string(rJSON))
}

