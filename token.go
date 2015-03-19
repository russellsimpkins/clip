package clip

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"math/rand"
	"time"
	"hash/crc32"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func GenerateToken() (token Token) {

	rndString := make([]byte, 10)
    for i:= 0 ; i < 10 ; i++ {
        rndString[i] = byte(randInt(65,90))
    }

	hash := sha256.New()
	hash.Write(rndString)
	md := hash.Sum(nil)
	token.StringValue = hex.EncodeToString(md)
	token.IntValue = crc32.ChecksumIEEE(md)
    return
}

//**********************************************************************
// HTTP HANDLER FUNCTIONS
//**********************************************************************

// handler to take requests from the interweb and return the token by name
func DeleteTokenHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		err   error
		token  Token
		vars  map[string]string
	)
	vars = mux.Vars(req)
	token = Token{}
	token.StringValue = vars["token"]
	token.Team = vars["team"]
	
	err = DeleteToken(&token)
	if err != nil {
		// for now, error out if we can't get the existing token
		str := fmt.Sprintf("Unable to delete the token: %s", err)
		SendError(500, str, writer)
		return
	}
	SendSuccess(writer)
	return
}

// handler to take requests from the interweb and return the token by name
func GetTokenHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		token  Token
		vars  map[string]string
	)
	//t := req.Header.Get("Authorization")

	vars = mux.Vars(req)
	token = Token{}
	token.StringValue = vars["token"]
	token.Team = vars["team"]
	err = GetToken(&token)
	if err != nil {
		str := fmt.Sprintf("Unable to fetch the token: %s", err)
		SendError(500, str, writer)
		return
	}
	body, err = json.Marshal(token)
	if err != nil {
		str := fmt.Sprintf("There was a problem encoding the token. Err: %s", err)
		SendError(500, str, writer)
		return
	}
	writer.Write(body)
	return
}

// handler to take request from the web and create a new token.
func CreateTokenHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		token Token
		team  Team
		vars  map[string]string
	)
	vars = mux.Vars(req)
	team = Team{}
	team.Name = vars["team"]
	//fmt.Println("Team: ", vars)
	err = GetTeamWithTeam(&team)
	if err != nil || len(team.Name) <= 0 {
		str := fmt.Sprintf("You Are requesting a token for a non-existant team. team: %s err: %s",
			team.Name, err)
		SendError(500, str, writer)
		return
	}
	// create a new token
	token = GenerateToken()
	token.Team = vars["team"]
	err = AddToken(&token)
	if err != nil {
		str := fmt.Sprintf("There was a problem creating the token. Err: %s", err)
		SendError(500, str, writer)
		return
	}

	// TODO <- Review when I have more time!
	if len(team.Token) == 0 {
		fmt.Println("HERE A")
		team.Token = make([]Token,1)
		team.Token[0] = token
	} else {
		fmt.Println("HERE B")
		t := make([]Token, len(team.Token), (cap(team.Token)+1)*2) // +1 in case cap(s) == 0
		copy(t, team.Token)
		team.Token = t
		team.Token = append(team.Token, token)
	}
	UpdateTeam(&team)

	
	body, err = json.Marshal(token)
	if err != nil {
		str := fmt.Sprintf("There was a problem Encoding the data. Err: %s", err)
		SendError(500, str, writer)
		return
	}
	writer.Write(body)	
	return
}

// handler to take request from the web and create a new token.
func UpdateTokenHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		token Token
		check Token
		vars  map[string]string
	)

	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		str := fmt.Sprintf("Unable to read in the body of the request: %s", body)
		SendError(500, str, writer)
		return
	}

	if body == nil || len(body) == 0 {
		str := fmt.Sprintf("No body in the request. We're expecting json of the token to create.")
		SendError(500, str, writer)
		return
	}

	err = json.Unmarshal(body, &token)

	if err != nil {
		str := fmt.Sprintf("There was a problem unmarshaling the json. error: %s", err)
		SendError(500, str, writer)
		return
	}
	vars = mux.Vars(req)
	check = Token{}
	check.Team = vars["team"]
	check.StringValue = vars["token"]
	err = GetToken(&check)

	if err != nil || len(check.StringValue) < 0 {
		str := fmt.Sprintf("You're updating that doesn't exists. error: %s", err)
		SendError(500, str, writer)
		return
	}

	// should we delete the old record?
	if token.StringValue != vars["token"] {
		// no
		str := fmt.Sprintf("You're tokens don't match. Panicing")
		SendError(500, str, writer)
	}
	err = UpdateToken(&token)
	if err != nil {
		str := fmt.Sprintf("There was a problem updating the token. Err: %s", err)
		SendError(500, str, writer)
		return
	}
	writer.Write(body)	
	return
}

//**********************************************************************
// DAO Methods
//**********************************************************************
func TokenKey(token *Token) (key string) {
	return fmt.Sprintf("%s:%s", token.Team, token.StringValue)
}

func AddToken(token *Token) (err error) {
	var (
		r     RedisHelper
		data  []byte
		check Token
		key   string
	)
	r, err = NewRedisHelper()
	if err != nil {
		return
	}
	defer r.Close()

	err = GetToken(token)
	if &check != nil && check.StringValue == token.StringValue {
		err = errors.New("You're trying to create a token that already exists.")
		return
	}
	data, err = json.Marshal(token)

	if err != nil {
		return
	}
	key = TokenKey(token)
	fmt.Printf("Key: %s", key)
	err = r.Store(key, data)
	return
}

func UpdateToken(token *Token) (err error) {
		var (
		r    RedisHelper
		data []byte
		key  string
	)
	r, err = NewRedisHelper()
	if err != nil {
		return
	}
	defer r.Close()
	data, err = json.Marshal(token)
	if err != nil {
		return
	}
	key = TokenKey(token)
	err = r.Store(key, data)
	return
}

func DeleteToken(token *Token) (err error) {
	var (
		r     RedisHelper
		key string
	)
	r, err = NewRedisHelper()
	if err != nil {
		return
	}
	defer r.Close()
	key = TokenKey(token)
	err = r.Delete(key)
	return
}

// Go to redis and get the token. 
func GetToken(token *Token) (err error) {
	var (
		r    RedisHelper
		data []byte
		key  string
	)
	r, err = NewRedisHelper()
	if err != nil {
		return
	}
	defer r.Close()
	key = TokenKey(token)
	data, _ = r.Fetch(key)
	json.Unmarshal(data, &token)
	return
}





/*
 * Taken right from the docs. This is how you populate a map of maps
 */
func add(m map[string]map[string]int, env, feature string, val int) {
    mm, ok := m[env]
    if !ok {
        mm = make(map[string]int)
        m[env] = mm
    }
    mm[feature] = val

	//var sample map[string]map[string]int
	//sample = make(map[string]map[string]int)
	//add(sample, "dev", "feature", 1)
	//fmt.Printf("The feature is on? %d\n", sample["dev"]["feature"])
}
