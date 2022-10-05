/**
 * String repeat
 * @param {string} str
 * @param {unknown} multiplication
 * @returns {string}
 */
export const str_repeat = (str: string, multiplication: unknown): string => {
  if (typeof multiplication === "number") {
    return str.repeat(multiplication);
  } else {
    throw new Error("Argument is must be number");
  }
};
