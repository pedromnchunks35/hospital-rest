package chaincodeConnection

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func post(contract *client.Contract, operation string, arg ...string) string {
	res, err := contract.SubmitTransaction(operation, arg...)
	if err != nil {
		fmt.Printf("error creating asset %v\n", err)
		_, err := contract.EvaluateTransaction(operation, arg...)
		return err.Error()
	}

	return string(res)
}

func read(contract *client.Contract, operation string, args ...string) string {
	res, err := contract.Evaluate(operation, client.WithArguments(args...))
	if err != nil {
		return err.Error()
	}
	return string(res)
}
