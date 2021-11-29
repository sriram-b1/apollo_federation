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

	"github.com/sriram-b1/covid19/graph/generated"
	"github.com/sriram-b1/covid19/graph/model"
)

type Covid_1 struct {
	Active string `json:"active"`

	Confirmed string `json:"confirmed"`

	Deaths string `json:"deaths"`

	Deltaconfirmed string `json:"deltaconfirmed"`

	Deltadeaths string `json:"deltadeaths"`

	Deltarecovered string `json:"deltarecovered"`

	Lastupdatedtime string `json:"lastupdatedtime"`

	Migratedother string `json:"migratedother"`

	Recovered string `json:"recovered"`

	State string `json:"state"`

	Statecode string `json:"statecode"`

	Statenotes string `json:"statenotes"`
}

type Covid_2 struct {
	Firstdoseadministered string `json:"firstdoseadministered"`
}

type CovidData struct {
	Cases_time_series []*model.Covid `json:"cases_time_series"`

	Statewise []Covid_1 `json:"statewise"`

	Tested []Covid_2 `json:"tested"`
}

func (r *queryResolver) Covid(ctx context.Context) ([]*model.Covid, error) {
	const endpoint = "https://data.covid19india.org/data.json"
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var routes CovidData
	json.Unmarshal(responseData, &routes)

	return routes.Cases_time_series, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
