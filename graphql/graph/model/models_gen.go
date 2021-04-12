// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Dog struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	IsGoodBoy bool   `json:"isGoodBoy"`
	User      *User  `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewDog struct {
	Name      string `json:"name"`
	IsGoodBoy bool   `json:"isGoodBoy"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshToken struct {
	Token string `json:"token"`
}

type User struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
