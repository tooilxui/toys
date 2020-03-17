package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shurcooL/graphql"
)

func main() {

	client := graphql.NewClient("http://192.1.1.115:8080/v1/graphql", http.DefaultClient)

	var query struct {
		Erp_token []struct {
			Token     graphql.String `json:"token"`
			Last_used graphql.String `json:"last_used"`
		} `json:"erp_token"`
	}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(query)
}
