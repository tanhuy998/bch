import { useParams } from "react-router-dom"
import BasicTab from "../components/basicTab";
import { createContext, memo, useContext, useEffect, useState } from "react";
import SingleCandidateUseCase from "../domain/usecases/singleCandidate.usecase";
import { candidate_model_t, candidate_signing_family_member_t } from "../domain/models/candidate.model";
import formatLocalDate from "../lib/formatLocalDate";


/**
 * @typedef {inport("../domain/models/candidate.model.js")} candidate_model_t
 */

//const SignedInformations = memo(_SignedInformations);
const FETCH_INTERVAL_BEAT = 1000 * 60 * 10;
const PageContext = createContext({
    candidateData: null,
})

export default function SingleCandidatePage({usecase}) {

    const {uuid} = useParams();
    const [isFetchingData, setIsFetchingData] = useState(false);
    const [ /**@type {candidate_model_t} */ candidateData, setCandidateData] = useState(null);

    if (!(usecase instanceof SingleCandidateUseCase)) {

        throw new Error('invalid usecase passed to SingleCandidatePage as prop');
    }

    function fetchData(uuid) {

        setIsFetchingData(true);
        console.log('fetch')
        usecase.read(uuid)
            .then(data => {

                setCandidateData(data);
                setIsFetchingData(false);
                
                if (typeof data !== 'object') {

                    return;
                }
                console.log(data)
                setInterval(fetchData, FETCH_INTERVAL_BEAT, uuid);
            })
            .catch(err => {

                alert(err?.messsage);
            });
    }   

    useEffect(() => {

        // if (typeof candidateData === 'object') {

        //     return;
        // }

        fetchData(uuid);

    }, []);

    useEffect(() => {

        if (
            typeof candidateData === 'object'
            || isFetchingData
        ) {
            
            return;
        }

        fetchData(uuid);
    })

    const tabs = {
        'Signed Informations': <_SignedInformations />
    }

    return (
        <>
            <PageContext.Provider value={{candidateData}}>
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

function _SignedInformations() {

    /**@type {candidate_model_t} */
    const candidateData = useContext(PageContext)?.candidateData;
    const signingInfo = candidateData?.signingInfo;
    const hasData = typeof signingInfo === 'object';
    
    if (!hasData) {

        return <></>;
    }

    const civilIdentity = signingInfo.civilIdentity;
    const education = signingInfo.education;
    const politic = signingInfo.politic;
    const family = signingInfo.family;
    
    console.log('signing info', hasData, signingInfo)
    return (
        <>
            <div className="card">
                <div className="card-body">
                    <Title>I. SƠ YẾU LÝ LỊCH</Title>
                    <div className="container">
                        <Row>
                            <Column>
                                Họ và Tên: {civilIdentity.name}
                            </Column>
                        </Row>
                        <Row>
                            <Column>Ngày sinh: {formatLocalDate(civilIdentity.dateOfBirth)}</Column>
                            <Column>Giới tính: {civilIdentity.male ? 'Name' : 'Nữ'}</Column>
                        </Row>
                        <Row>
                            <Column>Số Căn Cước Công Dân: {civilIdentity.idNumber}</Column>
                        </Row>
                        <Row>
                            <Column>Nơi Đăng Ký Khai sinh: {civilIdentity.birthPlace}</Column>
                        </Row>   
                        <Row>
                            <Column>Quê Quán: {civilIdentity.placeOfOrigin}</Column>
                        </Row>
                        <Row>
                            <Column>Dân Tôc: {civilIdentity.ethnicity}</Column>
                            <Column>Tôn Giáo: {typeof civilIdentity.religion === 'string' && civilIdentity.religion != '' ? civilIdentity.religion : 'Không'}</Column>
                            <Column>Quốc Tịch: {civilIdentity.nationality} </Column>
                        </Row>
                        <Row>
                            <Column>Nơi Thường Trú: {civilIdentity.permanentResident}</Column>
                        </Row> 
                        <Row>
                            <Column>Nơi Ở Hiện Tại: {civilIdentity.currentResident}</Column>
                        </Row>
                        <Row>
                            <Column>Thành Phần Gia Đình: </Column>
                            <Column>Bản Thân: </Column>
                        </Row>
                        <Row>
                            <Column>Trình Độ Văn Hóa: {education.highestGrade}</Column>
                            <Column>NĂm Tốt Nghiệp: {education.graduateAt}</Column>
                        </Row>
                        <Row>
                            <Column>Trình Độ Chuyên Môn: {education.college}</Column>
                            <Column>Chuyên Ngành Đào Tạo: {education.expertise}</Column>
                        </Row>
                        <Row>
                            <Column>Ngoại Ngữ: </Column>
                        </Row>
                        <Row>
                            <Column>Ngày Vào Đảng CSVN: {formatLocalDate(politic.partyJoinDate)}</Column>
                            <Column>Chính Thức: </Column>
                        </Row>
                        <Row>
                            <Column>Ngày Vào Đoàn TNCS Hồ Chí Minh: {politic.unionJoinDate}</Column>
                        </Row>
                        <Row>
                            <Column>Khen Thưởng: </Column>
                            <Column>Kỷ Luật: </Column>
                        </Row>
                        <Row>
                            <Column>Nghề Nghiệp: {signingInfo.job}</Column>
                            <Column>Lương: </Column>
                            <Column>Ngạch: </Column>
                            <Column>Bậc: </Column>
                        </Row>
                        <Row>
                            <Column>Nơi Làm Việc (Học Tập): {signingInfo.jobPlace}</Column>
                        </Row>
                        <Row>
                            <Column>Họ Tên Cha: {family.father.name}</Column>
                            <Column>Tình Trạng (Sống, Chết): {family.father.dead ? 'Chết' : 'Sống'}</Column>
                        </Row>
                        <Row>
                            <Column>Sinh Ngày: {formatLocalDate(family.father.dateOfBirth)}</Column>
                            <Column>Nghề Nghiệp: {family.father.job}</Column>
                        </Row>
                        <Row>
                            <Column>Họ Tên Mẹ: {family.mother.name}</Column>
                            <Column>Tình Trạng (Sống, Chết): {family.mother.dead ? 'Chết' : 'Sống'}</Column>
                        </Row>
                        <Row>
                            <Column>Sinh Ngày: {formatLocalDate(family.mother.dateOfBirth)}</Column>
                            <Column>Nghề Nghiệp: {family.mother.job}</Column>
                        </Row>

                    </div>
                </div>
            </div>
            <div className="card">
                <div className="card-body">
                    <div className="container">
                        <Title>II. TÌNH HÌNH KINH TẾ CHÍNH TRỊ CỦA GIA ĐÌNH</Title>
                        <PoliticDetail header="Cha" member={family.father} />
                        <br/>
                        <PoliticDetail header="Mẹ" member={family.mother} />
                        {
                            family.anothers.map(m => {

                                <>
                                    <PoliticDetail member={m}/>
                                    <br/>
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

    /**@type {candidate_model_t} */
    const candidateData = useContext(PageContext)?.candidateData;

    return (
        <>
            <h3>Candidate Detail</h3>
            <div className="container">
                <div className="row">
                    <div className="col-6 col-md-4">
                        Name: {candidateData?.name}
                    </div>
                    <div className="col-6 col-md-4">
                        ID Number: {candidateData?.idNumber}
                    </div>
                </div>
                <div className="row">
                    <div className="col-6 col-md-4">
                        Date Of Birth: {formatLocalDate(candidateData?.dateOfBirth)}
                    </div>
                    <div className="col-6 col-md-4">
                        address: {candidateData?.address}
                    </div>
                </div>
            </div>
        </>
    )
}

function Row({children}) {

    return (
        <div className="row">
            {children}
        </div>
    )
}

function Column({children}) {

    return (
        <div className="col">
            {children}
        </div>
    )
}

function Title({children}) {

    return (
        <div className="card-title" style={{textAlign: "center"}}>
            <h4>{children}</h4>
        </div>
    )
}

/**
 * 
 * 
 * @param {*} param0 
 */
function PoliticDetail({member, header}) {

    /**@type {candidate_signing_family_member_t} */
    const m = member;
    const h = m?.politic?.history;

    return (
        <>
            <Row>
                <Column>{header}: {m?.name}</Column>
                <Column>Ngày Sinh: {formatLocalDate(m?.dateOfBirth)}</Column>
            </Row>
            <Row>
                <Column>Trước 1975: {h?.beforeReunification}</Column>
            </Row>
            <Row>
                <Column>Sau 1975: {h?.afterReunification}</Column>
            </Row>
            <Row>
                <Column>Nghề Nghiệp: {m.job}</Column>
            </Row>
        </>
    )
}