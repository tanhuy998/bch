import { memo, useEffect, useState } from "react";
import Form from "../components/form";
import FormInput from "../components/formInput";
import { required } from "../components/lib/validator.";
import "react-datepicker/dist/react-datepicker.css";
import NewCandidateUseCase from "../domain/usecases/newCandidate.usecase";
import { Link, useParams } from "react-router-dom";

export default memo(NewCandidatePage);

function NewCandidatePage({ usecase }) {

    const {campaignUUID} = useParams();
    const [formData, setFormData] = useState(null);
    const [expireDate, setExpireDate] = useState();

    // useEffect(() => {

    //     if (formData instanceof campaign_model_t) {

    //         return;
    //     }

    // }, [formData])

    if (!(usecase instanceof NewCandidateUseCase)) {

        throw new Error('invalid usecase on NewCanpaignPage');
    }

    return (
        <>
            <div class="card">
                <div class="card-header">Form Validation</div>
                <div class="card-body">
                    <Link to={`/admin/campaign/${campaignUUID}`} class="btn btn-sm btn-outline-primary float-end"><i class="fas fa-arrow-circle-left"></i> Quay về</Link>
                    <h3 class="card-title">New Candidate</h3>
                    <br />
                    <Form delegate={usecase} className="needs-validation" novalidate="" accept-charset="utf-8">
                        <div class="mb-3">
                            <label for="address" className="form-label">Candidate Name</label>
                            <FormInput validate={required} type="text" className="form-control" name="title" required="true" />
                        </div>
                        <div class="mb-3">
                            <label for="address" class="form-label">ID Number</label>
                            <FormInput validate={required} className="form-control" name="description" />
                        </div>
                        <div class="mb-3">
                            <label for="address" class="form-label">Adress</label>
                            <FormInput validate={required} className="form-control" name="description" />
                        </div>
                        <div class="row g-2">
                            <div class="mb-3 col-md-4">
                                <label for="state" className="form-label" >Date Of Birth</label>
                                <br />
                                <FormInput validate={usecase.campaignExpireDateValidateFunc} name="expire" type="date" data-date-format="DD-MM-YYYY" className="form-control" required="true" />
                                {/* <DatePicker className="form-control"/> */}
                            </div>
                        </div>
                        <br />
                        <button type="submit" class="btn btn-primary"><i class="fas fa-save"></i> Save</button>
                    </Form>
                </div>
            </div>
        </>
    )
}