package clip

type WebResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Results string `json:"results,omitempty"`
}

type User struct {
	First string
	Last string
	Email string
	Groups []string
	Teams []Team
}

type Group struct {
	Name string
	Access uint8
	Users []User
	Teams []Team
}

type Team struct {
	Name string
	Groups []Group
	Token []Token
}

type Token struct {
	Team string
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
