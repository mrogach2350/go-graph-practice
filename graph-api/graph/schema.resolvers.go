package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/graph/generated"
	"main/graph/model"
	"net/http"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	response, err := http.Get(BaseUrl + UsersEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var usersResponse UsersResponse
	err = json.Unmarshal(body, &usersResponse)
	if err != nil {
		fmt.Println(err)
	}

	return usersResponse.Data, nil
}

func (r *queryResolver) User(ctx context.Context, input string) (*model.User, error) {
	response, err := http.Get(BaseUrl + UsersEndpoint + "/" + input)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var userResponse UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		fmt.Println(err)
	}

	return userResponse.Data, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	response, err := http.Get(BaseUrl + PostsEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var postResponse PostsResponse
	fmt.Println("data: ", string(body))
	err = json.Unmarshal(body, &postResponse)
	if err != nil {
		fmt.Println(err)
	}

	return postResponse.Data, nil
}

func (r *queryResolver) Post(ctx context.Context, input string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	response, err := http.Get(BaseUrl + CommentsEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var commentsResponse CommentsResponse
	fmt.Println("data: ", string(body))
	err = json.Unmarshal(body, &commentsResponse)
	if err != nil {
		fmt.Println(err)
	}

	return commentsResponse.Data, nil
}

func (r *queryResolver) Comment(ctx context.Context, input string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const BaseUrl = "http://localhost:7000/api"
const UsersEndpoint = "/users"
const PostsEndpoint = "/posts"
const CommentsEndpoint = "/comments"

type UsersResponse struct {
	Data []*model.User `json:"data"`
}
type UserResponse struct {
	Data *model.User `json:"data"`
}
type PostsResponse struct {
	Data []*model.Post `json:"data"`
}
type CommentsResponse struct {
	Data []*model.Comment `json:"data"`
}
