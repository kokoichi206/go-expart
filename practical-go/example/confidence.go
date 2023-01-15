package example

import (
	"encoding/json"
	"fmt"
)

// 次の2つの fmt パッケージのインタフェースを満たすといい！！
// type Stringer interface {
// 	String() string
// }
// type GoStringer interface {
// 	GoString() string
// }

type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

type CreditCard string

func (c CreditCard) String() string {
	// String() 関数を定義することで、機密情報をマスク！
	return "xxxx-xxxx-xxxx-xxxx"
}

func (c CreditCard) GoString() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func TestCredential() {
	c := ConfidentialCustomer{
		CustomerID: 0,
		CreditCard: "1234-1234-1234-1234",
	}

	fmt.Println(c)
	fmt.Printf("%v\n", c)

	bytes, _ := json.Marshal(c)
	// ここでは元通り利用可能。
	fmt.Println("JSON: ", string(bytes))
}
