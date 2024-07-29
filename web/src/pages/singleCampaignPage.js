import { Link, useNavigate, useParams } from "react-router-dom";
import PaginationTable from "../components/paginationTable";
import SingleCampaignUseCase from "../domain/usecases/singleCampaignUseCase.usecase";
import PaginationTableContext from "../contexts/paginationTable.context";
import { createContext, memo, useContext, useEffect, useReducer, useRef, useState } from "react";
import { tab } from "@testing-library/user-event/dist/tab";
import Tab from "../components/Tab";
import PillTab from "../components/pillTab";
import BasicTab from "../components/basicTab";
import Form from "../components/form";
import FormInput from "../components/formInput";
import { required } from "../components/lib/validator.";
import NewCandidateFormDelegator from "../domain/valueObject/newCandidateFormDelegator";
import PromptFormInput from "../components/promptFormInput";
import { ageAboveSixteenAndYoungerThanTwentySeven, validateIDNumber, validatePeopleName } from "../lib/validator";
import formatLocalDate, { strToLocalDate } from "../lib/formatLocalDate";
import CampaignProgressEndpoint from "../api/campaignProgress.api";

const CandidatesTabContext = createContext({
    formVisible: false,
    setFormVisible: null,
    refreshTab: () => { console.log('default refresh') },
})

const MemoNewCandidateForm = memo(NewCandidateForm);
const INIT_STATE = Symbol('_init_state_')
const COLUMN_TRANSFORM = {
    "dateOfBirth": strToLocalDate,
};

const PageLayoutContext = createContext({});

export default function SingleCampaignPage({ usecase }) {

    const { uuid } = useParams();
    const [/**@type {SingleCampaignRespnsePresenter?} */ campaignData, setCampaignData] = useState(null);
    
    if (!uuid) {

        throw new Error('invalid uuid')
    }

    if (!(usecase instanceof SingleCampaignUseCase)) {

        throw new Error('invalid use case passed to SingleCampaignPage');
    }

    usecase.newCandidateFormDelegator.setCampaignUUID(uuid);

    useEffect(() => {

        usecase.fetch(uuid)
            .then((data) => {
                console.log(data)
                setCampaignData(data)
            })

    }, [])

    // const candidateDisplayTable = (
    //     // <PaginationTable idField="uuid" exposedFields={['name', 'idNumber', 'address']} headers={['Tên', 'Số CCDD', 'Địa Chỉ']} endpoint={usecase.campaignCandidateListEndpoint} title="Canidates" />
    //     <PaginationTable />
    // )

    // const progressionTabs = {
    //     // All: //<CompactCampaignCandidateTable usecase={usecase} uuid={uuid} />,
    //     Signed: <CampaignCandidateProgression uuid={uuid} endpoint={usecase.campaignProgressEndpoint.candidate} />,
    //     Unsigned: <CampaignCandidateProgression uuid={uuid} endpoint={usecase.campaignCandidateListEndpoint} />,
    // }

    const mainTabs = {
        'Mô tả': `Lorem ipsum dolor sit amet. Eos neque dolor id accusantium reprehenderit et impedit ipsam quo illo praesentium. Ut porro amet aut autem ducimus At voluptas repellat. Est accusamus sequi est fuga voluptate ut aliquid minima 33 dolores nisi est maxime aspernatur qui sunt voluptatum. Ad sequi iure nam vero quis et aliquam repellat et eaque soluta et galisum quaerat ut rerum esse.

Qui nesciunt corporis in praesentium nemo ut vitae dolores et accusantium vitae. Ut nihil earum et iste quis id provident molestiae non voluptas veritatis ut ducimus similique et nulla pariatur et sint officia. In eius quibusdam qui facilis neque qui rerum consequatur. Qui reprehenderit possimus quo repellat quasi aut asperiores labore rem neque veritatis ut veniam distinctio est Quis dolor sit velit optio.

A nisi aspernatur non natus aliquam aut mollitia rerum. Non magnam aperiam quo eligendi veritatis sit eaque perferendis et atque modi aut cumque odio ex dolorum provident! Ut vero impedit ad voluptatem optio ut Quis velit. Sit suscipit dolor sit voluptatem voluptatum eum inventore tenetur est quas quia qui eligendi minima.`,
        Candidates: (
            <div class="card-body">
                {/* <h3 class="card-title">Candidates Detail</h3> */}

                {/* <PillTab tabs={candidateTabs} /> */}
                <CompactCampaignCandidateTable pageUsecase={usecase} formDelegator={usecase.newCandidateFormDelegator} endpoint={usecase.campaignCandidateListEndpoint} uuid={uuid} />
            </div>
        ),
        'Tiến Độ': (
            // <div class="card-body">
            //     <dib>
            //         Total candidate count:
            //     </dib>
            //     <PillTab tabs={progressionTabs} />
            // </div>
            <CampaignProgressionTab uuid={uuid} endpoint={usecase.campaignProgressEndpoint}/>
        )
    };



    const pageLayout = {
        mainTab: useRef(),
    }

    return (
        <>
            <PageLayoutContext.Provider value={pageLayout}>
                <div class="container">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="row">
                                <div class="col-md-12">
                                    <div class="card">
                                        <h5 class="card-header">Campaign Management</h5>
                                        <div class="card-body">
                                            <h1 class="card-title">{campaignData?.title || 'Unknown'}</h1>
                                            <br />
                                            {/* <div class="container"> */}
                                            <div class="row">
                                                <div class="col">
                                                    <h4 class="card-text">Ngày Bắt Đầu: {campaignData?.issueTime || ''}</h4>
                                                </div>
                                                <div class="col">
                                                    <h4 class="card-text">Ngày Kết Thúc: {campaignData?.expiredTime || ''}</h4>
                                                </div>
                                            </div>
                                            {/* </div> */}
                                            <br />

                                            <Link to={`/admin/campaign/edit/${uuid}`} class="btn btn-primary">Chỉnh sửa</Link>
                                            <br />
                                            <br />
                                            <div ref={pageLayout.mainTab}>
                                                <BasicTab tabs={mainTabs} initTabIndex={0} />
                                            </div>
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
                </div>
            </PageLayoutContext.Provider>
        </>
    )
}

function increaseAddedCandidateCount(state) {
    console.log('reducer', state)
    return state + 1;
}

function CompactCampaignCandidateTable({ uuid, endpoint, pageUsecase, formDelegator }) {

    const [formVisible, setFormVisible] = useState(false);
    const [submissionSuccess, setSubmissionSuccess] = useState(INIT_STATE);
    const [candidateAddedCount, addOneCandidate] = useReducer(increaseAddedCandidateCount, 0);
    //const addCandidateForm = useRef();
    const { mainTab } = useContext(PageLayoutContext);

    if (!(formDelegator instanceof NewCandidateFormDelegator)) {

        throw new Error('formDelegator must be instance of NewCandidateFormDelegator');
    }

    if (!(pageUsecase instanceof SingleCampaignUseCase)) {

        throw new Error('pageUsecase must be instance of SingleCampaignUseCase');
    }

    const allCandidateExtraContextValues = {
        EXTRA_FETCH_ARGS: [uuid]
    }

    const defaultTableContext = {
        columnTransform: COLUMN_TRANSFORM,
        idField: "uuid",
        exposedFields: ['name', 'dateOfBirth', 'idNumber', 'address'],
        headers: ['Tên', 'Ngày Sinh', 'Số CCDD', 'Địa Chỉ'],
        endpoint: endpoint,
        title: "Canidates",
        rowManipulator: endpoint,
    }

    // useEffect(() => {

    //     if (submissionSuccess === INIT_STATE) {

    //         return;
    //     }

    //     if (submissionSuccess === false) {

    //         return;
    //     }

    //     //setSubmissionSuccess(false);
    //     addOneCandidate();
    //     console.log('new candidate added', candidateAddedCount)
    // }, [submissionSuccess])

    return (
        <>
            <div>
                <div className={"collapse-wrapper" + (formVisible ? ' is-open' : '')} style={{ "background-color": '#E0E0E0', borderRadius: 6, }} onTransitionEnd={(e) => { e.propertyName === 'grid-template-rows' && formVisible && mainTab?.current.scrollIntoView({ behavior: "smooth", block: 'start' }); }}>
                    {/* {formVisible && ( */}
                    {(
                        <>
                            {/* {display: !formVisible ? 'none' : undefined } */}
                            <div className="collapse-content ">
                                <div className="card-body">
                                    <h3 class="card-title">New Candidate</h3>
                                    <br />
                                    <CandidatesTabContext.Provider value={{ formVisible, setFormVisible, refreshTab: addOneCandidate }}>
                                        <NewCandidateForm formDelegator={formDelegator} />
                                    </CandidatesTabContext.Provider>
                                </div>
                            </div>
                        </>
                    )}
                </div>
                {/* {!formVisible && (
                <>
                    
                </>
            )} */}
                <h5>
                    {formVisible && <br />}
                    All Candidates
                    <button style={{ display: formVisible ? 'none' : undefined }} onClick={() => { toggleCompactTableForm(formVisible, setFormVisible); }} class="btn btn-sm btn-outline-primary float-end"><i class="fas fa-plus-circle"></i> Thêm mới</button>
                    {/* <a href="users.html" class="btn btn-sm btn-outline-info float-end me-1"><i class="fas fa-angle-left"></i> <span class="btn-header">Return</span></a> */}
                </h5>
                <br />
                {/* {formVisible && <h5>All Candidate</h5>} */}
                <PaginationTableContext.Provider value={{ ...defaultTableContext, ...allCandidateExtraContextValues }}>
                    <PaginationTable key={candidateAddedCount} refresh={candidateAddedCount} rowManipulator={pageUsecase.candidateListTableRowManipulator} candidateAdded={candidateAddedCount} />
                </PaginationTableContext.Provider>
            </div>
        </>
    )
}

 function NewCandidateForm({ formDelegator }) {

    const { formVisible, setFormVisible, refreshTab } = useContext(CandidatesTabContext);

    const thisYear = (new Date()).getFullYear();
    const minDate = `1-1-${thisYear - 17}`;
    const maxDate = `12-31-${thisYear + 10}`

    formDelegator.setRefreshEmitter(refreshTab);

    return (
        <Form delegate={formDelegator} className="needs-validation" novalidate="" accept-charset="utf-8">
            {/* <div class="container" > */}
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    {/* <label for="address" className="form-label">Candidate Name</label> */}
                    <PromptFormInput
                        validate={validatePeopleName}
                        label="Candidate Name"
                        invalidMessage="Tên chỉ chứa ký tự!"
                        type="text"
                        className="form-control"
                        name="name"
                        required="true"
                    />
                </div>
                <div class="mb-3 col-md-6">
                    {/* <label for="address" class="form-label">ID Number</label> */}
                    <PromptFormInput
                        validate={validateIDNumber}
                        label="ID Number"
                        invalidMessage="Số CCCD không hợp lệ!"
                        type="text"
                        className="form-control"
                        name="idNumber"
                    />
                </div>
            </div>
            <br />
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    
                    <PromptFormInput
                        name="dateOfBirth"
                        label="Date Of Birth"
                        validate={ageAboveSixteenAndYoungerThanTwentySeven}
                        invalidMessage="Ngày sinh không hợp lệ!"
                        type="date"
                        value={minDate}
                        min={minDate}
                        max={maxDate}
                        data-date-format="DD-MM-YYYY"
                        className="form-control"
                        required="true"
                    />
                </div>
                <div class="mb-3 col-md-6">
                    {/* <label for="address" class="form-label">Adress</label> */}
                    <PromptFormInput
                        validate={required}
                        label="Address"
                        invalidMessage="Addrress is required!s"
                        type="text"
                        className="form-control"
                        name="address"
                    />
                </div>
            </div>
            <br />
            {/* </div> */}
            <button type="submit" class="btn btn-primary"><i class="fas fa-save"></i> Save</button>
            <button type="button" onClick={() => { toggleCompactTableForm(formVisible, setFormVisible) }} className="btn btn-sm btn-outline-primary" style={{ margin: '5px', paddingTop: "7px", paddingBottom: "7px" }}>Đóng</button>
        </Form>
    )
}

function CampaignCandidateProgression({ uuid, endpoint }) {
    
    const allCandidateExtraContextValues = {
        EXTRA_FETCH_ARGS: [`/${uuid}`]
    }

    const defaultTableContext = {
        columnTransform: COLUMN_TRANSFORM,
        idField: "uuid",
        exposedFields: ['name', 'dateOfBirth', 'idNumber', 'address'],
        headers: ['Tên', 'Ngay Sinh', 'Số CCDD', 'Địa Chỉ'],
        endpoint: endpoint,
        title: "Canidates",
        //rowManipulator: endpoint,
    }

    return (
        <PaginationTableContext.Provider value={{ ...defaultTableContext, ...allCandidateExtraContextValues }}>
            <PaginationTable />
        </PaginationTableContext.Provider>
    )
}

function CampaignProgressionTab({uuid, endpoint}) {

    const [progressionData, setProgressionData] = useState();

    useEffect(() => {

        endpoint.fetchReport(uuid)
        .then((res) => {
            
            setProgressionData(res.data);

        })
        .catch((e) => {})

    }, []); 

    const progressionTabs = {
        // All: //<CompactCampaignCandidateTable usecase={usecase} uuid={uuid} />,
        Signed: <CampaignCandidateProgression key="signed" uuid={uuid} endpoint={endpoint.signedCandidates} />,
        Unsigned: <CampaignCandidateProgression key="unsigned" uuid={uuid} endpoint={endpoint.unSignedCandidates} />,
    }

    const percentage = ((progressionData?.signedCount || 0) / (progressionData?.totalCount || 1)) * 100

    return (
        <div class="card-body">
            <dib>
                Total candidate count: {progressionData?.totalCount || 'N/A'}
            </dib>
            <div>
                Signed candidate count: {progressionData?.signedCount || 'N/A'}
            </div>
            <div>
                Percentage: {percentage?.toFixed(2) || 'N/A'}%
            </div>
            <br/>
            <PillTab tabs={progressionTabs} />
        </div>
    )    
}

function toggleCompactTableForm(state, setter) {

    setter(!state);
}

