package config

import (
	"github.com/rayne22/bulk_sms_api/controllers"
)
func(server *Server) Routes()  {
	// Home Route
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

}
