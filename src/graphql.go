package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Defining the structure of the GraphQL Query
var helloWorldQuery = graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		// "hello" field to return a greeting message
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello, World!", nil // Returning "Hello, World!"
			},
		},
		// "name" field to return a custom name
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hi - Edwin Allauca!", nil 
			},
		},
	},
}

var schema, schemaErr = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(helloWorldQuery),
})

func main() {
	// Verificar si hubo un error al crear el esquema
	if schemaErr != nil {
		log.Fatalf("Error al crear el esquema: %v", schemaErr)
	}

	// Check if there was an error creating the schema
	h := handler.New(&handler.Config{
		Schema:  &schema,
		Pretty:  true,  // Print in JSON
		GraphiQL: true, // Enable the GraphiQL interface
	})

	// Configure the route for the GraphQL server
	http.Handle("/graphql", h)

	// Server in port 8081
	address := ":8081"
	fmt.Printf("Servidor en ejecución en http://localhost%s/graphql\n", address)

	// Start the server and handle errors if they occur
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
