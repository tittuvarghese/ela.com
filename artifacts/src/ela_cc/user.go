package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// createUser - Creates a User Records and callback representing the transactionID
// @params user_id, name, email, phone_number, profile_image, role
// @params name {first_name, last_name}
func (t *elaChainCode) createUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 arguments to create user record")
	}

	var userRecords User
	var userID string

	userID = args[0]

	// Checking whether user is already created
	userRecordsCheck, err := stub.GetState(userID)

	if err != nil || userRecordsCheck != nil {
		return shim.Error("UserID is already exist in the system")
	}

	json.Unmarshal([]byte(args[1]), &userRecords)

	userAsBytes, _ := json.Marshal(userRecords)

	stub.PutState(userID, userAsBytes)

	// Transaction Response
	logger.Infof("Create User Response:%s\n", string(userAsBytes))
	return shim.Success(nil)
}

// queryUser - Query a user record using search key (userID)
// @params searchKey
func (t *elaChainCode) queryUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var searchKey string
	var jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a search key")
	}

	searchKey = args[0]

	// Get the state from the ledger
	userRecords, err := stub.GetState(searchKey)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + searchKey + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", searchKey))
	}

	if userRecords == nil {
		jsonResp = "{\"Error\":\"No data found for " + searchKey + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", searchKey))
	}

	jsonResp = "{\"Search Key\":\"" + searchKey + "\",\"Data\":\"" + string(userRecords) + "\"}"
	logger.Infof("Query Response: %s\n", jsonResp)
	return shim.Success(userRecords)
}

// updateUser - Updates User Record callback representing the transactionID
// @params searchKey, {name, password, country, phone_number, profile_image, role}, userEmail
// @params name {first_name, last_name}
func (t *elaChainCode) updateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3 arguments to update user record, user_id, data and email")
	}

	var userRecords User
	var updateUserRecords UpdateUser
	var searchKey string
	var userEmail string

	searchKey = args[0]
	userEmail = args[2]

	// Get the state from the ledger
	userasBytes, err := stub.GetState(searchKey)

	if err != nil || userasBytes == nil {
		return shim.Error(fmt.Sprintf("No data found for the key %s", searchKey))
	}

	json.Unmarshal(userasBytes, &userRecords)
	json.Unmarshal([]byte(args[1]), &updateUserRecords)

	// Updating fileds
	if userRecords.Email == userEmail {
		updateUserAsBytes, _ := json.Marshal(updateUserRecords)
		json.Unmarshal(updateUserAsBytes, &userRecords)
	} else {
		jsonResp := "{\"Error\":\"No access to the record of " + searchKey + "\"}"
		logger.Infof("Share Access Response:%s\n", jsonResp)
		return shim.Error("No access to the record")
	}

	userAsBytes, _ := json.Marshal(userRecords)

	//stub.PutState(searchKey, patientAsBytes)
	stub.PutState(searchKey, userAsBytes)

	// Transaction Response
	logger.Infof("Update User Response:%s\n", string(userAsBytes))
	return shim.Success(nil)
}
