package twitch

// UserResponse represents a response containing metadata from a particular Twitch User
type UserResponse struct {
	Data       UserData           `json:"data"`
	Extensions ResponseExtensions `json:"extensions"`
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
