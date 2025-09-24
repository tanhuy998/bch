package mongoStorage

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

// type (
// 	IMongoDebugLoggableContext interface {
// 		GetMongoDebogLog() interface{}
// 	}
// )

type (
	aggregate_debug_log_t struct {
		QueryType string      `json:"query_type"`
		Detail    interface{} `json:"detail,omitempty"`
		Query     interface{} `json:"query,omitempty"`
	}
)

type (
	collection_filterable_operation_log_t struct {
		//AppliedFilter interface{} `json:"applied_filter"`
		AppliedFilter filter_debug_log_t `json:"applied_filter"`
	}
)

type (
	_reversed_json_marshaller_filter_log_t filter_debug_log_t

	filter_debug_log_t struct {
		OriginFilter interface{}    `json:"origin_form"`
		BsonLogForm  _bson_log_form `json:"bson_form"`
		Err          error          `json:"error"`
	}
)

func (this *filter_debug_log_t) MarshalJSON() ([]byte, error) {

	this.resolveLog()

	return json.Marshal(_reversed_json_marshaller_filter_log_t(*this))
}

type (
	_bson_log_form bson.Raw
)

func (this _bson_log_form) MarshalJSON() ([]byte, error) {

	return bson.MarshalExtJSONIndent(
		bson.Raw(this), true, true, "", "\t",
	)
}

func NewfilterableOpLogLine(filter interface{}) *collection_filterable_operation_log_t {

	ret := &collection_filterable_operation_log_t{
		filter_debug_log_t{
			OriginFilter: filter,
		},
	}

	return ret
}

func (this *filter_debug_log_t) resolveLog() {

	// raw, err := bson.MarshalExtJSONIndent(this.JsonForm, true, true, "", "\t")

	// this.MarshalledForm = string(raw)
	// this.Error = err

	r, err := bson.Marshal(this.OriginFilter)

	if err != nil {

		this.Err = err
		return
	}

	this.BsonLogForm = _bson_log_form(r)
}

// func (this filter_debug_log_t) MarshalJSON() ([]byte, error) {

// 	//var exported _inverse_filter_log_t = _inverse_filter_log_t(this)

// 	rawForm, err := bson.Marshal(this.JsonForm)

// 	if err != nil {

// 		return nil, err
// 	}

// 	logMap := make(map[string]interface{})

// 	r_jsonForm, _ := json.Marshal(this.JsonForm)
// 	logMap["json_form"] = string(r_jsonForm)

// 	r_rawForm, _ := json.Marshal(bson.Raw(rawForm))
// 	logMap["raw_form"] = string(r_rawForm)

// 	return json.Marshal(logMap)
// }
