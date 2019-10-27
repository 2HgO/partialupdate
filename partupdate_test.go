package partupdate

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"fmt"
)

type testModel struct {
	FirstName string 	`json:"firstName,omitempty"`
	LastName  string 	`json:"lastName,omitempty"`
	Age       float64 `json:"age,omitempty"`
}

type testcase struct {
	desc		string
	test  	testModel
	answer 	map[string]interface{}
}

func TestPartialUpdate(t *testing.T) {
	testcases := []testcase{
		{
			"Full update parameters",
			testModel{FirstName: "oghogho", LastName: "odemwingie", Age: 22},
			map[string]interface{}{"firstName": "oghogho", "lastName": "odemwingie", "age": float64(22)},
		},
		{
			"Two parameter update",
			testModel{FirstName: "oghogho", Age: 22},
			map[string]interface{}{"firstName": "oghogho", "age": float64(22)},
		},
		{
			"One string parameter update",
			testModel{LastName: "odemwingie"},
			map[string]interface{}{"lastName": "odemwingie"},
		},
		{
			"One number parameter upadate",
			testModel{Age: 22},
			map[string]interface{}{"age": float64(22)},
		},
	}

	for _, val := range testcases {
		var (
			model testModel
		)
		buff, _ := json.Marshal(val.test)

		if got, err := PartialUpdate(model, bytes.NewReader(buff)); err != nil || !reflect.DeepEqual(got, val.answer) {
			t.Errorf("\tGot: %v;\n\t\t\tExpected: %v\n", got, val.answer)
		} else {
			fmt.Println(val.desc, "passed")
		}
	}
}
