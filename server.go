package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var graphqlHandler, graphqlLoginHandler *relay.Handler

type Opt struct{
	debug bool
	listen string
}

var opt Opt

func init(){
	debug := false
	if os.Getenv("DEBUG") != ""{
		debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	}
	flag.BoolVar(&opt.debug, "debug", debug, "debug")
	port, _ := strconv.ParseInt(os.Getenv("DEMO_PORT"), 10, 64)
	if port < 1{
		port = 8787
	}
	flag.StringVar(&opt.listen, "listen", fmt.Sprint(":", port), "usage :port")

	s, err := getSchema("./graphql-files/schema.graphql")
	if err != nil{
		log.WithFields(log.Fields{"time": time.Now()}).Info(err)
	}
	schema := graphql.MustParseSchema(s, &Resolver{app: DefaultApp}, graphql.UseStringDescriptions())

	sLogin, err := getSchema("./graphql-files/login.graphql")
	if err != nil{
		log.WithFields(log.Fields{"time": time.Now()}).Info(err)
	}
	schemaLogin := graphql.MustParseSchema(sLogin, &GraphqlLogin{app: DefaultApp}, graphql.UseStringDescriptions())
	graphqlHandler = &relay.Handler{Schema: schema}
	graphqlLoginHandler = &relay.Handler{Schema: schemaLogin}

}

func main(){
	flag.Parse()
	DefaultApp.Init(true)
	err := DefaultApp.GnerateFakeData()
	if err != nil{
		log.Println(err)
	}
	log.Printf("debug: %v", opt.debug)
	log.Printf("listen port%v", opt.listen)

	http.Handle("/login/graphql", logged(graphqlLoginHandler))
	if DefaultApp.debug{
		http.Handle("/query", logged(graphqlHandler))
	}else{
		http.Handle("/query", logged(authToken(graphqlHandler)))
	}
	http.HandleFunc("/login", CreateTokenEndpoint)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(opt.listen, nil))
}



