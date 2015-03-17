package clip

import (
	"encoding/json"
//	"fmt"
)

func FetchUser(email string) (user User, err error) {
	var (
		r RedisHelper
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
		r RedisHelper
		data []byte
	)
	
	r = NewRedisHelper()
	defer r.Close()

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
		r RedisHelper
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
	
