package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/MatsuoTakuro/starwars/graph/generated"
	"github.com/MatsuoTakuro/starwars/graph/model"
)

func (r *mutationResolver) CreateReview(ctx context.Context, episode model.Episode, review model.Review) (*model.Review, error) {
	review.Time = time.Now()
	time.Sleep(1 * time.Second)
	r.reviews[episode] = append(r.reviews[episode], &review)
	return &review, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
