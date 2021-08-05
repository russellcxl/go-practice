package models

import "github.com/graphql-go/graphql"

type Tutorial struct {
	Title    string
	Id       int
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

var (



	// this must follow the struct type and its fields
	CommentType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"Body": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	AuthorType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	})

	TutorialType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Id": &graphql.Field{
				Type: graphql.Int,
			},
			"Author": &graphql.Field{
				Type: AuthorType,
			},
			"Comments": &graphql.Field{
				Type: graphql.NewList(CommentType),
			},
		},
	})

)

