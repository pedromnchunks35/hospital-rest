# Info
- This project is a basic REST API to call grpc method for the 
hyperledger fabric network

# Project details

## [How to use mocks](./docs/mocks.md)

## [Project Structure](./docs/structure.md)

## ROUTE EXAMPLE
```
http://SA-SRVAIDA77:30030/0/10
```
- Check more routes [here](./routes/chaincode.go)

## Build the image
```
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t pedrosilvamnchunks/rest-chaincode:latest --push .
```
