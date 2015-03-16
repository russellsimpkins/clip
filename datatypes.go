package "clip"

type User struct {
	id int
	First string
	Last string
	Email string
}

type Department struct {
	id int
	Name string
	Members []User
	Applications []string
	Token []Token
}

type Token struct {
	IntValue uint32
	StringValue string
	Department Department
}

type FeatureFlag struct {
	Token string             `json:token`
	Application string       `json:application`
	Features map[string]Flag `json:features`
}

type Flag struct {
	Development int `json:dev`
	Staging int     `json:stg`
	Integration int `json:int`
	Production int  `json:prd`
}
