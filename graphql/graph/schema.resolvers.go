package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/generated"
	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/model"
	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/repository"
)

func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	return repository.Save(input), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (string, error) {
	return repository.RegisterUser(input), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input *model.RefreshToken) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	return repository.FindByID(id), nil
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return repository.AllDogs(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
