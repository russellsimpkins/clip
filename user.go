package clip

import (
	"encoding/json"
)

// this file holds all the logic for handling users in the system

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
	r.Store(user.Email, data)
	return
}

func DeleteUser(user *User) (err error) {
	var r RedisHelper
	r = NewRedisHelper()
	_, err = r.Conn.Do("DEL", user.Email)

	return
}



