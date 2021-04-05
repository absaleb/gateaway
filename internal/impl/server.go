package impl

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	AddressHttp string
	AddressGrpc string
}

func (srv *Server) Run() {
	log.Infof("gateway service starting on http://%s/query ...", srv.AddressHttp)

	http.Handle("/query", getGraphqlHandler())
	log.Fatal(http.ListenAndServe(srv.AddressHttp, nil))
}
