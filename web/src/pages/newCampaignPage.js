import { memo, useEffect, useState } from "react"
import { campaign_model_t } from "../domain/models/campaign.model";
import Form from "../components/form";
import FormInput from "../components/formInput";
import NewCampaignUseCase from "../domain/usecases/newCampaign.usecase";
import { required } from "../components/lib/validator.";

export default memo(NewCampaignPage);

function NewCampaignPage({ usecase }) {

    const [formData, setFormData] = useState(null);
    const [expireDate, setExpireDate] = useState();

    // useEffect(() => {

    //     if (formData instanceof campaign_model_t) {

    //         return;
    //     }

    // }, [formData])

    if (!(usecase instanceof NewCampaignUseCase)) {

        throw new Error('invalid usecase on NewCanpaignPage');
    }

    const expireDateThreshold = new Date();
    expireDateThreshold.setDate(expireDateThreshold.getDate() + 1);

    return (
        <div class="card">
                <div class="card-header">Form Validation</div>
                <div class="card-body">
                    <h3 class="card-title">Launch New Campaign</h3>
                    <br />
                    <Form interceptSubmit={(data) => {console.log('submit', data)}} dataModel={usecase.dataModel} className="needs-validation" novalidate="" accept-charset="utf-8">    
                        <div class="mb-3">
                            <label for="address" className="form-label">Campaign Title</label>
                            <FormInput validate={required} type="text" className="form-control" name="title" required="true" />
                        </div>
                        <div class="mb-3">
                            <label for="address" class="form-label">Description</label>
                            <FormInput validate={required} textArea={true} className="form-control" name="description"/>
                        </div>
                        <div class="row g-2">
                            <div class="mb-3 col-md-4">
                                <label for="state" className="form-label" >End Date</label>
                                <FormInput validate={usecase.campaignExpireDateValidateFunc} name="expire" type="date" className="form-control" value={expireDateThreshold} required="true" min={expireDateThreshold} />
                            </div>
                        </div>
                        <br />
                        <button type="submit" class="btn btn-primary"><i class="fas fa-save"></i> Save</button>
                    </Form>
                </div>
        </div>
    )
}