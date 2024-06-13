import { createContext } from "react"

const TableRowContext = createContext({
    columnTransform: null,
    rowManipulator: null,
})

export default TableRowContext