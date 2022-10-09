/**
 * str_repeat — Repeat a string
 * @param str - The string to be repeated.
 * @param times - Number of time the `str` string should be repeated. `times` has to be greater than or equal to 0. If the `times` is set to 0, the function will return an empty string.
 * @returns Returns string repeated times.
 */
export const str_repeat = (str: string, times: number): string => {
  if (typeof times !== "number") throw new Error("Argument must be a number");
  if (times < 0) throw new Error("Invalid count value");
  return str.repeat(times);
};
