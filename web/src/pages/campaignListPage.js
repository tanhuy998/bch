import { useEffect } from "react";
import BasicTab from "../components/basicTab";
import PaginationTable from "../components/paginationTable";
import PillTab from "../components/pillTab";
import Tab from "../components/Tab";
import PaginationTableContext from "../contexts/paginationTable.context";
import TabContext, { basicTabStyle, pillTabStyle } from "../contexts/tab.context";
import CampaignListUseCase from "../domain/usecases/campaignListUseCase.usecase";
import formatLocalDate from "../lib/formatLocalDate";
import { Link } from "react-router-dom";

const COLUMN_TRANSFORM = {
    issueTime: transformDateString,
    expire: transformDateString,
}

function transformDateString(value) {

    return formatLocalDate(new Date(value));
}

export default function ({ usecase }) {

    if (!(usecase instanceof CampaignListUseCase)) {

        throw new Error('invalid usecase for CampaignList page');
    }

    useEffect(() => {


    }, []);

    const tableContext = {
        columnTransform: COLUMN_TRANSFORM,
        rowManipulator: usecase.tableRowManipulator,
        idField: "uuid",
        exposedFields: ['title', 'issueTime', 'expire'],
        headers: ['Campaign Name', 'Issue Time', 'Expires'],
        title: "Campaigns",
    }

    const tabs = {
        All: (
            <PaginationTableContext.Provider value={tableContext}>
                <PaginationTable endpoint={usecase}/>
            </PaginationTableContext.Provider>
        ),
        Completed: '',
        InCompleted: '',
    }

    return (
        <>
            <div class="page-title">

                <h3>
                    Campaign List
                    <Link to="/admin/campaign/new" class="btn btn-sm btn-outline-primary float-end"><i class="fas fa-plus-circle"></i> Thêm mới</Link>
                    {/* <a href="users.html" class="btn btn-sm btn-outline-info float-end me-1"><i class="fas fa-angle-left"></i> <span class="btn-header">Return</span></a> */}
                </h3>
            </div>
            <div class="row">
                <div class="col-md-12 col-lg-12">
                    <div class="card">
                        <div class="card-header">Basic DataTables Table</div>
                        <div class="card-body">
                            
                            <BasicTab tabs={tabs} initTabKey={"All"}/>
                        </div>
                        {/* <div class="card-body">
                            <p class="card-title"></p>
                            <PaginationTable columnTransform={COLUMN_TRANSFORM} rowManipulator={usecase.tableRowManipulator} idField={"uuid"} endpoint={usecase} exposedFields={['title', 'issueTime', 'expire']} headers={['Campaign Name', 'Issue Time', 'Expires']} title="Campaigns" />
                        </div> */}
                    </div>
                </div>
            </div>
        </>
    )
}