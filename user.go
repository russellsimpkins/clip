package clip

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchUser(email string) (user User, err error) {
	var (
		r    RedisHelper
		data []byte
	)
	r = NewRedisHelper()
	defer r.Close()
	data, err = r.Fetch(email)
	if err != nil {
		return
	}
	json.Unmarshal(data, &user)
	return
}

// store a new user, using the email as the key
func AddUser(user *User) (err error) {
	var (
		r     RedisHelper
		data  []byte
		check User
	)

	r = NewRedisHelper()
	defer r.Close()

	check, err = FetchUser(user.Email)

	if &check != nil && check.Email == user.Email {
		err = errors.New("You're trying to create a user that already exists.")
		return
	}
	data, err = json.Marshal(user)

	if err != nil {
		return
	}
	err = r.Store(user.Email, data)
	return
}

// since email is the key, this is all we need
func DeleteUser(email string) (err error) {
	var r RedisHelper
	r = NewRedisHelper()
	_, err = r.Conn.Do("DEL", email)
	return
}

// pass in the original email, in case we're changing the email
func UpdateUser(origEmail string, user *User) (err error) {
	var (
		r    RedisHelper
		data []byte
	)
	r = NewRedisHelper()
	defer r.Close()
	data, err = json.Marshal(user)
	if err != nil {
		return
	}
	if origEmail != user.Email {
		_, err = r.Conn.Do("DEL", origEmail)
		if err != nil {
			return
		}
	}
	err = r.Store(user.Email, data)
	return
}

// create user expects json in the post body
func CreateUserHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body []byte
		vars map[string]string
		err  error
		user User
	)

	vars = mux.Vars(req)
	body, err = ioutil.ReadAll(req.Body)

	if err != nil {
		str := fmt.Sprintf("Unable to read in the body of the request: %s", body)
		SendError(500, str, writer)
		return
	}

	if body == nil || len(body) == 0 {
		str := fmt.Sprintf("No body in the request. We're expecting json of the user")
		SendError(500, str, writer)
		return
	}

	err = json.Unmarshal(body, &user)

	if err != nil {
		str := fmt.Sprintf("There was a problem unmarshaling the json. error: %s", err)
		SendError(500, str, writer)
		return
	}

	if vars["email"] != user.Email {

	}
	err = AddUser(&user)

	if err != nil {
		str := fmt.Sprintf("There was a problem adding the user. error: %s", err)
		SendError(500, str, writer)
		return
	}

	body, err = json.Marshal(user)
	writer.Write(body)
}

// create user expects json in the post body
func UpdateUserHandler(writer http.ResponseWriter, req *http.Request) {
	var (
		body  []byte
		err   error
		user  User
		check User
		vars  map[string]string
	)
	vars = mux.Vars(req)
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		str := fmt.Sprintf("Unable to read in the body of the request: %s", body)
		SendError(500, str, writer)
		return
	}

	if body == nil || len(body) == 0 {
		str := fmt.Sprintf("No body in the request. We're expecting json of the user")
		SendError(500, str, writer)
		return
	}

	err = json.Unmarshal(body, &user)

	if err != nil {
		str := fmt.Sprintf("There was a problem unmarshaling the json. error: %s", err)
		SendError(500, str, writer)
		return
	}

	check, err = FetchUser(vars["email"])

	if err != nil || check.Email == "" {
		str := fmt.Sprintf("Unable to load user. error: %s", err)
		SendError(500, str, writer)
		return
	}

	err = UpdateUser(vars["email"], &user)
	if err != nil {
		str := fmt.Sprintf("There was a problem updating the user. Err: %s", err)
		SendError(500, str, writer)
		return
	}
	body, err = json.Marshal(user)
	writer.Write(body)
}

func SendError(status int, data string, writer http.ResponseWriter) {
	what := WebResponse{}
	what.Status = status
	what.Message = data
	body, _ := json.Marshal(what)
	out := string(body)
	log.Print(out)
	http.Error(writer, out, status)
	return
}
