package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/kukgini/cckit2/examples/cars"
)

func main() {
	cc := cars.NewWithoutAccessControl()
	if err := shim.Start(cc); err != nil {
		fmt.Printf("Error starting Cars chaincode: %s", err)
	}
}
