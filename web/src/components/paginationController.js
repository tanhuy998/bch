import { createContext, useContext, useEffect, useState } from 'react';
import { DEFAULT_PAGINATION_LIMIT } from '../api/constant';
import PaginationTableContext, { EXTRA_FETCH_ARGS } from '../contexts/paginationTable.context';
import HttpEndpoint from '../backend/endpoint';

const PaginationControllerContext = createContext({
    isLastDataPage: null, 
    pageCounter: null, 
    debounce: null, 
    setQuery: null,
})

function PaginationNavButton({
    pageReducer,
    label,
    isPrevious,
    navigationQuery,
}) {

    const { isLastDataPage, pageCounter, debounce, setQuery } = useContext(PaginationControllerContext);

    let exposedButton;

    function emitQuery() {

        if (debounce) {

            return;
        }

        pageReducer();
        setQuery(navigationQuery);
    }

    
    const direction = isPrevious ? "previous" : "next";
    const tagClass = `paginate_button page-item ` + direction;
    const tagId = `dataTables-example_` + direction;

    if (
        typeof navigationQuery !== 'object' || 
        typeof pageCounter !== 'number' ||
        pageCounter === 1 && isPrevious
        || isLastDataPage && !isPrevious
    ) {

        exposedButton = <></>;
    }
    else {

        exposedButton = <button onClick={() => { debounce === false && emitQuery() }} aria-controls="dataTables-example" data-dt-idx="0" tabindex="0" class="page-link">{label}</button>;
    }

    return (
        <li class={tagClass} id={tagId}>
            {/* <a href={endpoint} aria-controls="dataTables-example" data-dt-idx="0" tabindex="0" class="page-link">{label}</a> */}
            {exposedButton}
        </li>
    )
}

export default function PaginationController({ dataTotalCount, currentPageNumber, navigator, setEndpointData, endpoint }) {

    if (!(endpoint instanceof HttpEndpoint)) {

        throw new Error('invalid endpoint object to for fetching data');
    }

    const [query, setQuery] = useState(null)
    const [debounce, setDebounce] = useState(false);
    const [pageCounter, setPageCounter] = useState(1);

    const tableContext = useContext(PaginationTableContext);
    const extra_fetch_args = tableContext?.[EXTRA_FETCH_ARGS];


    function fetchData() {

        endpoint.fetch(query || {}, ...(extra_fetch_args || []))
            .then((data) => {
                setDebounce(false);
                setEndpointData(data)
            })
            .catch(err => {

                alert(err?.message || err);
            });

        setDebounce(true);
    }

    useEffect(() => {

        fetchData();

    }, [query])
    
    const isLastDataPage = calculatePage(dataTotalCount, DEFAULT_PAGINATION_LIMIT) === pageCounter
    const context = {
        isLastDataPage,
        pageCounter,
        debounce,
        setQuery,
    }

    return (
        <PaginationControllerContext.Provider value={context}>
            <div class="row">
                <div class="col-sm-12 col-md-5">
                    <div class="dataTables_info" id="dataTables-example_info" role="status"
                        aria-live="polite"></div>
                </div>

                <div class="col-sm-12 col-md-7">
                    <div class="dataTables_paginate paging_simple_numbers" id="dataTables-example_paginate">
                        <ul class="pagination">
                            
                            <PaginationNavButton  pageReducer={() => { setPageCounter(pageCounter - 1) }} navigationQuery={navigator?.previous} isPrevious={true} label="Trước" />
                            {/* {pageCounterPlaceHolder} */}
                            {pageCounter > 0 &&
                                <li class="paginate_button page-item active">
                                    <a href="#" aria-controls="dataTables-example" data-dt-idx="1" tabindex="0" class="page-link">
                                        {pageCounter}
                                    </a>
                                </li>
                            }
                            <PaginationNavButton pageReducer={() => { setPageCounter(pageCounter + 1) }} navigationQuery={navigator?.next} label="sau" />
                            
                        </ul>
                    </div>
                </div>
            </div>
        </PaginationControllerContext.Provider>
    )
}

function calculatePage(totalCount, pageLimit) {

    const odd = (totalCount % pageLimit) > 0 ? 1 : 0;
    const even = Math.floor(totalCount / pageLimit);
    
    return even + odd;
}