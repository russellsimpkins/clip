package clip

import (
	"testing"
)

func TestAddUser(t *testing.T) {

	var user User
	user = User{}
	user.First = "Russ"
	user.Last = "Simpkins"
	user.Email = "russell.simpkins@nytimes.com"

	err := AddUser(&user)
	if err != nil {
		t.Log("There was an error: ", err)
		t.Fail()
	}
	err = DeleteUser(&user)
	if err != nil {
		t.Log("There was an error: ", err)
		t.Fail()
	}

}
