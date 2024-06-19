import { useParams } from "react-router-dom"
import BasicTab from "../components/basicTab";
import { memo } from "react";

const SignedInformations = memo(_SignedInformations);

export default function SingleCandidatePage({usecase}) {

    const {uuid} = useParams();

    const tabs = {
        'Signed Informations': <SignedInformations />
    }

    return (
        <>
            <div className="card">
                <div className="card-header">
                    Candidate Detail
                </div>
                <div className="card-body">
                    <CandidateDetail />
                </div>
                <div className="card-body">
                    <BasicTab tabs={tabs}/>
                </div>
            </div>
        </>
    )
}

function _SignedInformations() {

    return (
        <>
        </>
    )
}

function CandidateDetail() {

    return (
        <>
            <h3>Candidate Detail</h3>
            <div className="container">
                <div className="row">
                    <div className="col-6 col-md-4">
                        Name
                    </div>
                    <div className="col-6 col-md-4">
                        ID Number
                    </div>
                </div>
                <div className="row">
                    <div className="col-6 col-md-4">
                        Date Of Birth
                    </div>
                    <div className="col-6 col-md-4">
                        address
                    </div>
                </div>
            </div>
        </>
    )
}