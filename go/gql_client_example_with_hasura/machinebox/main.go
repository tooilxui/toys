package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/machinebox/graphql"
)

func main() {
	client := graphql.NewClient("http://192.1.1.115:8080/v1/graphql")

	client.Log = func(s string) {
		fmt.Printf("[time]: %v, [message]: %v\n", time.Now().Format("2006-01-02 15:04:05"), s)
	}

	var res result
	ctx := context.Background()

	// 1. query with parameter
	// - query string in hasura :
	// query ($spec: jsonb ){
	//  	erp_json_test(where: {data: {
	//		_contains: $spec
	//  	}}) {
	//		data
	// 	}
	// }
	// - query variable in hasura:
	// {
	//	"spec": {
	//    	"B": {
	//        "C": "ccc",
	//        "D": "ddd"
	//      }
	//	 }
	// }
	// query result in hasura :
	// {
	//   "data": {
	//     "erp_json_test": [
	//       {
	//         "data": {
	//           "A": "aaa",
	//           "B": {
	//             "C": "ccc",
	//             "D": "ddd"
	//           }
	//         }
	//       }
	//     ]
	//   }
	// }
	type variable struct {
		A string            `json:",omitempty"`
		B map[string]string `json:",omitempty"`
	}
	req := graphql.NewRequest("query ($spec: jsonb) { erp_json_test(where: {data: {_contains: $spec}}) { data } }")
	req.Var("spec", variable{B: map[string]string{"C": "ccc", "D": "ddd"}})
	req.Header.Add("content-type", "application/json")

	if err := client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("I'm Result 1: %s\n", res)
}

type result struct {
	Data []erp_json_test `json:"erp_json_test"`
}

type erp_json_test struct {
	Data struct {
		A string `json:"A"`
		B struct {
			C string `json:"C"`
			D string `json:"D"`
		} `json:"B"`
	} `json:"data"`
}
