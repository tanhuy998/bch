import { useEffect, useState } from "react"
import { campaign_model_t } from "../domain/models/campaign.model";
import Form from "../components/form";
import FormInput from "../components/formInput";

export default function NewCampaignPage({ usecase }) {

    const [formData, setFormData] = useState(null);
    const [expireDate, setExpireDate] = useState();

    useEffect(() => {

        if (formData instanceof campaign_model_t) {

            return;
        }

    }, [formData])

    const expireDateThreshold = new Date();
    expireDateThreshold.setDate(expireDateThreshold.getDate() + 1);

    return (
        <div class="card">
                <div class="card-header">Form Validation</div>
                <div class="card-body">
                    <h3 class="card-title">Launch New Campaign</h3>
                    <br />
                    <Form handleFormData={() => {}} className="needs-validation" novalidate="" accept-charset="utf-8">    
                        <div class="mb-3">
                            <label for="address" className="form-label">Campaign Title</label>
                            <FormInput type="text" className="form-control" name="title" required="true" />
                            {/* <div class="valid-feedback">Looks good!</div>
                            <div class="invalid-feedback">Please enter your address.</div> */}
                        </div>
                        <div class="mb-3">
                            <label for="address" class="form-label">Description</label>
                            <textarea className="form-control" name="description">
                            </textarea>
                            {/* <div class="valid-feedback">Looks good!</div>
                            <div class="invalid-feedback">Please enter your address.</div> */}
                        </div>
                        <div class="row g-2">
                            <div class="mb-3 col-md-4">
                                <label for="state" className="form-label" >End Date</label>
                                <FormInput onChange={() => { }} name="expire" type="date" className="form-control" value={expireDateThreshold} required="true" min={expireDateThreshold} />
                            </div>
                        </div>
                        <br />
                        <button type="submit" class="btn btn-primary"><i class="fas fa-save"></i> Save</button>
                    </Form>
                </div>
        </div>
    )
}