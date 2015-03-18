package clip
import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func FetchTeam(name string) (team Team, err error) {
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

func GetTeamHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		team  Team
		vars  map[string]string
	)
	//t := req.Header.Get("Authorization")

	vars = mux.Vars(req)
	team, err = FetchTeam(vars["name"])
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


func CreateTeamHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		team  Team
		check Team
		key   string
	)
	//t := req.Header.Get("Authorization")

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

	check, err = FetchTeam(team.Name)

	if err != nil || len(check.Name) > 0 {
		str := fmt.Sprintf("You're creating a team that already exists. error: %s %s", err, check.Name)
		SendError(500, str, writer)
		return
	}

	key = fmt.Sprintf("%s:%s", "team:", team.Name)
	err = AddTeam(key, &team)
	if err != nil {
		str := fmt.Sprintf("There was a problem creating the team. Err: %s", err)
		SendError(500, str, writer)
		return
	}
	writer.Write(body)	
	return
}



func AddTeam(key string, team *Team) (err error) {
	var (
		r     RedisHelper
		data  []byte
		check Team
	)
	r = NewRedisHelper()
	defer r.Close()

	check, err = FetchTeam(key)
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
	return
}

func DeleteTeam(team Team) (err error) {
	return
}

