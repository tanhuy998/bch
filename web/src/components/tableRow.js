function Button({url, icon}) {

    return (
        <a href={url} class="btn btn-outline-info btn-rounded"><i class={"fas " + icon}></i></a>
    )
}

function RowModificationPanel({ endpoint, rowData, detailUrl, modfifyUrl, deleteUrl }) {

    if (typeof crud != 'object') {

        return <></>
    }

    const buttons = [];

    if (typeof detailUrl === 'string') {

        buttons.push(<Button url={detailUrl} icon="fa-info-circle"/>)
    }

    if (typeof modfifyUrl === 'string') {

        buttons.push(<Button url={modfifyUrl} icon="fa-pen"/>)
    }

    if (typeof deleteUrl === 'string') {

        buttons.push(<Button url={deleteUrl} icon="fa-trash"/>)
    }

    return (
        <td class="text-end">
            {/* <a href={detailUrl} class="btn btn-outline-info btn-rounded"><i class="fas fa-info-circle"></i></a>
            <a href={modfifyUrl} class="btn btn-outline-info btn-rounded"><i class="fas fa-pen"></i></a>
            <a href={deleteUrl} class="btn btn-outline-danger btn-rounded"><i class="fas fa-trash"></i></a> */}
            {buttons}
        </td>
    )
}

export default function TableRow({ exposedFields, dataObject, crud , endpoint}) {

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
                <RowModificationPanel crud={crud} endpoint={endpoint}/>
            </tr>
        </>
    )
}