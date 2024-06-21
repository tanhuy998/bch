import { useEffect, useState, Component, memo, createContext, useContext } from "react";
import TableRow from "./tableRow";
import PaginationController from "./paginationController";
import HttpEndpoint from "../backend/endpoint";
import PaginationTableContext from "../contexts/paginationTable.context";
import TableRowContext from "../contexts/tableRow.context";
import TableRowManipulator from "./lib/tableRowDataAction";

const DEFAULT_TABLE_BODY_HEIGHT = 400

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

export default function PaginationTable({ height, refresh, columnTransform, rowManipulator, idField, endpoint, exposedFields, headers, title }) {

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

    const hasRowManipulator = rowManipulator instanceof TableRowManipulator;

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
            <div>
                <table class="table table-hover table-striped" id="dataTables-example" width="100%">
                    <thead style={{
                        display: 'table',
                        width: '100%',
                        'table-layout': 'fixed',
                    }}>
                        <tr>
                            {exposedHeaders}
                            {hasRowManipulator && <td></td>}
                        </tr>
                    </thead>
                    <tbody style={{
                        display: 'block',
                        'max-height': height || DEFAULT_TABLE_BODY_HEIGHT,
                        'overflow-y': 'scroll'
                    }}>
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