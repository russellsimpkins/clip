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
		DeleteTokenHandler).Methods("DELETE")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}",
		GetTokenHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}/{app:[a-zA-Z0-9 \\%\\.\\-_]+}",
		GetTokenAppHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}/{app:[a-zA-Z0-9 \\%\\.\\-_]+}/{feature:[a-zA-Z0-9 \\%\\.\\-_]+}",
		GetTokenAppFeatureHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}/{app:[a-zA-Z0-9 \\%\\.\\-_]+}/{feature:[a-zA-Z0-9 \\%\\.\\-_]+}/{env:(sbx|dev|int|stg|prd)}",
		GetTokenAppFeatureEnvHandler).Methods("GET")

	

	

}
