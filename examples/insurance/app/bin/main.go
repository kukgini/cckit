package main

import (
	"fmt"

	"github.com/op/go-logging"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/kukgini/cckit2/examples/insurance/app"
)

var logger = logging.MustGetLogger("main")

func main() {
	//logger.SetLevel(shim.LogInfo)

	err := shim.Start(new(app.SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
