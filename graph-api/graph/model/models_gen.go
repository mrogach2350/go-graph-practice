// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	UUID   string `json:"UUID"`
	Body   string `json:"Body"`
	Name   string `json:"Name"`
	PostID string `json:"PostId"`
}

type NewComment struct {
	Body   string `json:"Body"`
	Name   string `json:"Name"`
	PostID string `json:"PostId"`
}

type NewPost struct {
	Body   string `json:"Body"`
	Title  string `json:"Title"`
	UserID string `json:"UserId"`
}

type NewUser struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}

type Post struct {
	UUID   string `json:"UUID"`
	Body   string `json:"Body"`
	Title  string `json:"Title"`
	UserID string `json:"UserId"`
}

type User struct {
	UUID      string `json:"UUID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}
