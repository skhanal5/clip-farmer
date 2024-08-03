package model

type UserResponse struct {
	Data       UserData   `json:"data"`
	Extensions Extensions `json:"extensions"`
}

type UserData struct {
	User User `json:"user"`
}

type User struct {
	Id    string `json:"id"`
	Clips Clips  `json:"clips"`
}

type Clips struct {
	Edges []Edges `json:"edges"`
}

type Edges struct {
	Node Node `json:"node"`
}

type Node struct {
	Slug string `json:"slug"`
}
