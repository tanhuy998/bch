import { createContext } from "react";
import TableRowManipulator from "../components/lib/tableRowDataAction";

/**@type {TableRowManipulator} */
const rowManipulator = null 

const PaginationTableContext =  createContext({ 
    idField: '', 
    endpoint: null,   
    exposedFields: [], 
    headers: [], 
    title: '',
    rowManipulator,
    columnTransform: null,
});

export const EXTRA_FETCH_ARGS = 'EXTRA_FETCH_ARGS'

export const PaginationController = createContext();

export default PaginationTableContext;