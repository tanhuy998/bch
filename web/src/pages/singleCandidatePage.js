import { Link, useHref, useLocation, useNavigate, useParams, useResolvedPath } from "react-router-dom"
import BasicTab from "../components/basicTab";
import { createContext, memo, useContext, useEffect, useState } from "react";
import SingleCandidateUseCase from "../domain/usecases/singleCandidate.usecase";
import { candidate_model_t, candidate_signing_family_member_t } from "../domain/models/candidate.model";
import formatLocalDate from "../lib/formatLocalDate";
import QRCode from "react-qr-code";
import useBrowserLocation from "../hooks/useBrowserLocation";
import AppContext from "../contexts/app.context";
import useNotFoundRedirection from "../hooks/useNotFoundRedirection";
import ErrorResponse from "../backend/error/errorResponse";
import { useEndpointResponseErrorHandler } from "../hooks/handleEndpointError.lib";
import { type } from "@testing-library/user-event/dist/type";

/**
 * @typedef {inport("../domain/models/candidate.model.js")} candidate_model_t
 */

//const SignedInformations = memo(_SignedInformations);
const FETCH_INTERVAL_BEAT = 1000 * 60 * 10;
const PageContext = createContext({
    candidateData: null,
})

export default function SingleCandidatePage({ usecase }) {

    const { uuid } = useParams();
    const [currentDataFetchingInterval, setCurrentDataFetchingInterval] = useState(0);
    const [ /**@type {candidate_model_t} */ candidateData, setCandidateData] = useState(null);
    const responseErrorHandleFunc = useEndpointResponseErrorHandler();

    if (!(usecase instanceof SingleCandidateUseCase)) {

        throw new Error('invalid usecase passed to SingleCandidatePage as prop');
    }

    function fetchData(uuid) {

        setCurrentDataFetchingInterval(true);
        console.log('fetch')
        return usecase.read(uuid)
            .then(data => {

                setCandidateData(data);
                

                if (typeof data !== 'object') {

                    return;
                }
                console.log(data)
            })
            .catch(err => {

                if (
                    err instanceof ErrorResponse
                ) {

                    responseErrorHandleFunc(err);
                }

                alert(err?.messsage);
            });
    }

    useEffect(() => {

        fetchData(uuid)
        .then(() => {

            const interval = setInterval(fetchData, FETCH_INTERVAL_BEAT, uuid);
            setCurrentDataFetchingInterval(interval);
        })

        return () => {

            clearInterval(currentDataFetchingInterval);
        }

    }, []);

    const tabs = {
        'Signed Informations': <_SignedInformations usecase={usecase} candidateUUID={uuid}/>
    }

    return (
        <>
            <PageContext.Provider value={{ candidateData, usecase: usecase }}>
                <div className="card">
                    <div className="card-header">
                        Candidate Detail
                    </div>

                    <div className="card-body">
                        <CandidateDetail />
                    </div>
                    <div className="card-body">
                        <BasicTab tabs={tabs} />
                    </div>
                </div>
            </PageContext.Provider>
        </>
    )
}

function _SignedInformations({usecase, candidateUUID}) {

    // /**@type {candidate_model_t} */
    // const candidateData = useContext(PageContext)?.candidateData;
    // const signingInfo = candidateData?.signingInfo;

    const [signingInfo, setSigningInfo] = useState(undefined);
    const hasData = typeof signingInfo === 'object';

    useEffect(() => {

        usecase.getCandidateSigingInfo(candidateUUID)
                .then(data => {
                    setSigningInfo(data)
                })
                .catch(e => {


                })

    }, [])

    if (!hasData) {

        return (
            <>
                <h5>Candidate hasn't signed yet</h5>
            </>
        );
    }

    const civilIdentity = signingInfo?.civilIdentity;
    const education = signingInfo?.education;
    const politic = signingInfo?.politic;
    const family = signingInfo?.family;

    //console.log('signing info', hasData, signingInfo)
    return (
        <>
            <div className="card">
                <div className="card-body">
                    <Title>I. SƠ YẾU LÝ LỊCH</Title>
                    <br />
                    <div className="container">
                        <Row>
                            <Column>
                                Họ và Tên: <SigningInfo>{civilIdentity.name}</SigningInfo>
                            </Column>
                        </Row>
                        <Row>
                            <Column>Ngày sinh: <SigningInfo>{formatLocalDate(civilIdentity.dateOfBirth)}</SigningInfo></Column>
                            <Column>Giới tính: <SigningInfo>{civilIdentity.male ? 'Name' : 'Nữ'}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column>Số Căn Cước Công Dân: <SigningInfo>{civilIdentity.idNumber}</SigningInfo></Column>
                        </Row>
                        <br />
                        <Row>
                            <Column>Nơi Đăng Ký Khai sinh: <signingInfo>{civilIdentity.birthPlace}</signingInfo></Column>
                        </Row>
                        <Row>
                            <Column>Quê Quán: <SigningInfo>{civilIdentity.placeOfOrigin}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column>Dân Tôc: <SigningInfo>{civilIdentity.ethnicity}</SigningInfo></Column>
                            <Column>Tôn Giáo: <SigningInfo>{typeof civilIdentity.religion === 'string' && civilIdentity.religion != '' ? civilIdentity.religion : 'Không'}</SigningInfo></Column>
                            <Column>Quốc Tịch: <SigningInfo>{civilIdentity.nationality}</SigningInfo> </Column>
                        </Row>
                        <br />
                        <Row>
                            <Column>Nơi Thường Trú: <SigningInfo>{civilIdentity.permanentResident}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column>Nơi Ở Hiện Tại: <SigningInfo>{civilIdentity.currentResident}</SigningInfo></Column>
                        </Row>
                        <LineSeperator />
                        <Row>
                            <Column>Thành Phần Gia Đình: </Column>
                            <Column>Bản Thân: </Column>
                        </Row>
                        <br />
                        <Row>
                            <Column>Trình Độ Văn Hóa: <SigningInfo>{education.highestGrade}/12</SigningInfo></Column>
                            <Column>NĂm Tốt Nghiệp:<SigningInfo> {education.graduateAt}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column>Trình Độ Chuyên Môn: <SigningInfo>{education.college}</SigningInfo></Column>
                            <Column>Chuyên Ngành Đào Tạo: <SigningInfo>{education.expertise}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column>Ngoại Ngữ: </Column>
                        </Row>
                        <br />
                        <Row>
                            <Column>Ngày Vào Đảng CSVN: <SigningInfo>{formatLocalDate(politic.partyJoinDate)}</SigningInfo></Column>
                            <Column>Chính Thức: </Column>
                        </Row>
                        <Row>
                            <Column>Ngày Vào Đoàn TNCS Hồ Chí Minh: <SigningInfo>{politic.unionJoinDate}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column>Khen Thưởng: </Column>
                            <Column>Kỷ Luật: </Column>
                        </Row>
                        <LineSeperator />
                        <Row>
                            <Column>Nghề Nghiệp: <SigningInfo>{signingInfo.job}</SigningInfo></Column>
                            <Column>Lương: </Column>
                            <Column>Ngạch: </Column>
                            <Column>Bậc: </Column>
                        </Row>
                        <Row>
                            <Column>Nơi Làm Việc (Học Tập): <SigningInfo>{signingInfo.jobPlace}</SigningInfo></Column>
                        </Row>
                        <LineSeperator />
                        <Row>
                            <Column>Họ Tên Cha: <SigningInfo>{family?.father?.name}</SigningInfo></Column>
                            <Column>Tình Trạng (Sống, Chết): <SigningInfo>{family?.father?.dead ? 'Chết' : 'Sống'}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column>Sinh Ngày: <SigningInfo>{formatLocalDate(family?.father?.dateOfBirth)}</SigningInfo></Column>
                            <Column>Nghề Nghiệp: <SigningInfo>{family?.father?.job}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column><br /></Column>
                            <Column />
                            <Column />
                        </Row>
                        <Row>
                            <Column>Họ Tên Mẹ: <SigningInfo>{family?.mother?.name}</SigningInfo></Column>
                            <Column>Tình Trạng (Sống, Chết): <SigningInfo>{family?.mother?.dead ? 'Chết' : 'Sống'}</SigningInfo></Column>
                        </Row>
                        <Row>
                            <Column>Sinh Ngày: <SigningInfo>{formatLocalDate(family?.mother?.dateOfBirth)}</SigningInfo></Column>
                            <Column>Nghề Nghiệp: <SigningInfo>{family?.mother?.job}</SigningInfo></Column>
                        </Row>

                    </div>
                </div>
            </div>
            <div className="card">
                <div className="card-body">
                    <div className="container">
                        <Title>II. TÌNH HÌNH KINH TẾ CHÍNH TRỊ CỦA GIA ĐÌNH</Title>
                        <br />
                        <PoliticDetail header="Cha" member={family?.father} />
                        <br />
                        <PoliticDetail header="Mẹ" member={family?.mother} />
                        {
                            family?.anothers?.map(m => {

                                <>
                                    <PoliticDetail member={m} />
                                    <br />
                                </>
                            })
                        }
                    </div>
                </div>
            </div>
        </>
    )
}

function CandidateDetail() {

    const { uuid } = useParams();
    /**@type {candidate_model_t} */
    const candidateData = useContext(PageContext)?.candidateData;

    
    return (
        <>
            <h3>Candidate Detail</h3>
            <div className="container">
                <div className="row">
                    <div className="mb-3 col-md-4">
                        Name: <SigningInfo>{candidateData?.name}</SigningInfo>
                    </div>
                    <div className="mb-3 col-md-4">
                        ID Number: <SigningInfo>{candidateData?.idNumber}</SigningInfo>
                    </div>
                </div>
                <div className="row">
                    <div className="mb-3 col-md-4">
                        Date Of Birth: <SigningInfo>{formatLocalDate(candidateData?.dateOfBirth)}</SigningInfo>
                    </div>
                    <div className="mb-3 col-md-4">
                        address: <SigningInfo>{candidateData?.address}</SigningInfo>
                    </div>
                </div>
                <CandidateSigningSuppliment campaignUUID={candidateData?.campaignUUID} candidateUUID={uuid} />
            </div>
        </>
    )
}

function CandidateSigningSuppliment({campaignUUID, candidateUUID}) {
    
    const pageUsecase = useContext(PageContext)?.usecase;
    const navigate = useNavigate();
    const [isOpenSigning, setIsOpenSigning] = useState(false);
    const signingURL = useSigningURLGeneration(campaignUUID, candidateUUID);
    
    useEffect(() => {

        if (typeof campaignUUID !== 'string') {

            return;
        }

        pageUsecase?.isOpenSigning(campaignUUID, candidateUUID)
                    .then(openState => setIsOpenSigning(openState));

    }, [campaignUUID])

    if (!(pageUsecase instanceof SingleCandidateUseCase)) {

        return <></>
    }

    return (
        <>
        {
                isOpenSigning &&
                <>
                    <br />
                    <div style={{ height: "auto", margin: "0 auto", maxWidth: 256, width: "100%" }}>
                        <QRCode
                            style={{ height: "auto", maxWidth: "100%", width: "100%" }}
                            value={signingURL}
                            viewBox={`0 0 256 256`}
                        />
                    </div>
                    <br />
                    <button className="btn btn-sm btn-outline-primary" onClick={() => { navigate(`/signing/campaign/${campaignUUID}/candidate/${candidateUUID}`) }}>
                        Go To Signing Page <i className="fas fa-arrow-right"></i>
                    </button>
                    <div className="row">

                    </div>
                </>
        }
        </>
    )
}

function useSigningURLGeneration(campaignUUID, candidateUUID) {

    const {APP_HOST_NAME, APP_HTTP_PROTOCOL, APP_PORT} = useContext(AppContext);
    
    const location = window.location;
    const hostName = APP_HOST_NAME || location.hostname;
    const protocol = APP_HTTP_PROTOCOL || location.protocol;
    let port = APP_PORT || location.port
    port = port ? `:${port}` : "";

    return `${protocol}//${hostName}${port}/signing/campaign/${campaignUUID}/candidate/${candidateUUID}`
}

function Row({ children }) {

    return (
        <div className="row">
            {children}
        </div>
    )
}

function Column({ children }) {

    return (
        <div className="col">
            {children}
        </div>
    )
}

function Title({ children }) {

    return (
        <div className="card-title" style={{ textAlign: "center" }}>
            <h5>{children}</h5>
        </div>
    )
}

/**
 * 
 * 
 * @param {*} param0 
 */
function PoliticDetail({ member, header }) {

    /**@type {candidate_signing_family_member_t} */
    const m = member;
    const h = m?.politic?.history;

    return (
        <>
            <Row>
                <Column>{header}: <SigningInfo>{m?.name}</SigningInfo></Column>
                <Column>Ngày Sinh: <SigningInfo>{formatLocalDate(m?.dateOfBirth)}</SigningInfo></Column>
            </Row>
            <Row>
                <Column>Trước 1975: <SigningInfo>{h?.beforeReunification}</SigningInfo></Column>
            </Row>
            <Row>
                <Column>Sau 1975: <SigningInfo>{h?.afterReunification}</SigningInfo></Column>
            </Row>
            <Row>
                <Column>Nghề Nghiệp: <SigningInfo>{m?.job}</SigningInfo></Column>
            </Row>
        </>
    )
}

function LineSeperator() {

    return (
        <>
            <br />
            <div className="line"></div>
            <br />
        </>
    )
}

function SigningInfo({ children }) {

    return (
        <strong style={{
            fontSize: "20",
            textTransform: "uppercase"
        }}
        >
            {children}
        </strong>
    )
}