package services

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"hospital-rest/utils/chaincodeConnection"
)

type FabricServiceImpl struct{}

func (f *FabricServiceImpl) ReadInFabric(operation string, args ...string) (string, *client.Gateway) {
	return chaincodeConnection.ReadInFabric(operation, args...)
}

func (f *FabricServiceImpl) PostInFabric(operation string, args ...string) (string, *client.Gateway) {
	return chaincodeConnection.ReadInFabric(operation, args...)
}

func NewFabricServiceImpl() *FabricServiceImpl {
	return &FabricServiceImpl{}
}
