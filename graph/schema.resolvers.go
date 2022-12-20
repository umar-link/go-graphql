package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/umar-link/go-graphql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   string(len(r.TodoList) + 1),
		Text: input.Text,
		Done: true,
		User: &model.User{
			ID:   input.UserID,
			Name: "sdfsdf",
		},
	}

	r.TodoList = append(r.TodoList, todo)
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.TodoList, nil
	// return []*model.Todo{
	// 	&model.Todo{
	// 		ID:   "1",
	// 		Text: "testing",
	// 		Done: true,
	// 		User: &model.User{
	// 			ID:   "1",
	// 			Name: "tes user",
	// 		},
	// 	},
	// }, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
