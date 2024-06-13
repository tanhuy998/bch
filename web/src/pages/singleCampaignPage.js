import { Link, useParams } from "react-router-dom";
import PaginationTable from "../components/paginationTable";
import SingleCampaignUseCase from "../domain/usecases/singleCampaignUseCase.usecase";
import PaginationTableContext from "../contexts/paginationTableContext";
import { useEffect, useState } from "react";
import { tab } from "@testing-library/user-event/dist/tab";
import Tab from "../components/Tab";



export default function SingleCampaignPage({ usecase }) {

    const { uuid } = useParams();
    const [/**@type {SingleCampaignRespnsePresenter?} */ campaignData, setCampaignData] = useState(null);

    if (!uuid) {

        throw new Error('invalid uuid')
    }

    if (!(usecase instanceof SingleCampaignUseCase)) {

        throw new Error('invalid use case passed to SingleCampaignPage');
    }

    useEffect(() => {

        usecase.fetch(uuid)
            .then((data) => {

                setCampaignData(data)
            })

    }, [])

    const contextValues = {
        EXTRA_FETCH_ARGS: [uuid]
    }

    const tabs = {
        All: (<PaginationTableContext.Provider value={contextValues}>
            <PaginationTable idField="uuid" exposedFields={['name', 'idNumber', 'address']} headers={['Tên', 'Số CCDD', 'Địa Chỉ']} endpoint={usecase.campaignCandidateListEndpoint} title="Canidates" />
        </PaginationTableContext.Provider>
        ),
        Signed: '',
        Unsigned: ''
    }

    console.log('usecase', usecase.campaignCandidateListEndpoint)
    return (
        <>
            <div class="row">
                <div class="col-md-12">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="card">
                                <h5 class="card-header">Campaign Management</h5>
                                <div class="card-body">
                                    <h1 class="card-title">{campaignData?.title || 'Unknown'}</h1>
                                    <br />
                                    <h4 class="card-text">Ngày Bắt Đầu: {campaignData?.issueTime || ''} Ngày Kết Thúc: {campaignData?.expire || ''}</h4>
                                    <p class="card-text">This is some text within a card body...</p>
                                    <Link to="#" class="btn btn-primary">Chỉnh sửa</Link>
                                    <br />

                                </div>

                                <div class="card-body">
                                    <h3 class="card-title">Candidates Detail</h3>
                                    <Tab tabs={tabs} />
                                </div>

                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}