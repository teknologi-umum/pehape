/**
 * str_repeat — Repeat a string
 * @param {string} str
 * @param {number} times
 * @returns {string} Returns string repeated times times.
 */
export const str_repeat = (str: string, times: number): string => {
  if (typeof times !== "number") throw new Error("Argument must be a number");
  return str.repeat(times);
};
