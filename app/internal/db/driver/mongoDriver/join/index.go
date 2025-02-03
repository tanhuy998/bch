package join

func NewJoinQueryBuilder(
	targetCollection string,
) *join_query_builder {

	return &join_query_builder{
		target_collection: targetCollection,
	}
}
