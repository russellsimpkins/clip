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

func TeamNames() (teams []string) {
	teams = make([]string, 5)
	teams[0] = "Data Universe"
	teams[1] = "MobileWeb"
	teams[2] = "WebTech"
	teams[3] = "IOS"
	teams[4] = "Search"
	return
}

func CleanTeams() {
	teams := TeamNames();
	var router *mux.Router
	router = mux.NewRouter()
	SetRoutes(router)
	response := httptest.NewRecorder()

	for t := range teams {
		request, _ := http.NewRequest("DELETE", fmt.Sprintf("/svc/clip/team/%s", teams[t]), nil)
		router.ServeHTTP(response, request)
	}
}

func TestCreateTeam(t *testing.T) {
	var router *mux.Router
	router = mux.NewRouter()
	SetRoutes(router)
	
	CleanTeams()
	teams := TeamNames();
	
	for idx := range teams {
		team := Team{}
		team.Name = teams[idx]
		data, _ := json.Marshal(team)
		request, _ := http.NewRequest("POST", "/svc/clip/team", strings.NewReader(string(data)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Log(response.Body.String())
			t.Fail()
		}
	}
}

func TestGetTeam(t *testing.T) {
	var router *mux.Router
	router = mux.NewRouter()
	SetRoutes(router)
	teams := TeamNames();
	for idx := range teams {
		team := Team{}
		request, _ := http.NewRequest("GET", fmt.Sprintf("/svc/clip/team/%s", teams[idx]), nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Log(response.Body.String())
			t.Fail()
		}
		_ = json.Unmarshal([]byte(response.Body.String()), &team)
		if team.Name != teams[idx] {
			t.Errorf("FAIL: Team %s was not fetched", teams[idx])
		}
	}
}

func TestUpdateTeam(t *testing.T) {
	var router *mux.Router
	router = mux.NewRouter()
	SetRoutes(router)
	teams := TeamNames();
	
	for idx := range teams {
		team := Team{}
		request, _ := http.NewRequest("GET", fmt.Sprintf("/svc/clip/team/%s", teams[idx]), nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Log(response.Body.String())
			t.Fail()
		}
		_ = json.Unmarshal([]byte(response.Body.String()), &team)
		team.Users = make([]User, 1)
		user := User{}
		user.Email = fmt.Sprintf("%s@yntimes.com", team.Name)
		team.Users[0] = user

		
		data, _ := json.Marshal(team)
		request, _ = http.NewRequest("PUT", fmt.Sprintf("/svc/clip/team/%s", teams[idx]), strings.NewReader(string(data)))
		response = httptest.NewRecorder()
		router.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Log(response.Body.String())
			t.Fail()
		}
		
		_ = json.Unmarshal([]byte(response.Body.String()), &team)
		if len(team.Users) != 1 {
			t.Errorf("FAIL: Team %s was not updated correctly", teams[idx])
		}
	}
}

func TestGetTeamList(t *testing.T) {
	var router *mux.Router
	router = mux.NewRouter()
	SetRoutes(router)
	request, _ := http.NewRequest("GET", "/svc/clip/teams", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Log(response.Body.String())
		t.Fail()
	}
	s := response.Body.String()
	if s == "" {
		t.Errorf("FAIL: We didn't get any teams back")
	}
		
	t.Log(response.Body.String())
}

func TestCleanTeams(t *testing.T) {
	//CleanTeams()
}
