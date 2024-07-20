import { required } from "../components/lib/validator.";
import { ageAboveSixteenAndYoungerThanTwentySeven, validateIDNumber, validatePeopleName } from "../lib/validator";
import Form from "../components/form";
import PromptFormInput from "../components/promptFormInput";
import EditSingleCandidateUseCase from "../domain/usecases/editSingleCandidate.usecase";

export default function EditSingleCandidatePage({usecase}) {

    if (!(usecase instanceof EditSingleCandidateUseCase)) {

        throw new Error("invalid usecase passed to EditSingleCandidatePage that is not instance of EditSingleCandidateUseCase");
    }

    return (
        <>
            <div className="card">
                <div className="card-header">
                    Edit Candidate
                </div>

                <div className="card-body">
                    <CandidateInfoForm formDelegator={usecase.formDelegator} />
                </div>
            </div>
            
        </>
    )
}

function CandidateInfoForm({ formDelegator }) {

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
        </Form>
    )
}