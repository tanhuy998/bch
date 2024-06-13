import { DEFAULT_PAGINATION_LIMIT } from "../constant";

export default function preprocessPaginationQuery(query = {}) {

    return {
        p_pivot: query.p_pivot || '',
        p_limit: query.p_limit || DEFAULT_PAGINATION_LIMIT,
        p_prev: query.p_prev || false,
    }
}