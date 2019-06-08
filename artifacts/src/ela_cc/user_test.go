package main

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const userID = "ELU-001"
const userEmail = "user@ela.com"

// Request - Create User
const createUserRequest = `{
  "user_id" : "ELU-001",
  "name": {
    "first_name": "Adam",
    "last_name": "Smith"
  },
  "email" : "user@ela.com",
  "phone_number" : "+91-9876543210",
  "profile_image" : "profile_image_url_ipfs",
  "role" : "producer"
}`

// Response - Query User
const queryUserResponse = `{"user_id":"ELU-001","name":{"first_name":"Adam","last_name":"Smith"},"email":"user@ela.com","phone_number":"+91-9876543210","profile_image":"profile_image_url_ipfs","role":"producer"}`

// Request - Update User
const updateUserRequest = `{
  "name": {
    "first_name": "Adam",
    "last_name": "Johns"
  },
  "phone_number" : "+91-9876543211",
  "role" : "producer"
}`

// Response - Updated User
const queryUpdatedUserResponse = `{"user_id":"ELU-001","name":{"first_name":"Adam","last_name":"Johns"},"email":"user@ela.com","phone_number":"+91-9876543211","profile_image":"profile_image_url_ipfs","role":"producer"}`

func Test_CreateUser(t *testing.T) {
	chaincodeToInvoke := CCName
	scc := new(elaChainCode)
	stub := shim.NewMockStub(CCName, scc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	// Mocking peer chaincode
	stub.MockPeerChaincode(chaincodeToInvoke, stub)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("Event"), []byte("1")})

	checkInvoke(t, stub, [][]byte{[]byte("createUser"), []byte(userID), []byte(createUserRequest)})
	checkQuery(t, stub, "queryUser", userID, queryUserResponse)
}

func Test_UpdateUser(t *testing.T) {
	chaincodeToInvoke := CCName
	scc := new(elaChainCode)
	stub := shim.NewMockStub(CCName, scc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	// Mocking peer chaincode
	stub.MockPeerChaincode(chaincodeToInvoke, stub)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("Event"), []byte("1")})

	checkInvoke(t, stub, [][]byte{[]byte("createUser"), []byte(userID), []byte(createUserRequest)})
	checkQuery(t, stub, "queryUser", userID, queryUserResponse)
	checkInvoke(t, stub, [][]byte{[]byte("updateUser"), []byte(userID), []byte(updateUserRequest), []byte(userEmail)})
	checkQuery(t, stub, "queryUser", userID, queryUpdatedUserResponse)
}
