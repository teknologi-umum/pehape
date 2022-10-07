/**
 * chr — Generate a single-byte string from a number.
 * @param ascii - An integer between 0 and 255.
 * @returns Returns a single-character string containing the specified byte.
 */
export const chr = (ascii: number): string => {
  return String.fromCharCode(ascii);
};
