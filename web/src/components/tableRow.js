import { Link, redirect, useNavigate } from "react-router-dom";
import CampaignListUseCase from "../domain/usecases/campaignListUseCase.usecase";
import { useContext } from "react";
import TableRowContext from "../contexts/tableRow.context";
import PaginationTableContext from "../contexts/paginationTable.context";
import TableRowManipulator from "./lib/tableRowDataAction";

function NavigateButton({navigate, url, icon}) {

    if (typeof url !== 'string' || url === '') {

        return <></>
    }

    return (
        <button onClick={() => {navigate(url, {replace: true})}} class="btn btn-outline-info btn-rounded"><i class={"fas " + icon}></i></button>
    )
}

function DeleteButton({endpointUrl}) {

    if (typeof endpointUrl !== 'string' || endpointUrl === '') {

        return <></>
    }

    return (
        <button class="btn btn-outline-info btn-rounded"><i class="fas fa-trash"></i></button>
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
            {detailUrl && <NavigateButton navigate={navigate} url={detailUrl} icon="fa-file"/>}
            {modfifyUrl &&  <NavigateButton navigate={navigate} url={modfifyUrl} icon="fa-pen"/>}
            {/* {deleteUrl && <Button navigate={navigate} url={deleteUrl} icon="fa-trash"/>} */}
            {deleteUrl && <DeleteButton endpointUrl={deleteUrl} />}
        </td>
    )
}

export default function TableRow({ idField, exposedFields, dataObject, crud , endpoint}) {

    const {columnTransform} = useContext(TableRowContext);
    exposedFields = Array.isArray(exposedFields) ? exposedFields : [];

    return (
        <>
            <tr>
                {/* <td>1</td>
                <td>Dakota Rice</td>
                <td>$36,738</td>
                <td>United States</td>
                <td>Oud-Turnhout</td> */}
                {
                    exposedFields.map(header => {

                        const value = dataObject?.[header];

                        // return <td>{dataObject?.[header]}</td>
                        const transform = columnTransform?.[header];

                        return <td>{typeof transform === 'function' ? transform(value) : value}</td>
                    })
                }
                <RowManipulator idField={idField} rowData={dataObject} crud={crud} endpoint={endpoint}/>
            </tr>
        </>
    )
}