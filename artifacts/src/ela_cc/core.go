package main

import (
	"fmt"
	"time"

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
	PhoneNumber  string `json:"phone_number,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"` // Do we need this ?
	Role         string `json:"role,omitempty"`
}

// UpdateUser data structure for updating user info.
type UpdateUser struct {
	Name         Name   `json:"name,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"` // Do we need this ?
	Role         string `json:"role,omitempty"`
}

//ProductData data structure for keeping product data
type ProductData struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	Unit          string `json:"unit,omitempty"`
	Price         string `json:"price,omitempty"`
	DateOfHarvest string `json:"date_of_harvest,omitempty"`
	WeightLeft    int    `json:"weight_left,omitempty"`
	FarmerID      string `json:"farmer_id,omitempty"`
}

//UpdateProductData
type UpdateProductData struct {
	Name          string `json:"name,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	Unit          string `json:"unit,omitempty"`
	Price         string `json:"price,omitempty"`
	DateOfHarvest string `json:"date_of_harvest,omitempty"`
}

//PreserverBuy data structure for keeping buy transaction data
type Buy struct {
	ID           string    `json:"id,omitempty"`
	BuyerType    string    `json:"buyer_type,omitempty"`
	BuyerID      string    `json:"buyer_id,omitempty"`
	BuyerName    string    `json:"buyer_name,omitempty"`
	SellerID     string    `json:"seller_id,omitempty"`
	SellerName   string    `json:"seller_name,omitempty"`
	ProductID    string    `json:"product_id,omitempty"`
	ProductPrice string    `json:"product_price,omitempty"`
	Quantity     int       `json:"quantity,omitempty"`
	Unit         string    `json:"unit,omitempty"`
	Status       bool      `json:"status"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

//PreserverBuy data structure for keeping buy transaction data
type UpdatePreserverBuy struct {
	Status bool `json:"status,omitempty"`
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
	} else if function == "createProduct" {
		return t.createProduct(stub, args)
		// createProduct - to Create Product Record
		// Creates product Record with ProductID as key.
	} else if function == "queryProduct" {
		return t.queryProduct(stub, args)
		// queryProduct - to Query Product Record
		// Query queryProduct Record with ProductID as key.
	} else if function == "updateProduct" {
		return t.updateProduct(stub, args)
		// updateProduct - to update product record
		// Updates user records with ProductID as key.
	} else if function == "createTransaction" {
		return t.createTransaction(stub, args)
		// createTransaction - to Create Transaction Record
		// Creates transaction Record with TransactionID as key.
	} else if function == "queryTransaction" {
		return t.queryTransaction(stub, args)
		// queryTransaction - to Query Transaction Record
		// Query queryTransaction Record with TransactionID as key.
	} else if function == "updateTransaction" {
		return t.updateTransaction(stub, args)
		// updateTransaction - to update Transaction record
		// Updates transaction records with TransactionID as key.
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
