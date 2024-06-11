import { useEffect, useState } from 'react';
import {Link} from 'react-router-dom'

async function fetchData(endpoint, query, setEndpointData, setDebounce) {

    console.log('table fetch')
    try {
        
        const res = await endpoint.fetch(query || {});
        console.log(res)
        setDebounce(false);
        setEndpointData(res);
    }
    catch (e) {
        
        console.log(e)
    }
}

function PaginationNavButton({debounce, setEndpointData, endpoint, setDebounce, label, isPrevious, navigationQuery, setQuery}) {

    const direction = isPrevious ? "previous" : "next";
    const tagClass = `paginate_button page-item ` + direction;
    const tagId = `dataTables-example_` + direction;

    if (typeof navigationQuery !== 'object') {

        return <></>
    }

    function retrieveData() {
        //console.log(navigationQuery)
        fetchData(endpoint, navigationQuery, setEndpointData, setDebounce);
        setDebounce(true)
    }

    return (
        <li class={tagClass} id={tagId}>
            {/* <a href={endpoint} aria-controls="dataTables-example" data-dt-idx="0" tabindex="0" class="page-link">{label}</a> */}
            <button onClick={() => { console.log('click', label, debounce); debounce === false && retrieveData()}} aria-controls="dataTables-example" data-dt-idx="0" tabindex="0" class="page-link">{label}</button>
        </li>
    )
}

export default function PaginationController({ initDebounce, currentPageNumber, navigator, endpointData, setEndpointData, endpoint}) {

    const [query, setQuery] = useState(null)
    const [debounce, setDebounce] = useState(false);

    useEffect(() => {

        fetchData(endpoint, {}, setEndpointData, setDebounce);
    }, [query])

    currentPageNumber = currentPageNumber || 1
    
    return (
        <div class="row">
            <div class="col-sm-12 col-md-5">
                <div class="dataTables_info" id="dataTables-example_info" role="status"
                    aria-live="polite"></div>
            </div>
            <div class="col-sm-12 col-md-7">
                <div class="dataTables_paginate paging_simple_numbers"
                    id="dataTables-example_paginate">
                    {typeof endpointData !== 'object' && setQuery({})}
                    <ul class="pagination">
                        <PaginationNavButton endpoint={endpoint} setEndpointData={setEndpointData} setDebounce={setDebounce} debounce={debounce} setQuery={setQuery} navigationQuery={navigator?.previous}  isPrevious={true} label="Trước"/>
                        <li class="paginate_button page-item active"><a href="#"
                            aria-controls="dataTables-example" data-dt-idx="1"
                            tabindex="0" class="page-link">{currentPageNumber}</a></li>
                        <PaginationNavButton endpoint={endpoint} setEndpointData={setEndpointData} setDebounce={setDebounce} debounce={debounce} setQuery={setQuery} navigationQuery={navigator?.next} label="sau"/>
                    </ul>
                </div>
            </div>
        </div>
    )
}