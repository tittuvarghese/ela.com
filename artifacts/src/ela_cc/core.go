package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger(CCName)

// elaChainCode Chaincode implementation
type elaChainCode struct {
}

// Name data struct for Name
type Name struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// User data structure for storing user info.
type User struct {
	UserID       string `json:"user_id,omitempty"`
	Name         Name   `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"` // Do we need this ?
	Country      string `json:"country,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"` // Do we need this ?
	Role         string `json:"role,omitempty"`
}

// UpdateUser data structure for updating user info.
type UpdateUser struct {
	Name         Name   `json:"name,omitempty"`
	Password     string `json:"password,omitempty"` // Do we need this ?
	Country      string `json:"country,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"` // Do we need this ?
	Role         string `json:"role,omitempty"`
}

// Function to initialize SmartContract in the network
func (t *elaChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### Ela CC Init ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	fmt.Println("Initilization Completed")
	return shim.Success(nil)
}

// Function to perform Invoke Operations on the chain
func (t *elaChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	logger.Info("########### Ela Invoke - " + function + " ###########")

	if function == "createUser" {
		return t.createUser(stub, args)
		// createUser - to Create User Record
		// Creates User Record with UserID as key.
	} else if function == "queryUser" {
		return t.queryUser(stub, args)
		// queryUser - to Query User Record
		// Query queryUser Record with UserID as key.
	} else if function == "updateUser" {
		return t.updateUser(stub, args)
		// updateUser - to update user record
		// Updates user records with UserID as key and userEmail for ownership.
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'createUser', 'queryUser', 'updateUser'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'createUser', 'queryUser', 'updateUser'. But got: %v", args[0]))
}

// main function to start smart contract
func main() {
	err := shim.Start(new(elaChainCode))
	if err != nil {
		logger.Errorf("Error starting chaincode for rxmed: %s", err)
	}
}
