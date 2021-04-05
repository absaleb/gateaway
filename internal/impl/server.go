package impl

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

type Server struct {
	AddressHttp string
	AddressGrpc string
}

func (srv *Server) Run() {
	http.Handle("/query", getGraphqlHandler())
	log.Fatal(http.ListenAndServe(srv.AddressHttp, nil))
}

func getGraphqlHandler() http.Handler {
	schema, err := graphql.NewSchema()
}
