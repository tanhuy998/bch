import { required } from "../components/lib/validator.";

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

    return required(val) && REGEEX_PEOPLE_NAME.test(val);
}

export const validateFormalName = validatePeopleName;

/**
 * 
 * @param {Date} inputDate 
 * @returns {boolean}
 */
export function ageAboveSixteenAndYoungerThanTwentySeven(inputDate) {
    
    if (!(inputDate instanceof Date)) {
        
        //return false;
        inputDate = new Date(inputDate);
    }
    
    const thisYear = (new Date()).getFullYear();
    const age = thisYear - inputDate.getFullYear();

    return age >= 17 && age < 27;
}

/**
 * 
 * @param {Date|string} inputVal
 * @returns {boolean} 
 */
export function dayAfterNow(inputVal) {

    if (!(inputVal instanceof Date)) {

        //return false;
        inputVal = new Date(inputVal);
    }

    const today = new Date;

    return (
        // date.getDate() - today.getDate() >= 1
        // || date.getMonth() - today.getMonth() > 0
        // || date.getFullYear() - today.getFullYear() > 0

        inputVal.getFullYear() - today.getFullYear() > 0
        || inputVal.getMonth() - today.getMonth() > 0
        || inputVal.getDate() - today.getDate() > 0
    ) 
}