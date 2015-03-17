package clip

import (
	"github.com/garyburd/redigo/redis"
)

// this file holds all the logic for handling users in the system
type ClipUser interface {
	AddUser(user *User) (err error)
	UpdateUser(user *User) (err error)
	DeleteUser(user User) (err error)
}



func getNextUserId() (err error, id uint32) {
	sql := "INCR sys.users"
	var (
		conn redis.Conn
		reply []interface{}
	)
	conn, err = redis.Dial("tcp", ":6379")
	if err != nil {
		return
	}
	defer conn.Close()
	reply, err = redis.Values(conn.Do(sql))

	if err != nil {
		return
	}
	
	_, err = redis.Scan(reply, &id)
	return
}


