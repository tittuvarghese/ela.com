package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// createProduct - Creates a Product Record and callback representing the transactionID
// @params id, name, quantity, unit, price, date_of_harvest, farmer_id
func (t *elaChainCode) createProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 arguments to create product record")
	}

	var productRecords ProductData
	var productID string

	productID = args[0]

	// Checking whether product is already created
	productRecordsCheck, err := stub.GetState(productID)

	if err != nil || productRecordsCheck != nil {
		return shim.Error("ProductID is already exist in the system")
	}

	json.Unmarshal([]byte(args[1]), &productRecords)
	productRecords.WeightLeft = productRecords.Quantity

	productAsBytes, _ := json.Marshal(productRecords)

	stub.PutState(productID, productAsBytes)

	// Transaction Response
	logger.Infof("Create Product Response:%s\n", string(productAsBytes))
	return shim.Success(nil)
}

// queryProduct - Query a product record using search key (productID)
// @params searchKey
func (t *elaChainCode) queryProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var searchKey string
	var jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a search key")
	}

	searchKey = args[0]

	// Get the state from the ledger
	productRecords, err := stub.GetState(searchKey)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + searchKey + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", searchKey))
	}

	if productRecords == nil {
		jsonResp = "{\"Error\":\"No data found for " + searchKey + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", searchKey))
	}

	jsonResp = "{\"Search Key\":\"" + searchKey + "\",\"Data\":\"" + string(productRecords) + "\"}"
	logger.Infof("Query Response: %s\n", jsonResp)
	return shim.Success(productRecords)
}

// updateProduct - Updates Product Record callback representing the transactionID
// @params searchKey, {name, quantity, unit, price, date_of_harvest,weight_left, farmer_id}
func (t *elaChainCode) updateProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 arguments to update product record, id, data")
	}

	var productRecords ProductData
	var updateProductRecords UpdateProductData
	var searchKey string
	var currentQuantity int

	searchKey = args[0]

	// Get the state from the ledger
	productasBytes, err := stub.GetState(searchKey)

	if err != nil || productasBytes == nil {
		return shim.Error(fmt.Sprintf("No data found for the key %s", searchKey))
	}

	json.Unmarshal(productasBytes, &productRecords)
	json.Unmarshal([]byte(args[1]), &updateProductRecords)

	// Updating fileds
	// if userRecords.Email == userEmail {
	currentQuantity = productRecords.Quantity

	updateProductAsBytes, _ := json.Marshal(updateProductRecords)
	json.Unmarshal(updateProductAsBytes, &productRecords)

	productRecords.Quantity += currentQuantity
	productRecords.WeightLeft += updateProductRecords.Quantity

	fmt.Println(productRecords.Quantity)
	fmt.Println(productRecords.WeightLeft)

	// } else {
	// 	jsonResp := "{\"Error\":\"No access to the record of " + searchKey + "\"}"
	// 	logger.Infof("Share Access Response:%s\n", jsonResp)
	// 	return shim.Error("No access to the record")
	// }

	productAsBytes, _ := json.Marshal(productRecords)

	//stub.PutState(searchKey, patientAsBytes)
	stub.PutState(searchKey, productAsBytes)

	// Transaction Response
	logger.Infof("Update Product Response:%s\n", string(productAsBytes))
	return shim.Success(nil)
}
