package model

const BaseUrl = "http://localhost:7000/api"
const UsersEndpoint = "/users"
const PostsEndpoint = "/posts"
const CommentsEndpoint = "/comments"

type UsersResponse struct {
	Data []*User `json:"data"`
}
type UserResponse struct {
	Data *User `json:"data"`
}
type PostsResponse struct {
	Data []*Post `json:"data"`
}
type CommentsResponse struct {
	Data []*Comment `json:"data"`
}