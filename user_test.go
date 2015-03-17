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
	user.First = "larry"
	err = UpdateUser(user.Email, &user)
	if err != nil {
		t.Log("There was an error: ", err)
		t.Fail()
	}
	
	if user.First != "larry" {
		t.Log("We expected the first name to change and it did not")
		t.Fail()
	}
	user = User{}
	user, err = FetchUser("russell.simpkins@nytimes.com")

	if err != nil {
		t.Log("We expected to pull data from redis, it failed.", err)
		t.Fail()
	}

	if user.Email != "russell.simpkins@nytimes.com" {
		t.Log("We expected to pull data from redis, it failed.", user)
		t.Fail()
	}
	
	err = DeleteUser(user.Email)
	if err != nil {
		t.Log("There was an error: ", err)
		t.Fail()
	}

}
