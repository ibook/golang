package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Token struct {
	Owner			string	`json:"Owner"`
	TotalSupply 	uint	`json:"TotalSupply"`
	TokenName 		string	`json:"TokenName"`
	TokenSymbol 	string	`json:"TokenSymbol"`
	BalanceOf		map[string]uint	`json:"BalanceOf"`
}

func (token *Token) initialSupply(){
	token.BalanceOf[token.Owner] = token.TotalSupply;
}

func (token *Token) transfer (_from string, _to string, _value uint){
	if(token.BalanceOf[_from] >= _value){
		token.BalanceOf[_from] -= _value;
		token.BalanceOf[_to] += _value;
	}
}

func (token *Token) balance (_from string) uint{
	return token.BalanceOf[_from]
}

func (token *Token) burn(_value uint) {
	if(token.BalanceOf[token.Owner] >= _value){
		token.BalanceOf[token.Owner] -= _value;
		token.TotalSupply -= _value;
	}
}

func (token *Token) burnFrom(_from string, _value uint) {
	if(token.BalanceOf[_from] >= _value){
		token.BalanceOf[_from] -= _value;
		token.TotalSupply -= _value;
	}
}

func (token *Token) mint(_value uint) {
	
	token.BalanceOf[token.Owner] += _value;
	token.TotalSupply += _value;
	
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) initLedger(stub shim.ChaincodeStubInterface) sc.Response {
	
	token := &Token{
		Owner: "netkiller",
		TotalSupply: 10000,
		TokenName: "代币通正",
		TokenSymbol: "COIN",
		BalanceOf: map[string]uint{}}
	
	token.initialSupply()

	tokenAsBytes, _ := json.Marshal(token)
	stub.PutState("Token", tokenAsBytes)
	fmt.Println("Added", tokenAsBytes)
	
	return shim.Success(nil)
}

func (s *SmartContract) transferToken(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	tokenAsBytes, _ := stub.GetState(args[0])
	token := Token{}

	json.Unmarshal(tokenAsBytes, &token)
	token.transfer(args[1],args[2],args[3])

	tokenAsBytes, _ = json.Marshal(token)
	stub.PutState(args[0], tokenAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) balanceToken(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	tokenAsBytes, _ := stub.GetState(args[0])
	token := Token{}

	json.Unmarshal(tokenAsBytes, &token)
	amount := token.balance(args[1])

	return shim.Success(amount)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "balanceToken" {
		return s.balanceToken(stub, args)
	} else if function == "initLedger" {
		return s.initLedger(stub)
	} else if function == "transferToken" {
		return s.transferToken(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
