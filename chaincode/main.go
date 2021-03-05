package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/leyle/corpid_poc/chaincode/contract"
)

func main() {
	cc, err := contractapi.NewChaincode(&contract.Contract{})
	if err != nil {
		fmt.Println("create new chaincode failed, ", err.Error())
		return
	}

	err = cc.Start()
	if err != nil {
		fmt.Println("start chaincode failed, ", err.Error())
		return
	}

	fmt.Println("start chaincode success")
}
