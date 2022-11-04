/* eslint-disable no-console */

/**
 * echo — Output one or more strings.
 * @param expressions - One or more string expressions to output, separated by commas. Non-string values will be coerced to strings, even when the strict_types directive is enabled.
*/
export const echo = (...expressions: string[]): void => {
  console.log(...expressions);
};