package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

type AssetRequest struct {
	ID string `json:"id"`
}

type BlockRequest struct {
	Block uint32 `json:"block"`
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func poolsByAsset(w http.ResponseWriter, r *http.Request) {

	var assetID AssetRequest

	err := json.NewDecoder(r.Body).Decode(&assetID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("asset ID: ", assetID)
	x := fmt.Sprintf(`{
		tokens(where: {
		  id: "%s"
		  }) {
		  whitelistPools {
			id
			token0 {
			  name
			}
			token1 {
			  name
			}
			swaps (where: {
				timestamp_gt:"1622519505"
				timestamp_lt:"1622861244"
			  }){
				token0 {
				  name
				}
				token1{
				  name
				}
				timestamp
			  }
		
		  }
		}
	  }
	  `, assetID.ID)

	apiKey := goDotEnvVariable("API_KEY")
	url := fmt.Sprintf("https://gateway.thegraph.com/api/%s/subgraphs/id/0x9bde7bf4d5b13ef94373ced7c8ee0be59735a298-2", apiKey)
	graphqlClient := graphql.NewClient(url)
	graphqlRequest := graphql.NewRequest(x)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	payload, err := json.Marshal(graphqlResponse)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

}

func swapsByBlock(w http.ResponseWriter, r *http.Request) {

	var blockNum BlockRequest

	err := json.NewDecoder(r.Body).Decode(&blockNum)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("block Number: ", blockNum.Block)
	x := fmt.Sprintf(`{
		swaps(block: {number: %d}){
		  transaction{
			id
		  }
		  token0{
			name
		  }
		  token1{
			name
		  }
		}
	}`, blockNum.Block)

	apiKey := goDotEnvVariable("API_KEY")
	url := fmt.Sprintf("https://gateway.thegraph.com/api/%s/subgraphs/id/0x9bde7bf4d5b13ef94373ced7c8ee0be59735a298-2", apiKey)
	graphqlClient := graphql.NewClient(url)
	graphqlRequest := graphql.NewRequest(x)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	payload, err := json.Marshal(graphqlResponse)
	fmt.Println("hello")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Here are all the endpoints you can hit to query the uniswap subgraph")
}

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/assetPools", poolsByAsset).Methods("GET")
	myRouter.HandleFunc("/blockSwaps", swapsByBlock).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", myRouter))
}
