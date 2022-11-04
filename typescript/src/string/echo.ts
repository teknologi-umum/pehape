/* eslint-disable no-console */

/**
 * echo — Strip whitespace (or other characters) from the end of a string.
 * @param str - The input string.
 * @param characters - You can also specify the characters you want to strip, by means of the characters parameter. Simply list all characters that you want to be stripped. With .. you can specify a range of characters.
 * @returns string with whitespace (or other characters) stripped from the end of string.
 *   Without the second parameter, rtrim() will strip these characters:
 *   " " (ASCII 32 (0x20)), an ordinary space.
 *   "\t" (ASCII 9 (0x09)), a tab.
 *   "\n" (ASCII 10 (0x0A)), a new line (line feed).
 *   "\r" (ASCII 13 (0x0D)), a carriage return.
 *   "\0" (ASCII 0 (0x00)), the NULL-byte.
 *   "\v" (ASCII 11 (0x0B)), a vertical tab.
*/
export const echo = (...expressions: string[]): void => {
  console.log(...expressions);
};