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
	response, err := http.Get(model.BaseUrl + model.UsersEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var usersResponse model.UsersResponse
	err = json.Unmarshal(body, &usersResponse)
	if err != nil {
		fmt.Println(err)
	}

	return usersResponse.Data, nil
}

func (r *queryResolver) User(ctx context.Context, input string) (*model.User, error) {
	response, err := http.Get(model.BaseUrl + model.UsersEndpoint + "/" + input)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var userResponse model.UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		fmt.Println(err)
	}

	return userResponse.Data, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	response, err := http.Get(model.BaseUrl + model.PostsEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var postResponse model.PostsResponse
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
	response, err := http.Get(model.BaseUrl + model.CommentsEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var commentsResponse model.CommentsResponse
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
