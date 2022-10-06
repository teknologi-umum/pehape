/**
 * Replacing the single quote, double quote with backslash
 * @param {string} text
 * @returns {string}
 */
export const addslashes = (text: string): string => text.replace(/'|"/g, "\\'");