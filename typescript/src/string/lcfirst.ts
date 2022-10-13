/**
 * Make a string's first character lowercase.
 * @param {String} text The input text.
 * @returns {String} a string with the first character of string lowercased.
 */
export const lcfirst = (text: string): string => text.charAt(0).toLowerCase() + text.slice(1);