package paginateOutput

import (
	libCommon "app/internal/lib/common"
	"encoding/json"
)

const (
	optional_query_struct_key = "OptionalQuery"
)

type (
	// Because NavigatorDir implement json.Marshaler.
	// This type is defined in order to prevent recursion
	// when passing directly the value of NavigatorDir type
	// to json.Marshal() without converting it into another similar type which
	// does not implement json.Marshaler.
	_nav_dir_t[Optional_Query_T any] NavigatorDir[Optional_Query_T]

	NavigatorDir[Optional_Query_T any] struct {
		Cursor        interface{}       `json:"p_cursor,omitempty"`
		PageSize      uint64            `json:"p_size,omitempty"`
		Offset        int64             `json:"p_page,omitempty"`
		OptionalQuery *Optional_Query_T `json:",omitempty"`
	}
)

func (this *NavigatorDir[Optional_Query_T]) init(cursor interface{}, pageSize uint64, offset int64) {

	this.Cursor = cursor
	this.PageSize = pageSize
	this.Offset = libCommon.Ternary(offset < 0, 0, offset)
}

func (this NavigatorDir[Optional_Query_T]) MarshalJSON() (bytes []byte, err error) {

	bytes, err = json.Marshal(
		_nav_dir_t[Optional_Query_T](this),
	)

	if this.OptionalQuery == nil {

		return
	}

	var mapOfThis map[string]json.RawMessage
	_ = json.Unmarshal(bytes, &mapOfThis)

	bytes, _ = json.Marshal(*this.OptionalQuery)

	var mapOfOption map[string]json.RawMessage
	_ = json.Unmarshal(bytes, &mapOfOption)

	delete(mapOfThis, optional_query_struct_key)

	for k, v := range mapOfOption {

		mapOfThis[k] = v
	}

	return json.Marshal(mapOfThis)
}
