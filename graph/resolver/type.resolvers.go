package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/MatsuoTakuro/starwars/graph/generated"
	"github.com/MatsuoTakuro/starwars/graph/model"
)

func (r *droidResolver) Friends(ctx context.Context, obj *model.Droid) ([]model.Character, error) {
	return r.resolveCharacters(ctx, obj.FriendIds)
}

func (r *droidResolver) FriendsConnection(ctx context.Context, obj *model.Droid, first *int, after *string) (*model.FriendsConnection, error) {
	return r.resolveFriendConnection(ctx, obj.FriendIds, first, after)
}

func (r *friendsConnectionResolver) Edges(ctx context.Context, obj *model.FriendsConnection) ([]*model.FriendsEdge, error) {
	friends, err := r.resolveCharacters(ctx, obj.Ids)
	if err != nil {
		return nil, err
	}

	edges := make([]*model.FriendsEdge, obj.To-obj.From)
	for i := range edges {
		edges[i] = &model.FriendsEdge{
			Cursor: model.EncodeCursor(obj.From + i),
			Node:   friends[obj.From+i],
		}
	}
	return edges, nil
}

func (r *friendsConnectionResolver) Friends(ctx context.Context, obj *model.FriendsConnection) ([]model.Character, error) {
	return r.resolveCharacters(ctx, obj.Ids)
}

func (r *humanResolver) Friends(ctx context.Context, obj *model.Human) ([]model.Character, error) {
	return r.resolveCharacters(ctx, obj.FriendIds)
}

func (r *humanResolver) FriendsConnection(ctx context.Context, obj *model.Human, first *int, after *string) (*model.FriendsConnection, error) {
	return r.resolveFriendConnection(ctx, obj.FriendIds, first, after)
}

func (r *humanResolver) Starships(ctx context.Context, obj *model.Human) ([]*model.Starship, error) {
	var result []*model.Starship
	for _, id := range obj.StarshipIds {
		char, err := r.Query().Starship(ctx, id)
		if err != nil {
			return nil, err
		}
		if char != nil {
			result = append(result, char)
		}
	}
	return result, nil
}

func (r *starshipResolver) Length(ctx context.Context, obj *model.Starship, unit *model.LengthUnit) (float64, error) {
	switch *unit {
	case model.LengthUnitMeter, "":
		return obj.Length, nil
	case model.LengthUnitFoot:
		return obj.Length * 3.28084, nil
	default:
		return 0, errors.New("invalid unit")
	}
}

// Droid returns generated.DroidResolver implementation.
func (r *Resolver) Droid() generated.DroidResolver { return &droidResolver{r} }

// FriendsConnection returns generated.FriendsConnectionResolver implementation.
func (r *Resolver) FriendsConnection() generated.FriendsConnectionResolver {
	return &friendsConnectionResolver{r}
}

// Human returns generated.HumanResolver implementation.
func (r *Resolver) Human() generated.HumanResolver { return &humanResolver{r} }

// Starship returns generated.StarshipResolver implementation.
func (r *Resolver) Starship() generated.StarshipResolver { return &starshipResolver{r} }

type droidResolver struct{ *Resolver }
type friendsConnectionResolver struct{ *Resolver }
type humanResolver struct{ *Resolver }
type starshipResolver struct{ *Resolver }
