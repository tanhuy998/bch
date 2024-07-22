import { required } from "../components/lib/validator.";
import { ageAboveSixteenAndYoungerThanTwentySeven, validateIDNumber, validatePeopleName } from "../lib/validator";
import Form from "../components/form";
import PromptFormInput from "../components/promptFormInput";
import EditSingleCandidateUseCase from "../domain/usecases/editSingleCandidate.usecase";
import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { candidate_model_t } from "../domain/models/candidate.model";

const NOT_FOUND_MODEL = Symbol('NOT_FOUND_MODEL');

export default function EditSingleCandidatePage({usecase}) {

    const { candidateUUID } = useParams();
    const [isDisabled, setIsDisabled] = useState(true);

    if (!(usecase instanceof EditSingleCandidateUseCase)) {

        throw new Error("invalid usecase passed to EditSingleCandidatePage that is not instance of EditSingleCandidateUseCase");
    }

    usecase.setCandidateUUID(candidateUUID);
    const candidateObj = useFetchCandidate(usecase);

    //const isDisabled = (!(candidateObj instanceof candidate_model_t));

    useEffect(() => {

        if (!(candidateObj instanceof candidate_model_t)) {

            return;
        }

        const dataModel = usecase.formDelegator.dataModel;
        
        
        Object.assign(dataModel, candidateObj)
        
        dataModel.uuid = undefined;
        dataModel.campaignUUID = undefined;
        dataModel.signingInfo = undefined;
        dataModel.ObjectID = undefined;

        console.log("model", dataModel)
        setIsDisabled(false);
        console.log('update edit page')

    }, [candidateObj])

    useEffect(() => {

        return () => {
            
            usecase.formDelegator.reset();
        }

    }, []);

    return (
        <>
            <div className="card">
                <div className="card-header">
                    Edit Candidate
                </div>

                <div className="card-body">
                    {!isDisabled && <CandidateInfoForm formDelegator={usecase.formDelegator} isDisabled={isDisabled} />}
                </div>
            </div>
            
        </>
    )
}

/**
 * 
 * @param {EditSingleCandidateUseCase} usecase 
 * @returns {candidate_model_t?}
 */
function useFetchCandidate(usecase) {

    const [candidate, setCandidate] = useState();

    useEffect(() => {

        //usecase.read(usecase.candidateUUID)
        usecase.fetchCandidate()
        .then((model) => {
            
            setCandidate(model)
        })
        .catch(() => {

            setCandidate(NOT_FOUND_MODEL)
        })
        
    }, [])

    return candidate;
}

function CandidateInfoForm({ formDelegator, isDisabled }) {

    const thisYear = (new Date()).getFullYear();
    const minDate = `1-1-${thisYear - 17}`;
    const maxDate = `12-31-${thisYear + 10}`

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
                        disable={isDisabled}
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
                        disable={isDisabled}
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
                        disable={isDisabled}
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
                        disable={isDisabled}
                    />
                </div>
            </div>
            <br />
            {/* </div> */}
            <button type="submit" class="btn btn-primary"><i class="fas fa-save"></i> Save</button>
        </Form>
    )
}