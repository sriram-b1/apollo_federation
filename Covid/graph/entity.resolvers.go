package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sriram-b1/covid19/graph/generated"
	"github.com/sriram-b1/covid19/graph/model"
)

func (r *entityResolver) FindCovidByDate(ctx context.Context, date *string) (*model.Covid, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
