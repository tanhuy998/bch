import { Link } from "react-router-dom";
import CampaignListUseCase from "../domain/usecases/campaignListUseCase.usecase";

function Button({url, icon}) {

    if (typeof url !== 'string' || url === '') {

        return <></>
    }

    return (
        <Link to={url} class="btn btn-outline-info btn-rounded"><i class={"fas " + icon}></i></Link>
    )
}

function RowModificationPanel({endpoint, crud, rowData, idField}) {

    if (
        typeof idField !== 'string' 
        || idField === '' 
        || !(endpoint instanceof CampaignListUseCase)
    ) {
        console.log("nothing", idField)
        return <></>
    }

    const detailUrl = endpoint.generateGetSingleCampaignURL(rowData[idField]);
    const modfifyUrl = endpoint.generateModifySingleCampaignURL(rowData[idField]);
    const deleteUrl = endpoint.generateDeleteSingleCampaignURL();

    return (
        <td class="text-end">
            {/* <a href={detailUrl} class="btn btn-outline-info btn-rounded"><i class="fas fa-info-circle"></i></a>
            <a href={modfifyUrl} class="btn btn-outline-info btn-rounded"><i class="fas fa-pen"></i></a>
            <a href={deleteUrl} class="btn btn-outline-danger btn-rounded"><i class="fas fa-trash"></i></a> */}
            <Button url={detailUrl} icon="fa-info-circle"/>
            <Button url={modfifyUrl} icon="fa-pen"/>
            <Button url={deleteUrl} icon="fa-trash"/>
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
                <RowModificationPanel idField={idField} rowData={dataObject} crud={crud} endpoint={endpoint}/>
            </tr>
        </>
    )
}