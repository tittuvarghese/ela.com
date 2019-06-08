package main

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const transactionID = "ELT-001"

// Request - Create Transaction
const createTransactionRequest = `{
  "id" : "ELT-001",
  "buyer_id": "ELB-001",
  "buyer_name" : "Asha",
  "seller_id" : "ELF-001",
  "seller_name" : "Elton",
  "product_id" : "ELP-001",
  "product_price": "10",
  "quantity" : 10,
  "unit" : "Nos",
  "buyer_type" : "Preserver Buy"
}`

// Response - Query Transaction
const queryTransactionResponse = `{"id":"ELT-001","buyer_type":"Preserver Buy","buyer_id":"ELB-001","buyer_name":"Asha","seller_id":"ELF-001","seller_name":"Elton","product_id":"ELP-001","product_price":"10","quantity":10,"unit":"Nos","status":false,"timestamp":"0001-01-01T00:00:00Z"}`

// Request - Update Product
const updateTransactionRequest = `{
  "status" : true
}`

// Response - Updated Product
const queryUpdatedTransactionResponse = `{"id":"ELT-001","buyer_type":"Preserver Buy","buyer_id":"ELB-001","buyer_name":"Asha","seller_id":"ELF-001","seller_name":"Elton","product_id":"ELP-001","product_price":"10","quantity":10,"unit":"Nos","status":true,"timestamp":"0001-01-01T00:00:00Z"}`

func Test_CreateTransaction(t *testing.T) {
	chaincodeToInvoke := CCName
	scc := new(elaChainCode)
	stub := shim.NewMockStub(CCName, scc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	// Mocking peer chaincode
	stub.MockPeerChaincode(chaincodeToInvoke, stub)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("Event"), []byte("1")})

	checkInvoke(t, stub, [][]byte{[]byte("createTransaction"), []byte(transactionID), []byte(createTransactionRequest)})
	checkQuery(t, stub, "queryTransaction", transactionID, queryTransactionResponse)
}

func Test_UpdateTransaction(t *testing.T) {
	chaincodeToInvoke := CCName
	scc := new(elaChainCode)
	stub := shim.NewMockStub(CCName, scc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	// Mocking peer chaincode
	stub.MockPeerChaincode(chaincodeToInvoke, stub)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("Event"), []byte("1")})

	checkInvoke(t, stub, [][]byte{[]byte("createTransaction"), []byte(transactionID), []byte(createTransactionRequest)})
	checkQuery(t, stub, "queryTransaction", transactionID, queryTransactionResponse)
	checkInvoke(t, stub, [][]byte{[]byte("updateTransaction"), []byte(transactionID), []byte(updateTransactionRequest)})
	checkQuery(t, stub, "queryTransaction", transactionID, queryUpdatedTransactionResponse)
}
