import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";

export default function JobSectionFormm({ delegator }) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput 
                        label="Nghề nghiệp" 
                        name="job" 
                        noticeText={"Công việc hiện tại, nếu còn đi học thì nhập \"Học sinh\" "}
                        textArea={true} 
                    />
                </div>
                <div className="col-md-6">
                    <PromptFormInput 
                        label="Nơi Làm Việc" 
                        name="jobPlace" 
                        noticeText={"Nơi làm việc hiện tại, nếu còn đi học thì nhập vào tên trường hiện tại đang theo học"}
                        textArea={true} 
                    />
                </div>
            </div>
        </Form>
    )
}