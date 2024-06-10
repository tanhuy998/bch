import {Link} from 'react-router-dom'

function PaginationNavButton({label, endpoint, isPrevious}) {

    const direction = isPrevious ? "previous" : "next";
    const tagClass = `paginate_button page-item ` + direction;
    const tagId = `dataTables-example_` + direction;

    return (
        <li class={tagClass} id={tagId}>
            {/* <a href={endpoint} aria-controls="dataTables-example" data-dt-idx="0" tabindex="0" class="page-link">{label}</a> */}
            <button to={endpoint} aria-controls="dataTables-example" data-dt-idx="0" tabindex="0" class="page-link">{label}</button>
        </li>
    )
}

export default function PaginationPanel({currentPageNumber}) {

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
                    <ul class="pagination">
                        <PaginationNavButton isPrevious={true} label="Trước"/>
                        <li class="paginate_button page-item active"><a href="#"
                            aria-controls="dataTables-example" data-dt-idx="1"
                            tabindex="0" class="page-link">{currentPageNumber}</a></li>
                        <PaginationNavButton label="sau"/>
                    </ul>
                </div>
            </div>
        </div>
    )
}