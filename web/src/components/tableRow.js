import { Link, redirect, useNavigate } from "react-router-dom";
import CampaignListUseCase from "../domain/usecases/campaignListUseCase.usecase";
import { useContext } from "react";
import TableRowContext from "../contexts/tableRow.context";
import PaginationTableContext from "../contexts/paginationTable.context";
import TableRowManipulator from "./lib/tableRowDataAction";

function Button({navigate, url, icon}) {

    if (typeof url !== 'string' || url === '') {

        return <></>
    }

    return (
        <button onClick={() => {navigate(url, {replace: true})}} class="btn btn-outline-info btn-rounded"><i class={"fas " + icon}></i></button>
    )
}

function RowManipulator({ endpoint, crud, rowData, idField}) {

    const {rowManipulator} = useContext(TableRowContext);
    const navigate = useNavigate();

    if (
        typeof idField !== 'string' 
        || idField === '' 
        //|| !(endpoint instanceof CampaignListUseCase)  
    ) {
        
        return <></>
    }

    if (!(rowManipulator instanceof TableRowManipulator)) {

        return <></>
    }

    const id = rowData[idField];

    const detailUrl = rowManipulator.generateRowDetailPath(id);
    const modfifyUrl = rowManipulator.generateRowModificationPath(id);
    const deleteUrl = rowManipulator.generateRowDeletePath(id);
    console.log(detailUrl, modfifyUrl, deleteUrl)
    return (
        <td class="text-end">
            {/* <a href={detailUrl} class="btn btn-outline-info btn-rounded"><i class="fas fa-info-circle"></i></a>
            <a href={modfifyUrl} class="btn btn-outline-info btn-rounded"><i class="fas fa-pen"></i></a>
            <a href={deleteUrl} class="btn btn-outline-danger btn-rounded"><i class="fas fa-trash"></i></a> */}
            {detailUrl && <Button navigate={navigate} url={detailUrl} icon="fa-info-circle"/>}
            {modfifyUrl &&  <Button navigate={navigate} url={modfifyUrl} icon="fa-pen"/>}
            {deleteUrl && <Button navigate={navigate} url={deleteUrl} icon="fa-trash"/>}
        </td>
    )
}

export default function TableRow({ idField, exposedFields, dataObject, crud , endpoint}) {

    exposedFields = Array.isArray(exposedFields) ? exposedFields : [];

    return (
        <>
            <tr>
                {/* <td>1</td>
                <td>Dakota Rice</td>
                <td>$36,738</td>
                <td>United States</td>
                <td>Oud-Turnhout</td> */}
                {exposedFields.map(header => <td>{dataObject?.[header]}</td>)}
                <RowManipulator idField={idField} rowData={dataObject} crud={crud} endpoint={endpoint}/>
            </tr>
        </>
    )
}