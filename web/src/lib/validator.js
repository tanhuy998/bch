const REGEX_ID_NUMBER = /^\d{12}$/;
const REGEEX_PEOPLE_NAME = /^[\p{Letter}\s]{3,}$/u ///^[a-zA-Z](([a-zA-Z])| ){5,}$/;

/**
 * 
 * @param {string} val 
 * @returns {boolean}
 */
export function validateIDNumber(val) {

    return REGEX_ID_NUMBER.test(val);
}

/**
 * 
 * @param {string} val 
 * @returns {boolean}
 */
export function validatePeopleName(val) {

    return REGEEX_PEOPLE_NAME.test(val);
}

export const validateFormalName = validatePeopleName;

/**
 * 
 * @param {Date} date 
 * @returns {boolean}
 */
export function ageAboveSixteenAndYoungerThanTwentySeven(date) {
    
    if (!(date instanceof Date)) {

        return false;
    }
    
    const thisYear = (new Date()).getFullYear();
    const age = thisYear - date.getFullYear();

    return age >= 17 && age < 27;
}

/**
 * 
 * @param {Date} date
 * @returns {boolean} 
 */
export function dayAfterNow(date) {

    if (!(date instanceof Date)) {

        return false;
    }

    const today = new Date;

    return (
        date.getDate() - today.getDate() >= 1
        || date.getMonth() - today.getMonth() > 0
        || date.getFullYear() - today.getFullYear() > 0
    ) 
}