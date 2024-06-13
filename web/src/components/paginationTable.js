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

    const exposedHeaders = (Array.isArray(headers) ? headers : []).map(header => <th>{header}</th>)
    const [endpointData, setEndpointData] = useState(null);

    return (
        <>
            <table class="table table-hover table-striped" id="dataTables-example" width="100%">
                <thead>
                    <tr>
                        {exposedHeaders}
                    </tr>
                </thead>
                <tbody>
                    {
                        (endpointData?.data || []).map(row => {

                            return <TableRow idField={idField} endpoint={endpoint} exposedFields={exposedFields} dataObject={row} />
                        })
                    }
                </tbody>
                <br />
                <PaginationController dataTotalCount={endpointData?.dataTotalCount} endpoint={endpoint} navigator={endpointData?.navigation} setEndpointData={setEndpointData} />
            </table>
        </>
    )
}