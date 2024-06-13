
/**
 * 
 * @param {Date} date 
 */
export default function formatLocalDate(date) {

    return `${date.getDay()}-${date.getMonth()}-${date.getFullYear()}`
}