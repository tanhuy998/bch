package model

import "encoding/json"

type JsonObjectValue[Object_Value_T any] map[string]Object_Value_T

type (
	OperationKind = string

	JsonPatchRawValueOperation struct {
		Op    string      `json:"op" bson:"op"`
		Path  string      `json:"path" bson:"path"`
		Value interface{} `json:"value" bson:"value"`
	}

	JsonPatchGenericValueOperation[Op_Value_T any] struct {
		Op    string     `json:"op" bson:"op"`
		Path  string     `json:"path" bson:"path"`
		Value Op_Value_T `json:"value" bson:"value"`
	}
)

func ConvertJsonPatchOperation[Op_Value_T any](
	target *JsonPatchRawValueOperation,
) (*JsonPatchGenericValueOperation[Op_Value_T], error) {

	jsonPatchOpValue := target.Value

	val, ok := jsonPatchOpValue.(Op_Value_T)

	if ok {

		return &JsonPatchGenericValueOperation[Op_Value_T]{
			Op:    target.Op,
			Path:  target.Path,
			Value: val,
		}, nil
	}

	jsonPatchOpValue, err := convertJsonPatchOperationValueFromRaw[Op_Value_T](jsonPatchOpValue)

	if err != nil {

		return nil, err
	}

	return &JsonPatchGenericValueOperation[Op_Value_T]{
		Op:    target.Op,
		Path:  target.Path,
		Value: val,
	}, nil
}

func convertJsonPatchOperationValueFromRaw[Op_Value_T any](val interface{}) (*Op_Value_T, error) {

	var rawValue []byte

	switch val.(type) {
	case string:
		rawValue = []byte(val.(string))
	case []byte:
		rawValue = val.([]byte)
	}

	var ret *Op_Value_T = new(Op_Value_T)

	err := json.Unmarshal(rawValue, &ret)

	if err != nil {

		return nil, err
	}

	return ret, nil
}
