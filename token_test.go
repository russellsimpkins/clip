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

func TestTokenCrud(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	//**********************************************************************
	// ROUTES
	//----------------------------------------------------------------------
	r.HandleFunc("/svc/clip/team", CreateTeamHandler).Methods("POST")
	r.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\%\\.\\-_]+}", DeleteTeamHandler).Methods("DELETE")
	r.HandleFunc("/svc/clip/team/{name:[a-zA-Z0-9 \\%\\.\\-_]+}", GetTeamHandler).Methods("GET")


	r.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}",
		GetTokenHandler).Methods("GET")
	r.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\.\\-_]+}/token",
		CreateTokenHandler).Methods("POST")
	r.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\.\\-_]+}/token/{token:[a-zA-Z0-9]+}",
		UpdateTokenHandler).Methods("PUT")
	r.HandleFunc("/svc/clip/team/{team:[a-zA-Z0-9 \\%\\.\\-_]+}/token/{token:[a-zA-Z0-9]+}",
		DeleteTokenHandler).Methods("DELETE")

	

	var request *http.Request

	// Create a team
	team := Team{}
	team.Name = "DU"
	data, _ := json.Marshal(team)

	
	if 1 == 1 {
		request, _ = http.NewRequest("POST", "/svc/clip/team", strings.NewReader(string(data)))
		r.ServeHTTP(w, request)

		if w.Code != 200 {
			t.Log("ERROR: ", w.Body.String())
			t.Fail()
		}
		t.Log(w.Body.String())
	} else {
		request, _ = http.NewRequest("GET", "/svc/clip/team/DU", nil)
		r.ServeHTTP(w, request)
		if w.Code != 200 {
			t.Log("ERROR: ", w.Body.String())
			t.Fail()
		}
		t.Log(w.Body.String())
		_ = json.Unmarshal([]byte(w.Body.String()), &team)
	}



	
	// Create a token
	w = httptest.NewRecorder()	
	request, _ = http.NewRequest("POST", "/svc/clip/team/DU/token", nil)
	r.ServeHTTP(w, request)
	t.Log(w.Body.String())
	if w.Code != 200 {
		t.Log("ERROR: ", w.Body.String())
		t.Fail()
	}


	// Update a token 
	if 1 == 1 {
		w = httptest.NewRecorder()
		request, _ = http.NewRequest("DELETE", "/svc/clip/team/DU", nil)
		r.ServeHTTP(w, request)
	}
}
