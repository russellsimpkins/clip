package clip

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setRoutes(router *mux.Router) {

	router.HandleFunc("/svc/clip/teams", GetTeamsHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team", CreateTeamHandler).Methods("POST")
	router.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\.\\-_]+}", UpdateTeamHandler).Methods("PUT")
	router.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\%\\.\\-_]+}", GetTeamHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\%\\.\\-_]+}", DeleteTeamHandler).Methods("DELETE")
	// token routes
	router.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\.\\-_]+}/token", CreateTokenHandler).Methods("POST")
	router.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\.\\-_]+}/token/{token:[a-zA-Z0-9]+}", UpdateTokenHandler).Methods("PUT")
	router.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}", GetTokenHandler).Methods("GET")
	router.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}", DeleteTokenHandler).Methods("DELETE")
	return
}


func TestToken(t *testing.T) {
	token := GenerateToken()
	fmt.Printf("Token.IntValue    = %d\n", token.IntValue)
	fmt.Printf("Token.StringValue = %s\n", token.StringValue)
}

func setupTeam() (team Team) {

	var (
		router *mux.Router
		request *http.Request
	)
	
	router = mux.NewRouter()
	setRoutes(router)
	team = Team{}
	team.Name = "Test"
	data, _ := json.Marshal(team)
	request, _ = http.NewRequest("POST", "/svc/clip/team", strings.NewReader(string(data)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return
}


func TestTokenCrud(t *testing.T) {
	
	var (
		router *mux.Router
		request *http.Request
	)
	
	router = mux.NewRouter()
	SetRoutes(router)
	
	// Create a team
	team := setupTeam()
		
	w := httptest.NewRecorder()	
	request, _ = http.NewRequest("POST", "/svc/clip/team/DU/token", nil)
	router.ServeHTTP(w, request)
	t.Log(w.Body.String())
	if w.Code != 200 {
		t.Log("ERROR: ", w.Body.String())
		t.Fail()
	}
	
	_ = json.Unmarshal([]byte(w.Body.String()), &team)
	
	//team.Token[0].Applications = make(map[string]Feature, 2)
	//f := Feature{}
	//f.Flags = make(map[string]Flag, 2)
	//fl := f.Flags["chocolate"]
	//fl.Sandbox = 1
	//fl.Development = 1
	//fl.Attribs = make(map[string]int, 1)
	//fl.Attribs["displayFE"] = 1
	//f.Flags["chocolate"] = fl
	//team.Token[0].Applications["doughnuts"] = f
	
	
	data, _ := json.Marshal(team)
	t.Log(string(data))
	cleanUp(&team)
}

func cleanUp(team *Team) {

	var router *mux.Router
	router = mux.NewRouter()
	SetRoutes(router)
	response := httptest.NewRecorder()

	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/svc/clip/team/%s", team.Name), nil)
	router.ServeHTTP(response, request)
	
}
