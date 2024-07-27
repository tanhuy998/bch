import { memo, useEffect, useState } from "react"
import { campaign_model_t } from "../domain/models/campaign.model";
import Form from "../components/form";
import FormInput from "../components/formInput";
import NewCampaignUseCase from "../domain/usecases/newCampaign.usecase";
import { required } from "../components/lib/validator.";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import PromptFormInput from "../components/promptFormInput";
import EditSingleCampaignUseCase from "../domain/usecases/editSingleCampaign.usecase";
import { useNavigate, useParams } from "react-router-dom";
import LoadingCircle from "../components/loadingCircle";
import { dayAfterNow } from "../lib/validator";
import BinaryCheckBoxFormInput from "../components/binaryCheckBoxFormInput";

const NOT_FOUND_MODEL = Symbol('not_found_model');

export default function EditSingleCampaignPage({ usecase }) {

    const navigator = useNavigate();
    const { campaignUUID } = useParams();

    if (!(usecase instanceof EditSingleCampaignUseCase)) {

        throw new Error('invalid usecase on NewCanpaignPage');
    }

    usecase.setCampaignUUID(campaignUUID);
    usecase.setNavigator(navigator);

    const campaignObject = useFetchCampaign(usecase);

    // useEffect(() => {

    //     if (
    //         !campaignObject ||
    //         campaignObject === NOT_FOUND_MODEL
    //     ) {

    //         return
    //     }

    //     const dataModel = usecase.dataModel;
        
    //     Object.assign(dataModel, campaignObject);
        
    // }, [campaignObject])

    return (
        <div class="card">
            <div class="card-header">Form Validation</div>
            <div class="card-body">
                {
                    (
                        !campaignObject && 
                        <>
                            <LoadingCircle size={50}/>
                        </>
                    )
                    ||
                    (
                        campaignObject === NOT_FOUND_MODEL &&
                        <h4>Not Found</h4>
                    )
                    ||
                    campaignObject instanceof campaign_model_t &&
                    <EditSingleCampaignForm usecase={usecase}/>
                }
            </div>
        </div>
    )
}


/**
 * 
 * @param {EditSingleCampaignUseCase} usecase 
 * @returns {candidate_model_t?}	
 */
function useFetchCampaign(usecase) {

    const [candidate, setCandidate] = useState();

    useEffect(() => {

        usecase.fetchCampaign()
            .then((model) => {
                
                if (model instanceof campaign_model_t) {

                    const dataModel = usecase.dataModel;

                    Object.assign(dataModel, model);
                }

                setCandidate(model)
            })
            .catch(() => {

                setCandidate(NOT_FOUND_MODEL)
            })

        return () => {

            usecase.reset();
        }

    }, [])

    return candidate;
}

function EditSingleCampaignForm({usecase}) {

    const expireDateThreshold = new Date();

    expireDateThreshold.setDate(expireDateThreshold.getDate() + 1);

    return (
        <>
            <h3 class="card-title">Modify Campaign</h3>
            <br />
            <Form delegate={usecase} className="needs-validation" novalidate="" accept-charset="utf-8">
                <div class="mb-3">
                    {/* <label for="address" className="form-label">Campaign Title</label> */}
                    <PromptFormInput
                        label="Campaign Tilte"
                        validate={required}
                        invalidMessage="Campaign title is required!"
                        type="text"
                        className="form-control"
                        name="title"
                        required="true"
                    />
                </div>
                <div class="row g-2">
                    <div class="mb-3 col-md-4">
                        <PromptFormInput
                            label="Issue Date"
                            validate={issueMustBePastOfExpire.bind(usecase)}
                            invalidMessage="Ngay bat dau phai la ngay truoc ngay ket thuc"
                            name="issueTime"
                            type="date"
                            className="form-control"
                        />
                    </div>
                    <div class="mb-3 col-md-4">
                        {/* <label for="state" className="form-label" >End Date</label> */}

                        <PromptFormInput
                            label="Deadline"
                            validate={expireMustBeFutureOfIssue.bind(usecase)}//{usecase.campaignExpireDateValidateFunc}
                            invalidMessage="Ngay ket thuc phai la ngay sau ngay bat dau"
                            name="expire"
                            type="date"
                            data-date-format="DD-MM-YYYY"
                            className="form-control"
                            value={expireDateThreshold}
                            required="true"
                            min={expireDateThreshold}
                        />
                        {/* <DatePicker className="form-control"/> */}
                    </div>
                    <div>
                        <BinaryCheckBoxFormInput className="form-check-input" />
                    </div>
                </div>
                <div class="mb-3">
                    {/* <label for="address" class="form-label">Description</label> */}
                    <PromptFormInput
                        label="Description"
                        textArea={true}
                        className="form-control"
                        name="description"
                    />
                </div>                
                <br />
                <button type="submit" class="btn btn-primary"><i class="fas fa-save"></i> Save</button>
            </Form>
        </>
    )
}

/**
 * @this EditSingleCampaignUseCase
 * @param {Date|string} expireValue 
 */
function expireMustBeFutureOfIssue(expireValue) {

    let issueTime = this.dataModel.issueTime;
    issueTime = issueTime instanceof Date ? issueTime : new Date(issueTime);
    expireValue = expireValue instanceof Date ? expireValue : new Date(expireValue)

    return (

        expireValue.getFullYear() - issueTime.getFullYear() > 0
        || expireValue.getMonth() - issueTime.getMonth() > 0
        || expireValue.getDate() - issueTime.getDate() > 0
    )
}

/**
 * @this EditSingleCampaignUseCase
 * @param {Date|string} issueValue 
 */
function issueMustBePastOfExpire(issueValue) {

    let expireTime = this.dataModel.expire;
    expireTime = expireTime instanceof Date ? expireTime : new Date(expireTime);
    issueValue = issueValue instanceof Date ? issueValue : new Date(issueValue)

    //const currentYear = (new Date()).getFullYear();

    return (

        issueValue.getFullYear() - expireTime.getFullYear() < 0
        || issueValue.getMonth() - expireTime.getMonth() < 0
        || issueValue.getDate() - expireTime.getDate() < 0
    )
}