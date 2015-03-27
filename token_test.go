
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

type IntMap map[string]int

func makeMap(name string, value int) ( v *IntMap ) {
	m := make(IntMap)
	v = &m
	m[name] = value
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
	team.Name = "Test"
	team, _ = GetTeam(team.Name)
	setupTeam(&team)
	t.Log("INFO: team: ", team)

	w := httptest.NewRecorder()
	tok := Token{}
	tok.Name = "FrontendAPIs"
	data, _ := json.Marshal(tok)
	url := fmt.Sprintf("/svc/clip/team/%s/token", team.Name)
	//t.Log("INFO: url: ", url)
	request, _ = http.NewRequest("POST", url, strings.NewReader(string(data)))

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
	if 1 == 2 {
		return
	}

	team.Token[0].Applications = make(map[string]Feature, 2)
	if len(team.Token) > 1 {
		team.Token[1].Applications = make(map[string]Feature, 2)
	}

	f := Feature{}
	team.Token[0].Applications["MobileWeb"] = f

	f.Flags = make(map[string]Flag, 2)
	fl := f.Flags["usePapiForBlogs"]
 	fl.Sandbox = 1
	fl.Development = 1
	fl.Attribs = make(map[string]bool, 2)
	fl.Attribs["exportToBrowser"] = false
	f.Flags["usePapiForBlogs"] = fl


	fl = f.Flags["useAmazonDirectMatch"]
	fl.Attribs = make(map[string]bool, 2)
	fl.Attribs["exportToBrowser"] = true
	fl.Sandbox = 1
	fl.Development = 1
	f.Flags["useAmazonDirectMatch"] = fl
	fmt.Println("TEAM: ", team.Token)

	team.Token[0].Applications["MobileWeb"] = f
	if len(team.Token) > 1 {
	team.Token[1].Applications["MobileWeb"] = f
	}

	data, _ = json.Marshal(team)

	t.Log(string(data))

	url = fmt.Sprintf("/svc/clip/team/%s", team.Name)
	request, _ = http.NewRequest("PUT", url, strings.NewReader(string(data)))
	router.ServeHTTP(w, request)
	t.Log(w.Body.String())
	//cleanUp(&team)
}

func cleanUp(team *Team) {

	var router *mux.Router
	router = mux.NewRouter()
	SetRoutes(router)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/svc/clip/team/%s", team.Name), nil)
	router.ServeHTTP(response, request)

}
