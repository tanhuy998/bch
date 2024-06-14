
/**
 * 
 * @param {Date} date 
 */
export default function formatLocalDate(date) {

    return `${date.getDate()}-${date.getMonth()}-${date.getFullYear()}`
}