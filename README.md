# MessariChallenge

How to run code: <br>

`docker build -t messari-challenge . && docker run -p 3000:3000 -it messari-challenge`

Postman Collection: <br>

https://www.getpostman.com/collections/ee4685cbf025939989cf 


Given an asset ID:<br>
    - What pools exist that include it?
    <br>
    - What is the total volume of that asset swapped in a given time range?
    <br>
BONUS! Given a block number:<br>
    - What swaps occurred during that specific block?
    <br>
    - List all assets swapped during that specific block
    <br>

There are four endpoints available for each request mentioned above:<br>
    - /assetPools -- This endpoint expects assetId and returns the id of the pools that include the asset along with the token names 
    <br>
    - /assetVolume -- This endpoint expects assetId, and the timestamp range (lower, upper) and returns the amount of all the swaps that happened in the given time range in USD 
    <br>
    - /blockSwaps -- This endpoint expects block number and returns all the transactionId of all the swaps happened in that particular block 
    <br>
    - /allAssetsSwapped --  This endpoint expects block number and returns the name of the tokens that were swapped in that particular block
    <br>

NOTE: I have added .env file on purpose, so that it is easier to query Uniswap subgraph