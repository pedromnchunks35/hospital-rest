package chaincodeConnection

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
	"os"
)

func getContract() *client.Contract {
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
	defer gateway.Close()

	channel := gateway.GetNetwork("channel1")

	return channel.GetContract("basic")
}

func ReadInFabric(operation string, args ...string) string {
	contract := getContract()

	return read(contract, operation, args...)
}

func PostInFabric(operation string, args ...string) string {
	contract := getContract()

	return post(contract, operation, args...)
}
