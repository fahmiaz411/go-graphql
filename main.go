package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.Fields{
		"Welcome": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "TO GRAPHQL", nil
			},
		},
	}

	abcd := graphql.ObjectConfig{
		Name: "ABCD",
		Fields: fields,
	}

	xyz := graphql.SchemaConfig{Query: graphql.NewObject(abcd)}
	schema, err := graphql.NewSchema(xyz)
	if err != nil {
		log.Fatalln("error create schema", err.Error())
	}

	query := `
		{
			Welcome
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	a := graphql.Do(params)
	if len(a.Errors) > 0 {
		for _, v := range a.Errors {
			log.Fatalln(v)
		}
	}

	rJson, _ := json.Marshal(a)
	fmt.Println(string(rJson))
}