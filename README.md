# MessariChallenge

How to run code: 

`docker build -t messari-challenge . && docker run -p 3000:3000 -it messari-challenge`

Postman Collection:

https://www.getpostman.com/collections/ee4685cbf025939989cf 


Given an asset ID:
    - What pools exist that include it?
    - What is the total volume of that asset swapped in a given time range?
BONUS! Given a block number:
    - What swaps occurred during that specific block?
    - List all assets swapped during that specific block

There are four endpoints available for each request mentioned above:
    - /assetPools -- This endpoint expects assetId and returns the id of the pools that include the asset along with the token names 
    - /assetVolume -- This endpoint expects assetId, and the timestamp range (lower, upper) and returns the amount of all the swaps that happened in the given time range in USD 
    - /blockSwaps -- This endpoint expects block number and returns all the transactionId of all the swaps happened in that particular block 
    - /allAssetsSwapped --  This endpoint expects block number and returns the name of the tokens that were swapped in that particular block

NOTE: I have added .env file on purpose, so that it is easier to query Uniswap subgraph