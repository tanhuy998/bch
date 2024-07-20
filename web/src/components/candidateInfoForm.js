// import PromptFormInput from "./promptFormInput";


// export default function CandidateInfoForm({ formDelegator }) {

//     const { formVisible, setFormVisible, refreshTab } = useContext(CandidatesTabContext);

//     const thisYear = (new Date()).getFullYear();
//     const minDate = `1-1-${thisYear - 17}`;
//     const maxDate = `12-31-${thisYear + 10}`

//     formDelegator.setRefreshEmitter(refreshTab);

//     return (
//         <Form delegate={formDelegator} className="needs-validation" novalidate="" accept-charset="utf-8">
//             {/* <div class="container" > */}
//             <div class="row g-2">
//                 <div class="mb-3 col-md-6">
//                     {/* <label for="address" className="form-label">Candidate Name</label> */}
//                     <PromptFormInput
//                         validate={validatePeopleName}
//                         label="Candidate Name"
//                         invalidMessage="Tên chỉ chứa ký tự!"
//                         type="text"
//                         className="form-control"
//                         name="name"
//                         required="true"
//                     />
//                 </div>
//                 <div class="mb-3 col-md-6">
//                     {/* <label for="address" class="form-label">ID Number</label> */}
//                     <PromptFormInput
//                         validate={validateIDNumber}
//                         label="ID Number"
//                         invalidMessage="Số CCCD không hợp lệ!"
//                         type="text"
//                         className="form-control"
//                         name="idNumber"
//                     />
//                 </div>
//             </div>
//             <br />
//             <div class="row g-2">
//                 <div class="mb-3 col-md-6">

//                     <PromptFormInput
//                         name="dateOfBirth"
//                         label="Date Of Birth"
//                         validate={ageAboveSixteenAndYoungerThanTwentySeven}
//                         invalidMessage="Ngày sinh không hợp lệ!"
//                         type="date"
//                         value={minDate}
//                         min={minDate}
//                         max={maxDate}
//                         data-date-format="DD-MM-YYYY"
//                         className="form-control"
//                         required="true"
//                     />
//                 </div>
//                 <div class="mb-3 col-md-6">
//                     {/* <label for="address" class="form-label">Adress</label> */}
//                     <PromptFormInput
//                         validate={required}
//                         label="Address"
//                         invalidMessage="Addrress is required!s"
//                         type="text"
//                         className="form-control"
//                         name="address"
//                     />
//                 </div>
//             </div>
//             <br />
//             {/* </div> */}
//             <button type="submit" class="btn btn-primary"><i class="fas fa-save"></i> Save</button>
//             <button type="button" onClick={() => { toggleCompactTableForm(formVisible, setFormVisible) }} className="btn btn-sm btn-outline-primary" style={{ margin: '5px', paddingTop: "7px", paddingBottom: "7px" }}>Đóng</button>
//         </Form>
//     )
// }