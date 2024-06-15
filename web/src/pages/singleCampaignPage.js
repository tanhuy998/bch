import { Link, useParams } from "react-router-dom";
import PaginationTable from "../components/paginationTable";
import SingleCampaignUseCase from "../domain/usecases/singleCampaignUseCase.usecase";
import PaginationTableContext from "../contexts/paginationTable.context";
import { useEffect, useState } from "react";
import { tab } from "@testing-library/user-event/dist/tab";
import Tab from "../components/Tab";
import PillTab from "../components/pillTab";
import BasicTab from "../components/basicTab";



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

    const defaultTableContext = {
        idField: "uuid",
        exposedFields: ['name', 'idNumber', 'address'],
        headers: ['Tên', 'Số CCDD', 'Địa Chỉ'],
        endpoint: usecase.campaignCandidateListEndpoint,
        title: "Canidates",
        rowManipulator: usecase.candidateListTableRowManipulator
    }

    const allCandidateExtraContextValues = {
        EXTRA_FETCH_ARGS: [uuid]
    }

    const candidateDisplayTable = (
        // <PaginationTable idField="uuid" exposedFields={['name', 'idNumber', 'address']} headers={['Tên', 'Số CCDD', 'Địa Chỉ']} endpoint={usecase.campaignCandidateListEndpoint} title="Canidates" />
        <PaginationTable />
    )

    const tabs = {
        All: (
            <PaginationTableContext.Provider value={{ ...defaultTableContext, ...allCandidateExtraContextValues }}>
                <PaginationTable />
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
                                    <div class="container">
                                        <div class="row">
                                            <div class="col">
                                                <h4 class="card-text">Ngày Bắt Đầu: {campaignData?.issueTime || ''}</h4>
                                            </div>
                                            <div class="col">
                                                <h4 class="card-text">Ngày Kết Thúc: {campaignData?.expire || ''}</h4>
                                            </div>
                                        </div>
                                    </div>
                                    <br/>
                                    <BasicTab tabs={{'Mô tả': `

Lorem ipsum dolor sit amet. Eos neque dolor id accusantium reprehenderit et impedit ipsam quo illo praesentium. Ut porro amet aut autem ducimus At voluptas repellat. Est accusamus sequi est fuga voluptate ut aliquid minima 33 dolores nisi est maxime aspernatur qui sunt voluptatum. Ad sequi iure nam vero quis et aliquam repellat et eaque soluta et galisum quaerat ut rerum esse.

Qui nesciunt corporis in praesentium nemo ut vitae dolores et accusantium vitae. Ut nihil earum et iste quis id provident molestiae non voluptas veritatis ut ducimus similique et nulla pariatur et sint officia. In eius quibusdam qui facilis neque qui rerum consequatur. Qui reprehenderit possimus quo repellat quasi aut asperiores labore rem neque veritatis ut veniam distinctio est Quis dolor sit velit optio.

A nisi aspernatur non natus aliquam aut mollitia rerum. Non magnam aperiam quo eligendi veritatis sit eaque perferendis et atque modi aut cumque odio ex dolorum provident! Ut vero impedit ad voluptatem optio ut Quis velit. Sit suscipit dolor sit voluptatem voluptatum eum inventore tenetur est quas quia qui eligendi minima.
`}}/>
                                    <br/>
                                    
                                    <Link to="#" class="btn btn-primary">Chỉnh sửa</Link>
                                    <br />

                                </div>

                                <div class="card-body">
                                    <h3 class="card-title">Candidates Detail</h3>
                                    <PillTab tabs={tabs} />
                                </div>

                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}