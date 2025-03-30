package services

import "github.com/hyperledger/fabric-gateway/pkg/client"

type FabricService interface {
	ReadInFabric(operation string, args ...string) (string, *client.Gateway)
	PostInFabric(operation string, args ...string) (string, *client.Gateway)
}
