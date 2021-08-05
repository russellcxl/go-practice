package main

import (
	"encoding/json"
	"fmt"
	"git.garena.com/russell.chanxl/personal/Graphql_2/mutations"
	"git.garena.com/russell.chanxl/personal/Graphql_2/queries"
	"github.com/graphql-go/graphql"
	"log"
)



func main() {

	// overarching schema; includes all query and mutation methods
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        queries.QueryType,
		Mutation:     mutations.MutationType,
		Subscription: nil,
		Types:        nil,
		Directives:   nil,
		Extensions:   nil,
	})
	if err != nil {
		log.Fatalln(err)
	}


	mutate := `mutation {create_tutorial(title:"Tutorial by Quentin") { Title } } `
	query := `{all_tutorials { Title } }`

	// mutation
	r := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  mutate,
	})

	if len(r.Errors) > 0 {
		log.Fatalln(r.Errors)
	}

	rJSON, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("%s\n", rJSON)


	// query
	r = graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})

	if len(r.Errors) > 0 {
		log.Fatalln(r.Errors)
	}

	rJSON, _ = json.MarshalIndent(r, "", "  ")
	fmt.Printf("%s\n", rJSON)
}

