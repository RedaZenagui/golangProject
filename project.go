package main

import (
	"./types"
	"encoding/json"
	"net/http"
	"github.com/SalahEddineBC/api"
	"context"
	"time"
	
	"github.com/graphql-go/graphql"
)

//Creating our possible query structure
var releaseNote = graphql.NewObject(graphql.ObjectConfig{
	Name: "releaseNote",
	Fields: graphql.Fields{
		"date"        : &graphql.Field{ Type : graphql.String},
		"product"     : &graphql.Field{ Type : graphql.String},
		"tagline"     : &graphql.Field{ Type : graphql.String},
		"text"        : &graphql.Field{ Type : graphql.String},
		"product_lead": &graphql.Field{ Type : graphql.String},
			
		},
	},
)
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getReleaseNotes" : &graphql.Field{ 
		Type : graphql.NewList(releaseNote),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	        defer cancel()
	        r, err := c.GetReleaseNotes(ctx, &api.Empty{})

	  	},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{ //Defining our graphql shema which contains our Query structure

    Query: rootQuery,
    
})


type QueryS struct { //the struct which contains our string request
	Query string
}


func resolve(query string, shema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	return result
}


func main() {
   
	http.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		var q QueryS
		if req.Method == "Post" {
			json.NewDecoder(req.Body).Decode(&q)
			result := resolve(q.Query, schema)
			json.NewEncoder(res).Encode(result)
		}

	})

	http.ListenAndServe(":8080", nil)
}
