/**
 * ltrim — Strip whitespace (or other characters) from the beginning of a string.
 * @param str - The input string.
 * @param characters - You can also specify the characters you want to strip, by means of the characters parameter. Simply list all characters that you want to be stripped. With .. you can specify a range of characters.
 * @returns string with whitespace stripped from the beginning of string
 *   Without the second parameter, ltrim() will strip these characters:
 *   " " (ASCII 32 (0x20)), an ordinary space.
 *   "\t" (ASCII 9 (0x09)), a tab.
 *   "\n" (ASCII 10 (0x0A)), a new line (line feed).
 *   "\r" (ASCII 13 (0x0D)), a carriage return.
 *   "\0" (ASCII 0 (0x00)), the NUL-byte.
 *   "\v" (ASCII 11 (0x0B)), a vertical tab.
*/
export const ltrim = (str: string, characters = " \n\r\t\v\x00"): string => {
  let charactersEscaped = "";
  if (!characters.includes("..")) {
    charactersEscaped = characters;
  } else {
    const [start, end] = characters.split("..");
    for (let i = start.charCodeAt(0); i <= end.charCodeAt(0); i++) {
      charactersEscaped += String.fromCharCode(i);
    }
  }

  charactersEscaped = charactersEscaped.replace(/([\\^$.|?*+()[\]{}])/g, "\\$1");
  const charactersRegex = new RegExp(`^[${charactersEscaped}]+`);
  return str.replace(charactersRegex, "");
};