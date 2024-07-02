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
     *  @type {string}
     */
    this.idNumber;

    /**
     *  @type {Date}
     */
    this.dateOfBirth;
    
    /**
     *  @type {string}
     */
    this.name;

    /**
     *  @type {string}
     */
    this.address;

    /**
     * @type {candidate_signing_info_t}
     */
    this.signingInfo;
}

function civil_identity_t() {

    /**
     * @type {string}
     */
    this.idNumber;

    /**
     * @type {string}
     */
    this.name;

    /**
     * @type {Date}
     */
    this.dateOfBirth;

    /**
     * @type {string}
     */
    this.birthPlace;

    /**
     * @type {string}
     */
    this.ethnicity;

    /**
     *  @type {string}
     */
    this.religion;

    /**
     * @type {string}
     */
    this.permanentResident;

    /**
     * @type {string}
     */
    this.temporaryResident;

    /**
     * @type {string}
     */
    this.currentResident;

    /**
     * @type {boolean}
     */
    this.male = true;

    /**
     *  @type {string}
     */
    this.nationality;

    /**
     * @type {string}
     */
    this.placeOfOrigin;
}

function candidate_signing_education_t() {

    /**
     * @type {string}
     */
    this.primarySchool;

    /**
     * @type {string}
     */
    this.secondarySchool;

    /**
     * @type {string}
     */
    this.highSchool;

    /**
     * @type {string}
     */
    this.highestGrade;

    /**
     * @type {string}
     */
    this.college;

    /**
     * @type {Date}
     */
    this.graduateAt;

    /**
     * @type {string}
     */
    this.expertise;
}

function candidate_signing_politic_history_t() {

    /**
     *  @type {string}
     */
    this.afterReunification;

    /**
     *  @type {string}
     */
    this.beforeReunification;
}

function candidate_signing_politic_t() {

    /**
     * @type {Date?}
     */
    this.unionJoinDate;

    /**
     * @type {Date?}
     */
    this.partyJoinDate;

    /**@type {candidate_signing_politic_history_t} */
    this.history;
}


function candidate_signing_info_t() {

    /**
     * @type {string}
     */
    this.jobPlace;

    /**
     * @type {string}
     */
    this.job;

    /**
     * @type {civil_identity_t}
     */
    this.civilIdentity;

    /**
     * @type {candidate_signing_education_t}
     */
    this.education;

    /**
     * @type {candidate_signing_politic_t}
     */
    this.politic;

    /**
     * @type {candidate_signing_family_t}
     */
    this.family;
}

function candidate_signing_family_member_t() {

    /**
     * @type {string}
     */
    this.job;

    /**
     * @type {string}
     */
    this.name;

    /**
     * @type {Date}
     */
    this.dateOfBirth;

    /**
     * @type {boolean}
     */
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