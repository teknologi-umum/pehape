import { count_chars } from "./count_chars";

const testString = "Testingjkjk123@.#";
const charByteObject: { [key: string]: number; } = {
    "T": 1, "e": 1, "s": 1, "t": 1, "i": 1, "n": 1, "g": 1,
    "j": 2, "k": 2, "1": 1, "2": 1, "3": 1, "@": 1, ".": 1, "#": 1
};

describe("count_chars", () => {
  it("should return all character counts as array on mode 0", () => {
    const result = count_chars(testString, 0);
    expect(result).toBeInstanceOf(Array);
    expect(result).toHaveLength(255);
    if (!Array.isArray(result)) return;
    result.forEach((count: number, i: number) => {
      const charFromIndex = String.fromCharCode(i);
      if (!charByteObject[charFromIndex]) {
        expect(count).toStrictEqual(0);
      } else {
        expect(count).toStrictEqual(charByteObject[charFromIndex]);
      }
    });
  });

  it("should return only character counts that are greater than 0 on mode 1", () => {
    const result = count_chars(testString, 1);
    Object.keys(result).forEach((key: string) => {
      const charFromIndex = String.fromCharCode(Number(key));
      expect(result[Number(key)]).toStrictEqual(charByteObject[charFromIndex]);
    });
  });

  it("should return only character counts that are equal to 0 on mode 2", () => {
    const result = count_chars(testString, 2);
    for (let i = 0; i < 255; i++) {
      const charFromIndex = String.fromCharCode(i);
      if (charByteObject[charFromIndex]) {
        expect(result[i]).toBeUndefined();
      } else {
        expect(result[i]).toStrictEqual(0);
      }
    }
  });

  it("should return a string containing all unique characters on mode 3", () => {
    expect(count_chars(testString, 3)).toStrictEqual("#.123@Tegijknst");
  });

  it("should return a string containing all unused characters on mode 4", () => {
    const result = count_chars(testString, 4);
    testString.split("").forEach(char => {
      expect(result).not.toMatch(char);
    });
  });

  it("should return empty array when given empty string on mode 0", () => {
    const result = count_chars("", 0);
    expect(result).toBeInstanceOf(Array);
    expect(result).toHaveLength(255);
    if (!Array.isArray(result)) return;
    result.forEach((count: number) => {
      expect(count).toStrictEqual(0);
    });
  });

  it("should return empty object when given empty string on mode 1", () => {
    expect(Object.keys(count_chars("", 1))).toHaveLength(0);
  });

  it("should return an object with 255 keys when given empty string on mode 2", () => {
    expect(Object.keys(count_chars("", 2))).toHaveLength(255);
  });

  it("should return an empty string when given empty string on mode 3", () => {
    expect(count_chars("", 3)).toStrictEqual("");
  });

  it("should return a string with 255 indices when given empty string on mode 4", () => {
    expect(count_chars("", 4)).toHaveLength(255);
  });

  it("should throw error on invalid mode values", () => {
    expect(() => count_chars(testString, 5)).toThrow(RangeError);
    expect(() => count_chars(testString, -1)).toThrow(RangeError);
    expect(() => count_chars(testString, 0.75)).toThrow(RangeError);
  });
});