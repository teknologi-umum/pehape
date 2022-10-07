/**
 * chr — Generate a single-byte string from a number
 * @param ascii - The number between 0 and 255.
 * @returns Returns a one-character string containing the character specified by interpreting ascii as an unsigned integer.
 */
export const chr = (ascii: number): string => {
  return String.fromCharCode(ascii);
};
