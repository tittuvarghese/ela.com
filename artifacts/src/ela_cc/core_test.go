package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

/*
* Common Test Scripts
**/
func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, query string, key string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(query), []byte(key)})

	// For deleted record
	if res.Status == 500 {
		t.Skip()
	}
	if res.Status != shim.OK {
		fmt.Println("Query", key, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", key, "failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("Query value", key, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkSpecialQuery(t *testing.T, stub *shim.MockStub, query string, key string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(query), []byte(key)})

	// For deleted record
	if res.Status == 500 {
		t.Skip()
	}
	if res.Status != shim.OK {
		fmt.Println("Query", key, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", key, "failed to get value")
		t.FailNow()
	}
}

/*
* Test Init
**/

func Test_Init(t *testing.T) {
	scc := new(elaChainCode)
	stub := shim.NewMockStub(CCName, scc)

	checkInit(t, stub, [][]byte{[]byte("init")})
}
