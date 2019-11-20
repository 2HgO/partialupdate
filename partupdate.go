package partialupdate

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strings"
)

/*
 **	PartialUpdate generates a string -> interface mapping of the fields sent in the body
 ** of a PATCH/PUT request based on the structure of the collection model
 ** [TODO] proper documentation of function
 */
func PartialUpdate(model interface{}, in io.Reader) (map[string]interface{}, error) {
	if reflect.ValueOf(model).Kind() != reflect.Struct {
		return nil, errors.New("Model type must be a struct")
	}
	fields := reflect.TypeOf(model)
	structCount := fields.NumField()

	var newFields = make([]reflect.StructField, structCount)
	for i := 0; i < structCount; i++ {
		name := fields.Field(i).Name
		typ := reflect.PtrTo(fields.Field(i).Type)
		var tag reflect.StructTag
		if iftag, ok := fields.Field(i).Tag.Lookup("json"); ok {
			tag = reflect.StructTag(`json:"` + strings.Split(iftag, ",")[0] + `,omitempty"`)
		} else {
			tag = reflect.StructTag(`json:"` + name + `,omitempty"`)
		}
		newFields[i] = reflect.StructField{
			Name: name,
			Type: typ,
			Tag:  tag,
		}
	}

	newStruct := reflect.New(reflect.StructOf(newFields)).Interface()
	err := json.NewDecoder(in).Decode(newStruct)
	if err != nil {
		return nil, err
	}

	var endMap map[string]interface{}
	p, err := json.Marshal(newStruct)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(p, &endMap); err != nil {
		return nil, err
	}

	return endMap, nil
}
