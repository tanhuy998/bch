import { useEffect, useState, Component, memo, createContext, useContext } from "react";
import TableRow from "./tableRow";
import PaginationController from "./paginationController";
import HttpEndpoint from "../backend/endpoint";
import PaginationTableContext from "../contexts/paginationTable.context";
import TableRowContext from "../contexts/tableRow.context";

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



//export default memo(PaginationTable);

export default function PaginationTable({ refresh, columnTransform, rowManipulator, idField, endpoint, exposedFields, headers, title }) {

    const context = useContext(PaginationTableContext);

    idField = idField || context?.idField;
    endpoint = endpoint || context?.endpoint;
    exposedFields = exposedFields || context?.exposedFields;
    headers = headers || context?.headers;
    title = title || context?.title;
    rowManipulator = rowManipulator || context?.rowManipulator;
    columnTransform = columnTransform || context?.columnTransform;

    const exposedHeaders = (Array.isArray(headers) ? headers : []).map(header => <th>{header}</th>)
    const [endpointData, setEndpointData] = useState(null);
    const [lastRefreshSignal, setLastRefreshSignal] = useState(refresh);

    useEffect(() => {

        if (refresh === lastRefreshSignal) {

            return;
        }

        setLastRefreshSignal(refresh);

        // if (!refresh) {

        //     return;
        // }

        setEndpointData(null);

    }, []);

    return (
        <>
            <div style={{ height: "400px", overflow: "auto" }}>
                <table class="table table-hover table-striped" id="dataTables-example" width="100%">
                    <thead style={{position: "sticky"}}>
                        <tr>
                            {exposedHeaders}
                        </tr>
                    </thead>
                    <tbody>
                        <TableRowContext.Provider value={{ rowManipulator, columnTransform }}>
                            {
                                (endpointData?.data || []).map(row => {

                                    return <TableRow idField={idField} endpoint={endpoint} exposedFields={exposedFields} dataObject={row} />
                                })
                            }
                        </TableRowContext.Provider>
                    </tbody>
                </table>
            </div>
            <br />
            <PaginationController dataTotalCount={endpointData?.dataTotalCount} endpoint={endpoint} navigator={endpointData?.navigation} setEndpointData={setEndpointData} />
        </>
    )
}