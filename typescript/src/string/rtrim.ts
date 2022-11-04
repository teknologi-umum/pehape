/**
 * rtrim — Strip whitespace (or other characters) from the end of a string.
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
export const rtrim = (str: string, characters = " \n\r\t\v\x00"): string => {
  let charactersEscaped = "";
  const ranges = characters.match(/.\.\../g);
  if (!ranges) {
    charactersEscaped = characters;
  } else {
    ranges.forEach(range => {
      const [start, end] = range.split("..");
      let startCode = start.charCodeAt(0);
      let endCode = end.charCodeAt(0);
  
      if (startCode > endCode) {
        [startCode, endCode] = [endCode, startCode];
      }
      for (let i = startCode; i <= endCode; i++) {
        charactersEscaped += String.fromCharCode(i);
      }
    });
  }

  charactersEscaped = charactersEscaped.replace("\\", "\\\\").replace("]", "\\]");
  const charactersRegex = new RegExp(`[${charactersEscaped}]+$`);
  return str.replace(charactersRegex, "");
};