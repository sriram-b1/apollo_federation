package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sriram-b1/apollo_federation/Route/graph/generated"
	"github.com/sriram-b1/apollo_federation/Route/graph/model"
)

func (r *entityResolver) FindRouteByID(ctx context.Context, id *string) (*model.Route, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
