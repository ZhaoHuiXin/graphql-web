package main

import (
	"time"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)
//var schema, schemaLogin *graphql.Schema
var graphqlHandler, graphqlLoginHandler *relay.Handler

type Resolver struct {
	app *App
}

func init(){
	DefaultApp.Init(true)
	s, err := getSchema("./graphql-files/schema.graphql")
	if err != nil{
		log.WithFields(log.Fields{"time": time.Now()}).Info(err)
	}
	schema := graphql.MustParseSchema(s, &Resolver{app: DefaultApp}, graphql.UseStringDescriptions())

	sLogin, err := getSchema("./graphql-files/login.graphql")
	if err != nil{
		log.WithFields(log.Fields{"time": time.Now()}).Info(err)
	}
	schemaLogin := graphql.MustParseSchema(sLogin, &GraphqlLogin{}, graphql.UseStringDescriptions())
	graphqlHandler = &relay.Handler{Schema: schema}
	graphqlLoginHandler = &relay.Handler{Schema: schemaLogin}
	err = DefaultApp.GnerateFakeData()
	if err != nil{
		log.Println(err)
	}
}

func main(){
	http.Handle("/login/graphql", logged(graphqlLoginHandler))
	http.Handle("/query", logged(authToken(graphqlHandler)))
	http.HandleFunc("/login", CreateTokenEndpoint)
	log.Fatal(http.ListenAndServe(":8787", nil))

}



