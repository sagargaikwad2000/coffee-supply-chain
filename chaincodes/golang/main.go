package main

import (
	"beanchain/contracts"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	beanChainContract := new(contracts.BeanChainContract)
	chaincode, err := contractapi.NewChaincode(
		beanChainContract,
	)
	if err != nil {
		fmt.Printf("Error creating  chaincode: %s", err.Error())
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting  chaincode: %s", err.Error())
	}

}
