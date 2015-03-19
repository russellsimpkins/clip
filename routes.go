package clip

import (
	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {

	router.HandleFunc("/svc/clip/user/{email:[a-zA-Z0-9\\.\\-\\@]+}", CreateUserHandler).Methods("POST")
	router.HandleFunc("/svc/clip/user/{email:[a-zA-Z0-9\\.\\-\\@]+}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/svc/clip/user/{email:[a-zA-Z0-9\\.\\-\\@]+}", FetchUserHandler).Methods("GET")

	router.HandleFunc("/svc/clip/team", CreateTeamHandler).Methods("POST")
	router.HandleFunc("/svc/clip/teams", GetTeamsHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\.\\-_]+}", UpdateTeamHandler).Methods("PUT")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}", GetTeamHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}", DeleteTeamHandler).Methods("DELETE")

	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\.\\-_]+}/token",
		CreateTokenHandler).Methods("POST")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\.\\-_]+}/token/{token:[a-zA-Z0-9]+}",
		UpdateTokenHandler).Methods("PUT")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}",
		GetTokenHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}",
	DeleteTokenHandler).Methods("DELETE")

}
