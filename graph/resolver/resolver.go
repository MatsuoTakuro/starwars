//go:generate rm -rf generated
//go:generate go run ../../testdata/gqlgen.go

package resolver

import (
	"context"
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"github.com/MatsuoTakuro/starwars/graph/generated"
	"github.com/MatsuoTakuro/starwars/graph/model"
)

type Resolver struct {
	humans    map[string]model.Human
	droid     map[string]model.Droid
	starships map[string]model.Starship
	reviews   map[model.Episode][]*model.Review
}

func (r *Resolver) Droid() generated.DroidResolver {
	return &droidResolver{r}
}

func (r *Resolver) FriendsConnection() generated.FriendsConnectionResolver {
	return &friendsConnectionResolver{r}
}

func (r *Resolver) Human() generated.HumanResolver {
	return &humanResolver{r}
}

func (r *Resolver) Starship() generated.StarshipResolver {
	return &starshipResolver{r}
}

func (r *Resolver) resolveCharacters(ctx context.Context, ids []string) ([]model.Character, error) {
	result := make([]model.Character, len(ids))
	for i, id := range ids {
		char, err := r.Query().Character(ctx, id)
		if err != nil {
			return nil, err
		}
		result[i] = char
	}
	return result, nil
}

type droidResolver struct{ *Resolver }

func (r *droidResolver) Friends(ctx context.Context, obj *model.Droid) ([]model.Character, error) {
	return r.resolveCharacters(ctx, obj.FriendIds)
}

func (r *droidResolver) FriendsConnection(ctx context.Context, obj *model.Droid, first *int, after *string) (*model.FriendsConnection, error) {
	return r.resolveFriendConnection(ctx, obj.FriendIds, first, after)
}

type friendsConnectionResolver struct{ *Resolver }

func (r *Resolver) resolveFriendConnection(_ context.Context, ids []string, first *int, after *string) (*model.FriendsConnection, error) {
	from := 0
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return nil, err
		}
		from = i
	}

	to := len(ids)
	if first != nil {
		to = from + *first
		if to > len(ids) {
			to = len(ids)
		}
	}

	return &model.FriendsConnection{
		Ids:  ids,
		From: from,
		To:   to,
	}, nil
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

type humanResolver struct{ *Resolver }

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

type starshipResolver struct{ *Resolver }

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

func NewResolver() generated.Config {
	r := Resolver{}
	r.humans = map[string]model.Human{
		"1000": {
			CharacterFields: model.CharacterFields{
				ID:        "1000",
				Name:      "Luke Skywalker",
				FriendIds: []string{"1002", "1003", "2000", "2001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			HeightMeters: 1.72,
			Mass:         77,
			StarshipIds:  []string{"3001", "3003"},
		},
		"1001": {
			CharacterFields: model.CharacterFields{
				ID:        "1001",
				Name:      "Darth Vader",
				FriendIds: []string{"1004"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			HeightMeters: 2.02,
			Mass:         136,
			StarshipIds:  []string{"3002"},
		},
		"1002": {
			CharacterFields: model.CharacterFields{
				ID:        "1002",
				Name:      "Han Solo",
				FriendIds: []string{"1000", "1003", "2001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			HeightMeters: 1.8,
			Mass:         80,
			StarshipIds:  []string{"3000", "3003"},
		},
		"1003": {
			CharacterFields: model.CharacterFields{
				ID:        "1003",
				Name:      "Leia Organa",
				FriendIds: []string{"1000", "1002", "2000", "2001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			HeightMeters: 1.5,
			Mass:         49,
		},
		"1004": {
			CharacterFields: model.CharacterFields{
				ID:        "1004",
				Name:      "Wilhuff Tarkin",
				FriendIds: []string{"1001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope},
			},
			HeightMeters: 1.8,
			Mass:         0,
		},
	}

	r.droid = map[string]model.Droid{
		"2000": {
			CharacterFields: model.CharacterFields{
				ID:        "2000",
				Name:      "C-3PO",
				FriendIds: []string{"1000", "1002", "1003", "2001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			PrimaryFunction: "Protocol",
		},
		"2001": {
			CharacterFields: model.CharacterFields{
				ID:        "2001",
				Name:      "R2-D2",
				FriendIds: []string{"1000", "1002", "1003"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			PrimaryFunction: "Astromech",
		},
	}

	r.starships = map[string]model.Starship{
		"3000": {
			ID:   "3000",
			Name: "Millennium Falcon",
			History: [][]int{
				{1, 2},
				{4, 5},
				{1, 2},
				{3, 2},
			},
			Length: 34.37,
		},
		"3001": {
			ID:   "3001",
			Name: "X-Wing",
			History: [][]int{
				{6, 4},
				{3, 2},
				{2, 3},
				{5, 1},
			},
			Length: 12.5,
		},
		"3002": {
			ID:   "3002",
			Name: "TIE Advanced x1",
			History: [][]int{
				{3, 2},
				{7, 2},
				{6, 4},
				{3, 2},
			},
			Length: 9.2,
		},
		"3003": {
			ID:   "3003",
			Name: "Imperial shuttle",
			History: [][]int{
				{1, 7},
				{3, 5},
				{5, 3},
				{7, 1},
			},
			Length: 20,
		},
	}

	r.reviews = map[model.Episode][]*model.Review{}

	return generated.Config{
		Resolvers: &r,
	}
}
