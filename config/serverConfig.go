package config

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Server struct {
	Router *mux.Router
	Client *mongo.Client
}

func (server *Server) Initialize()  {

	LogConf()

	server.DBConection()

	server.Router = mux.NewRouter()

	server.Routes()

}

func (server *Server) Run(addr string) {
	//if err := http.ListenAndServe("0.0.0.0:5000", l.LogRequest("",handlers.LoggingHandler(os.Stdout, server.Router))); err != nil {
	//	panic(err)
	//}
	fmt.Println("Listening to port 5500")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}