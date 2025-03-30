package chaincodeConnection

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
	"os"
)

func getContract() (*client.Contract, *client.Gateway) {
	tlsCertPath := os.Getenv("TLS_CERT_PATH")
	peerAddress := os.Getenv("PEER_TO_CONNECT")
	clientCertPath := os.Getenv("HLF_CLIENT_CERT_PATH")
	mspId := os.Getenv("CLIENT_MSP")
	keystorePath := os.Getenv("CLIENT_KEYSTORE")

	gateway, err := createConnection(
		tlsCertPath,
		peerAddress,
		clientCertPath,
		mspId,
		keystorePath,
	)

	if err != nil {
		log.Fatalf("%v", err)
	}

	if gateway == nil {
		log.Fatal("Failed to create gateway connection")
	}

	channel := gateway.GetNetwork("channel1")

	return channel.GetContract("form"), gateway
}

func ReadInFabric(operation string, args ...string) (string, *client.Gateway) {
	contract, gateway := getContract()

	return read(contract, operation, args...), gateway
}

func PostInFabric(operation string, args ...string) (string, *client.Gateway) {
	contract, gateway := getContract()
	defer gateway.Close()

	return post(contract, operation, args...), gateway
}
