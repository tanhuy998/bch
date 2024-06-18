import { Link, useParams } from "react-router-dom";
import PaginationTable from "../components/paginationTable";
import SingleCampaignUseCase from "../domain/usecases/singleCampaignUseCase.usecase";
import PaginationTableContext from "../contexts/paginationTable.context";
import { useEffect, useState } from "react";
import { tab } from "@testing-library/user-event/dist/tab";
import Tab from "../components/Tab";
import PillTab from "../components/pillTab";
import BasicTab from "../components/basicTab";
import Form from "../components/form";
import FormInput from "../components/formInput";
import { required } from "../components/lib/validator.";
import NewCandidateFormDelegator from "../domain/valueObject/newCandidateFormDelegator";



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



    // const candidateDisplayTable = (
    //     // <PaginationTable idField="uuid" exposedFields={['name', 'idNumber', 'address']} headers={['Tên', 'Số CCDD', 'Địa Chỉ']} endpoint={usecase.campaignCandidateListEndpoint} title="Canidates" />
    //     <PaginationTable />
    // )

    const candidateTabs = {
        // All: //<CompactCampaignCandidateTable usecase={usecase} uuid={uuid} />,
        Signed: <CampaignCandidateProgression uuid={uuid} endpoint={usecase.campaignCandidateListEndpoint} />,
        Unsigned: <CampaignCandidateProgression uuid={uuid} endpoint={usecase.campaignCandidateListEndpoint} />,
    }

    const mainTabs = {
        'Mô tả': `Lorem ipsum dolor sit amet. Eos neque dolor id accusantium reprehenderit et impedit ipsam quo illo praesentium. Ut porro amet aut autem ducimus At voluptas repellat. Est accusamus sequi est fuga voluptate ut aliquid minima 33 dolores nisi est maxime aspernatur qui sunt voluptatum. Ad sequi iure nam vero quis et aliquam repellat et eaque soluta et galisum quaerat ut rerum esse.

Qui nesciunt corporis in praesentium nemo ut vitae dolores et accusantium vitae. Ut nihil earum et iste quis id provident molestiae non voluptas veritatis ut ducimus similique et nulla pariatur et sint officia. In eius quibusdam qui facilis neque qui rerum consequatur. Qui reprehenderit possimus quo repellat quasi aut asperiores labore rem neque veritatis ut veniam distinctio est Quis dolor sit velit optio.

A nisi aspernatur non natus aliquam aut mollitia rerum. Non magnam aperiam quo eligendi veritatis sit eaque perferendis et atque modi aut cumque odio ex dolorum provident! Ut vero impedit ad voluptatem optio ut Quis velit. Sit suscipit dolor sit voluptatem voluptatum eum inventore tenetur est quas quia qui eligendi minima.`,
        Candidates: (
            <div class="card-body">
                {/* <h3 class="card-title">Candidates Detail</h3> */}

                {/* <PillTab tabs={candidateTabs} /> */}
                <CompactCampaignCandidateTable formDelegator={usecase.newCandidateFormDelegator} endpoint={usecase.campaignCandidateListEndpoint} uuid={uuid} />
            </div>
        ),
        Progression: (
            <div class="card-body">
                <PillTab tabs={candidateTabs}/>
            </div>
        )
    };

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
                                    <br />

                                    <Link to="#" class="btn btn-primary">Chỉnh sửa</Link>
                                    <br />
                                    <br />
                                    <BasicTab tabs={mainTabs} initTabIndex={0} />
                                </div>

                                {/* <div class="card-body">
                                    <h3 class="card-title">Candidates Detail</h3>

                                    <PillTab tabs={candidateTabs} />
                                </div> */}

                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

function CompactCampaignCandidateTable({ uuid, endpoint, formDelegator }) {

    if (!(formDelegator instanceof NewCandidateFormDelegator)) {

        throw new Error('formDelegator must be instance of NewCandidateFormDelegator');
    }

    const [formVisible, setFormVisible] = useState(false);

    const allCandidateExtraContextValues = {
        EXTRA_FETCH_ARGS: [uuid]
    }

    const defaultTableContext = {
        idField: "uuid",
        exposedFields: ['name', 'idNumber', 'address'],
        headers: ['Tên', 'Số CCDD', 'Địa Chỉ'],
        endpoint: endpoint,
        title: "Canidates",
        rowManipulator: endpoint,
    }

    return (
        <>
            {!formVisible && (
                <>
                    <h5>
                        All Candidates
                        <button onClick={() => { toggleCompactTableForm(formVisible, setFormVisible) }} class="btn btn-sm btn-outline-primary float-end"><i class="fas fa-plus-circle"></i> Thêm mới</button>
                        {/* <a href="users.html" class="btn btn-sm btn-outline-info float-end me-1"><i class="fas fa-angle-left"></i> <span class="btn-header">Return</span></a> */}
                    </h5>
                    <br/>
                </>
            )}

            {formVisible && (
                <>
                    <div class="card-body" style={{ "background-color": '#E0E0E0' }}>
                        <h3 class="card-title">New Candidate</h3>
                        <br />
                        <Form delegate={formDelegator} className="needs-validation" novalidate="" accept-charset="utf-8">
                            <div class="container" >
                                <div class="row">
                                    <div class="mb-6 col">
                                        <label for="address" className="form-label">Candidate Name</label>
                                        <FormInput validate={required} type="text" className="form-control" name="title" required="true" />
                                    </div>
                                    <div class="mb-6 col">
                                        <label for="address" class="form-label">ID Number</label>
                                        <FormInput validate={required} className="form-control" name="description" />
                                    </div>
                                </div>
                                <br />
                                <div class="row">
                                    <div class="col mb-6">
                                        <div >
                                            <label for="state" className="form-label" >Date Of Birth</label>
                                            <FormInput name="expire" type="date" data-date-format="DD-MM-YYYY" className="form-control" required="true" />
                                            {/* <DatePicker className="form-control"/> */}
                                        </div>
                                    </div>
                                    <div class="col mb-6">
                                        <label for="address" class="form-label">Adress</label>
                                        <FormInput validate={required} className="form-control" name="description" />
                                    </div>
                                </div>
                                <br />
                            </div>
                            <button type="submit" class="btn btn-primary"><i class="fas fa-save"></i> Save</button>
                            <button onClick={() => {toggleCompactTableForm(formVisible, setFormVisible)}} className="btn btn-sm btn-outline-primary" style={{margin: 5}}>Hủy</button>
                        </Form>
                    </div>
                    <br />
                </>
            )}
            {formVisible && <h5>All Candidate</h5>}
            <PaginationTableContext.Provider value={{ ...defaultTableContext, ...allCandidateExtraContextValues }}>
                <PaginationTable />
            </PaginationTableContext.Provider>
        </>
    )
}

function CampaignCandidateProgression({uuid, endpoint}) {

    const allCandidateExtraContextValues = {
        EXTRA_FETCH_ARGS: [uuid]
    }

    const defaultTableContext = {
        idField: "uuid",
        exposedFields: ['name', 'idNumber', 'address'],
        headers: ['Tên', 'Số CCDD', 'Địa Chỉ'],
        endpoint: endpoint,
        title: "Canidates",
        rowManipulator: endpoint,
    }

    return (
        <PaginationTableContext.Provider value={{ ...defaultTableContext, ...allCandidateExtraContextValues }}>
            <PaginationTable />
        </PaginationTableContext.Provider>
    )
}

function toggleCompactTableForm(state, setter) {

    setter(!state);
}

