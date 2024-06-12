import { useEffect, useState } from 'react';
import { DEFAULT_PAGINATION_LIMIT } from '../api/constant';

function PaginationNavButton({
    isLastDataPage, 
    pageReducer, 
    pageCounter, 
    debounce, 
    label, 
    isPrevious, 
    navigationQuery, 
    setQuery,
}) {
   
    if (
        typeof navigationQuery !== 'object' ||
        typeof pageCounter !== 'number' ||
        pageCounter === 1 && isPrevious 
        || isLastDataPage && !isPrevious
    ) {

        return <></>
    }

    const direction = isPrevious ? "previous" : "next";
    const tagClass = `paginate_button page-item ` + direction;
    const tagId = `dataTables-example_` + direction;

    function emitQuery() {

        if (debounce) {

            return;
        }

        pageReducer();
        setQuery(navigationQuery);
    }

    return (
        <li class={tagClass} id={tagId}>
            {/* <a href={endpoint} aria-controls="dataTables-example" data-dt-idx="0" tabindex="0" class="page-link">{label}</a> */}
            <button onClick={() => { debounce === false && emitQuery()} } aria-controls="dataTables-example" data-dt-idx="0" tabindex="0" class="page-link">{label}</button>
        </li>
    )
}

export default function PaginationController({dataTotalCount, currentPageNumber, navigator, setEndpointData, endpoint}) {

    const [query, setQuery] = useState(null)
    const [debounce, setDebounce] = useState(false);
    const [pageCounter, setPageCounter] = useState(dataTotalCount > 0 ? 1 : null);

    function fetchData() {

        endpoint.fetch(query || {})
        .then((data) => {
            setDebounce(false);
            setEndpointData(data)
        });

        setDebounce(true);
    }

    useEffect(() => {

       fetchData();

    }, [query])

    useEffect(() => {

        if (
            dataTotalCount > 0 
            && pageCounter === null
        ) {

            setPageCounter(1);
        }
    })

    const isLastDataPage = calculatePage(dataTotalCount, DEFAULT_PAGINATION_LIMIT) === pageCounter

    return (
        <div class="row">
            <div class="col-sm-12 col-md-5">
                <div class="dataTables_info" id="dataTables-example_info" role="status"
                    aria-live="polite"></div>
            </div>
            <div class="col-sm-12 col-md-7">
                <div class="dataTables_paginate paging_simple_numbers" id="dataTables-example_paginate">
                    <ul class="pagination">
                        <PaginationNavButton pageCounter={pageCounter} pageReducer={() => {setPageCounter(pageCounter - 1)}} debounce={debounce} setQuery={setQuery} navigationQuery={navigator?.previous}  isPrevious={true} label="Trước"/>
                        {/* {pageCounterPlaceHolder} */}
                        { pageCounter > 0 && 
                            <li class="paginate_button page-item active">
                                <a href="#"aria-controls="dataTables-example" data-dt-idx="1" tabindex="0" class="page-link">
                                    {pageCounter}
                                </a>
                            </li>
                        }
                        <PaginationNavButton isLastDataPage={isLastDataPage} pageCounter={pageCounter} pageReducer={() => {setPageCounter(pageCounter + 1)}} debounce={debounce} setQuery={setQuery} navigationQuery={navigator?.next} label="sau"/>
                    </ul>
                </div>
            </div>
        </div>
    )
}

function calculatePage(totalCount, pageLimit) {

    const odd = (totalCount % pageLimit) > 0 ? 1 : 0;
    const even = totalCount % pageLimit;

    return even + odd;
}