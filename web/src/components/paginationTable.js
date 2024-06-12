import { useEffect, useState, Component, memo, createContext } from "react";
import TableRow from "./tableRow";
import PaginationController from "./paginationController";
import HttpEndpoint from "../backend/endpoint";

/**
 * 
 * @param {Function} navigator 
 * @returns {Function}
 */
function dispatchNavigate(navigator) {

    return
}

async function fetchData(setDataFunc, endpoint, query) {


}



export default memo(PaginationTable);

function PaginationTable({ idField, endpoint, exposedFields, headers, title }) {

    if (!(endpoint instanceof HttpEndpoint)) {

        throw new Error('invalid endpoint object to for fetching data');
    }

    const exposedHeaders = (Array.isArray(headers) ? headers : []).map(header => <th>{header}</th>)
    const [endpointData, setEndpointData] = useState(null);

    return (
        <>
            <div class="page-title">

                <h3>{title}
                    <a href="roles.html" class="btn btn-sm btn-outline-primary float-end"><i class="fas fa-plus-circle"></i> Thêm mới</a>
                    {/* <a href="users.html" class="btn btn-sm btn-outline-info float-end me-1"><i class="fas fa-angle-left"></i> <span class="btn-header">Return</span></a> */}
                </h3>
            </div>
            <div class="row">
                <div class="col-md-12 col-lg-12">
                    <div class="card">
                        <div class="card-header">Basic DataTables Table</div>
                        <div class="card-body">
                            <p class="card-title"></p>
                            <table class="table table-hover" id="dataTables-example" width="100%">
                                <thead>
                                    <tr>
                                        {exposedHeaders}
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        (endpointData?.data || []).map( row => {

                                            return <TableRow idField={idField} endpoint={endpoint} exposedFields={exposedFields} dataObject={row}/>
                                        })
                                    }
                                </tbody>
                            </table>
                        </div>
                        <PaginationController dataTotalCount={endpointData?.dataTotalCount} endpoint={endpoint} navigator={endpointData?.navigation} setEndpointData={setEndpointData} />
                    </div>
                </div>
            </div>
        </>
    )
}