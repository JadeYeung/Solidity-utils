/**
 * @Author: yangsuiyu
 * @Date: 2020/4/4 11:49
 * solidity-utils
 */

package abiutils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestSolidityDataType(t *testing.T) {
	fmt.Println(GetSolidityTypeByName("int"))        //256
	fmt.Println(GetSolidityTypeByName("uint"))       //256
	fmt.Println(GetSolidityTypeByName("string"))     //0
	fmt.Println(GetSolidityTypeByName("address"))    //160
	fmt.Println(GetSolidityTypeByName("bytes6"))     //48
	fmt.Println(GetSolidityTypeByName("bytes6[3]"))  //144
	fmt.Println(GetSolidityTypeByName("int8[10]"))   //80
	fmt.Println(GetSolidityTypeByName("uint32[10]")) //320

}

func TestTransformTxData(t *testing.T) {
	const abi = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"genesisId\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"brief\",\"type\":\"string\"},{\"name\":\"detail\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"string\"},{\"name\":\"associationId\",\"type\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"updateActivity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voluntaryId\",\"type\":\"uint256\"}],\"name\":\"getVoluntaryById\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"activityId\",\"type\":\"uint256\"},{\"name\":\"userAddr\",\"type\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"previousItemId\",\"type\":\"uint256\"},{\"name\":\"vhours\",\"type\":\"uint256\"},{\"name\":\"generation\",\"type\":\"uint256\"},{\"name\":\"enable\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"adminAddr\",\"type\":\"address\"}],\"name\":\"getAdminIdByAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicityId\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"brief\",\"type\":\"string\"},{\"name\":\"detail\",\"type\":\"string\"},{\"name\":\"creator\",\"type\":\"address\"},{\"name\":\"publicityType\",\"type\":\"uint256\"}],\"name\":\"postPublicity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"associationAddr\",\"type\":\"address\"}],\"name\":\"getEnrollmentByAssociation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"activityId\",\"type\":\"uint256\"}],\"name\":\"getActivityItemById\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"brief\",\"type\":\"string\"},{\"name\":\"detail\",\"type\":\"string\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"previousItemId\",\"type\":\"uint256\"},{\"name\":\"genesisId\",\"type\":\"uint256\"},{\"name\":\"generation\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"string\"},{\"name\":\"creator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"uint256\"}],\"name\":\"disableUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllPublicity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"adminAddr\",\"type\":\"address\"}],\"name\":\"adminRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"userRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicityId\",\"type\":\"uint256\"}],\"name\":\"getPublicityById\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"brief\",\"type\":\"string\"},{\"name\":\"detail\",\"type\":\"string\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getEnrollmentByUser\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"associationAddr\",\"type\":\"address\"}],\"name\":\"getActivityByAssociationAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"genesisId\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"brief\",\"type\":\"string\"},{\"name\":\"detail\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"string\"},{\"name\":\"associationId\",\"type\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"createActivity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enrollmentId\",\"type\":\"uint256\"},{\"name\":\"isSuccessful\",\"type\":\"bool\"}],\"name\":\"updateEnrollment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAddressById\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"associationId\",\"type\":\"uint256\"}],\"name\":\"disableAssociation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"revokeVoluntary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"associationAddr\",\"type\":\"address\"},{\"name\":\"userAddr\",\"type\":\"address\"},{\"name\":\"activityId\",\"type\":\"uint256\"},{\"name\":\"vhours\",\"type\":\"uint256\"}],\"name\":\"issueVhours\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"traceVoluntaryRecord\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enrollmentId\",\"type\":\"uint256\"}],\"name\":\"getEnrollmentById\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"activityId\",\"type\":\"uint256\"},{\"name\":\"userAddr\",\"type\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"updateTime\",\"type\":\"uint256\"},{\"name\":\"isSuccessful\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"genesisId\",\"type\":\"uint256\"}],\"name\":\"getVoluntaryByActivityId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"activityId\",\"type\":\"uint256\"}],\"name\":\"getEnrollmentByActivity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"associationAddr\",\"type\":\"address\"}],\"name\":\"getAssociationIdByAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"activityId\",\"type\":\"uint256\"},{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"userEnroll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"associationAddr\",\"type\":\"address\"}],\"name\":\"associationRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllActivity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getVhours\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getUserIdByAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"genesisId\",\"type\":\"uint256\"}],\"name\":\"traceActivity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		t.Errorf("fail to connent eth client:%v", err.Error())
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Errorf("fail to get the latest block:%v", err.Error())
	}
	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		t.Errorf("fail to get a block by the number of block:%v", err.Error())
	}
	tx, err := client.TransactionInBlock(context.Background(), block.Hash(), 0)
	if err != nil {
		t.Errorf("fail to get transactions of block,%v", err.Error())
	}

	data, err := TransformTxData(abi, tx.Data())
	if err != nil {
		t.Errorf("fail to transform the input: %v", err.Error())
	} else {
		jsonByte, err := json.Marshal(data)
		if err != nil {
			t.Errorf("fail to marshal to json: %v", err.Error())
		}
		t.Log(string(jsonByte))
	}

}
