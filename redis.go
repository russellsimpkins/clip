package clip

import (
	"fmt"
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

func NewRedisHelper() (r RedisHelper, err error) {
	r = RedisHelper{}
	_, err = r.GetConn()
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

// utility to add an item to a set
func (h *RedisHelper) AddToSet(set string, item string) (err error) {
	_, err = h.Conn.Do("SADD", set, item);
	return
}

// utility to remove an item to a set
func (h *RedisHelper) RemFromSet(set string, item string) (err error) {
	_, err = h.Conn.Do("SREM", set, item);
	return
}

func (h *RedisHelper) GetMembers(set string) (items []string, err error) {
	var (
		reply interface{}
	)
	reply, err = h.Conn.Do("SMEMBERS", set);
	if err != nil {
		return
	}
	items, err = redis.Strings(reply, err)
	if err != nil {
		return
	}
	return
}

// this utility function attempts to store a key value pair
func (h *RedisHelper) Store(key string, data []byte) (err error) {
	if 1 == 2 {
		fmt.Printf("key: %s\n", key)
	}
	_, err = h.Conn.Do("SET", []byte(key), data)
	return
}

// utility to delete a key
func (h *RedisHelper) Delete(key string) (err error) {
	_, err = h.Conn.Do("DEL", []byte(key))
	return
}

// this utility function attempts to store a key value pair
func (h *RedisHelper) Fetch(key string) (data []byte, err error) {
	var reply interface{}
	var sniff string
	reply, err = h.Conn.Do("GET", key)
	if err != nil {
		return
	}
	sniff, err = redis.String(reply, err)
	if err != nil {
		return
	}
	data = []byte(sniff)
	return
}
