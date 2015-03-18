package clip

import (
	"github.com/garyburd/redigo/redis"
	"hash/fnv"
)

type RedisHelper struct {
	Conn redis.Conn
}

type Helper interface {
	GetKey(key string) (hash uint64)
	GetConn() (conn redis.Conn, err error)
	GetNextVal(key string) (id uint32, err error)
	Close()
}

func NewRedisHelper() (r RedisHelper) {
	r = RedisHelper{}
	_, err := r.GetConn()
	if err != nil {
		return
	}
	return
}

func (h *RedisHelper) Close() {
	if h.Conn != nil {
		h.Conn.Close()
	}
}

func (h *RedisHelper) GetKey(key string) (hash uint64) {
	thash := fnv.New64()
	_ = thash.Sum([]byte(key))
	hash = thash.Sum64()
	return
}

func (h *RedisHelper) GetConn() (conn redis.Conn, err error) {
	conn, err = redis.Dial("tcp", ":6379")
	if err != nil {
		return
	}
	h.Conn = conn
	return
}

// this is my utility function for autoincrements
func (h *RedisHelper) GetNextVal(key string) (id int64, err error) {
	sql := "INCR"
	var (
		reply interface{}
	)
	reply, err = h.Conn.Do(sql, key)
	if err != nil {
		return
	}
	id, err = redis.Int64(reply, err)
	return
}

// this utility function attempts to store a key value pair
func (h *RedisHelper) Store(key string, data []byte) (err error) {
	_, err = h.Conn.Do("SET", []byte(key), data)
	return
}

func (h *RedisHelper) Delete(key string) (err error) {
	_, err = h.Conn.Do("DEL", []byte(key), data)
	return
}

// this utility function attempts to store a key value pair
func (h *RedisHelper) Fetch(key string) (data []byte, err error) {
	var reply interface{}
	var sniff string
	reply, err = h.Conn.Do("GET", key)
	sniff, err = redis.String(reply, err)
	data = []byte(sniff)
	return
}
