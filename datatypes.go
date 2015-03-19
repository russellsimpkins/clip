package clip

type WebResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Results string `json:"results,omitempty"`
}

type User struct {
	First string   `json:"first"`
	Last string    `json:"last"`
	Email string   `json:"email"`
}

type Team struct {
	Name string     `json:"name"`
	Users []User    `json:"users"`
	Token []Token   `json:"tokens"`
}

type Names struct {
	Name string     `json:"teamName"`
}

type Teams struct {
	Teams []Names `json:"teams"`
}

type Token struct {
	Team string      `json:"team"`
	IntValue uint32  `json:"crc32"`
	StringValue string `json:"token"`
	Applications map[string]Feature `json:"apps"`
}

type Feature struct {
	Flags map[string]Flag `json:"features"`
}

type Flag struct {
	Attribs map[string]int `json:"attributes"`
	Sandbox     int `json:"sbx"`
	Development int `json:"dev"`
	Staging     int `json:"stg"`
	Integration int `json:"int"`
	Production  int `json:"prd"`
}
