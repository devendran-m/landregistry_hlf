package main

import (
	"fmt"
	"landregistry-application-chaincode/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	landregistry := new(contracts.LandRegistry)
	//chaincode, err := contractapi.NewChaincode(new(SmartContract))
	chaincode, err := contractapi.NewChaincode(landregistry)
	if err != nil {
		fmt.Printf("Error create landreg chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting landreg chaincode: %s", err.Error())
	}
}
