package clip

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

//**********************************************************************
// HTTP HANDLER FUNCTIONS
//**********************************************************************

// handler to take requests from the interweb and return the team by name
func DeleteTeamHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		err   error
		team  Team
		vars  map[string]string
	)
	vars = mux.Vars(req)
	team = Team{}
	team.Name = vars["name"]
	
	err = DeleteTeam(&team)
	if err != nil {
		// for now, error out if we can't get the existing team
		str := fmt.Sprintf("Unable to delete the team: %s", err)
		SendError(500, str, writer)
		return
	}
	SendSuccess(writer)
	return
}

// handler to take requests from the interweb and return the team by name
func GetTeamHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		team  Team
		vars  map[string]string
	)
	//t := req.Header.Get("Authorization")

	vars = mux.Vars(req)
	team, err = GetTeam(vars["name"])
	if err != nil {
		str := fmt.Sprintf("Unable to fetch the team: %s", err)
		SendError(500, str, writer)
		return
	}
	body, err = json.Marshal(team)
	if err != nil {
		str := fmt.Sprintf("There was a problem encoding the team. Err: %s", err)
		SendError(500, str, writer)
		return
	}
	writer.Write(body)
	return
}

// handler to take request from the web and create a new team.
func CreateTeamHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		team  Team
		check Team
	)

	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		str := fmt.Sprintf("Unable to read in the body of the request: %s", body)
		SendError(500, str, writer)
		return
	}

	if body == nil || len(body) == 0 {
		str := fmt.Sprintf("No body in the request. We're expecting json of the team to create.")
		SendError(500, str, writer)
		return
	}

	err = json.Unmarshal(body, &team)

	if err != nil {
		str := fmt.Sprintf("There was a problem unmarshaling the json. error: %s", err)
		SendError(500, str, writer)
		return
	}

	check, err = GetTeam(team.Name)

	if err != nil || len(check.Name) > 0 {
		str := fmt.Sprintf("You're creating a team that already exists. error: %s %s", err, check.Name)
		SendError(500, str, writer)
		return
	}

	err = AddTeam(&team)
	if err != nil {
		str := fmt.Sprintf("There was a problem creating the team. Err: %s", err)
		SendError(500, str, writer)
		return
	}
	writer.Write(body)	
	return
}

// handler to take request from the web and create a new team.
func UpdateTeamHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		team  Team
		check Team
		vars  map[string]string
	)

	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		str := fmt.Sprintf("Unable to read in the body of the request: %s", body)
		SendError(500, str, writer)
		return
	}

	if body == nil || len(body) == 0 {
		str := fmt.Sprintf("No body in the request. We're expecting json of the team to create.")
		SendError(500, str, writer)
		return
	}

	err = json.Unmarshal(body, &team)

	if err != nil {
		str := fmt.Sprintf("There was a problem unmarshaling the json. error: %s", err)
		SendError(500, str, writer)
		return
	}
	vars = mux.Vars(req)
	check, err = GetTeam(vars["name"])

	if err != nil || len(check.Name) < 0 {
		str := fmt.Sprintf("You're updating that doesn't exists. error: %s %s", err, check.Name)
		SendError(500, str, writer)
		return
	}

	// should we delete the old record?
	if team.Name != vars["name"] {
		// yes
		DeleteTeam(&check)
		err = AddTeam(&team)
	} else {
		err = UpdateTeam(&team)
	}
	
	if err != nil {
		str := fmt.Sprintf("There was a problem updating the team. Err: %s", err)
		SendError(500, str, writer)
		return
	}
	writer.Write(body)	
	return
}

//**********************************************************************
// DAO Methods
//**********************************************************************
func TeamKey(team *Team) (key string) {
	return fmt.Sprintf("%s:%s", "team:", team.Name)
}

func AddTeam(team *Team) (err error) {
	var (
		r     RedisHelper
		data  []byte
		check Team
		key   string
	)
	r = NewRedisHelper()
	defer r.Close()
	key = TeamKey(team)
	check, err = GetTeam(key)
	if &check != nil && check.Name == team.Name {
		err = errors.New("You're trying to create a team that already exists.")
		return
	}
	data, err = json.Marshal(team)

	if err != nil {
		return
	}
	err = r.Store(key, data)
	return
}

func UpdateTeam(team *Team) (err error) {
		var (
		r    RedisHelper
		data []byte
		key  string
	)
	r = NewRedisHelper()
	defer r.Close()
	data, err = json.Marshal(team)
	if err != nil {
		return
	}
	key = TeamKey(team)
	err = r.Store(key, data)
	return
}

func DeleteTeam(team *Team) (err error) {
	var (
		r     RedisHelper
		key string
	)
	r = NewRedisHelper()
	defer r.Close()
	key = TeamKey(team)
	err = r.Delete(key)
	return
}

// Go to redis and get the team. 
func GetTeam(name string) (team Team, err error) {
	var (
		r    RedisHelper
		data []byte
		key  string
	)
	r = NewRedisHelper()
	defer r.Close()
	key = fmt.Sprintf("%s:%s", "team:", name)
	data, _ = r.Fetch(key)
	json.Unmarshal(data, &team)
	return
}
