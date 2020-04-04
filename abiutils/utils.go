/**
 * @Author: yangsuiyu
 * @Date: 2020/4/4 11:16
 * solidity-utils
 */

package abiutils

import (
	"github.com/astaxie/beego/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cast"
	"math/big"
	"strconv"
	"strings"
)

// 用于描述Solidity数据类型的结构体
type SolidityDataType struct {
	Name     string `json:"name"`     // 类型名，如：int32
	Type     Kind   `json:"type"`     // 值类型或引用类型
	IsNumber bool   `json:"isNumber"` // 是否数字
	IsArray  bool   `json:"isArray"`  // 是否数组
	BitSize  int    `json:"bitSize"`  // 以bit为单位的长度，用于以后补0，仅对值类型生效
}

type SolidityDataTypes []*SolidityDataType

// 用于封装Transaction中Data字节数组的结构体
type InputData struct {
	Param     SolidityDataType `json:"-"`         //参数的数据类型
	ParamName string           `json:"paramName"` //参数名
	Offset    int64            `json:"-"`         //偏移量，非值类型时为0
	Data      string           `json:"data"`      //数据内容
}

type InputDatas []InputData

// 用于描述Solidity函数的结构体
type SolidityMethod struct {
	Name       string     `json:"name"` // 方法标识，如：function foo(uint id, string name)
	InputDatas InputDatas `json:"params"`
}

// 数据类型的类型
type Kind int

const (
	VALUE    Kind = iota //值类型
	VARIABLE             //动态类型
)

var kindTypes = []string{
	VALUE:    "value",
	VARIABLE: "variable",
}

const (
	BOOL    = "bool"
	INT     = "int"
	UINT    = "uint"
	ADDRESS = "address"
	BYTES   = "bytes"
	STRING  = "string"
)

// 通过参数类型获得相关信息
//name = "uint"	=>	returns { uint, value, 256 }		// uint类型，值类型，256位
//name = "string"	=>	returns { string, variable, 0 }		// string类型，非值类型，无意义
//name = "byte3[10]"	=>	returns { bytes3[10], value, 0}	// bytes3数组类型，值类型，3*8*10=240位
func GetSolidityTypeByName(name string) *SolidityDataType {
	var bitSize = 0
	var typeName = VALUE
	var isNumber = false
	var isArray = false
	if name == BOOL { // bool
		bitSize = 8
	} else if name == ADDRESS {
		bitSize = 20 * 8
	} else if strings.Contains(name, UINT) {
		isNumber = true
		if len(name) == len(UINT) { // uint
			bitSize = 256
		} else if strings.Contains(name, "[") { //uint<M>[N]
			isArray = true
			base := 256
			left := strings.Index(name, "[")
			right := strings.Index(name, "]")
			length := cast.ToInt(name[left+1 : right]) // N
			if left > len(UINT) {                      //uint<M>[N]
				base = cast.ToInt(name[len(UINT):left]) // M
			}
			bitSize = base * length //M == 0 , uint[N]
		} else {
			bitSize = cast.ToInt(strings.Trim(name, UINT))
		}
	} else if strings.Contains(name, INT) {
		isNumber = true
		if len(name) == len(INT) { //int
			bitSize = 256
		} else if strings.Contains(name, "[") { //int[]
			isArray = true
			base := 256
			left := strings.Index(name, "[")
			right := strings.Index(name, "]")
			length := cast.ToInt(name[left+1 : right]) // N
			if left > len(INT) {                       //int<M>
				base = cast.ToInt(name[len(INT):left]) // M
			}
			bitSize = base * length //M == 0 , int[N]
		} else {
			bitSize = cast.ToInt(strings.Trim(name, INT))
		}
	} else if strings.Contains(name, BYTES) {
		if len(name) == len(BYTES) { //bytes非值类型
			typeName = VARIABLE
		} else if strings.Contains(name, "[") { //bytes[]
			isArray = true
			base := 1
			left := strings.Index(name, "[")
			right := strings.Index(name, "]")
			length := cast.ToInt(name[left+1 : right]) // N
			if left > len(BYTES) {                     //bytes<M>
				base = cast.ToInt(name[len(BYTES):left]) // M
			}
			bitSize = 8 * base * length //M == 0 , bytes[N]
		} else {
			bitSize = 8 * cast.ToInt(strings.Trim(name, BYTES))
		}
	} else if name == STRING {
		typeName = VARIABLE
	} else {
		bitSize = 0
		typeName = VALUE
		name = ""
	}
	return &SolidityDataType{name, typeName, isNumber, isArray, bitSize}
}

// 十六进制转换为十进制
func DecodeBigFromHex(hexStr string) (*big.Int, error) {
	hexStr = TrimLeadingZero(hexStr)
	return hexutil.DecodeBig(hexStr)
}

// 删除十六进制字符串前缀多余的0
func TrimLeadingZero(hexStr string) string {
	if strings.Contains(hexStr, "0x") {
		hexStr = strings.TrimPrefix(hexStr, "0x")
	}
	for {
		if !strings.HasPrefix(hexStr, "0") { //不再拥有“0”前缀时跳出
			break
		}
		hexStr = strings.TrimPrefix(hexStr, "0")
	}
	if hexStr == "" {
		return "0x" + "0"
	}
	return "0x" + hexStr
}

// 转换transaction中的input
func TransformTxData(abiStr string, input []byte) (*SolidityMethod, error) {
	abi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, err
	}
	method, err := abi.MethodById(input)
	if err != nil {
		return nil, err
	}
	inputDatas, err := unmarshalInput(method.Inputs, input)
	if err != nil {
		return nil, err
	}
	return &SolidityMethod{
		Name:       method.String(),
		InputDatas: inputDatas,
	}, nil
}

// 把inputs转换为InputDatas
func unmarshalInput(args abi.Arguments, input []byte) (InputDatas, error) {
	const head int64 = 4
	const bitSize int64 = 32
	var next = head

	var resultList InputDatas
	var inputData InputData
	for _, arg := range args {
		dataType := GetSolidityTypeByName(arg.Type.String())
		inputData.Param = *dataType
		inputData.ParamName = arg.Name
		inputData.Offset = 0
		if dataType.Type == VALUE { //值类型
			inputData.Data = hexutil.Encode(input[next : next+bitSize])
			if dataType.IsNumber {
				dataInt, err := DecodeBigFromHex(inputData.Data)
				if err != nil {
					return nil, err
				}
				inputData.Data = config.ToString(dataInt)
			}
		} else { //动态类型

			offsetHex := hexutil.Encode(input[next : next+bitSize])
			offsetInt, err := DecodeBigFromHex(offsetHex)
			if err != nil {
				return nil, err
			}
			inputData.Offset = offsetInt.Int64()

			dataHead := head + offsetInt.Int64()
			dataNext := dataHead + bitSize
			dataLenHex := hexutil.Encode(input[dataHead:dataNext])
			dataLen, err := DecodeBigFromHex(dataLenHex)
			if err != nil {
				return nil, err
			}
			dataHead = dataNext
			dataNext = dataHead + dataLen.Int64()
			inputData.Data = string(input[dataHead:dataNext])
		}
		next += bitSize
		resultList = append(resultList, inputData)
	}
	return resultList, nil
}

func (k Kind) String() string {
	if int(k) < len(kindTypes) {
		return kindTypes[k]
	}
	return "kind" + strconv.Itoa(int(k))
}

func (t SolidityDataType) String() string {
	return "{ name: " + t.Name + ", type: " + t.Type.String() + ", bitSize: " + cast.ToString(t.BitSize) + "}\n"
}
func (i InputData) String() string {
	return "{ param: " + i.Param.String() + ", paramName: " + i.ParamName + ", offset: " + cast.ToString(i.Offset) + ",data: " + string(i.Data) + "}\n"
}
