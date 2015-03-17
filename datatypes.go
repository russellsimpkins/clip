package clip

type User struct {
	First string
	Last string
	Email string
	Groups []string
	Teams []Team
}

type Group struct {
	id int
	Name string
	Access uint8
	Users []User
	Teams []Team
}

type Team struct {
	id int
	Name string
	Groups []Group
	Token []Token
}

type Token struct {
	IntValue uint32
	StringValue string
	Applications map[string]Feature `json:apps`
}

type Feature struct {
	Flags map[string]Flag `json:features`
}

type Flag struct {
	Development int `json:dev`
	Staging int     `json:stg`
	Integration int `json:int`
	Production int  `json:prd`
}
