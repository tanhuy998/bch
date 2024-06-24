import CandidateEducationFormDelegator from "../valueObject/candidateEducationFormDelegator";
import CandidateJobFormDelegator from "../valueObject/candidateJobFormDelegator";

export default class CandidateSigningUseCase {

    #identityFormDelegator;
    #candidateEducationFormDelegator = new CandidateEducationFormDelegator();
    #candidateJobFormDelegator = new CandidateJobFormDelegator();

    get candidateJobFormDelegator() {

        return this.#candidateJobFormDelegator;
    }

    get candidateEducationFormDelegator() {

        return this.#candidateEducationFormDelegator;
    }

    submit() {


    }
}