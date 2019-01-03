package main

import (
	"./types"
	"encoding/json"
	"net/http"
	
	"github.com/graphql-go/graphql"
)
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
