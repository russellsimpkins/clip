package clip

import (
	"github.com/garyburd/redigo/redis"
	"hash/fnv"
)

type RedisHelper struct {

}

type Helper interface {
	GetKey(key string) (hash uint64)
	GetConn() (conn redis.Conn, err error)
}

func (h *RedisHelper) GetKey(key string) (hash uint64) {
	thash := fnv.New64()
	_ = thash.Sum([]byte(key))
	hash = thash.Sum64()
	return
}

func (h *RedisHelper) GetConn() (conn redis.Conn, err error) {
	conn, err = redis.Dial("tcp", ":6379")
	return
}
