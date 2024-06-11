export {
    candidate_model_t,
    civil_identity_t,
    candidate_signing_education_t,
    candidate_signing_politic_history_t,
    candidate_signing_politic_t,
    candidate_signing_info_t,
    candidate_signing_family_member_t,
    candidate_signing_family_t,
}

function candidate_model_t() {

    /**
     * 
     */
    this.idNumber;

    /**
     * 
     */
    this.name;
    this.address;
}

function civil_identity_t() {


    this.idNumber;
    this.name;
    this.dateOfBirth;
    this.birthPlace;
    this.ethnicity;
    this.religion;
    this.permanentResident;
    this.temporaryResident;
    this.currentResident;
    this.male;
    this.nationality;
    this.placeOfOrigin;
}

function candidate_signing_education_t() {

    this.primarySchool;
    this.secondarySchool;
    this.highSchool;
    this.highestGrade;
    this.college;
    this.graduateAt;
    this.expertise;
}

function candidate_signing_politic_history_t() {

    this.afterReunification;
    this.beforeReunification;
}

function candidate_signing_politic_t() {

    this.unionJoinDate;
    this.partyJoinDate;

    /**@type {candidate_signing_politic_history_t} */
    this.history;
}


function candidate_signing_info_t() {

    this.jobPlace;
    this.job;

    /**
     * @type {civil_identity_t}
     */
    this.civilIdentity;
    /**
     * @type {}
     */
    this.education;
    this.politic;
}

function candidate_signing_family_member_t() {

    this.job;
    this.name;
    this.dateOfBirth;
    this.dead = false;

    /**
     * @type {candidate_signing_politic_t}
     */
    this.politic;
}

function candidate_signing_family_t() {

    /**
     * @type {candidate_signing_family_member_t}
     */
    this.mother;

    /**
     * @type {candidate_signing_family_member_t}
     */
    this.father;

    /**
     * @type {Array<candidate_signing_family_member_t>}
     */
    this.anothers = [];
}