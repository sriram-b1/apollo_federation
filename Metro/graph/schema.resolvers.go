package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sriram-b1/metro/graph/generated"
	"github.com/sriram-b1/metro/graph/model"
)

func (r *mutationResolver) AddMetro(ctx context.Context, input model.AddMetro) ([]*model.Metro, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Metro(ctx context.Context) ([]*model.Metro, error) {
	const endpoint = "https://api.metro.net/agencies/lametro/routes/"
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var routes JsonData
	json.Unmarshal(responseData, &routes)
	for _, item := range routes.Items {
		it := &model.Metro{ID: item.ID, DisplayName: item.DisplayName}
		fmt.Print("Item: ", it)
	}

	return routes.Items, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type JsonData struct {
	Items []*model.Metro `json:"items"`
}
