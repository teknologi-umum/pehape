export const explode = (
  separator: string,
  text: string,
  limit = Number.MAX_SAFE_INTEGER
): string[] => {
  if (separator === "") throw new Error("Argument limit must not be empty string");
  if (limit === 0) limit = 1;
  return text.split(separator).slice(0, limit);
};