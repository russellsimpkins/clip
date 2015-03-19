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

func TestToken(t *testing.T) {
	token := GenerateToken()
	fmt.Printf("Token.IntValue    = %d\n", token.IntValue)
	fmt.Printf("Token.StringValue = %s\n", token.StringValue)
}

func setupTeam(team *Team) {

	var (
		router *mux.Router
		request *http.Request
	)
	
	router = mux.NewRouter()
	SetRoutes(router)

	team.Name = "Test"
	data, _ := json.Marshal(team)
	//fmt.Println("DATA: ", string(data))
	request, _ = http.NewRequest("POST", "/svc/clip/team", strings.NewReader(string(data)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	//fmt.Println(response.Body.String())
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
	team := Team{}
	team.Name = "Data Universe"
	setupTeam(&team)
	t.Log("INFO: team: ", team)

	w := httptest.NewRecorder()
	url := fmt.Sprintf("/svc/clip/team/%s/token", team.Name)
	//t.Log("INFO: url: ", url)
	request, _ = http.NewRequest("POST", url, nil)
	
	router.ServeHTTP(w, request)
	t.Log(w.Body.String())
	
	if w.Code != 200 {
		t.Log("ERROR: ", w.Body.String())
		t.Fail()
	}
	
	url = fmt.Sprintf("/svc/clip/team/%s", team.Name)
	request, _ = http.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)
	team = Team{}
	err := json.Unmarshal(w.Body.Bytes(), &team)

	if err != nil {
		return
	}
	team.Token[0].Applications = make(map[string]Feature, 2)
	f := Feature{}
	f.Flags = make(map[string]Flag, 2)
	fl := f.Flags["chocolate"]
	fl.Sandbox = 1
	fl.Development = 1
	fl.Attribs = make(map[string]int, 1)
	fl.Attribs["displayFE"] = 1
	f.Flags["chocolate"] = fl
	team.Token[0].Applications["doughnuts"] = f
	
	
	data, _ := json.Marshal(team)
	
	t.Log(string(data))

	url = fmt.Sprintf("/svc/clip/team/%s", team.Name)
	request, _ = http.NewRequest("PUT", url, strings.NewReader(string(data)))
	router.ServeHTTP(w, request)
	t.Log(w.Body.String())
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
