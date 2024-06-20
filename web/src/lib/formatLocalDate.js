
/**
 * 
 * @param {Date} date 
 */
export default function formatLocalDate(date) {

    if (!(date instanceof Date)) {

        return '';
    }

    return `${date.getDate()}-${date.getMonth()}-${date.getFullYear()}`
}