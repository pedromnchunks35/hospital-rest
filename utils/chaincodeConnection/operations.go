package chaincodeConnection

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type Message struct {
	Id            string `json:"id"`
	TypeForm      string `json:"type_form"`
	Description   string `json:"description"`
	Timestamp     string `json:"timestamp"`
	InsertionType string `json:"insertion_type"`
	Hash          string `json:"hash"`
}

func post(contract *client.Contract, operation string, args ...string) string {
	res, err := contract.Submit(operation, client.WithArguments(args...))
	if err != nil {
		fmt.Printf("error creating asset %v\n", err)
		_, err := contract.Evaluate(operation, client.WithArguments(args...))
		return err.Error()
	}

	return string(res)
}

func read(contract *client.Contract, operation string, args ...string) string {
	res, err := contract.Evaluate(operation, client.WithArguments(args...))
	if err != nil {
		return err.Error()
	}

	msg := &Message{}
	err = json.Unmarshal(res, msg)
	if err != nil {
		return err.Error()
	}

	indentedMsg, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		return err.Error()
	}

	return string(indentedMsg)
}
