package graphqldelivery

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var DateTimeScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "DateTime",
	Description: "DateTime scalar type represents a date and time in RFC3339 format",
	Serialize: func(value interface{}) interface{} {
		switch v := value.(type) {
		case time.Time:
			return v.Format(time.RFC3339)
		case *time.Time:
			if v == nil {
				return nil
			}
			return v.Format(time.RFC3339)
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch v := value.(type) {
		case string:
			t, err := time.Parse(time.RFC3339, v)
			if err != nil {
				return nil
			}
			return t
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch v := valueAST.(type) {
		case *ast.StringValue:
			t, err := time.Parse(time.RFC3339, v.Value)
			if err != nil {
				return nil
			}
			return t
		default:
			return nil
		}
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"name":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"email":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"role":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"image":       &graphql.Field{Type: graphql.String},
		"imageFileId": &graphql.Field{Type: graphql.String},
		"createdAt":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"updatedAt":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
	},
})

var ArticleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"title":     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"content":   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"path":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"viewCount": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"likes":     &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"authorId":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"published": &graphql.Field{Type: graphql.NewNonNull(graphql.Boolean)},
		"author":    &graphql.Field{Type: UserType},
		"createdAt": &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"updatedAt": &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
	},
})

var EducationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Education",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"institution":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"degree":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"logo":         &graphql.Field{Type: graphql.String},
		"startDate":    &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"endDate":      &graphql.Field{Type: DateTimeScalar},
		"gpa":          &graphql.Field{Type: graphql.String},
		"achievements": &graphql.Field{Type: graphql.NewList(graphql.String)},
		"createdAt":    &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"updatedAt":    &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
	},
})

var ExperienceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Experience",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"company":     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"position":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"type":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"logo":        &graphql.Field{Type: graphql.String},
		"startDate":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"endDate":     &graphql.Field{Type: DateTimeScalar},
		"tags":        &graphql.Field{Type: graphql.NewList(graphql.String)},
		"description": &graphql.Field{Type: graphql.NewList(graphql.String)},
		"createdAt":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"updatedAt":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
	},
})

var ProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Project",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"title":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"subtitle":    &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"image":       &graphql.Field{Type: graphql.String},
		"tags":        &graphql.Field{Type: graphql.NewList(graphql.String)},
		"demoUrl":     &graphql.Field{Type: graphql.String},
		"githubUrl":   &graphql.Field{Type: graphql.String},
		"createdAt":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"updatedAt":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
	},
})

var GalleryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gallery",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"image":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"imageFileId": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"caption":     &graphql.Field{Type: graphql.String},
		"createdAt":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"updatedAt":   &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
	},
})

var SocialLinkType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SocialLink",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"title":     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"url":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"order":     &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"isActive":  &graphql.Field{Type: graphql.NewNonNull(graphql.Boolean)},
		"createdAt": &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
		"updatedAt": &graphql.Field{Type: graphql.NewNonNull(DateTimeScalar)},
	},
})
