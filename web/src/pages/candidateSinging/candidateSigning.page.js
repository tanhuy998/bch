import { createContext, useContext, useEffect, useReducer, useRef, useState } from "react";
//import CandidateSigningUseCase from "../../domain/usecases/candidateSigning.usecase"
import { useParams } from "react-router-dom";
import PillTab from "../../components/pillTab";
import Form from "../../components/form";
import FormInput from "../../components/formInput";
import BasicTab from "../../components/basicTab";
import Tab from "../../components/Tab";
import TabContext from "../../contexts/tab.context";
import { pillTabStyle } from "../../contexts/tab.context"
import TabEventContext from "../../contexts/tabEvent.contex";
import FormCollector from "../../components/formCollector";
import FormCollectorDispatchContext from "../../contexts/formCollectorDispatch.context";
import PromptFormInput from "../../components/promptFormInput";
import CandidateSigningUseCase from "../../domain/usecases/candidateSigning.usecase"; 
import IdentittySectionForm from "./forms/IdentitySectionForm.form";
import PoliticSectionForm from "./forms/politicSectionForm.form";
import EducationSectionForm from "./forms/educationSectionForm.form";
import JobSectionFormm from "./forms/jobSectionForm.form";
import PageController from "./pageController";



export default function CandidateSigningPage({ usecase }) {

    const { campaignUUID, candidateUUID } = useParams();

    if (!(usecase instanceof CandidateSigningUseCase)) {

        throw new Error('invalid usecase passed to CandidateSingingPage');
    }

    const pagePhases = {
        '1': (
            <div className="card">
                <div className="card-body">
                    <h4 className="card-title">Thông Tin Định Danh</h4>
                    <br />
                    <IdentittySectionForm name="1" />
                </div>
            </div>
        ),
        '2': (
            // <FormCollector>
                <div className="card">
                    <div className="card-body">
                        <h4 className="card-title">Thông Tin Học Vấn</h4>
                        <br />
                        <EducationSectionForm delegator={usecase.candidateEducationFormDelegator} name="2" />
                        <div className="row">
                            <div className="col">

                            </div>
                        </div>
                        <br />
                        <div className="line"></div>
                        <br />
                        <JobSectionFormm delegator={usecase.candidateJobFormDelegator}/>
                        <div className="row">
                            <div className="col">
                                <div className="line"></div>
                            </div>
                        </div>
                        <PoliticSectionForm name="2.1" />
                    </div>
                </div>
            // </FormCollector>
        )
    }

    return (
        <PageController pagePhases={pagePhases}/>
    )
}