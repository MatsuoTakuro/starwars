package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MatsuoTakuro/starwars/graph/generated"
	"github.com/MatsuoTakuro/starwars/graph/model"
)

func (r *mutationResolver) CreateReview(ctx context.Context, episode model.Episode, review model.ReviewInput) (*model.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
