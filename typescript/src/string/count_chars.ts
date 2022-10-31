/**
 * chr — Return information about characters used in a string.
 * @param string - The examined string.
 * @param mode - See return values.
 * @returns Depending on mode count_chars() returns one of the following:
 * 0 - an array with the byte-value as key and the frequency of every byte as value.
 * 1 - same as 0 but only byte-values with a frequency greater than zero are listed.
 * 2 - same as 0 but only byte-values with a frequency equal to zero are listed.
 * 3 - a string containing all unique characters is returned.
 * 4 - a string containing all not used characters is returned.
 */
export const count_chars = (string: string, mode: number): number[]|{[key: number]: number}|string => {
  switch (mode) {
    case 0: {
      return string.split("").reduce((prevArray: number[], char: string) => {
        const charByte = char.charCodeAt(0);
        prevArray[charByte]++;
        return prevArray;
      }, new Array(255).fill(0) as number[]);
    }
    case 1: {
      return string.split("").reduce((prevObject: {[key: number]: number}, char: string) => {
        const charByte = char.charCodeAt(0);
        prevObject[charByte] = !prevObject[charByte] ? 1 : prevObject[charByte] + 1;
        return prevObject;
      }, {});
    }
    case 2: {
      const byteObject: {[key: number]: number} = {};
      for (let i = 0; i < 255; i++) {
        byteObject[i] = 0;
      }
      string.split("").forEach((char: string) => {
          const charByte = char.charCodeAt(0);
          delete byteObject[charByte];
      });
      return byteObject;
    }
    case 3: {
      return Array.from(new Set(string)).sort((a, b) => a.charCodeAt(0) - b.charCodeAt(0)).join("");
    }
    case 4: {
      return new Array(255).fill(0).map((_, i) => {
        const char = String.fromCharCode(i);
        return string.includes(char) ? "" : char;
      }).join("");
    }
    default: {
      throw new RangeError("Argument #2 (mode) must be an integer between 1 and 4 (inclusive)");
    }
  } 
};