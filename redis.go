package clip

import (
	"github.com/garyburd/redigo/redis"
	// "hash"
	"hash/fnv"
)

type Helper interface {
	GetKey(key string) (hash uint64)
	GetConn() (conn Conn, err error)
}

func (h Helper) GetKey(key string) (hash uint64) {
	thash := fnv.Hash64()
	_ = thash.Sum([]byte(key))
	hash = thash.Sum64()
}

func (h Helper) GetConn() (conn Conn, err error) {
	conn, err = redis.Dial("tcp", ":6379")
}
