package main

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const productID = "ELP-001"

// Request - Create User
const createProductRequest = `{
  "id" : "ELP-001",
  "name": "Tomato",
  "quantity" : 10,
  "unit" : "Nos",
  "price" : "10",
  "date_of_harvest" : "27-05-2019",
  "farmer_id": "ELU-001"
}`

// Response - Query User
const queryProductResponse = `{"id":"ELP-001","name":"Tomato","quantity":10,"unit":"Nos","price":"10","date_of_harvest":"27-05-2019","weight_left":10,"farmer_id":"ELU-001"}`

// Request - Update Product
const updateProductRequest = `{
  "quantity" : 15
}`

// Response - Updated Product
const queryUpdatedProductResponse = `{"id":"ELP-001","name":"Tomato","quantity":25,"unit":"Nos","price":"10","date_of_harvest":"27-05-2019","weight_left":25,"farmer_id":"ELU-001"}`

func Test_CreateProduct(t *testing.T) {
	chaincodeToInvoke := CCName
	scc := new(elaChainCode)
	stub := shim.NewMockStub(CCName, scc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	// Mocking peer chaincode
	stub.MockPeerChaincode(chaincodeToInvoke, stub)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("Event"), []byte("1")})

	checkInvoke(t, stub, [][]byte{[]byte("createProduct"), []byte(productID), []byte(createProductRequest)})
	checkQuery(t, stub, "queryProduct", productID, queryProductResponse)
}

func Test_UpdateProduct(t *testing.T) {
	chaincodeToInvoke := CCName
	scc := new(elaChainCode)
	stub := shim.NewMockStub(CCName, scc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	// Mocking peer chaincode
	stub.MockPeerChaincode(chaincodeToInvoke, stub)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("Event"), []byte("1")})

	checkInvoke(t, stub, [][]byte{[]byte("createProduct"), []byte(productID), []byte(createProductRequest)})
	checkQuery(t, stub, "queryProduct", productID, queryProductResponse)
	checkInvoke(t, stub, [][]byte{[]byte("updateProduct"), []byte(productID), []byte(updateProductRequest)})
	checkQuery(t, stub, "queryProduct", productID, queryUpdatedProductResponse)
}
