package mongoQueryBuilder

type (
	aggregate_relation_debug_log struct {
		//QueryType       string               `json:"query_type"`
		LocalCollection string               `json:"local_collection"`
		Relations       []relation_debug_log `json:"relations"`
	}
)

type (
	relation_debug_log struct {
		RelationType      string `json:"relation_type"`
		ForeignCollection string `json:"foreign_collection"`
	}
)
