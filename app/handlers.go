package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"messarichallenge/app/models"

	"github.com/machinebox/graphql"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Post API")
	}
}

func (a *App) assetPoolsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := models.AssetPoolsRequest{}
		err := parse(w, r, &req)

		assetID := &models.AssetPoolsRequest{
			ID: req.ID,
		}

		if err != nil {
			log.Printf("Cannot parse body. err %v \n", err)
			sendReponse(w, r, nil, http.StatusBadRequest)
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

		sendReponse(w, r, graphqlResponse, http.StatusOK)
	}
}

func (a *App) assetVolumeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.AssetVolumeRequest{}
		err := parse(w, r, &req)

		assetID := &models.AssetVolumeRequest{
			ID:        req.ID,
			Timestamp: req.Timestamp,
		}

		if err != nil {
			log.Printf("Cannot parse body. err %v \n", err)
			sendReponse(w, r, nil, http.StatusBadRequest)
			return
		}

		fmt.Println("asset ID: ", assetID)
		x := fmt.Sprintf(`{
			tokens (where:{
			  id: "%s"
			}){
			  whitelistPools{
				swaps (where: {
				  timestamp_gt:"%d"
				  timestamp_lt:"%d"
				}){
				  token0 {
					name
				  }
				  token1{
					name
				  }
				  amountUSD
				}
			  }
			}
		  }
	  `, assetID.ID, assetID.Timestamp.Upper, assetID.Timestamp.Lower)

		apiKey := goDotEnvVariable("API_KEY")
		url := fmt.Sprintf("https://gateway.thegraph.com/api/%s/subgraphs/id/0x9bde7bf4d5b13ef94373ced7c8ee0be59735a298-2", apiKey)
		graphqlClient := graphql.NewClient(url)
		graphqlRequest := graphql.NewRequest(x)
		var graphqlResponse interface{}
		if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
			panic(err)
		}

		sendReponse(w, r, graphqlResponse, http.StatusOK)

	}

}

func (a *App) swapsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := models.BlockRequest{}
		err := parse(w, r, &req)

		blockNum := &models.BlockRequest{
			Block: req.Block,
		}

		if err != nil {
			log.Printf("Cannot parse body. err %v \n", err)
			sendReponse(w, r, nil, http.StatusBadRequest)
			return
		}

		fmt.Println("block Number: ", blockNum.Block)
		x := fmt.Sprintf(`{
		swaps(block: {number: %d}){
		  transaction{
			id
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

		sendReponse(w, r, graphqlResponse, http.StatusOK)
	}
}

func (a *App) allAssetsSwapped() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := models.BlockRequest{}
		err := parse(w, r, &req)

		blockNum := &models.BlockRequest{
			Block: req.Block,
		}

		if err != nil {
			log.Printf("Cannot parse body. err %v \n", err)
			sendReponse(w, r, nil, http.StatusBadRequest)
			return
		}

		fmt.Println("block Number: ", blockNum.Block)
		x := fmt.Sprintf(`{
		swaps(block: {number: %d}){
			token0 {
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

		sendReponse(w, r, graphqlResponse, http.StatusOK)
	}
}
