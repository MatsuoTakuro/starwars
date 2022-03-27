package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/MatsuoTakuro/starwars/graph/generated"
	"github.com/MatsuoTakuro/starwars/graph/model"
)

func (r *queryResolver) Hero(ctx context.Context, episode *model.Episode) (model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Reviews(ctx context.Context, episode model.Episode, since *time.Time) ([]*model.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Search(ctx context.Context, text string) ([]model.SearchResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Character(ctx context.Context, id string) (model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Droid(ctx context.Context, id string) (*model.Droid, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Human(ctx context.Context, id string) (*model.Human, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Starship(ctx context.Context, id string) (*model.Starship, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
