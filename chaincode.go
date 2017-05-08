/**
@author: Sushil Verma
@version: 1.0.0
@date: 07/04/2017
@Description: MedLab-Pharma chaincode
**/

package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Test_Suite_MedLabPharmaChaincode(t *testing.T) {
	fmt.Println("Entering into Test_Suite_MedLabPharmaChaincode test suite")

	stub := shim.NewMockStub("mock", new(MedLabPharmaChaincode))

	if stub == nil {
		fmt.Println("error")
	}

	params := []string{}
	//params2 := []string{"CON1"}

	params1 := []string{"MedLabPharma", "DHL", "Walmart", "Express Delivery", "{\"container_id\":\"CON1\",\"certified_by\":\"verizon\",\"elements\":{\"pallets\":[{\"pallet_id\":\"CON1-PAL1\",\"cases\":[{\"case_id\":\"CON1-PAL1-CASE1\",\"units\":[{\"unit_id\":\"CON1-PAL1-CASE1-UNIT1\", \"drug_id\":\"PARACETOMOL\", \"expiry_date\":\"13-APR-2020\",\"health_status\":\"healthy\",\"batch_number\":\"BBBBBB\",\"lot_number\":\"LLLLLL\",\"sale_status\":\"not sold\",\"consumer_name\":\"testuser\"},{\"unit_id\":\"CON1-PAL1-CASE1-UNIT2\"},{\"unit_id\":\"CON1-PAL1-CASE1-UNIT3\"}]},{\"case_id\":\"CON1-PAL1-CASE2\",\"units\":[{\"unit_id\":\"CON1-PAL1-CASE2-UNIT1\"},{\"unit_id\":\"CON1-PAL1-CASE2-UNIT2\"},{\"unit_id\":\"CON1-PAL1-CASE2-UNIT3\"}]},{\"case_id\":\"CON1-PAL1-CASE3\",\"units\":[{\"unit_id\":\"CON1-PAL1-CASE3-UNIT1\"},{\"unit_id\":\"CON1-PAL1-CASE3-UNIT2\"},{\"unit_id\":\"CON1-PAL1-CASE3-UNIT3\"}]}]},{\"pallet_id\":\"CON1-PAL2\",\"cases\":[{\"case_id\":\"CON1-PAL2-CASE1\",\"units\":[{\"unit_id\":\"CON1-PAL2-CASE1-UNIT1\"},{\"unit_id\":\"CON1-PAL2-CASE1-UNIT2\"},{\"unit_id\":\"CON1-PAL2-CASE1-UNIT3\"}]},{\"case_id\":\"CON1-PAL2-CASE2\",\"units\":[{\"unit_id\":\"CON1-PAL2-CASE2-UNIT1\"},{\"unit_id\":\"CON1-PAL2-CASE2-UNIT2\"},{\"unit_id\":\"CON1-PAL2-CASE2-UNIT3\"}]},{\"case_id\":\"CON1-PAL2-CASE3\",\"units\":[{\"unit_id\":\"CON1-PAL2-CASE3-UNIT1\"},{\"unit_id\":\"CON1-PAL2-CASE3-UNIT2\"},{\"unit_id\":\"CON1-PAL2-CASE3-UNIT3\"}]}]},{\"pallet_id\":\"CON1-PAL3\",\"cases\":[{\"case_id\":\"CON1-PAL3-CASE1\",\"units\":[{\"unit_id\":\"CON1-PAL3-CASE1-UNIT1\"},{\"unit_id\":\"CON1-PAL3-CASE1-UNIT2\"},{\"unit_id\":\"CON1-PAL3-CASE1-UNIT3\"}]},{\"case_id\":\"CON1-PAL3-CASE2\",\"units\":[{\"unit_id\":\"CON1-PAL3-CASE2-UNIT1\"},{\"unit_id\":\"CON1-PAL3-CASE2-UNIT2\"},{\"unit_id\":\"CON1-PAL3-CASE2-UNIT3\"}]},{\"case_id\":\"CON1-PAL3-CASE3\",\"units\":[{\"unit_id\":\"CON1-PAL3-CASE3-UNIT1\"},{\"unit_id\":\"CON1-PAL3-CASE3-UNIT2\"},{\"unit_id\":\"CON1-PAL3-CASE3-UNIT3\"}]}]}]}}"}
	checkInit(t, stub, "init", params)

	checkInvoke(t, stub, "ShipContainerUsingLogistics", params1)

//	checkQuery(t, stub, "GetContainerDetails", params2)
/*
	checkQuery(t, stub, "GetEmptyContainer", params)
	checkQuery(t, stub, "GetCurrentOwner", []string{"MedLabPharma"})
	checkQuery(t, stub, "GetMaxIDValue", []string{})
*/
	params11 := []string{"MedLabPharma", "DHL", "Walmart", "Express Delivery", "{\"container_id\":\"CON2\",\"certified_by\":\"verizon\",\"elements\":{\"pallets\":[{\"pallet_id\":\"CON1-PAL1\",\"cases\":[{\"case_id\":\"CON1-PAL1-CASE1\",\"units\":[{\"unit_id\":\"CON1-PAL1-CASE1-UNIT1\", \"drug_id\":\"PARACETOMOL\", \"expiry_date\":\"13-APR-2020\",\"health_status\":\"healthy\",\"batch_number\":\"BBBBBB\",\"lot_number\":\"LLLLLL\",\"sale_status\":\"not sold\",\"consumer_name\":\"testuser\"},{\"unit_id\":\"CON1-PAL1-CASE1-UNIT2\"},{\"unit_id\":\"CON1-PAL1-CASE1-UNIT3\"}]},{\"case_id\":\"CON1-PAL1-CASE2\",\"units\":[{\"unit_id\":\"CON1-PAL1-CASE2-UNIT1\"},{\"unit_id\":\"CON1-PAL1-CASE2-UNIT2\"},{\"unit_id\":\"CON1-PAL1-CASE2-UNIT3\"}]},{\"case_id\":\"CON1-PAL1-CASE3\",\"units\":[{\"unit_id\":\"CON1-PAL1-CASE3-UNIT1\"},{\"unit_id\":\"CON1-PAL1-CASE3-UNIT2\"},{\"unit_id\":\"CON1-PAL1-CASE3-UNIT3\"}]}]},{\"pallet_id\":\"CON1-PAL2\",\"cases\":[{\"case_id\":\"CON1-PAL2-CASE1\",\"units\":[{\"unit_id\":\"CON1-PAL2-CASE1-UNIT1\"},{\"unit_id\":\"CON1-PAL2-CASE1-UNIT2\"},{\"unit_id\":\"CON1-PAL2-CASE1-UNIT3\"}]},{\"case_id\":\"CON1-PAL2-CASE2\",\"units\":[{\"unit_id\":\"CON1-PAL2-CASE2-UNIT1\"},{\"unit_id\":\"CON1-PAL2-CASE2-UNIT2\"},{\"unit_id\":\"CON1-PAL2-CASE2-UNIT3\"}]},{\"case_id\":\"CON1-PAL2-CASE3\",\"units\":[{\"unit_id\":\"CON1-PAL2-CASE3-UNIT1\"},{\"unit_id\":\"CON1-PAL2-CASE3-UNIT2\"},{\"unit_id\":\"CON1-PAL2-CASE3-UNIT3\"}]}]},{\"pallet_id\":\"CON1-PAL3\",\"cases\":[{\"case_id\":\"CON1-PAL3-CASE1\",\"units\":[{\"unit_id\":\"CON1-PAL3-CASE1-UNIT1\"},{\"unit_id\":\"CON1-PAL3-CASE1-UNIT2\"},{\"unit_id\":\"CON1-PAL3-CASE1-UNIT3\"}]},{\"case_id\":\"CON1-PAL3-CASE2\",\"units\":[{\"unit_id\":\"CON1-PAL3-CASE2-UNIT1\"},{\"unit_id\":\"CON1-PAL3-CASE2-UNIT2\"},{\"unit_id\":\"CON1-PAL3-CASE2-UNIT3\"}]},{\"case_id\":\"CON1-PAL3-CASE3\",\"units\":[{\"unit_id\":\"CON1-PAL3-CASE3-UNIT1\"},{\"unit_id\":\"CON1-PAL3-CASE3-UNIT2\"},{\"unit_id\":\"CON1-PAL3-CASE3-UNIT3\"}]}]}]}}"}
	params12 := []string{"CON2","MedLabPharma", "DHL", "Walmart", "Express Delivery"}
	params13 := []string{"MedLabPharma"}
	checkInvoke(t, stub, "ShipContainerUsingLogistics", params11)
	checkInvoke(t,stub,"AcceptContainerbyLogistics", params12)
	checkQuery(t,stub,"GetContainerDetailsForOwner", params13)
	checkQuery(t,stub,"GetOwner", params13)
	checkQuery(t, stub, "GetCurrentOwner", []string{"MedLabPharma"})
	//checkQuery(t, stub, "GetMaxIDValue", []string{})

	/*
		params4 := []string{"UserManufacturer", "CON001"}
		checkInvoke(t, stub, "SetCurrentOwner", params4)

		params5 := []string{"UserManufacturer"}
		checkQuery(t, stub, "GetCurrentOwner", params5)

		checkInvoke(t, stub, "SetCurrentOwner", params4)
		checkQuery(t, stub, "GetCurrentOwner", params5)

		checkInvoke(t, stub, "SetCurrentOwner", []string{"UserManufacturer", "CON002"})
		checkInvoke(t, stub, "SetCurrentOwner", []string{"UserManufacturer", "CON003"})
		checkInvoke(t, stub, "SetCurrentOwner", []string{"UserManufacturer", "CON004"})
		checkInvoke(t, stub, "SetCurrentOwner", params4)

		checkInvoke(t, stub, "SetCurrentOwner", []string{"UserDistributor", "CON012"})
		checkInvoke(t, stub, "SetCurrentOwner", []string{"UserDistributor", "CON013"})
		checkInvoke(t, stub, "SetCurrentOwner", []string{"UserLogistics", "CON014"})

		checkQuery(t, stub, "GetCurrentOwner", []string{"UserManufacturer"})
		checkQuery(t, stub, "GetCurrentOwner", []string{"UserDistributor"})
		checkQuery(t, stub, "GetCurrentOwner", []string{"UserLogistics"})
	*/

	fmt.Println("Exiting from Test_Suite_MedLabPharmaChaincode test suite")
}

func checkInit(t *testing.T, stub *shim.MockStub, function string, args []string) {
	_, err := stub.MockInit("1", function, args)
	if err != nil {
		fmt.Println("Init failed", err)
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, function string, args []string) {

	_, err := stub.MockInvoke("1", function, args)
	if err != nil {
		fmt.Println("Invoke", args, "failed", err)
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, function string, args []string) {
	fmt.Println("enter checkquery")
	querybytes, err := stub.MockQuery(function, args)
	fmt.Println("finishing checkquery")

	if err != nil {
		fmt.Println("Query", function, "failed", err)
		t.FailNow()
	}
	if querybytes == nil {
		fmt.Println("Query", function, "failed to get value")
		t.FailNow()
	}
	str := string(querybytes)
	fmt.Println(str)
}
