import { createContext, useContext, useEffect, useReducer, useRef, useState } from "react";
import CandidateSigningUseCase from "../domain/usecases/candidateSigning.usecase"
import { useParams } from "react-router-dom";
import PillTab from "../components/pillTab";
import Form from "../components/form";
import FormInput from "../components/formInput";
import BasicTab from "../components/basicTab";
import Tab from "../components/Tab";
import TabContext from "../contexts/tab.context";
import { pillTabStyle } from "../contexts/tab.context"
import TabEventContext from "../contexts/tabEvent.contex";
import FormCollector from "../components/formCollector";
import FormCollectorDispatchContext from "../contexts/formCollectorDispatch.context";

const SUBMIT_PHASE = Infinity;

const customTabContextvalue = {
    ...pillTabStyle,
    li: {
        ...pillTabStyle.li,
        style: {
            display: 'inline-block',
        }
    },
    ul: {
        ...pillTabStyle.ul,
        style: {
            backGround: 'rgba(255, 255, 255, 0)',

        }
    }
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
        <FormCollector>
            <div className="card">
                <div className="card-body">
                    <h4 className="card-title">Thông Tin Học Vấn</h4>
                    <br />
                    <EducationSectionForm name="2" />
                    <div className="row">
                        <div className="col">
                            <div className="line"></div>
                        </div>
                    </div>
                    <PoliticSectionForm name="2.1" />
                </div>
            </div>
        </FormCollector>
    )
}

const pagePhaseKeys = Object.keys(pagePhases);
const pagePhasesCount = pagePhaseKeys.length;
const PageControllerContext = createContext();

function resolveNextPhaseKey(currentTabKey) {

    const currentPhaseIndex = pagePhaseKeys.indexOf(currentTabKey);

    if (currentPhaseIndex < 0) {

        throw new Error('unknown page phase');
    }

    if (currentPhaseIndex === pagePhasesCount - 1) {

        return SUBMIT_PHASE;
    }

    return pagePhaseKeys[currentPhaseIndex + 1];
}

function isFirstPagePhase(tabKey) {

    return pagePhaseKeys.indexOf(tabKey) === 0;
}

export default function CandidateSigningPage({ usecase }) {

    const { campaignUUID, candidateUUID } = useParams();

    return (
        <PageController />
    )
}

function EducationSectionForm() {


}

function PoliticSectionForm() {

    return (
        <Form>

        </Form>
    )
}

function IdentittySectionForm({ name }) {

    return (
        <Form name={name}>
            <div className="row">
                <div className="col">
                    <label >Họ Và Tên</label>
                    <FormInput className="form-control" name="name" />
                </div>
                <div className="col"></div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    <label for="email" class="form-label">Ngày Sinh</label>
                    <FormInput type="date" className="form-control" name="dateOfBirth" placeholder="Email" required="" />
                    <small class="form-text text-muted">Enter a valid email address.</small>
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invalid-feedback">Please enter your email address.</div>
                </div>
                <div class="mb-3 col-md-6">
                    <label for="password" class="form-label">Giới Tính</label>
                    <FormInput type="select" className="form-control" name="male" required="" />
                    <small class="form-text text-muted">Your password must be 8-20 characters long, contain letters and numbers only.</small>
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invcurrentTab, setCurrentTabalid-feedback">Please enter your password.</div>
                </div>
            </div>
            <div className="row g-2">
                <div className="col-6">
                    <label >Số Căn Cước Công DÂn</label>
                    <FormInput className="form-control" name="name" />
                </div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    <label for="city" class="form-label">Nơi Đăng Ký Khai Sinh</label>
                    <input type="text" class="form-control" name="city" placeholder="City" required="" />
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invalid-feedback">Please enter your City.</div>
                </div>
                <div class="mb-3 col-md-4">
                    <label for="state" class="form-label">Quê Quán</label>
                    <select name="state" class="form-select" required="">
                        <option value="" selected="">Choose...</option>
                        <option value="1">New York</option>
                        <option value="2">Los Angeles</option>
                    </select>
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invalid-feedback">Please select a State.</div>
                </div>
                <div class="mb-3 col-md-2">
                    <label for="zip" class="form-label">Dân Tộc</label>
                    <input type="text" class="form-control" name="zip" placeholder="00000" required="" />
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invalid-feedback">Please enter a Zip code.</div>
                </div>
            </div>
            <div className="line"><br /></div>
            <br />
            <div className="row">
                <div className="col">
                    <label >Nơi Thường Trú</label>
                    <FormInput className="form-control" name="name" />
                </div>
                <div className="col"></div>
            </div>
            <div className="row">
                <div className="col">
                    <label >Nơi Ở Hiện Tại</label>
                    <FormInput className="form-control" name="name" />
                </div>
                <div className="col"></div>
            </div>
        </Form>
    )
}

function NextPhaseButton() {

    const { currentTabKey, setCurrentTabKey, focusPoint } = useContext(PageControllerContext);
    const { formCollectorResponse, setEmitSignal, formCollectorHandShake, setFormCollectorHandShake } = useContext(FormCollectorDispatchContext) || {};

    const hasSignalEmit = typeof setEmitSignal === 'function';
    const hasHandShakeSetter = typeof setFormCollectorHandShake === 'function';

    const isHandShaked = formCollectorHandShake === true;

    function dispatchNextPhase() {

        const nextPhaseKey = resolveNextPhaseKey(currentTabKey);
        console.log('current phase', currentTabKey, 'next phase', nextPhaseKey)
        if (nextPhaseKey === SUBMIT_PHASE) {

            return;
        }

        focusPoint?.current?.scrollIntoView({ behavior: 'smooth', block: 'start' });
        setCurrentTabKey(nextPhaseKey);
    }

    function handleClick() {

        if (
            hasSignalEmit 
            && isHandShaked
        ) {

            setEmitSignal(currentTabKey);
            return;
        }

        dispatchNextPhase();
    }

    useEffect(() => {

        if (formCollectorResponse === true) {

            dispatchNextPhase();
            setFormCollectorHandShake(false);
        }

    }, [formCollectorResponse]);

    return (
        <button type="button" onClick={handleClick} class="btn btn-outline-primary mb-2" value={currentTabKey}>Tiếp Theo</button>
    )
}

function PreviousPhaseButton() {

    return (
        <button></button>
    )
}

function PageController({ children }) {

    const [currentTabKey, setCurrentTabKey] = useState(pagePhaseKeys[0]);
    const [formCollectorHandShake, setFormCollectorHandShake] = useState(false);
    const [formCollectorResponse, setFormCollectorResponse] = useState();
    const [emitSignal, setEmitSignal] = useState();
    //const [enumerateTabKey, nextPhase] = useReducer(nextPagePhase, pagePhaseKeys[0]);

    const pageMainTab = useRef();
    // if (!(usecase instanceof CandidateSigningUseCase)) {

    //     throw new Error('');
    // }

    // useEffect(() => {

    //     setCurrentTabKey(enumerateTabKey);

    // }, [enumerateTabKey])
    console.log('phase', currentTabKey)
    return (
        <>
            <PageControllerContext.Provider value={{ currentTabKey, setCurrentTabKey, focusPoint: pageMainTab, }}>
                <FormCollectorDispatchContext.Provider 
                    value={{ 
                        emitSignal: null, 
                        setEmitSignal, 
                        formCollectorResponse, 
                        setFormCollectorResponse,
                        formCollectorHandShake,
                        setFormCollectorHandShake, 
                    }}
                >
                    <div ref={pageMainTab}>
                        {/* <CustomTab initTabIndex={currentTab} tabs={tabContents} /> */}
                        <TabEventContext.Provider value={{ onTabSwitch: (oldTabKey, newTabKey) => { setCurrentTabKey(newTabKey) } }}>
                            <TabContext.Provider value={{ ...customTabContextvalue, currentTab: currentTabKey, }}>
                                <Tab initTabKey={currentTabKey} tabs={pagePhases} />
                                <br />

                            </TabContext.Provider>
                        </TabEventContext.Provider>
                        <NextPhaseButton />
                    </div>
                </FormCollectorDispatchContext.Provider>
            </PageControllerContext.Provider>
        </>
    )
}