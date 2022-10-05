export const str_repeat = (
  str: string,
  multiplication: number | unknown
): string => {
  if (typeof multiplication === "number") {
    return str.repeat(multiplication);
  } else {
    throw new Error("Argument is must be number");
  }
};
