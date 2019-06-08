package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// createTransaction - Creates a Transaction Record and callback representing the transactionID
// @params id, buyer_id, buyer_name, seller_id, seller_name, product_id, product_price, quantity, unit, buyer_type
func (t *elaChainCode) createTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 arguments to create product record")
	}

	var transactionRecords Buy
	var transactionID string

	transactionID = args[0]

	// Checking whether product is already created
	transactionRecordsCheck, err := stub.GetState(transactionID)

	if err != nil || transactionRecordsCheck != nil {
		return shim.Error("TransactionID is already exist in the system")
	}

	json.Unmarshal([]byte(args[1]), &transactionRecords)
	transactionRecords.Status = false

	transactionAsBytes, _ := json.Marshal(transactionRecords)

	stub.PutState(transactionID, transactionAsBytes)

	// Transaction Response
	logger.Infof("Create Transaction Response:%s\n", string(transactionAsBytes))
	return shim.Success(nil)
}

// queryProduct - Query a product record using search key (productID)
// @params searchKey
func (t *elaChainCode) queryTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var searchKey string
	var jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a search key")
	}

	searchKey = args[0]

	// Get the state from the ledger
	transactionRecords, err := stub.GetState(searchKey)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + searchKey + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", searchKey))
	}

	if transactionRecords == nil {
		jsonResp = "{\"Error\":\"No data found for " + searchKey + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", searchKey))
	}

	jsonResp = "{\"Search Key\":\"" + searchKey + "\",\"Data\":\"" + string(transactionRecords) + "\"}"
	logger.Infof("Query Response: %s\n", jsonResp)
	return shim.Success(transactionRecords)
}

// updateTransaction - Updates Product Record callback representing the transactionID
// @params searchKey, {status}
func (t *elaChainCode) updateTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 arguments to update product record, id, data")
	}

	var transactionRecords Buy
	var updateTransactionRecords UpdatePreserverBuy
	var searchKey string

	searchKey = args[0]

	// Get the state from the ledger
	transactionasBytes, err := stub.GetState(searchKey)

	if err != nil || transactionasBytes == nil {
		return shim.Error(fmt.Sprintf("No data found for the key %s", searchKey))
	}

	json.Unmarshal(transactionasBytes, &transactionRecords)
	json.Unmarshal([]byte(args[1]), &updateTransactionRecords)

	// Updating fileds
	// if userRecords.Email == userEmail {

	updateTransactionAsBytes, _ := json.Marshal(updateTransactionRecords)
	json.Unmarshal(updateTransactionAsBytes, &transactionRecords)

	// } else {
	// 	jsonResp := "{\"Error\":\"No access to the record of " + searchKey + "\"}"
	// 	logger.Infof("Share Access Response:%s\n", jsonResp)
	// 	return shim.Error("No access to the record")
	// }

	transactionAsBytes, _ := json.Marshal(transactionRecords)

	//stub.PutState(searchKey, patientAsBytes)
	stub.PutState(searchKey, transactionAsBytes)

	// Transaction Response
	logger.Infof("Update Transaction Response:%s\n", string(transactionAsBytes))
	return shim.Success(nil)
}

// Get History of a transaction by passing Key
func (s *elaChainCode) getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a search key")
	}

	searchKey := args[0]
	fmt.Printf("##### start History of Record: %s\n", searchKey)

	resultsIterator, err := stub.GetHistoryForKey(searchKey)
	if err != nil {
		fmt.Println(err)
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForProductID returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
