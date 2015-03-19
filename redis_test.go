package clip

import (
	"testing"
)

func TestPutRedis(t *testing.T) {
	var (
		r   RedisHelper
		id  int64
		err error
	)
	r, err = NewRedisHelper()
	if err != nil {
		return
	}
	defer r.Close()

	id, err = r.GetNextVal("sys.Token")
	if err != nil {
		t.Log("Something's wrong. We didn't get back a value", err)
		t.Fail()
		return
	}
	t.Log("We got a value %d", id)
	id, err = r.GetNextVal("test.Token")
	if err != nil {
		t.Log("Something's wrong. We didn't get back a value", err)
		t.Fail()
		return
	}
	// cleaning up after myself
	_, _ = r.Conn.Do("DEL", "sys.Token")
	t.Log("We got a value %d", id)

}
