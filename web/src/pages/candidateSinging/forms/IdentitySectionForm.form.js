import Form from "../../../components/form";
import FormInput from "../../../components/formInput";
import { required } from "../../../components/lib/validator.";
import PromptFormInput, { PrompEthinicitySelectInput, PromptSelectInput } from "../../../components/promptFormInput";
import SelectInput from "../../../components/selectInput";
import { ageAboveSixteenAndYoungerThanTwentySeven, validateFormalName, validateIDNumber, validatePeopleName } from "../../../lib/validator";

function transformMaleInput(inputVal) {
    console.log('male', inputVal)
    return inputVal === '1' ? true : false;
}

export default function IdentittySectionForm({ name, delegator }) {

    return (
        <Form delegate={delegator} name={name}>
            <div className="row">
                <div className="mb-3 col-md-6">
                    <PromptFormInput label='Họ Và Tên' type='text' name="name" validate={validateFormalName} invalidMessage="Tên chỉ chứa ký tự"/>
                </div>
                <div class="mb-3 col-md-3">
                    <PromptFormInput 
                        label="Ngày Sinh" 
                        type="date" 
                        name="dateOfBirth"
                        invalidMessage="Ngày sinh không hợp lệ"
                        validate={ageAboveSixteenAndYoungerThanTwentySeven}
                    />
                </div>
                <div class="mb-3 col-md-2">
                    {/* <PromptFormInput label="Giới Tính" name="male"/> */}
                    <PromptSelectInput 
                        name="male"
                        label="Giới Tính" 
                        className="form-control form-select" 
                        notNull={true} 
                        castedType={transformMaleInput}
                        defaultValue="1"
                    >
                        <option value="1" selected>Nam</option>
                        <option value="0">Nữ</option>
                    </PromptSelectInput>
                </div>
            </div>
            <div className="row g-2">
                <div className="mb-3 col-md-6">
                    <PromptFormInput 
                        label="Số Căn Cước Công Dân" 
                        invalidMessage="Số căn cước công dân không hợp lệ"
                        validate={validateIDNumber} 
                        name="idNumber" 
                        type="text"
                    />
                </div>
                <div className="mb-3 col-md-3">
                    <PromptFormInput
                        label="Quốc Tịch"
                        invalidMessage="Quốc tịch không bỏ trống"
                        validate={validateFormalName}
                        defaultValue="Việt Nam"
                        name="nationality"
                        type="text"
                    />
                </div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    <PromptFormInput 
                        label="Nơi Đăng Ký Khai Sinh" 
                        validate={required} 
                        name="birthPlace" 
                        type="text"
                        invalidMessage="Nơi đăng ký khai sinh là bắt buộc"
                    />
                </div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-4">
                    <PromptFormInput 
                        label="Quê quán" 
                        validate={validateFormalName}
                        invalidMessage="Quê quán không bỏ trống" 
                        name="placeOfOrigin" 
                        type="text" 
                    />
                </div>
                <div class="mb-3 col-md-1">
                    {/* <PromptFormInput 
                        label="Dân Tộc" 
                        validate={validateFormalName} 
                        name="ethnicity" 
                        type="text" 
                        invalidMessage="Dân tộc không bỏ trống"
                    /> */}
                    <PrompEthinicitySelectInput className="form-control form-select" />
                </div>
                <div class="mb-3 col-md-4">
                    <PromptFormInput 
                        label="Tôn Giáo"
                        name="religion" 
                        type="text" 
                    />
                </div>
            </div>
            <div className="line"><br /></div>
            <br />
            <div className="row">
                <div className="mb-3 col-md-6">
                    <PromptFormInput 
                        label="Nơi Thường Trú" 
                        validate={required} 
                        name="permanentResident" 
                        type="text"
                        invalidMessage="Nơi thường trú là bắt buộc"
                    />
                </div>
                <div className="mb-3 col-md-6">
                    <PromptFormInput label="Nơi Ở Hiện Tại" name="currentResident" type="text" />
                </div>
                {/* <div className="col"></div> */}
            </div>
            <div className="row">
                
                <div className="col"></div>
            </div>
        </Form>
    )
}