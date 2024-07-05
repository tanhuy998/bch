import { useEffect, useState } from "react";
import Form from "../../../components/form";
import PromptFormInput, { PromptSelectInput } from "../../../components/promptFormInput";
import SelectInput from "../../../components/selectInput";

export default function EducationSectionForm({ delegator }) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                {/* <div className="col-sm-2">Trình Độ Văn Hóa</div>
                <div className="col-md-1">
                    <PromptSelectInput 
                        className="form-select" 
                        name="highestGrade" 
                        //label="Trình Độ Văn Hóa" 
                        noticeText="Ví dụ: Hiện tại đang học lớp 11 thì trình độ văn hóa là 10."
                        defaultValue=""
                        castedType={Number}
                    >
                        <option></option>
                        {[1,2,3,4,5,6,7,8,9,10].map((grade) => <option value={grade}>{grade}</option>)}
                    </PromptSelectInput>
                </div> */}
                <HighestGradePromptSelectInput/>
            </div>
            <br/>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label="Trình Độ Chuyên Môn" name="expertise" type="text" />
                </div>
                <div className="col-md-6">
                    <PromptFormInput label="Chuyên Ngành Đào tạo" name="" type='text' />
                </div>
            </div>
        </Form>
    )
}

function HighestGradePromptSelectInput() {

    const [selected, setSelected] = useState(false);

    useEffect(() => {

        console.log(selected)

    }, [selected])

    return (
        <>
            <div className="col-sm-4">
                Trình Độ Văn Hóa
                <br/>
                {
                    !selected && <div class="invalid-feedback is-invalid">Trình độ học văn hóa là bắt buộc</div>
                }
                <small class="form-text text-muted">Ví dụ: Hiện tại đang học lớp 11 thì trình độ văn hóa là 10</small>
            </div>
            <div className="col-md-1">
                {/* <PromptFormInput label="Trình Độ Học Vấn" name="highestGrade" type="text" /> */}
                <SelectInput
                    className={"form-control form-select"}
                    style={!selected ? { borderColor: "#dc3545"} : {}}
                    name="highestGrade"
                    //label="Trình Độ Văn Hóa" 
                    defaultValue=""
                    onSelected={() => {setSelected(true)}}
                    onUnSelected={() => {setSelected(false)}}
                    castedType={Number}
                >
                    {[10, 9, 8, 7, 6, 5, 4, 3, 2, 1].map((grade) => <option value={grade}>{grade}</option>)}
                </SelectInput>
            </div>
        </>
    )
}