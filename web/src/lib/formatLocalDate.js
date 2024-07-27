
/**
 * 
 * @param {Date} date 
 */
export default function formatLocalDate(date) {

    if (!(date instanceof Date)) {

        return '';
    }

    return `${date.getDate()}-${date.getMonth()+1}-${date.getFullYear()}`
}

/**
 * 
 * @param {string} str 
 * @returns {string}
 */
export function strToLocalDate(str) {

    return formatLocalDate(new Date(str));
}