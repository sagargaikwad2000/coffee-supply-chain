package main

import (
	"fmt"

	"github.com/hyperledger/fabric-samples/chaincode/fabcar/go/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	beanchainChaincode := new(contracts.Beanchain)
	chaincode, err := contractapi.NewChaincode(
		beanchainChaincode,
	)
	if err != nil {
		fmt.Printf("Error creating  chaincode: %s", err.Error())
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting  chaincode: %s", err.Error())
	}

}
