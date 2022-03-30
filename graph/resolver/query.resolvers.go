package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"
	"time"

	"github.com/MatsuoTakuro/starwars/graph/generated"
	"github.com/MatsuoTakuro/starwars/graph/model"
)

func (r *queryResolver) Hero(ctx context.Context, episode *model.Episode) (model.Character, error) {
	if *episode == model.EpisodeEmpire {
		return r.humans["1000"], nil
	}
	return r.droid["2001"], nil
}

func (r *queryResolver) Reviews(ctx context.Context, episode model.Episode, since *time.Time) ([]*model.Review, error) {
	if since == nil {
		return r.reviews[episode], nil
	}

	var filtered []*model.Review
	for _, rev := range r.reviews[episode] {
		if rev.Time.After(*since) {
			filtered = append(filtered, rev)
		}
	}
	return filtered, nil
}

func (r *queryResolver) Search(ctx context.Context, text string) ([]model.SearchResult, error) {
	var l []model.SearchResult
	for _, h := range r.humans {
		if strings.Contains(h.Name, text) {
			l = append(l, h)
		}
	}
	for _, d := range r.droid {
		if strings.Contains(d.Name, text) {
			l = append(l, d)
		}
	}
	for _, s := range r.starships {
		if strings.Contains(s.Name, text) {
			l = append(l, s)
		}
	}
	return l, nil
}

func (r *queryResolver) Character(ctx context.Context, id string) (model.Character, error) {
	if h, ok := r.humans[id]; ok {
		return &h, nil
	}
	if d, ok := r.droid[id]; ok {
		return &d, nil
	}
	return nil, nil
}

func (r *queryResolver) Droid(ctx context.Context, id string) (*model.Droid, error) {
	if d, ok := r.droid[id]; ok {
		return &d, nil
	}
	return nil, nil
}

func (r *queryResolver) Human(ctx context.Context, id string) (*model.Human, error) {
	if h, ok := r.humans[id]; ok {
		return &h, nil
	}
	return nil, nil
}

func (r *queryResolver) Starship(ctx context.Context, id string) (*model.Starship, error) {
	if s, ok := r.starships[id]; ok {
		return &s, nil
	}
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
